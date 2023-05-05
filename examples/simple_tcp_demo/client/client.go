/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 20:34:11
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 00:50:45
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"github.com/liusuxian/nova/nclient"
	"github.com/liusuxian/nova/nlog"
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
				"127.0.0.1:8066",
				nclient.WithLockOSThread(true),
				nclient.WithHeartbeat(time.Duration(10000)*time.Millisecond),
				nclient.WithMaxHeartbeat(time.Duration(15000)*time.Millisecond),
			)
			// 设置当前 Client 的服务器人数超载检测器
			c.SetServerOverload()
			// 设置当前 Client 的心跳检测器
			c.SetHeartBeat(false)
			// 启动 Client
			c.Start()
			// 停止 Client
			<-ctx.Done()
			c.Stop()
		}(cancelCtx)
	}
	// 创建一个通道，用于接收信号
	sc := make(chan os.Signal, 1)
	// 注册信号接收器
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	sig := <-sc
	nlog.Info("Client Interrupt Signal", nlog.String("Signal", sig.String()))
	// 取消任务
	cancelFunc()
	// 等待一段时间
	<-time.After(1 * time.Second)
}
