/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 00:43:57
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nserver"
	"github.com/panjf2000/gnet/v2"
)

func main() {
	// 创建 Server
	s := nserver.NewServer(gnet.Options{
		Multicore: true,
		ReuseAddr: true,
		ReusePort: true,
		Ticker:    true,
	})
	// 启动服务器
	s.Start()
}
