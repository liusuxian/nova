/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 19:25:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/nserver"
	"github.com/panjf2000/gnet/v2"
)

func main() {
	s := nserver.NewTCPServer(8888, gnet.Options{
		Multicore:    true,
		ReuseAddr:    true,
		ReusePort:    true,
		Ticker:       true,
		TCPKeepAlive: nconf.MaxHeartbeat(),
	})
	s.Start()
}
