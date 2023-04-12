/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 13:23:48
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-12 18:26:28
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconn/connection.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconn

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/panjf2000/gnet/v2"
	"github.com/pkg/errors"
	"net"
	"sync"
	"time"
)

// Connection 连接结构
type Connection struct {
	conn             gnet.Conn                     // 当前连接的 Socket 套接字
	connID           int                           // 当前连接的 ID，也可以称作为 SessionID，ID 全局唯一
	msgHandler       niface.IMsgHandle             // 消息管理和对应处理方法的消息管理模块
	cancelCtx        context.Context               // 当前连接的 Cancel Context
	cancelFunc       context.CancelFunc            // 当前连接的 Cancel Func
	property         map[string]any                // 连接属性
	propertyLock     *sync.Mutex                   // 连接属性的互斥锁
	isClosed         bool                          // 当前连接的关闭状态
	connManager      niface.IConnManager           // 当前连接属于哪个 Connection Manager
	onConnStart      func(conn niface.IConnection) // 当前连接创建时的 Hook 函数
	onConnStop       func(conn niface.IConnection) // 当前连接断开时的 Hook 函数
	packet           niface.IDataPack              // 数据协议封包和拆包方式
	lastActivityTime time.Time                     // 最后一次活动时间
	lastActivityLock *sync.Mutex                   // 最后一次活动时间的互斥锁
	maxHeartbeat     time.Duration                 // 最长心跳检测间隔时间
	heartbeatChecker niface.IHeartBeatChecker      // 心跳检测器
}

// NewServerConn 创建一个 Server 服务端特性的连接
func NewServerConn(server niface.IServer, conn gnet.Conn, maxHeartbeat time.Duration) (c *Connection) {
	// 初始化 Connection 属性
	c = &Connection{
		conn:             conn,
		connID:           conn.Fd(),
		msgHandler:       server.GetMsgHandler(),
		property:         nil,
		propertyLock:     new(sync.Mutex),
		isClosed:         false,
		connManager:      server.GetConnManager(),
		onConnStart:      server.GetOnConnStart(),
		onConnStop:       server.GetOnConnStop(),
		packet:           server.GetPacket(),
		lastActivityLock: new(sync.Mutex),
		maxHeartbeat:     maxHeartbeat,
	}
	// 从当前 Server 克隆心跳检测器
	heartbeatChecker := server.GetHeartBeat()
	if heartbeatChecker != nil {
		// 绑定连接
		heartbeatChecker.Clone().BindConn(c)
	}
	// 将新创建的 Connection 添加到连接管理中
	server.GetConnManager().AddConn(c)
	return
}

// NewClientConn 创建一个 Client 客户端特性的连接
func NewClientConn(client niface.IClient, conn gnet.Conn, maxHeartbeat time.Duration) (c *Connection) {
	// 初始化 Connection 属性
	c = &Connection{
		conn:             conn,
		connID:           conn.Fd(),
		msgHandler:       client.GetMsgHandler(),
		property:         nil,
		propertyLock:     new(sync.Mutex),
		isClosed:         false,
		onConnStart:      client.GetOnConnStart(),
		onConnStop:       client.GetOnConnStop(),
		packet:           client.GetPacket(),
		lastActivityLock: new(sync.Mutex),
		maxHeartbeat:     maxHeartbeat,
	}
	// 从当前 Client 克隆心跳检测器
	heartbeatChecker := client.GetHeartBeat()
	if heartbeatChecker != nil {
		// 绑定连接
		heartbeatChecker.Clone().BindConn(c)
	}
	return c
}

// Start 启动连接
func (c *Connection) Start() {
	nlog.Info("Connection Start", nlog.Int("ConnID", c.connID))

	c.cancelCtx, c.cancelFunc = context.WithCancel(context.Background())
	// 调用连接创建时的 Hook 函数
	c.callOnConnStart()

	// 启动心跳检测
	if c.heartbeatChecker != nil {
		// 启动心跳检测
		c.heartbeatChecker.Start()
		// 更新连接活动时间
		c.UpdateActivity()
	}

	<-c.cancelCtx.Done()
	// 清理
	c.finalizer()
}

// Stop 停止连接
func (c *Connection) Stop() {
	c.cancelFunc()
}

// GetCancelCtx 返回 Cancel Context，用于用户自定义的 Goroutine 获取连接退出状态
func (c *Connection) GetCancelCtx() (ctx context.Context) {
	return c.cancelCtx
}

// GetConnection 从当前连接获取原始的 gnet.Conn
func (c *Connection) GetConnection() (conn gnet.Conn) {
	return c.conn
}

// GetConnID 获取当前 ConnID
func (c *Connection) GetConnID() (connID int) {
	return c.connID
}

// RemoteAddr 获取当前连接远程地址信息
func (c *Connection) RemoteAddr() (addr net.Addr) {
	return c.conn.RemoteAddr()
}

// LocalAddr 获取当前连接本地地址信息
func (c *Connection) LocalAddr() (addr net.Addr) {
	return c.conn.LocalAddr()
}

// Send 将数据发送给远程的对端
func (c *Connection) Send(data []byte, callback ...gnet.AsyncCallback) (err error) {
	// 判断当前连接的关闭状态
	if c.isClosed {
		err = errors.New("connection closed when send data")
		return
	}
	// 异步发送给客户端
	if len(callback) > 0 {
		err = c.conn.AsyncWrite(data, callback[0])
	} else {
		err = c.conn.AsyncWrite(data, nil)
	}
	return
}

// SendMsg 将 Message 数据发送给远程的对端
func (c *Connection) SendMsg(msgID uint16, data []byte, callback ...gnet.AsyncCallback) (err error) {
	// 判断当前连接的关闭状态
	if c.isClosed {
		err = errors.New("connection closed when send msg")
		return
	}
	// 封包
	buf := c.packet.Pack(npack.NewMsgPackage(msgID, data))
	// 异步发送给客户端
	if len(callback) > 0 {
		err = c.conn.AsyncWrite(buf, callback[0])
	} else {
		err = c.conn.AsyncWrite(buf, nil)
	}
	return
}

// SetProperty 设置当前连接属性
func (c *Connection) SetProperty(key string, value any) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if c.property == nil {
		c.property = make(map[string]any)
	}
	c.property[key] = value
}

// GetProperty 获取当前连接属性
func (c *Connection) GetProperty(key string) (value any, err error) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	var ok bool
	if value, ok = c.property[key]; !ok {
		err = errors.New("connection no property found")
		return
	}
	return
}

// RemoveProperty 移除当前连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

// IsAlive 判断当前连接是否存活
func (c *Connection) IsAlive() (isAlive bool) {
	if c.isClosed {
		return false
	}
	// 检查连接最后一次活动时间，如果超过心跳间隔，则认为连接已经死亡
	c.lastActivityLock.Lock()
	t := c.lastActivityTime
	c.lastActivityLock.Unlock()
	return time.Since(t) < c.maxHeartbeat
}

// UpdateActivity 更新连接活动时间
func (c *Connection) UpdateActivity() {
	c.lastActivityLock.Lock()
	defer c.lastActivityLock.Unlock()

	c.lastActivityTime = time.Now()
}

// SetHeartBeat 设置心跳检测器
func (c *Connection) SetHeartBeat(checker niface.IHeartBeatChecker) {
	c.heartbeatChecker = checker
}

// finalizer 清理器
func (c *Connection) finalizer() {
	// 调用连接断开时的 Hook 函数
	c.callOnConnStop()
	// 如果当前连接已经关闭
	if c.isClosed {
		return
	}
	// 关闭链接绑定的心跳检测器
	if c.heartbeatChecker != nil {
		c.heartbeatChecker.Stop()
	}
	// 关闭 Socket 连接
	_ = c.conn.Close()
	// 将当前连接从连接管理器中删除
	if c.connManager != nil {
		c.connManager.RemoveConn(c.connID)
	}
	// 设置当前连接的关闭状态
	c.isClosed = true
	nlog.Info("Connection Stop", nlog.Int("ConnID", c.connID))
}

// callOnConnStart 调用连接创建时的 Hook 函数
func (c *Connection) callOnConnStart() {
	if c.onConnStart != nil {
		nlog.Info("Connection CallOnConnStart...", nlog.Int("connID", c.connID))
		c.onConnStart(c)
	}
}

// callOnConnStop 调用连接断开时的 Hook 函数
func (c *Connection) callOnConnStop() {
	if c.onConnStop != nil {
		nlog.Info("Connection CallOnConnStop...", nlog.Int("connID", c.connID))
		c.onConnStop(c)
	}
}
