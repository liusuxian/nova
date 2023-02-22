/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:00:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 22:30:32
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/tcp_connection.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nnet

import (
	"context"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
	"time"
)

// TCPConnection TCP连接结构
type TCPConnection struct {
	TCPServer   niface.IServer    // 当前Conn属于哪个Server
	Conn        *net.TCPConn      // 当前连接的socket TCP套接字
	ConnID      uint32            // 当前连接的ID，也可以称作为SessionID，ID全局唯一
	MsgHandler  niface.IMsgHandle // 消息管理MsgID和对应处理方法的消息管理模块
	ctx         context.Context   // 告知该链接已经退出/停止的channel
	cancel      context.CancelFunc
	msgBuffChan chan []byte // 有缓冲管道，用于读、写两个goroutine之间的消息通信
	sync.RWMutex
	property     map[string]any // 链接属性
	propertyLock sync.Mutex     // 保护当前property的锁
	isClosed     bool           // 当前连接的关闭状态
}

// Start 启动连接
func (c *TCPConnection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	// 按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.TCPServer.CallOnConnStart(c)
	// 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.finalizer()
		return
	}
}

// Stop 停止连接
func (c *TCPConnection) Stop() {
	c.cancel()
}

// Context 返回ctx，用于用户自定义的go程获取连接退出状态
func (c *TCPConnection) Context() context.Context {
	return c.ctx
}

// GetConnection 从当前连接获取原始的socket Conn
func (c *TCPConnection) GetConnection() net.Conn {
	return c.Conn
}

// GetConnID 获取当前连接ID
func (c *TCPConnection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端地址信息
func (c *TCPConnection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 直接将数据发送给远程的客户端(无缓冲)
func (c *TCPConnection) SendMsg(msgID uint32, data []byte) (err error) {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		err = errors.New("Connection Closed When Send Msg")
		return
	}
	// 将data封包，并且发送
	dp := c.TCPServer.Packet()
	var msg []byte
	if msg, err = dp.Pack(npack.NewMsgPackage(msgID, data)); err != nil {
		nlog.Error(c.ctx, "Pack Error Msg", zap.Uint32("MsgID", msgID))
		err = errors.Wrap(err, "Pack Error Msg")
		return
	}
	// 写回客户端
	_, err = c.Conn.Write(msg)
	return
}

// SendBuffMsg 直接将数据发送给远程的客户端(有缓冲)
func (c *TCPConnection) SendBuffMsg(msgID uint32, data []byte) (err error) {
	c.RLock()
	defer c.RUnlock()
	if c.msgBuffChan == nil {
		c.msgBuffChan = make(chan []byte, nconf.GetUint32("server.maxMsgChanLen"))
		// 开启用于写回客户端数据流程的Goroutine
		// 此方法只读取MsgBuffChan中的数据没调用SendBuffMsg可以分配内存和启用协程
		go c.StartWriter()
	}
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()

	if c.isClosed == true {
		err = errors.New("Connection Closed When Send Buff Msg")
		return
	}
	// 将data封包，并且发送
	dp := c.TCPServer.Packet()
	var msg []byte
	if msg, err = dp.Pack(npack.NewMsgPackage(msgID, data)); err != nil {
		nlog.Error(c.ctx, "Pack Error Msg", zap.Uint32("MsgID", msgID))
		err = errors.Wrap(err, "Pack Error Msg")
		return
	}
	// 发送超时
	select {
	case <-idleTimeout.C:
		err = errors.New("Send Buff Msg Timeout")
		return
	case c.msgBuffChan <- msg:
		return
	}
	return
}

// SetProperty 设置链接属性
func (c *TCPConnection) SetProperty(key string, value any) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if c.property == nil {
		c.property = make(map[string]any)
	}

	c.property[key] = value
}

// GetProperty 获取链接属性
func (c *TCPConnection) GetProperty(key string) (any, error) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}

	return nil, errors.New("No Property Found")
}

// RemoveProperty 移除链接属性
func (c *TCPConnection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

// NewTCPConnection 创建TCP连接
func NewTCPConnection(server niface.IServer, conn *net.TCPConn, connID uint32, msgHandler niface.IMsgHandle) *TCPConnection {
	// 初始化Conn属性
	c := &TCPConnection{
		TCPServer:   server,
		Conn:        conn,
		ConnID:      connID,
		MsgHandler:  msgHandler,
		msgBuffChan: nil,
		property:    nil,
		isClosed:    false,
	}
	// 将新创建的Conn添加到链接管理中
	c.TCPServer.GetConnMgr().Add(c)
	return c
}

// StartWriter 写消息Goroutine，用户将数据发送给客户端
func (c *TCPConnection) StartWriter() {
	nlog.Info(c.ctx, "Writer Goroutine Is Running")
	defer nlog.Info(c.ctx, "Conn Writer Exit !!!", zap.String("RemoteAddr", c.RemoteAddr().String()))

	for {
		select {
		case data, ok := <-c.msgBuffChan:
			if ok {
				// 有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					nlog.Error(c.ctx, "Send Buff Data Error, Conn Writer Exit !!!", zap.Error(err))
					return
				}
			} else {
				nlog.Error(c.ctx, "Msgbuffchan Is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// StartReader 读消息Goroutine，用于从客户端中读取数据
func (c *TCPConnection) StartReader() {
	nlog.Info(c.ctx, "Reader Goroutine Is Running")
	defer nlog.Info(c.ctx, "Conn Reader Exit !!!", zap.String("RemoteAddr", c.RemoteAddr().String()))
	defer c.Stop()
	// 创建拆包解包的对象
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 读取客户端的Msg head
			headData := make([]byte, c.TCPServer.Packet().GetHeadLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				nlog.Error(c.ctx, "Read Msg Head Error", zap.Error(err))
				return
			}
			// 拆包，得到msgID和datalen放在msg中
			msg, err := c.TCPServer.Packet().Unpack(headData)
			if err != nil {
				nlog.Error(c.ctx, "Unpack Error", zap.Error(err))
				return
			}
			// 根据dataLen读取data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					nlog.Error(c.ctx, "Read Msg Data Error", zap.Error(err))
					return
				}
			}
			msg.SetData(data)
			// 得到当前客户端请求的Request数据
			req := Request{
				conn:  c,
				msg:   msg,
				index: 0,
			}
			if nconf.GetUint32("server.workerPoolSize") > 0 {
				// 已经启动工作池机制，将消息交给Worker处理
				c.MsgHandler.SendMsgToTaskQueue(&req)
			} else {
				// 从绑定好的消息和对应的处理方法中执行对应的Handle方法
				go c.MsgHandler.DoMsgHandler(&req)
			}
		}
	}
}

func (c *TCPConnection) finalizer() {
	// 如果用户注册了该连接的关闭回调业务，那么在此刻应该显示调用
	c.TCPServer.CallOnConnStop(c)

	c.Lock()
	defer c.Unlock()

	// 如果当前连接已经关闭
	if c.isClosed == true {
		return
	}

	nlog.Info(c.ctx, "Conn Stop()...", zap.Uint32("ConnID", c.ConnID))
	// 关闭socket连接
	_ = c.Conn.Close()

	// 将连接从连接管理器中删除
	c.TCPServer.GetConnMgr().Remove(c)

	// 关闭该连接全部管道
	if c.msgBuffChan != nil {
		close(c.msgBuffChan)
	}
	// 设置标志位
	c.isClosed = true
}
