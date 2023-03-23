/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 20:34:11
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 16:52:48
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
			)
			// 设置当前 Client 的心跳检测
			c.SetHeartBeat(nil)
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
	time.Sleep(5 * time.Second)
}
