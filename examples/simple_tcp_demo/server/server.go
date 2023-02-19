/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:44:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 12:06:33
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import "github.com/liusuxian/nova/nnet"

func main() {
	// 创建一个TCP服务器
	s := nnet.NewTCPServer("Nova v1.0.0", "0.0.0.0", 8989)
	// 启动server
	s.Server()
}
