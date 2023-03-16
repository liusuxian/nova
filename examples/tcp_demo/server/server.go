/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-16 17:56:18
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import "github.com/liusuxian/nova/nserver"

func main() {
	// 创建 Server
	s := nserver.NewServer(
		nserver.WithMulticore(true),
		nserver.WithReuseAddr(true),
		nserver.WithReusePort(true),
		nserver.WithReadBufferCap(1024),
		nserver.WithWriteBufferCap(1024),
		nserver.WithLockOSThread(true),
		nserver.WithTicker(true),
		nserver.WithSocketRecvBuffer(1024),
		nserver.WithSocketSendBuffer(1024),
	)
	// 启动服务器
	s.Start()
}
