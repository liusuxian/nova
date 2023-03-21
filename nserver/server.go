/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-21 16:43:23
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
	"github.com/liusuxian/nova/npack"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server 服务器结构
type Server struct {
	gnet.BuiltinEventEngine
	eng               gnet.Engine
	options           gnet.Options                  // 服务器 gnet 启动选项
	action            gnet.Action                   // gnet 事件完成后发生的动作，None 无任何操作，Close 关闭连接，Shutdown 停止整个gnet引擎
	serverConf        *ServerConfig                 // 服务器配置
	addr              string                        // 服务器绑定的地址
	ctx               context.Context               // 当前 Server 的根 Context
	msgHandler        niface.IMsgHandle             // 当前 Server 绑定的消息处理模块
	connMgr           niface.IConnManager           // 当前 Server 的连接管理模块
	onConnStart       func(conn niface.IConnection) // 当前 Server 的连接创建时的 Hook 函数
	onConnStop        func(conn niface.IConnection) // 当前 Server 的连接断开时的 Hook 函数
	packet            niface.IDataPack              // 当前 Server 绑定的数据协议封包方式
	heartbeatInterval time.Duration                 // 心跳检测间隔时间
	heartbeatChecker  *nheartbeat.HeartbeatChecker  // 心跳检测器
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name           string // 服务器应用名称，默认"Nova"
	Network        string // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	Port           int    // 服务器监听端口
	MaxHeartbeat   int    // 最长心跳检测间隔时间（单位:毫秒），默认 5000
	MaxConn        int    // 允许的客户端连接最大数量，默认 3
	WorkerPoolSize int    // 工作任务池最大工作 Goroutine 数量，默认 10
	MaxPacketSize  int    // 数据包的最大值（单位:字节），默认 4096
	PacketMethod   int    // 封包和拆包方式，1: 消息ID(2字节)-消息体长度(4字节)-消息内容，默认 1
	Endian         int    // 字节存储次序，1: 小端 2: 大端，默认 1
	MaxMsgChanLen  int    // 发送消息的缓冲最大长度（单位:字节），默认 1024
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
	heartbeatInterval := time.Duration(serCfg.MaxHeartbeat) * time.Millisecond
	s := &Server{
		action:            gnet.None,
		serverConf:        serCfg,
		addr:              fmt.Sprintf("%s://:%d", serCfg.Network, serCfg.Port),
		ctx:               ctx,
		msgHandler:        nmsghandler.NewMsgHandle(ctx, serCfg.WorkerPoolSize),
		connMgr:           nconn.NewConnManager(ctx),
		packet:            npack.NewPack(serCfg.PacketMethod, serCfg.Endian, serCfg.MaxPacketSize),
		heartbeatInterval: heartbeatInterval,
	}
	s.heartbeatChecker = nheartbeat.NewHeartbeatCheckerServer(ctx, s)
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// TODO Start 启动 Server
func (s *Server) Start() {
	nlog.Info(s.ctx, "Start Server......", zap.String("ServerName", s.serverConf.Name), zap.Int("Pid", os.Getpid()))

	go func(s *Server) {
		// 创建一个通道，用于接收信号
		c := make(chan os.Signal, 1)
		// 注册信号接收器
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		// 等待信号
		sig := <-c
		nlog.Info(s.ctx, "Server Interrupt Signal", zap.String("ServerName", s.serverConf.Name), zap.String("Signal", sig.String()))
		// 停止服务器
		s.Stop()
	}(s)

	if err := gnet.Run(s, s.addr, gnet.WithOptions(s.options)); err != nil {
		nlog.Fatal(s.ctx, "Start Server Error", zap.Error(err))
	}
}

// TODO Stop 停止 Server
func (s *Server) Stop() {
	nlog.Info(s.ctx, "Stop Server......", zap.String("ServerName", s.serverConf.Name), zap.Int("Pid", os.Getpid()))
	s.action = gnet.Shutdown
}

// AddRouter 给当前 Server 添加路由
func (s *Server) AddRouter(msgID uint16, router niface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
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

// SetHeartBeat 设置当前 Server 的心跳检测
func (s *Server) SetHeartBeat(option *niface.HeartBeatOption) {
	if option != nil {
		s.heartbeatChecker.SetHeartBeatMsgFunc(option.MakeMsg)
		s.heartbeatChecker.SetOnRemoteNotAlive(option.OnRemoteNotAlive)
		s.heartbeatChecker.BindRouter(option.MsgID, option.Router)
	}
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (s *Server) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info(s.ctx, "Server OnBoot", zap.String("listening", s.addr), zap.String("ServerName", s.serverConf.Name), zap.Any("options", s.options))
	s.eng = eng
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (s *Server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	nlog.Info(s.ctx, "Server OnClose", zap.String("RemoteAddr", c.RemoteAddr().String()), zap.Int("Connections", s.GetConnections()))
	// 删除连接
	s.connMgr.RemoveConn(c.Fd())
	return
}

// OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (s *Server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info(s.ctx, "Server OnOpen", zap.Int("connID", c.Fd()), zap.Int("Connections", s.GetConnections()))
	// 创建一个 Server 服务端特性的连接
	nconn.NewServerConn(s.ctx, s, c, s.heartbeatInterval)
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (s *Server) OnShutdown(eng gnet.Engine) {
	nlog.Info(s.ctx, "Server OnShutdown")
	return
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	nlog.Debug(s.ctx, "Server OnTick")
	// go s.heartbeatChecker.Check()
	return s.heartbeatInterval, s.action
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {
	for {
		msg, err := s.packet.UnPack(c)
		if err == npack.ErrIncompletePacket {
			break
		}
		if err != nil {
			nlog.Error(s.ctx, "Server OnTraffic Unpack Error", zap.Error(err))
			return gnet.Close
		}
		nlog.Debug(s.ctx, "Server OnTraffic", zap.Uint16("MsgID", msg.GetMsgID()), zap.Int("DataLen", msg.GetDataLen()), zap.ByteString("Data", msg.GetData()))
		conn, err := s.connMgr.GetConn(c.Fd())
		if err != nil {
			nlog.Error(s.ctx, "Server OnTraffic GetConn Error", zap.Error(err))
			return gnet.Close
		}
		// 更新连接活动时间
		conn.UpdateActivity()
	}
	return
}
