/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 20:34:11
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-21 17:19:30
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nclient"
	"github.com/panjf2000/gnet/v2"
	"log"
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
	sig := <-c
	log.Printf("收到退出信号 %s 客户端将退出\n", sig.String())
}
