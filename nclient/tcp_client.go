/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 19:43:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 20:21:10
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nclient/tcp_client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nclient

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
)

// TCPClient TCP 客户端结构
type TCPClient struct {
	gnet.BuiltinEventEngine
	client *gnet.Client    // gnet 客户端
	conn   gnet.Conn       // gnet Conn
	ctx    context.Context // TCP 客户端 Context
	ip     string          // TCP 服务器 ip
	port   uint16          // TCP 服务器端口
}

// NewTCPClient 创建 TCPClient
func NewTCPClient(ip string, port uint16, options gnet.Options) niface.IClient {
	c := &TCPClient{
		ctx:  context.Background(),
		ip:   ip,
		port: port,
	}
	client, err := gnet.NewClient(c, gnet.WithOptions(options))
	if err != nil {
		nlog.Fatal(c.ctx, "New TCP Client Error", zap.Error(err))
	}
	c.client = client

	return c
}

// Start 启动客户端
func (c *TCPClient) Start() {
	conn, err := c.client.Dial("tcp", fmt.Sprintf("%s:%d", c.ip, c.port))
	if err != nil {
		nlog.Fatal(c.ctx, "Start TCP Client Error", zap.Error(err))
	}
	c.conn = conn
	c.conn.Write([]byte("Start TCP Client"))
	c.client.Start()
}

// Stop 停止客户端
func (c *TCPClient) Stop() {
	c.client.Stop()
}
