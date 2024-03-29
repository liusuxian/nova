/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 01:45:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-06-22 11:18:09
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nserver

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/nconn"
	"github.com/liusuxian/nova/nerr"
	"github.com/liusuxian/nova/nheartbeat"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nmsghandler"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrequest"
	"github.com/liusuxian/nova/nrouter"
	"github.com/liusuxian/nova/nserveroverload"
	"github.com/panjf2000/gnet/v2"
	"time"
)

// Server 服务器结构
type Server struct {
	gnet.BuiltinEventEngine
	eng                   gnet.Engine
	serverConf            *ServerConfig                 // 服务器配置
	addr                  string                        // 服务器绑定的地址
	msgHandler            niface.IMsgHandle             // 当前 Server 绑定的消息处理模块
	connMgr               niface.IConnManager           // 当前 Server 的连接管理模块
	onConnStart           func(conn niface.IConnection) // 当前 Server 的连接创建时的 Hook 函数
	onConnStop            func(conn niface.IConnection) // 当前 Server 的连接断开时的 Hook 函数
	onStart               func(s niface.IServer)        // 当前 Server 启动时的 Hook 函数
	onStop                func(s niface.IServer)        // 当前 Server 停止时的 Hook 函数
	packet                niface.IDataPack              // 当前 Server 绑定的数据协议封包方式
	serverOverloadChecker niface.IServerOverloadChecker // 服务器人数超载检测器
	heartbeatChecker      niface.IHeartBeatChecker      // 心跳检测器
}

// ServerConfig 服务器配置
type ServerConfig struct {
	gnet.Options                         // 服务器 gnet 启动选项
	Name                   string        // 服务器应用名称，默认"Nova"
	Network                string        // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	Port                   int           // 服务器监听端口
	HeartBeat              time.Duration // 心跳发送间隔时间（一定要小于 maxHeartBeat 配置），默认 10秒
	MaxHeartBeat           time.Duration // 最长心跳检测间隔时间（一定要大于 heartBeat 配置），默认 15秒
	MaxConn                int           // 允许的客户端连接最大数量，默认 3
	WorkerPoolSize         int           // 工作任务池最大工作 Goroutine 数量，默认 10
	WorkerPoolSizeOverflow int           // 当处理任务超过工作任务池的容量时，增加的 Goroutine 数量，默认 5
	MaxPacketSize          int           // 数据包的最大值（单位:字节），默认 4096
	PacketMethod           int           // 封包和拆包方式，1: 消息ID(2字节)-消息体长度(4字节)-消息内容，默认 1
	Endian                 int           // 字节存储次序，1: 小端 2: 大端，默认 1
	SlowThreshold          time.Duration // 处理请求或执行操作时的慢速阈值
}

// ServerConfigOption 服务器配置选项
type ServerConfigOption func(sc *ServerConfig)

// 表示负载均衡算法的类型
const (
	RoundRobin       = gnet.RoundRobin       // 算法通过轮询事件循环列表将下一个被接受的连接分配到事件循环中
	LeastConnections = gnet.LeastConnections // 算法会将下一个被接受的连接分配到当前活跃连接数最少的事件循环中
	SourceAddrHash   = gnet.SourceAddrHash   // 算法通过哈希远程地址将下一个被接受的连接分配给事件循环
)

// TCP套接字选项的类型。用于设置 TCP 连接的 NoDelay 选项，该选项表示是否禁用 Nagle 算法。启用 NoDelay 选项可以降低延迟，但会增加网络负载
const (
	TCPNoDelay = gnet.TCPNoDelay // 不启用
	TCPDelay   = gnet.TCPDelay   // 启用
)

// NewServer 创建 Server
func NewServer(opts ...ServerConfigOption) (server niface.IServer) {
	// 获取服务器配置
	serCfg := &ServerConfig{}
	if err := nconf.StructKey("server", serCfg); err != nil {
		nlog.Fatal("New Server Error", nlog.Err(err))
	}
	// 初始化 Server 属性
	s := &Server{
		serverConf: serCfg,
		addr:       fmt.Sprintf("%s://:%d", serCfg.Network, serCfg.Port),
		msgHandler: nmsghandler.NewMsgHandle(serCfg.WorkerPoolSize, serCfg.WorkerPoolSizeOverflow, serCfg.SlowThreshold),
		connMgr:    nconn.NewConnManager(),
		packet:     npack.NewPack(serCfg.PacketMethod, serCfg.Endian, serCfg.MaxPacketSize),
	}
	// 处理服务器配置选项
	for _, opt := range opts {
		opt(s.serverConf)
	}
	// 添加全局组件
	s.Use(nrouter.RouterRecover) // 路由 Recover 处理器
	return s
}

// Start 启动 Server
func (s *Server) Start() {
	if err := gnet.Run(s, s.addr, gnet.WithOptions(s.serverConf.Options)); err != nil {
		nlog.Fatal("Start Server Error", nlog.Err(err))
	}
}

// Stop 停止 Server
func (s *Server) Stop() {
	s.eng.Stop(context.Background())
}

// AddRouter 添加业务处理器集合
func (s *Server) AddRouter(msgID uint16, handlers ...niface.RouterHandler) (router niface.IRouter) {
	return s.msgHandler.AddRouter(msgID, handlers...)
}

// Group 路由分组管理，并且会返回一个组管理器
func (s *Server) Group(startMsgID, endMsgID uint16, handlers ...niface.RouterHandler) (group niface.IGroupRouter) {
	return s.msgHandler.Group(startMsgID, endMsgID, handlers...)
}

// Use 添加全局组件
func (s *Server) Use(handlers ...niface.RouterHandler) (router niface.IRouter) {
	return s.msgHandler.Use(handlers...)
}

// GetConnManager 获取当前 Server 的连接管理
func (s *Server) GetConnManager() (connMgr niface.IConnManager) {
	return s.connMgr
}

// GetConnections 获取当前 Server 的活跃连接数
func (s *Server) GetConnections() (nums int) {
	return s.eng.CountConnections()
}

// SetOnConnStart 设置当前 Server 的连接创建时的 Hook 函数
func (s *Server) SetOnConnStart(f func(conn niface.IConnection)) {
	s.onConnStart = f
}

// SetOnConnStop 设置当前 Server 的连接断开时的 Hook 函数
func (s *Server) SetOnConnStop(f func(conn niface.IConnection)) {
	s.onConnStop = f
}

// GetOnConnStart 获取当前 Server 的连接创建时的 Hook 函数
func (s *Server) GetOnConnStart() (f func(conn niface.IConnection)) {
	return s.onConnStart
}

// GetOnConnStop 获取当前 Server 的连接断开时的 Hook 函数
func (s *Server) GetOnConnStop() (f func(conn niface.IConnection)) {
	return s.onConnStop
}

// SetOnStart 设置当前 Server 启动时的 Hook 函数
func (s *Server) SetOnStart(f func(s niface.IServer)) {
	s.onStart = f
}

// SetOnStop 设置当前 Server 停止时的 Hook 函数
func (s *Server) SetOnStop(f func(s niface.IServer)) {
	s.onStop = f
}

// SetPacket 设置当前 Server 绑定的数据协议封包和拆包方式
func (s *Server) SetPacket(packet niface.IDataPack) {
	s.packet = packet
}

// GetPacket 获取当前 Server 绑定的数据协议封包和拆包方式
func (s *Server) GetPacket() (packet niface.IDataPack) {
	return s.packet
}

// GetMsgHandler 获取当前 Server 绑定的消息处理模块
func (s *Server) GetMsgHandler() (handler niface.IMsgHandle) {
	return s.msgHandler
}

// SetServerOverload 设置当前 Server 的服务器人数超载检测器
func (s *Server) SetServerOverload(option ...*niface.ServerOverloadOption) {
	checker := nserveroverload.NewServerOverloadChecker()
	// 用户自定义
	if len(option) > 0 {
		opt := option[0]
		checker.SetServerOverloadMsgFunc(opt.MakeMsg)
		checker.SetMsgID(opt.MsgID)
	}
	s.serverOverloadChecker = checker
}

// GetServerOverload 获取当前 Server 的服务器人数超载检测器
func (s *Server) GetServerOverload() (checker niface.IServerOverloadChecker) {
	return s.serverOverloadChecker
}

// SetHeartBeat 设置当前 Server 的心跳检测器
func (s *Server) SetHeartBeat(initiate bool, option ...*niface.HeartBeatOption) {
	// 创建心跳检测器
	checker := nheartbeat.NewHeartBeatChecker(s.serverConf.HeartBeat, initiate)
	// 用户自定义
	if len(option) > 0 {
		opt := option[0]
		checker.SetHeartBeatMsgFunc(opt.MakeMsg)
		checker.SetOnRemoteNotAlive(opt.OnRemoteNotAlive)
		checker.SetMsgID(opt.MsgID)
	}
	s.heartbeatChecker = checker
}

// GetHeartBeat 获取当前 Server 的心跳检测器
func (s *Server) GetHeartBeat() (checker niface.IHeartBeatChecker) {
	return s.heartbeatChecker
}

// AddInterceptor 添加拦截器
func (s *Server) AddInterceptor(interceptor niface.IInterceptor) {
	s.msgHandler.AddInterceptor(interceptor)
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (s *Server) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info("Server OnBoot", nlog.String("listening", s.addr), nlog.Reflect("ServerConf", s.serverConf))
	s.eng = eng
	// 调用 Server 启动时的 Hook 函数
	s.callOnStart()
	// 启动 Worker 工作池
	s.msgHandler.StartWorkerPool()
	// 打印所有路由
	s.msgHandler.PrintRouters()
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (s *Server) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info("Server OnClose", nlog.Int("connID", conn.Fd()), nlog.String("RemoteAddr", conn.RemoteAddr().String()), nlog.Int("Connections", s.GetConnections()))
	// 通过 ConnID 获取连接
	iConn, _ := s.connMgr.GetConn(conn.Fd())
	// 停止连接
	if iConn != nil {
		iConn.Stop()
	}
	return
}

// OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (s *Server) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info("Server OnOpen", nlog.Int("connID", conn.Fd()), nlog.Int("Connections", s.GetConnections()))
	// 服务器人数超载检测
	if s.serverOverloadChecker != nil {
		if s.serverOverloadChecker.Check(s, s.serverConf.MaxConn) {
			// 立即发送服务器人数超载消息
			serverOverloadMsg, err := s.serverOverloadChecker.GetMessage()
			if err != nil {
				nlog.Error("Get ServerOverloadMsg Error", nlog.Err(err))
				return nil, gnet.Close
			}
			out = s.packet.Pack(serverOverloadMsg)
			return
		}
	}
	// 创建一个 Server 服务端特性的连接
	serverConn := nconn.NewServerConn(s, conn, s.serverConf.MaxHeartBeat)
	// 启动连接
	serverConn.Start()
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (s *Server) OnShutdown(eng gnet.Engine) {
	nlog.Info("Server OnShutdown")
	// 调用 Server 停止时的 Hook 函数
	s.callOnStop()
	// 停止 Worker 工作池
	s.msgHandler.StopWorkerPool()
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	delay = time.Duration(1000) * time.Millisecond
	return
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (s *Server) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	for {
		// 获取包头长度(字节数)
		headLen := s.packet.GetHeadLen()
		// 读消息头
		headBuf, _ := conn.Peek(headLen)
		// 拆包头
		msg, err := s.packet.UnPackHead(headBuf)
		if err == nerr.ErrIncompletePacket {
			break
		}
		if err != nil {
			nlog.Error("Server OnTraffic UnPackHead Error", nlog.Uint16("MsgID", msg.GetMsgID()), nlog.Int("DataLen", msg.GetDataLen()), nlog.Err(err))
			return gnet.Close
		}
		// 读取整个消息数据
		msgLen := headLen + msg.GetDataLen()
		if conn.InboundBuffered() < msgLen {
			break
		}
		msgBuf, _ := conn.Peek(msgLen)
		_, _ = conn.Discard(msgLen)
		// 通过 ConnID 获取连接
		iConn, isExist := s.connMgr.GetConn(conn.Fd())
		if !isExist {
			return
		}
		// 更新连接活动时间
		iConn.UpdateActivity()
		// 拷贝整个消息数据
		copyMsgBuf := make([]byte, len(msgBuf))
		copy(copyMsgBuf, msgBuf)
		// 将请求提交给 Worker 工作池处理
		s.msgHandler.SubmitToWorkerPool(nrequest.NewRequestFunc(iConn, func() {
			// 拆包体
			s.packet.UnPackBody(copyMsgBuf, msg)
			nlog.Debug("Server OnTraffic", nlog.Int("connID", conn.Fd()), nlog.Uint16("MsgID", msg.GetMsgID()), nlog.Int("DataLen", msg.GetDataLen()))
			// 得到当前客户端请求的 Request 数据
			request := nrequest.NewRequest(iConn, msg)
			// 处理请求消息
			s.msgHandler.Execute(request)
		}))
	}
	return
}

// callOnStart 调用 Server 启动时的 Hook 函数
func (s *Server) callOnStart() {
	if s.onStart != nil {
		nlog.Info("Server CallOnStart...")
		s.onStart(s)
	}
}

// callOnStop 调用 Server 停止时的 Hook 函数
func (s *Server) callOnStop() {
	if s.onStop != nil {
		nlog.Info("Server CallOnStop...")
		s.onStop(s)
	}
}
