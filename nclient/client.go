/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 19:43:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 17:01:43
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nclient/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nclient

import (
	"context"
	"github.com/liusuxian/nova/nconn"
	"github.com/liusuxian/nova/nheartbeat"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nmsghandler"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrequest"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
	"time"
)

// Client 客户端结构
type Client struct {
	gnet.BuiltinEventEngine
	eng               gnet.Engine
	client            *gnet.Client                  // gnet 客户端
	options           gnet.Options                  // 客户端 gnet 启动选项
	network           string                        // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	addr              string                        // 服务器地址
	ctx               context.Context               // 当前 Client 的 Context
	conn              niface.IConnection            // Client 连接
	msgHandler        niface.IMsgHandle             // 当前 Client 绑定的消息处理模块
	onConnStart       func(conn niface.IConnection) // 当前 Client 的连接创建时的 Hook 函数
	onConnStop        func(conn niface.IConnection) // 当前 Client 的连接断开时的 Hook 函数
	packet            niface.IDataPack              // 当前 Client 绑定的数据协议封包方式
	heartbeatInterval time.Duration                 // 心跳检测间隔时间
	heartbeatChecker  *nheartbeat.HeartbeatChecker  // 心跳检测器
}

// NewClient 创建 Client
func NewClient(network, addr string, opts ...Option) niface.IClient {
	// 初始化 Client 属性
	ctx := context.Background()
	c := &Client{
		network:           network,
		addr:              addr,
		ctx:               ctx,
		msgHandler:        nmsghandler.NewMsgHandle(0),
		packet:            npack.NewPack(npack.DefaultPacketMethod, npack.LittleEndian, 4096),
		heartbeatInterval: time.Duration(5000) * time.Millisecond,
	}
	// 处理服务选项
	for _, opt := range opts {
		opt(c)
	}
	// 创建 Client
	client, err := gnet.NewClient(c, gnet.WithOptions(c.options))
	if err != nil {
		nlog.Fatal(ctx, "New Client Error", zap.Error(err))
	}
	c.client = client
	return c
}

// Start 启动 Client
func (c *Client) Start() {
	nlog.Info(c.ctx, "Start Client......")
	// 启动 Client
	if err := c.client.Start(); err != nil {
		nlog.Fatal(c.ctx, "Start Client Error", zap.Error(err))
	}
}

// Stop 停止 Client
func (c *Client) Stop() {
	nlog.Info(c.ctx, "Stop Client......")
	_ = c.client.Stop()
}

// AddRouter 给当前 Client 添加路由
func (c *Client) AddRouter(msgID uint16, router niface.IRouter) {
	c.msgHandler.AddRouter(msgID, router)
}

// Conn 当前 Client 的连接信息
func (c *Client) Conn() niface.IConnection {
	return c.conn
}

// SetOnConnStart 设置当前 Client 的连接创建时的 Hook 函数
func (c *Client) SetOnConnStart(hookFunc func(niface.IConnection)) {
	c.onConnStart = hookFunc
}

// SetOnConnStop 设置当前 Client 的连接断开时的 Hook 函数
func (c *Client) SetOnConnStop(hookFunc func(niface.IConnection)) {
	c.onConnStop = hookFunc
}

// GetOnConnStart 获取当前 Client 的连接创建时的 Hook 函数
func (c *Client) GetOnConnStart() func(niface.IConnection) {
	return c.onConnStart
}

// GetOnConnStop 获取当前 Client 的连接断开时的 Hook 函数
func (c *Client) GetOnConnStop() func(niface.IConnection) {
	return c.onConnStop
}

// SetPacket 设置当前 Client 绑定的数据协议封包和拆包方式
func (c *Client) SetPacket(packet niface.IDataPack) {
	c.packet = packet
}

// GetPacket 获取当前 Client 绑定的数据协议封包和拆包方式
func (c *Client) GetPacket() niface.IDataPack {
	return c.packet
}

// GetMsgHandler 获取当前 Client 绑定的消息处理模块
func (c *Client) GetMsgHandler() niface.IMsgHandle {
	return c.msgHandler
}

// SetHeartBeat 设置当前 Client 的心跳检测
func (c *Client) SetHeartBeat(option *niface.HeartBeatOption) {
	checker := nheartbeat.NewHeartbeatCheckerClient(c)
	// 用户自定义
	if option != nil {
		checker.SetHeartBeatMsgFunc(option.MakeMsg)
		checker.SetOnRemoteNotAlive(option.OnRemoteNotAlive)
		checker.BindRouter(option.MsgID, option.Router)
	}
	// 添加心跳检测的路由
	c.AddRouter(checker.GetMsgID(), checker.GetRouter())
	c.heartbeatChecker = checker
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (c *Client) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info(c.ctx, "Client OnBoot", zap.String("Network", c.network), zap.String("Addr", c.addr), zap.Any("options", c.options))
	c.eng = eng
	if _, err := c.client.Dial(c.network, c.addr); err != nil {
		nlog.Fatal(c.ctx, "Client OnBoot Error", zap.Error(err))
	}
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (c *Client) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info(c.ctx, "Client OnClose", zap.Int("connID", conn.Fd()), zap.String("LocalAddr", conn.LocalAddr().String()))
	// 停止连接
	c.conn.Stop()
	return
}

// OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (c *Client) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info(c.ctx, "Client OnOpen", zap.Int("connID", conn.Fd()))
	// 创建一个 Client 客户端特性的连接
	c.conn = nconn.NewClientConn(c, conn, c.heartbeatInterval)
	// 启动连接
	go c.conn.Start()
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (c *Client) OnShutdown(eng gnet.Engine) {
	nlog.Info(c.ctx, "Client OnShutdown")
	return
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (c *Client) OnTick() (delay time.Duration, action gnet.Action) {
	if c.heartbeatChecker != nil {
		go c.heartbeatChecker.Check()
	}
	delay = c.heartbeatInterval
	return
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (c *Client) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	for {
		msg, err := c.packet.UnPack(conn)
		if err == npack.ErrIncompletePacket {
			break
		}
		if err != nil {
			nlog.Error(c.ctx, "Client OnTraffic Unpack Error", zap.Error(err))
			return
		}
		nlog.Debug(c.ctx, "Client OnTraffic", zap.Int("connID", conn.Fd()), zap.Uint16("MsgID", msg.GetMsgID()), zap.Int("DataLen", msg.GetDataLen()), zap.ByteString("Data", msg.GetData()))
		// 更新连接活动时间
		c.conn.UpdateActivity()
		// 得到当前客户端请求的 Request 数据
		request := nrequest.NewRequest(c.conn, msg)
		// 处理请求消息
		c.msgHandler.HandleRequest(request)
	}
	return
}
