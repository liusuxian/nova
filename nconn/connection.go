/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:00:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 22:18:22
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconn/connection.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconn

import (
	"context"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrequest"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
	"time"
)

// Connection 连接结构
type Connection struct {
	conn             net.Conn                      // 当前连接的 Socket 套接字
	connID           uint64                        // 当前连接的 ID，也可以称作为 SessionID，ID 全局唯一
	msgHandler       niface.IMsgHandle             // 消息管理和对应处理方法的消息管理模块
	ctx              context.Context               // 当前连接的 Context
	cancel           context.CancelFunc            // 当前连接的 Cancel Context
	msgBuffChan      chan []byte                   // 有缓冲的 Channel，用于读、写两个 Goroutine 之间的消息通信
	msgLock          sync.RWMutex                  // 收发消息的并发读写锁
	property         map[string]any                // 连接属性
	propertyLock     sync.Mutex                    // 连接属性的并发锁
	isClosed         bool                          // 当前连接的关闭状态
	connManager      niface.IConnManager           // 当前连接属于哪个 Connection Manager
	onConnStart      func(conn niface.IConnection) // 当前连接创建时的 Hook 函数
	onConnStop       func(conn niface.IConnection) // 当前连接断开时的 Hook 函数
	packet           niface.IDataPack              // 数据报文封包方式
	lastActivityTime time.Time                     // 最后一次活动时间
}

// newServerConn 创建一个 Server 服务端特性的连接
func newServerConn(server niface.IServer, conn net.Conn, connID uint64) *Connection {
	// 初始化 Connection 属性
	c := &Connection{
		conn:        conn,
		connID:      connID,
		msgBuffChan: nil,
		property:    nil,
		isClosed:    false,
	}
	c.packet = server.GetPacket()
	c.onConnStart = server.GetOnConnStart()
	c.onConnStop = server.GetOnConnStop()
	c.msgHandler = server.GetMsgHandler()
	// 将当前的 Connection 与 Server 的 ConnManager 绑定
	c.connManager = server.GetConnManager()
	// 将新创建的 Connection 添加到连接管理中
	server.GetConnManager().AddConn(c)
	return c
}

// newClientConn 创建一个 Client 客户端特性的连接
func newClientConn(client niface.IClient, conn net.Conn) *Connection {
	// 初始化 Connection 属性
	c := &Connection{
		conn:        conn,
		msgBuffChan: nil,
		property:    nil,
		isClosed:    false,
	}
	c.packet = client.GetPacket()
	c.onConnStart = client.GetOnConnStart()
	c.onConnStop = client.GetOnConnStop()
	c.msgHandler = client.GetMsgHandler()
	return c
}

// StartWriter 写消息 Goroutine，将数据发送给客户端
func (c *Connection) StartWriter() {
	nlog.Info(c.ctx, "Connection Writer Is Running")
	defer nlog.Info(c.ctx, "Connection Writer Exit !!!", zap.String("RemoteAddr", c.RemoteAddr().String()))

	for {
		select {
		case data, ok := <-c.msgBuffChan:
			if ok {
				// 有数据，发送给客户端
				if _, err := c.conn.Write(data); err != nil {
					nlog.Error(c.ctx, "Connection Writer Send Buff Data Error", zap.Error(err))
					return
				}
				// 发送给客户端成功, 更新连接活动时间
				c.updateActivity()
			} else {
				nlog.Error(c.ctx, "Connection Writer MsgBuffChan Is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// StartReader 读消息 Goroutine，从客户端读取数据
func (c *Connection) StartReader() {
	nlog.Info(c.ctx, "Connection Reader Is Running")
	defer nlog.Info(c.ctx, "Connection Reader Exit !!!", zap.String("RemoteAddr", c.RemoteAddr().String()))
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 读取客户端的消息头
			msgHead := make([]byte, c.packet.GetHeadLen())
			if _, err := io.ReadFull(c.conn, msgHead); err != nil {
				nlog.Error(c.ctx, "Connection Reader Read Msg Head Error", zap.Error(err))
				return
			}
			nlog.Debug(c.ctx, "Connection Reader Read Msg Head", zap.ByteString("MsgHead", msgHead))
			// 更新连接活动时间
			c.updateActivity()
			// 拆包
			msg, err := c.packet.Unpack(msgHead)
			if err != nil {
				nlog.Error(c.ctx, "Connection Reader Unpack Error", zap.Error(err))
				return
			}
			// 读取客户端的消息体
			var msgData []byte
			if msg.GetDataLen() > 0 {
				msgData = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.conn, msgData); err != nil {
					nlog.Error(c.ctx, "Connection Reader Read Msg Data Error", zap.Error(err))
					return
				}
			}
			msg.SetData(msgData)
			nlog.Debug(c.ctx, "Connection Reader Read Msg Data", zap.ByteString("MsgData", msgData))
			// 创建当前客户端请求的 Request 数据
			req := nrequest.NewRequest(c, msg)
			// 处理消息
			if nconf.WorkerPoolSize() > 0 {
				// 将消息交给 WorkerPool，由 Worker 进行处理
				c.msgHandler.SendMsgToWorkerPool(req)
			} else {
				// 马上以非阻塞方式处理消息
				go c.msgHandler.DoMsgHandler(req)
			}
		}
	}
}

// Start 启动连接
func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	// 执行创建连接时需要处理的业务
	c.callOnConnStart()
	// 启动读消息 Goroutine，从客户端读取数据
	go c.StartReader()

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

// Context 返回 Context，用于用户自定义的 Goroutine 获取连接退出状态
func (c *Connection) Context() context.Context {
	return c.ctx
}

// GetConnection 从当前连接获取原始的 Socket Conn
func (c *Connection) GetConnection() net.Conn {
	return c.conn
}

// GetConnID 获取当前 ConnID
func (c *Connection) GetConnID() uint64 {
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

// SendMsg 直接将 Message 数据发送给远程的客户端(无缓冲)
func (c *Connection) SendMsg(msgID uint32, data []byte) (err error) {
	c.msgLock.RLock()
	defer c.msgLock.RUnlock()
	// 判断当前连接的关闭状态
	if c.isClosed == true {
		err = errors.New("Connection Closed When Send Msg")
		return
	}
	// 封包
	var msg []byte
	if msg, err = c.packet.Pack(npack.NewMsgPackage(msgID, data)); err != nil {
		nlog.Error(c.ctx, "Connection Pack Msg Error", zap.Uint32("MsgID", msgID), zap.Error(err))
		err = errors.Wrap(err, "Connection Pack Msg Error")
		return
	}
	// 发送给客户端
	if _, err = c.conn.Write(msg); err != nil {
		nlog.Error(c.ctx, "Connection Send Msg Error", zap.Uint32("MsgID", msgID), zap.Error(err))
		return
	}
	// 发送给客户端成功, 更新连接活动时间
	c.updateActivity()
	return
}

// SendBuffMsg 直接将 Message 数据发送给远程的客户端(有缓冲)
func (c *Connection) SendBuffMsg(msgID uint32, data []byte) (err error) {
	c.msgLock.RLock()
	defer c.msgLock.RUnlock()
	// 启动写消息 Goroutine，将数据发送给客户端
	if c.msgBuffChan == nil {
		c.msgBuffChan = make(chan []byte, nconf.MaxMsgChanLen())
		go c.StartWriter()
	}
	// 创建定时器
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()
	// 判断当前连接的关闭状态
	if c.isClosed == true {
		err = errors.New("Connection Closed When Send Buff Msg")
		return
	}
	// 封包
	var msg []byte
	if msg, err = c.packet.Pack(npack.NewMsgPackage(msgID, data)); err != nil {
		nlog.Error(c.ctx, "Connection Pack Msg Error", zap.Uint32("MsgID", msgID), zap.Error(err))
		err = errors.Wrap(err, "Connection Pack Msg Error")
		return
	}
	// 发送超时
	select {
	case <-idleTimeout.C:
		err = errors.New("Connection Send Buff Msg Timeout")
		return
	case c.msgBuffChan <- msg:
		return
	}
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
	return time.Now().Sub(c.lastActivityTime) < nconf.HeartbeatMax()
}

// finalizer 清理器
func (c *Connection) finalizer() {
	// 执行连接断开时需要处理的业务
	c.callOnConnStop()

	c.msgLock.Lock()
	defer c.msgLock.Unlock()

	// 如果当前连接已经关闭
	if c.isClosed {
		return
	}
	// 关闭 Socket 连接
	_ = c.conn.Close()
	// 将当前连接从连接管理器中删除
	if c.connManager != nil {
		c.connManager.RemoveConn(c)
	}
	// 关闭当前连接全部管道
	if c.msgBuffChan != nil {
		close(c.msgBuffChan)
	}
	// 设置当前连接的关闭状态
	c.isClosed = true
	nlog.Info(c.ctx, "Connection Stop", zap.Uint64("ConnID", c.connID))
}

// callOnConnStart 调用连接创建时的 Hook 函数
func (c *Connection) callOnConnStart() {
	if c.onConnStart != nil {
		nlog.Info(c.ctx, "Connection CallOnConnStart...")
		c.onConnStart(c)
	}
}

// callOnConnStop 调用连接断开时的 Hook 函数
func (c *Connection) callOnConnStop() {
	if c.onConnStop != nil {
		nlog.Info(c.ctx, "Connection CallOnConnStop...")
		c.onConnStop(c)
	}
}

// updateActivity 更新连接活动时间
func (c *Connection) updateActivity() {
	c.lastActivityTime = time.Now()
}
