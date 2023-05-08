/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 22:19:14
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 21:22:13
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/serveroverload"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/unmarshalmsg"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nserver"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建 Server
	s := nserver.NewServer(
		nserver.WithMulticore(true),
		nserver.WithReuseAddr(true),
		nserver.WithReusePort(true),
		nserver.WithLockOSThread(true),
	)
	// 设置当前 Server 的服务器人数超载检测器
	serveroverload.SetServerOverload(s)
	// 设置当前 Server 的心跳检测器
	heartbeat.SetHeartBeat(s, true)
	// 添加解析消息拦截器
	unmarshalmsg.AddInterceptor(s)
	go func() {
		// 创建一个通道，用于接收信号
		sc := make(chan os.Signal, 1)
		// 注册信号接收器
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		// 等待信号
		sig := <-sc
		nlog.Info("Server Interrupt Signal", nlog.String("Signal", sig.String()))
		// 停止服务器
		s.Stop()
	}()
	// 启动服务器
	s.Start()
	// 等待一段时间
	<-time.After(1 * time.Second)
}
