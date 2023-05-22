/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 22:19:14
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-22 17:38:08
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/redisdb"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/router"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/serveroverload"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/unmarshalmsg"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nserver"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	// 创建 Server
	s := nserver.NewServer(func(sc *nserver.ServerConfig) {
		sc.NumEventLoop = runtime.NumCPU() * 2
		sc.ReuseAddr = true
		sc.ReusePort = true
		sc.LockOSThread = true
		sc.TCPKeepAlive = time.Second * 30
	})
	// 设置当前 Server 启动时的 Hook 函数
	s.SetOnStart(func(s niface.IServer) {
		redisdb.Start()
	})
	// 设置当前 Server 停止时的 Hook 函数
	s.SetOnStop(func(s niface.IServer) {
		redisdb.Close()
	})
	// 设置当前 Server 的服务器人数超载检测器
	serveroverload.SetServerOverload(s)
	// 设置当前 Server 的心跳检测器
	heartbeat.SetHeartBeat(s, true)
	// 启动路由
	router.StartRouter(s)
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
