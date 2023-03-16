/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 23:09:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-16 19:00:58
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nclient"
	"github.com/panjf2000/gnet/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := nclient.NewClient("tcp", "127.0.0.1:8888", gnet.Options{})
			c.Start()
		}()
	}
	wg.Wait()

	// 创建一个通道，用于接收信号
	c := make(chan os.Signal, 1)
	// 注册信号接收器
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	<-c
}
