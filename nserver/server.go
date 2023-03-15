/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 20:35:24
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
	"time"
)

// Server 服务器结构
type Server struct {
	gnet.BuiltinEventEngine
	eng               gnet.Engine
	options           gnet.Options                  // 服务器 Net 启动选项
	serverConf        *ServerConfig                 // 服务器配置
	addr              string                        // 服务器绑定的地址
	ctx               context.Context               // 当前 Server 的 Context
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
	Port           uint16 // 服务器监听端口（uint16）
	MaxHeartbeat   uint32 // 最长心跳检测间隔时间（单位: 毫秒 uint32），默认 5000
	MaxConn        uint32 // 允许的客户端连接最大数量，默认 3（uint32）
	WorkerPoolSize uint32 // 工作任务池最大工作 Goroutine 数量，默认 10（uint32）
	MaxPacketSize  uint32 // 数据包的最大值，默认 4096（单位:字节 uint32）
	PacketMethod   uint8  // 封包和拆包方式，默认 1，1: 消息ID(2字节)-消息体长度(4字节)-消息内容（单位:字节 uint8）
	Endian         uint8  // 字节存储次序，默认小端，1: 小端 2: 大端（单位:字节 uint8）
	MaxMsgChanLen  uint32 // SendBuffMsg发送消息的缓冲最大长度，默认 1024（单位:字节 uint32）
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
		serverConf:        serCfg,
		addr:              fmt.Sprintf("%s://:%d", serCfg.Network, serCfg.Port),
		ctx:               ctx,
		msgHandler:        nmsghandler.NewMsgHandle(),
		connMgr:           nconn.NewConnManager(),
		packet:            npack.Factory().NewPack(serCfg.PacketMethod, serCfg.Endian, serCfg.MaxPacketSize),
		heartbeatInterval: heartbeatInterval,
	}
	s.heartbeatChecker = nheartbeat.NewHeartbeatCheckerServer(s)
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// TODO Start 启动服务器
func (s *Server) Start() {
	if err := gnet.Run(s, s.addr, gnet.WithOptions(s.options)); err != nil {
		nlog.Fatal(s.ctx, "Start Server Error", zap.Error(err))
	}
}

// TODO Stop 停止服务器
func (s *Server) Stop() {
	nlog.Info(s.ctx, "Stop Server......", zap.String("ServerName", s.serverConf.Name))
	// 清除并停止当前所有连接
	s.connMgr.ClearAllConn()
	s.eng.Stop(s.ctx)
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
func (s *Server) AddRouter(msgID uint16, router niface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
}

// GetConnManager 获取连接管理
func (s *Server) GetConnManager() niface.IConnManager {
	return s.connMgr
}

// 获取当前活跃的连接数
func (s *Server) GetConnections() uint32 {
	return uint32(s.eng.CountConnections())
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

// SetPacket 设置当前 Server 绑定的数据协议封包方式
func (s *Server) SetPacket(packet niface.IDataPack) {
	s.packet = packet
}

// GetPacket 获取当前 Server 绑定的数据协议封包方式
func (s *Server) GetPacket() niface.IDataPack {
	return s.packet
}

// 获取当前 Server 绑定的消息处理模块
func (s *Server) GetMsgHandler() niface.IMsgHandle {
	return s.msgHandler
}

// SetHeartBeat 设置心跳检测
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

// TODO OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (s *Server) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info(s.ctx, "Server OnClose", zap.String("RemoteAddr", conn.RemoteAddr().String()), zap.Uint32("Connections", s.GetConnections()))
	return
}

// TODO OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (s *Server) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info(s.ctx, "Server OnOpen", zap.Int("connID", conn.Fd()), zap.Uint32("Connections", s.GetConnections()))
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (s *Server) OnShutdown(eng gnet.Engine) {
	nlog.Info(s.ctx, "Server OnShutdown")
	return
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	nlog.Debug(s.ctx, "Server OnTick", zap.Uint32("Connections", s.GetConnections()))
	go s.heartbeatChecker.Check()
	delay = s.heartbeatInterval
	return
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (s *Server) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	for {
		msg, err := s.packet.Unpack(conn)
		if err == npack.ErrIncompletePacket {
			break
		}
		if err != nil {
			nlog.Error(s.ctx, "Server OnTraffic Error", zap.Error(err))
			action = gnet.Close
			return
		}
		nlog.Debug(s.ctx, "Server OnTraffic", zap.Uint16("MsgID", msg.GetMsgID()), zap.Uint32("DataLen", msg.GetDataLen()), zap.ByteString("Data", msg.GetData()))
	}
	return
}
