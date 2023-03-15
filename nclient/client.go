/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 19:43:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 14:51:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nclient/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nclient

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/panjf2000/gnet/v2"
	"go.uber.org/zap"
)

// Client 客户端结构
type Client struct {
	gnet.BuiltinEventEngine
	client  *gnet.Client    // gnet 客户端
	conn    gnet.Conn       // gnet Conn
	ctx     context.Context // 客户端 Context
	network string          // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	addr    string          // 服务器地址
}

// NewClient 创建 Client
func NewClient(network, addr string, options gnet.Options) niface.IClient {
	c := &Client{
		ctx:     context.Background(),
		network: network,
		addr:    addr,
	}
	client, err := gnet.NewClient(c, gnet.WithOptions(options))
	if err != nil {
		nlog.Fatal(c.ctx, "New Client Error", zap.Error(err))
	}
	c.client = client

	return c
}

// Start 启动客户端
func (c *Client) Start() {
	conn, err := c.client.Dial(c.network, c.addr)
	if err != nil {
		nlog.Fatal(c.ctx, "Start Client Error", zap.Error(err))
	}
	c.conn = conn
	// 创建一个封包对象
	dp := npack.Factory().NewPack(1, 1, 4096)
	// 封装一个msg1包
	msg1 := npack.NewMsgPackage(1, []byte("hello"))
	sendData1, _ := dp.Pack(msg1)
	// 封装一个msg2包
	msg2 := npack.NewMsgPackage(2, []byte("world!!"))
	sendData2, _ := dp.Pack(msg2)
	// 将sendData1和sendData2拼接一起，组成粘包
	sendData1 = append(sendData1, sendData2...)
	c.conn.Write(sendData1)
	c.client.Start()
}

// Stop 停止客户端
func (c *Client) Stop() {
	c.client.Stop()
}
