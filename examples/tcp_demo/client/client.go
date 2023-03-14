/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 23:09:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 20:25:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/client/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/nclient"
	"github.com/panjf2000/gnet/v2"
)

func main() {
	c := nclient.NewTCPClient("127.0.0.1", 8888, gnet.Options{})
	c.Start()
}
