/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:44:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 12:05:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Tcp Client Start...")
	time.Sleep(1 * time.Second)
	// 连接服务器，得到一个conn连接
	var conn net.Conn
	var err error
	if conn, err = net.Dial("tcp", "127.0.0.1:8989"); err != nil {
		fmt.Printf("Tcp Client Start Error: %v\n", err)
		return
	}
	// 连接调用Write写数据
	for {
		if _, err = conn.Write([]byte("Hello Nova v1.0.0")); err != nil {
			fmt.Printf("Write Conn Error: %v\n", err)
			return
		}
		// 读取
		buf := make([]byte, 512)
		var cnt int
		if cnt, err = conn.Read(buf); err != nil {
			fmt.Printf("Recv Buf Error: %v\n", err)
			return
		}
		fmt.Printf("Tcp Server Call Back: %s Cnt: %d\n", buf, cnt)
		// 阻塞
		time.Sleep(1 * time.Second)
	}
}
