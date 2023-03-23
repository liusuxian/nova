/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 18:10:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 16:00:44
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/server/server.go
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
		nserver.WithLockOSThread(true),
		nserver.WithTicker(true),
	)
	// 设置当前 Server 的心跳检测
	s.SetHeartBeat(nil)
	// 启动服务器
	s.Start()
}
