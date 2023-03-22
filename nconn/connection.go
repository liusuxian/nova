/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:00:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-22 21:35:35
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
	"go.uber.org/zap"
	"net"
	"sync"
	"time"
)

// Connection 连接结构
type Connection struct {
	conn              gnet.Conn                     // 当前连接的 Socket 套接字
	connID            int                           // 当前连接的 ID，也可以称作为 SessionID，ID 全局唯一
	msgHandler        niface.IMsgHandle             // 消息管理和对应处理方法的消息管理模块
	rootCtx           context.Context               // 当前 Server/Client 的根 Context
	ctx               context.Context               // 当前连接的 Context
	cancel            context.CancelFunc            // 当前连接的 Cancel Context
	sendMsgErrChan    chan error                    // 将 Message 数据发送给远程的对端时报错
	property          map[string]any                // 连接属性
	propertyLock      sync.Mutex                    // 连接属性的并发锁
	isClosed          bool                          // 当前连接的关闭状态
	connManager       niface.IConnManager           // 当前连接属于哪个 Connection Manager
	onConnStart       func(conn niface.IConnection) // 当前连接创建时的 Hook 函数
	onConnStop        func(conn niface.IConnection) // 当前连接断开时的 Hook 函数
	packet            niface.IDataPack              // 数据协议封包和拆包方式
	heartbeatInterval time.Duration                 // 心跳检测间隔时间
	lastActivityTime  time.Time                     // 最后一次活动时间
}

// NewServerConn 创建一个 Server 服务端特性的连接
func NewServerConn(ctx context.Context, server niface.IServer, conn gnet.Conn, heartbeatInterval time.Duration) *Connection {
	// 初始化 Connection 属性
	c := &Connection{
		conn:              conn,
		connID:            conn.Fd(),
		msgHandler:        server.GetMsgHandler(),
		rootCtx:           ctx,
		sendMsgErrChan:    make(chan error, 1),
		property:          nil,
		isClosed:          false,
		connManager:       server.GetConnManager(),
		onConnStart:       server.GetOnConnStart(),
		onConnStop:        server.GetOnConnStop(),
		packet:            server.GetPacket(),
		heartbeatInterval: heartbeatInterval,
	}
	// 将新创建的 Connection 添加到连接管理中
	server.GetConnManager().AddConn(c)
	return c
}

// NewClientConn 创建一个 Client 客户端特性的连接
func NewClientConn(ctx context.Context, client niface.IClient, conn gnet.Conn, heartbeatInterval time.Duration) *Connection {
	// 初始化 Connection 属性
	c := &Connection{
		conn:              conn,
		connID:            conn.Fd(),
		msgHandler:        client.GetMsgHandler(),
		rootCtx:           ctx,
		sendMsgErrChan:    make(chan error, 1),
		property:          nil,
		isClosed:          false,
		onConnStart:       client.GetOnConnStart(),
		onConnStop:        client.GetOnConnStop(),
		packet:            client.GetPacket(),
		heartbeatInterval: heartbeatInterval,
	}
	return c
}

// Start 启动连接
func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	// 调用连接创建时的 Hook 函数
	c.callOnConnStart()

	select {
	case <-c.ctx.Done():
		c.finalizer()
		return
	}
}

// Stop 停止连接
func (c *Connection) Stop() {
	c.cancel()
}

// GetCtx 返回 Context，用于用户自定义的 Goroutine 获取连接退出状态
func (c *Connection) GetCtx() context.Context {
	return c.ctx
}

// GetConnection 从当前连接获取原始的 gnet.Conn
func (c *Connection) GetConnection() gnet.Conn {
	return c.conn
}

// GetConnID 获取当前 ConnID
func (c *Connection) GetConnID() int {
	return c.connID
}

// RemoteAddr 获取当前连接远程地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// LocalAddr 获取当前连接本地地址信息
func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// SendMsg 将 Message 数据发送给远程的对端
func (c *Connection) SendMsg(msgID uint16, data []byte, callback ...gnet.AsyncCallback) (err error) {
	// 判断当前连接的关闭状态
	if c.isClosed {
		err = errors.New("Connection Closed When Send Msg")
		return
	}
	// 封包
	buf := c.packet.Pack(npack.NewMsgPackage(msgID, data))
	// 异步发送给客户端
	go func() {
		if len(callback) > 0 {
			c.sendMsgErrChan <- c.conn.AsyncWrite(buf, callback[0])
		} else {
			c.sendMsgErrChan <- c.conn.AsyncWrite(buf, nil)
		}
	}()
	// 定时器
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()
	// 接收错误
	select {
	case <-idleTimeout.C:
		nlog.Error(c.rootCtx, "Connection Send Msg Timeout", zap.Int("connID", c.connID), zap.Uint16("MsgID", msgID))
		err = errors.New("Connection Send Msg Timeout")
		return
	case err = <-c.sendMsgErrChan:
		if err != nil {
			nlog.Error(c.rootCtx, "Connection Send Msg Error", zap.Int("connID", c.connID), zap.Uint16("MsgID", msgID), zap.Error(err))
			return
		}
	case <-c.ctx.Done():
		nlog.Info(c.rootCtx, "Connection Closed When Send Msg", zap.Int("connID", c.connID))
		err = errors.New("Connection Closed When Send Msg")
		return
	}
	// 发送给客户端成功, 更新连接活动时间
	c.UpdateActivity()
	nlog.Debug(c.rootCtx, "Connection Send Msg Succeed", zap.Int("connID", c.connID))
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
	if value, ok = c.property[key]; ok {
		return
	}

	err = errors.New("Connection No Property Found")
	return
}

// RemoveProperty 移除当前连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

// IsAlive 判断当前连接是否存活
func (c *Connection) IsAlive() bool {
	if c.isClosed {
		return false
	}
	// 检查连接最后一次活动时间，如果超过心跳间隔，则认为连接已经死亡
	return time.Now().Sub(c.lastActivityTime) < c.heartbeatInterval
}

// UpdateActivity 更新连接活动时间
func (c *Connection) UpdateActivity() {
	c.lastActivityTime = time.Now()
}

// finalizer 清理器
func (c *Connection) finalizer() {
	// 调用连接断开时的 Hook 函数
	c.callOnConnStop()
	// 如果当前连接已经关闭
	if c.isClosed {
		return
	}
	// 关闭 Socket 连接
	_ = c.conn.Close()
	// 将当前连接从连接管理器中删除
	if c.connManager != nil {
		c.connManager.RemoveConn(c.connID)
	}
	// 关闭当前连接全部管道
	if c.sendMsgErrChan != nil {
		close(c.sendMsgErrChan)
	}
	// 设置当前连接的关闭状态
	c.isClosed = true
	nlog.Info(c.rootCtx, "Connection Stop", zap.Int("ConnID", c.connID))
}

// callOnConnStart 调用连接创建时的 Hook 函数
func (c *Connection) callOnConnStart() {
	if c.onConnStart != nil {
		nlog.Info(c.rootCtx, "Connection CallOnConnStart...", zap.Int("connID", c.connID))
		c.onConnStart(c)
	}
}

// callOnConnStop 调用连接断开时的 Hook 函数
func (c *Connection) callOnConnStop() {
	if c.onConnStop != nil {
		nlog.Info(c.rootCtx, "Connection CallOnConnStop...", zap.Int("connID", c.connID))
		c.onConnStop(c)
	}
}
