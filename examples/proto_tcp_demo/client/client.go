/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 22:43:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 16:00:53
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package main

import (
	"context"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/router"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/serveroverload"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/unmarshalmsg"
	"github.com/liusuxian/nova/nclient"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
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
			c := nclient.NewClient(func(cc *nclient.ClientConfig) {
				cc.LockOSThread = true
				cc.Network = "tcp"
				cc.Addr = "127.0.0.1:8099"
				cc.HeartBeat = time.Duration(10000) * time.Millisecond
				cc.MaxHeartBeat = time.Duration(15000) * time.Millisecond
				cc.PacketMethod = npack.DefaultPacketMethod
				cc.Endian = npack.LittleEndian
				cc.MaxPacketSize = 4096
			})
			// 设置当前 Client 的服务器人数超载检测器
			serveroverload.SetServerOverload(c)
			// 设置当前 Client 的心跳检测器
			heartbeat.SetHeartBeat(c, false)
			// 启动路由
			r := router.StartRouter(c)
			// 添加解析消息拦截器
			unmarshalmsg.AddInterceptor(c, r)
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
