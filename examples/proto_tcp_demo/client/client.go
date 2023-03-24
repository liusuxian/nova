/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 22:19:14
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-24 14:59:08
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/overload"
	"github.com/liusuxian/nova/nclient"
	"github.com/liusuxian/nova/nlog"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	clientNum := 4
	for i := 0; i < clientNum; i++ {
		go func(ctx context.Context) {
			// 创建 Client
			c := nclient.NewClient(
				"tcp",
				"05807165157c4471.natapp.cc:1688",
				nclient.WithLockOSThread(true),
				nclient.WithTicker(true),
				nclient.WithHeartbeat(time.Duration(3000)*time.Millisecond),
				nclient.WithMaxHeartbeat(time.Duration(5000)*time.Millisecond),
			)
			// 设置当前 Client 的服务器人数超载消息
			overload.SetOverLoadMsg(c)
			// 设置当前 Client 的心跳检测
			heartbeat.SetHeartBeat(c, false)
			// 启动 Client
			c.Start()
			select {
			case <-ctx.Done():
				c.Stop()
				return
			}
		}(cancelCtx)
	}
	// 创建一个通道，用于接收信号
	sc := make(chan os.Signal, 1)
	// 注册信号接收器
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	sig := <-sc
	nlog.Info(cancelCtx, "Client Interrupt Signal", zap.String("Signal", sig.String()))
	// 取消任务
	cancelFunc()
	// 等待一段时间
	idleTimeout := time.NewTimer(2 * time.Second)
	defer idleTimeout.Stop()
	select {
	case <-idleTimeout.C:
	}
}
