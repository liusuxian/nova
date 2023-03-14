/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 19:39:58
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nserver/tcp_server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nserver

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
	"time"
)

// TCPServer TCP 服务器结构
type TCPServer struct {
	gnet.BuiltinEventEngine
	eng     gnet.Engine
	options gnet.Options    // TCP 服务器启动选项
	ctx     context.Context // TCP 服务器 Context
	port    uint16          // TCP 服务器 端口
}

// NewTCPServer 创建 TCPServer
func NewTCPServer(port uint16, options gnet.Options) niface.IServer {
	s := &TCPServer{
		options: options,
		ctx:     context.Background(),
		port:    port,
	}

	return s
}

// Start 启动服务器
func (s *TCPServer) Start() {
	if err := gnet.Run(s, fmt.Sprintf("tcp://:%d", s.port), gnet.WithOptions(s.options)); err != nil {
		nlog.Fatal(s.ctx, "Start TCP Server Error", zap.Error(err))
	}
}

// Stop 停止服务器
func (s *TCPServer) Stop() {
	_ = s.eng.Stop(s.ctx)
}

// OnBoot 在引擎准备好接受连接时触发。参数 engine 包含信息和各种实用工具。
func (s *TCPServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	nlog.Info(s.ctx, "TCPServer OnBoot", zap.String("listening", fmt.Sprintf("tcp://:%d", s.port)), zap.Any("options", s.options))
	s.eng = eng
	return
}

// OnClose 在连接关闭时触发。参数 err 是最后已知的连接错误。
func (s *TCPServer) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	nlog.Info(s.ctx, "TCPServer OnClose", zap.String("RemoteAddr", conn.RemoteAddr().String()))
	return
}

// OnOpen 在新连接打开时触发。参数 out 是将要发送回对等方的返回值。
func (s *TCPServer) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	nlog.Info(s.ctx, "TCPServer OnOpen")
	return
}

// OnShutdown 在引擎被关闭时触发，它在所有事件循环和连接关闭后立即调用。
func (s *TCPServer) OnShutdown(eng gnet.Engine) {
	nlog.Info(s.ctx, "TCPServer OnShutdown")
	return
}

// OnTick 在引擎启动后立即触发，并在 delay 返回值指定的持续时间后再次触发。
func (s *TCPServer) OnTick() (delay time.Duration, action gnet.Action) {
	nlog.Debug(s.ctx, "TCPServer OnTick")
	delay = nconf.MaxHeartbeat()
	return
}

// OnTraffic 在本地套接字从对等方接收数据时触发。
func (s *TCPServer) OnTraffic(conn gnet.Conn) (action gnet.Action) {
	var buf []byte
	conn.Read(buf)
	nlog.Info(s.ctx, "TCPServer OnTraffic", zap.ByteString("data", buf))
	return
}
