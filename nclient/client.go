/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 19:43:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-16 19:25:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nclient/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nclient

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/npack"
	"github.com/panjf2000/gnet/v2"
	"github.com/pkg/errors"
)

// Client 客户端结构
type Client struct {
	gnet.BuiltinEventEngine
	client  *gnet.Client // gnet 客户端
	conn    gnet.Conn    // gnet Conn
	network string       // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	addr    string       // 服务器地址
}

// NewClient 创建 Client
func NewClient(network, addr string, options gnet.Options) niface.IClient {
	c := &Client{
		network: network,
		addr:    addr,
	}
	client, err := gnet.NewClient(c, gnet.WithOptions(options))
	if err != nil {
		panic(errors.Wrapf(err, "New Client Error"))
	}
	c.client = client

	return c
}

// Start 启动客户端
func (c *Client) Start() {
	conn, err := c.client.Dial(c.network, c.addr)
	if err != nil {
		panic(errors.Wrapf(err, "Start Client Error"))
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

// 给当前 Client 添加路由
func (c *Client) AddRouter(msgID uint16, router niface.IRouter) {
}

// 当前 Client 的连接信息
func (c *Client) Conn() niface.IConnection {
	return nil
}

// 设置当前 Client 的连接创建时Hook函数
func (c *Client) SetOnConnStart(f func(niface.IConnection)) {
}

// 设置当前 Client 的连接断开时的Hook函数
func (c *Client) SetOnConnStop(f func(niface.IConnection)) {
}

// 获取当前 Client 的连接创建时Hook函数
func (c *Client) GetOnConnStart() func(niface.IConnection) {
	return nil
}

// 获取当前 Client 的连接断开时的Hook函数
func (c *Client) GetOnConnStop() func(niface.IConnection) {
	return nil
}

// 设置当前 Client 绑定的数据协议封包方式
func (c *Client) SetPacket(pack niface.IDataPack) {
}

// 获取当前 Client 绑定的数据协议封包方式
func (c *Client) GetPacket() niface.IDataPack {
	return nil
}

// 获取当前 Client 绑定的消息处理模块
func (c *Client) GetMsgHandler() niface.IMsgHandle {
	return nil
}

// 设置心跳检测
func (c *Client) SetHeartBeat(option *niface.HeartBeatOption) {
}
