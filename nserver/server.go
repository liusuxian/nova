/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-24 14:23:32
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nserver/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nserver

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/nconn"
	"github.com/liusuxian/nova/nheartbeat"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nmsghandler"
	"github.com/liusuxian/nova/noverload"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrequest"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
	"os"
	"time"
)

// Server 服务器结构
type Server struct {
	gnet.BuiltinEventEngine
	eng              gnet.Engine
	options          gnet.Options                  // 服务器 gnet 启动选项
	serverConf       *ServerConfig                 // 服务器配置
	addr             string                        // 服务器绑定的地址
	ctx              context.Context               // 当前 Server 的 Context
	msgHandler       niface.IMsgHandle             // 当前 Server 绑定的消息处理模块
	connMgr          niface.IConnManager           // 当前 Server 的连接管理模块
	onConnStart      func(conn niface.IConnection) // 当前 Server 的连接创建时的 Hook 函数
	onConnStop       func(conn niface.IConnection) // 当前 Server 的连接断开时的 Hook 函数
	packet           niface.IDataPack              // 当前 Server 绑定的数据协议封包方式
	overLoadMsg      *noverload.OverLoadMsg        // 服务器人数超载消息
	heartbeat        time.Duration                 // 心跳发送间隔时间
	heartbeatChecker *nheartbeat.HeartbeatChecker  // 心跳检测器
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name           string // 服务器应用名称，默认"Nova"
	Network        string // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	Port           int    // 服务器监听端口
	Heartbeat      int    // 心跳发送间隔时间（单位:毫秒，一定要小于 maxHeartbeat 配置），默认 3000
	MaxHeartbeat   int    // 最长心跳检测间隔时间（单位:毫秒，一定要大于 heartbeat 配置），默认 5000
	MaxConn        int    // 允许的客户端连接最大数量，默认 3
	WorkerPoolSize int    // 工作任务池最大工作 Goroutine 数量，默认 10
	MaxPacketSize  int    // 数据包的最大值（单位:字节），默认 4096
	PacketMethod   int    // 封包和拆包方式，1: 消息ID(2字节)-消息体长度(4字节)-消息内容，默认 1
	Endian         int    // 字节存储次序，1: 小端 2: 大端，默认 1
}

// NewServer 创建 Server
func NewServer(opts ...Option) niface.IServer {
	// 获取服务器配置
	serCfg := &ServerConfig{}
	ctx := context.Background()
	if err := nconf.StructKey("server", serCfg); err != nil {
		nlog.Fatal(ctx, "New Server Error", zap.Error(err))
	}
	// 初始化 Server 属性
	heartbeat := time.Duration(serCfg.Heartbeat) * time.Millisecond
	s := &Server{
		serverConf: serCfg,
		addr:       fmt.Sprintf("%s://:%d", serCfg.Network, serCfg.Port),
		ctx:        ctx,
		msgHandler: nmsghandler.NewMsgHandle(serCfg.WorkerPoolSize),
		connMgr:    nconn.NewConnManager(ctx),
		packet:     npack.NewPack(serCfg.PacketMethod, serCfg.Endian, serCfg.MaxPacketSize),
		heartbeat:  heartbeat,
	}
	// 处理服务选项
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start 启动 Server
func (s *Server) Start() {
	nlog.Info(s.ctx, "Start Server......", zap.String("ServerName", s.serverConf.Name), zap.Int("Pid", os.Getpid()))
	if err := gnet.Run(s, s.addr, gnet.WithOptions(s.options)); err != nil {
		nlog.Fatal(s.ctx, "Start Server Error", zap.Error(err))
	}
}

// Stop 停止 Server
func (s *Server) Stop() {
	nlog.Info(s.ctx, "Stop Server......", zap.String("ServerName", s.serverConf.Name), zap.Int("Pid", os.Getpid()))
	s.eng.Stop(s.ctx)
}

// AddRouter 给当前 Server 添加路由
func (s *Server) AddRouter(msgID uint16, router niface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
}

// GetCtx 获取当前 Server 的 Context
func (s *Server) GetCtx() context.Context {
	return s.ctx
}

// GetConnManager 获取当前 Server 的连接管理
func (s *Server) GetConnManager() niface.IConnManager {
	return s.connMgr
}

// GetConnections 获取当前 Server 的活跃连接数
func (s *Server) GetConnections() int {
	return s.eng.CountConnections()
}

// SetOnConnStart 设置当前 Server 的连接创建时的 Hook 函数
func (s *Server) SetOnConnStart(hookFunc func(niface.IConnection)) {
	s.onConnStart = hookFunc
}

// SetOnConnStop 设置当前 Server 的连接断开时的 Hook 函数
func (s *Server) SetOnConnStop(hookFunc func(niface.IConnection)) {
	s.onConnStop = hookFunc
}

// GetOnConnStart 获取当前 Server 的连接创建时的 Hook 函数
func (s *Server) GetOnConnStart() func(niface.IConnection) {
	return s.onConnStart
}

// GetOnConnStop 获取当前 Server 的连接断开时的 Hook 函数
func (s *Server) GetOnConnStop() func(niface.IConnection) {
	return s.onConnStop
}

// SetPacket 设置当前 Server 绑定的数据协议封包和拆包方式
func (s *Server) SetPacket(packet niface.IDataPack) {
	s.packet = packet
}

// GetPacket 获取当前 Server 绑定的数据协议封包和拆包方式
func (s *Server) GetPacket() niface.IDataPack {
	return s.packet
}

// GetMsgHandler 获取当前 Server 绑定的消息处理模块
func (s *Server) GetMsgHandler() niface.IMsgHandle {
	return s.msgHandler
}

// SetOverLoadMsg 设置当前 Server 的服务器人数超载消息
func (s *Server) SetOverLoadMsg(option *niface.OverLoadMsgOption) {
	overLoadMsg := noverload.NewOverLoadMsgServer()
	// 用户自定义
	if option != nil {
		overLoadMsg.SetOverLoadMsgFunc(option.MakeMsg)
		overLoadMsg.BindRouter(option.MsgID, option.Router)
	}
	s.overLoadMsg = overLoadMsg
}

// SetHeartBeat 设置当前 Server 的心跳检测
func (s *Server) SetHeartBeat(option *niface.HeartBeatOption, initiate bool) {
	checker := nheartbeat.NewHeartbeatCheckerServer(s, initiate)
	// 用户自定义
	if option != nil {
		checker.SetHeartBeatMsgFunc(option.MakeMsg)
		checker.SetOnRemoteNotAlive(option.OnRemoteNotAlive)
		checker.BindRouter(option.MsgID, option.Router)
	}
	// 添加心跳检测的路由
	s.AddRouter(checker.GetMsgID(), checker.GetRouter())
	s.heartbeatChecker = checker
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (s *Server) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info(s.ctx, "Server OnBoot", zap.String("listening", s.addr), zap.String("ServerName", s.serverConf.Name), zap.Any("options", s.options))
	s.eng = eng
	// 启动 Worker 工作池
	s.msgHandler.StartWorkerPool()
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (s *Server) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info(s.ctx, "Server OnClose", zap.Int("connID", conn.Fd()), zap.String("RemoteAddr", conn.RemoteAddr().String()), zap.Int("Connections", s.GetConnections()))
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
	nlog.Info(s.ctx, "Server OnOpen", zap.Int("connID", conn.Fd()), zap.Int("Connections", s.GetConnections()))
	// 检测允许的客户端连接最大数量
	if s.GetConnections() > s.serverConf.MaxConn {
		out = s.packet.Pack(npack.NewMsgPackage(s.overLoadMsg.GetMsgID(), s.overLoadMsg.GetMsgData()))
		// 踢连接
		go s.doKickConn(conn)
		return
	}
	// 创建一个 Server 服务端特性的连接
	serverConn := nconn.NewServerConn(s, conn, time.Duration(s.serverConf.MaxHeartbeat)*time.Millisecond)
	// 启动连接
	go serverConn.Start()
	// 发送心跳
	out = s.packet.Pack(npack.NewMsgPackage(s.heartbeatChecker.GetMsgID(), s.heartbeatChecker.GetMsgData()))
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (s *Server) OnShutdown(eng gnet.Engine) {
	nlog.Info(s.ctx, "Server OnShutdown")
	// 停止 Worker 工作池
	s.msgHandler.StopWorkerPool()
	return
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	if s.heartbeatChecker != nil {
		go s.heartbeatChecker.Check()
	}
	delay = s.heartbeat
	return
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (s *Server) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	for {
		msg, err := s.packet.UnPack(conn)
		if err == npack.ErrIncompletePacket {
			break
		}
		if err != nil {
			nlog.Error(s.ctx, "Server OnTraffic Unpack Error", zap.Error(err))
			return gnet.Close
		}
		nlog.Debug(s.ctx, "Server OnTraffic", zap.Int("connID", conn.Fd()), zap.Uint16("MsgID", msg.GetMsgID()))
		iConn, err := s.connMgr.GetConn(conn.Fd())
		if err != nil {
			return gnet.Close
		}
		// 更新连接活动时间
		iConn.UpdateActivity()
		// 得到当前客户端请求的 Request 数据
		request := nrequest.NewRequest(iConn, msg)
		// 处理请求消息
		s.msgHandler.HandleRequest(request)
	}
	return
}

// doKickConn 踢连接
func (s *Server) doKickConn(conn gnet.Conn) {
	select {
	case <-time.After(10 * time.Millisecond):
		_ = conn.Close()
		return
	}
}
