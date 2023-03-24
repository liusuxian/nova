/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 18:10:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-24 14:26:16
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nserver"
	"go.uber.org/zap"
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
		nserver.WithTicker(true),
	)
	// 设置当前 Server 的服务器人数超载消息
	s.SetOverLoadMsg(nil)
	// 设置当前 Server 的心跳检测
	s.SetHeartBeat(nil, true)
	go func() {
		// 创建一个通道，用于接收信号
		sc := make(chan os.Signal, 1)
		// 注册信号接收器
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		// 等待信号
		sig := <-sc
		nlog.Info(s.GetCtx(), "Server Interrupt Signal", zap.String("Signal", sig.String()))
		// 停止服务器
		s.Stop()
	}()
	// 启动服务器
	s.Start()
	// 等待一段时间
	idleTimeout := time.NewTimer(2 * time.Second)
	defer idleTimeout.Stop()
	select {
	case <-idleTimeout.C:
	}
}
