/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 10:49:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 21:33:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nclient/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nclient

import (
	"github.com/liusuxian/nova/nconn"
	"github.com/liusuxian/nova/nheartbeat"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nmsghandler"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrequest"
	"github.com/liusuxian/nova/nserveroverload"
	"github.com/panjf2000/gnet/v2"
	"time"
)

// Client 客户端结构
type Client struct {
	gnet.BuiltinEventEngine
	client                *gnet.Client                  // gnet 客户端
	options               gnet.Options                  // 客户端 gnet 启动选项
	network               string                        // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	addr                  string                        // 服务器地址
	conn                  niface.IConnection            // Client 连接
	msgHandler            niface.IMsgHandle             // 当前 Client 绑定的消息处理模块
	onConnStart           func(conn niface.IConnection) // 当前 Client 的连接创建时的 Hook 函数
	onConnStop            func(conn niface.IConnection) // 当前 Client 的连接断开时的 Hook 函数
	packet                niface.IDataPack              // 当前 Client 绑定的数据协议封包方式
	serverOverloadChecker niface.IServerOverloadChecker // 服务器人数超载检测器
	heartbeat             time.Duration                 // 心跳发送间隔时间
	maxHeartbeat          time.Duration                 // 最长心跳检测间隔时间
	heartbeatChecker      niface.IHeartBeatChecker      // 心跳检测器
}

// NewClient 创建 Client
func NewClient(network, addr string, opts ...Option) (client niface.IClient) {
	// 初始化 Client 属性
	c := &Client{
		network:    network,
		addr:       addr,
		msgHandler: nmsghandler.NewMsgHandle(0),
		packet:     npack.NewPack(npack.DefaultPacketMethod, npack.LittleEndian, npack.DefaultMaxPacketSize),
	}
	// 处理服务选项
	for _, opt := range opts {
		opt(c)
	}
	// 创建 Client
	cli, err := gnet.NewClient(c, gnet.WithOptions(c.options))
	if err != nil {
		nlog.Fatal("New Client Error", nlog.Err(err))
	}
	c.client = cli
	return c
}

// Start 启动 Client
func (c *Client) Start() {
	// 启动 Client
	if err := c.client.Start(); err != nil {
		nlog.Fatal("Start Client Error", nlog.Err(err))
	}
}

// Stop 停止 Client
func (c *Client) Stop() {
	_ = c.client.Stop()
}

// AddRouter 给当前 Client 添加路由
func (c *Client) AddRouter(msgID uint16, router niface.IRouter) {
	c.msgHandler.AddRouter(msgID, router)
}

// Conn 当前 Client 的连接信息
func (c *Client) Conn() (conn niface.IConnection) {
	return c.conn
}

// SetOnConnStart 设置当前 Client 的连接创建时的 Hook 函数
func (c *Client) SetOnConnStart(f func(conn niface.IConnection)) {
	c.onConnStart = f
}

// SetOnConnStop 设置当前 Client 的连接断开时的 Hook 函数
func (c *Client) SetOnConnStop(f func(conn niface.IConnection)) {
	c.onConnStop = f
}

// GetOnConnStart 获取当前 Client 的连接创建时的 Hook 函数
func (c *Client) GetOnConnStart() (f func(conn niface.IConnection)) {
	return c.onConnStart
}

// GetOnConnStop 获取当前 Client 的连接断开时的 Hook 函数
func (c *Client) GetOnConnStop() (f func(conn niface.IConnection)) {
	return c.onConnStop
}

// SetPacket 设置当前 Client 绑定的数据协议封包和拆包方式
func (c *Client) SetPacket(packet niface.IDataPack) {
	c.packet = packet
}

// GetPacket 获取当前 Client 绑定的数据协议封包和拆包方式
func (c *Client) GetPacket() (packet niface.IDataPack) {
	return c.packet
}

// GetMsgHandler 获取当前 Client 绑定的消息处理模块
func (c *Client) GetMsgHandler() (handler niface.IMsgHandle) {
	return c.msgHandler
}

// SetServerOverload 设置当前 Client 的服务器人数超载检测器
func (c *Client) SetServerOverload(option ...*niface.ServerOverloadOption) {
	checker := nserveroverload.NewServerOverloadChecker(true)
	// 用户自定义
	if len(option) > 0 {
		opt := option[0]
		checker.SetServerOverloadMsgFunc(opt.MakeMsg)
		checker.BindRouter(opt.MsgID, opt.Router)
	}
	// 添加服务器人数超载消息的路由
	c.AddRouter(checker.GetMsgID(), checker.GetRouter())
	c.serverOverloadChecker = checker
}

// SetHeartBeat 设置当前 Client 的心跳检测器
func (c *Client) SetHeartBeat(initiate bool, option ...*niface.HeartBeatOption) {
	checker := nheartbeat.NewHeartbeatChecker(c.heartbeat, initiate)
	// 用户自定义
	if len(option) > 0 {
		opt := option[0]
		checker.SetHeartBeatMsgFunc(opt.MakeMsg)
		checker.SetOnRemoteNotAlive(opt.OnRemoteNotAlive)
		checker.BindRouter(opt.MsgID, opt.Router)
	}
	// 添加心跳检测的路由
	c.AddRouter(checker.GetMsgID(), checker.GetRouter())
	c.heartbeatChecker = checker
}

// GetHeartBeat 获取当前 Client 的心跳检测器
func (c *Client) GetHeartBeat() (checker niface.IHeartBeatChecker) {
	return c.heartbeatChecker
}

// AddInterceptor 添加拦截器
func (c *Client) AddInterceptor(interceptor niface.IInterceptor) {
	c.msgHandler.AddInterceptor(interceptor)
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (c *Client) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info("Client OnBoot", nlog.String("Network", c.network), nlog.String("Addr", c.addr), nlog.Reflect("options", c.options))
	// 连接服务器
	if _, err := c.client.Dial(c.network, c.addr); err != nil {
		nlog.Fatal("Client OnBoot Error", nlog.Err(err))
	}
	// 打印所有路由
	c.msgHandler.PrintRouters()
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (c *Client) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info("Client OnClose", nlog.Int("connID", conn.Fd()), nlog.String("RemoteAddr", conn.RemoteAddr().String()))
	// 停止连接
	c.conn.Stop()
	return
}

// OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (c *Client) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info("Client OnOpen", nlog.Int("connID", conn.Fd()))
	// 创建一个 Client 客户端特性的连接
	c.conn = nconn.NewClientConn(c, conn, c.maxHeartbeat)
	// 启动连接
	go c.conn.Start()
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (c *Client) OnShutdown(eng gnet.Engine) {
	nlog.Info("Client OnShutdown")
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (c *Client) OnTick() (delay time.Duration, action gnet.Action) {
	delay = time.Duration(1000) * time.Millisecond
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
			nlog.Error("Client OnTraffic Unpack Error", nlog.Err(err))
			return
		}
		nlog.Debug("Client OnTraffic", nlog.Int("connID", conn.Fd()), nlog.Uint16("MsgID", msg.GetMsgID()))
		// 更新连接活动时间
		c.conn.UpdateActivity()
		// 得到当前客户端请求的 Request 数据
		request := nrequest.NewRequest(c.conn, msg)
		// 处理请求消息
		c.msgHandler.Execute(request)
	}
	return
}
