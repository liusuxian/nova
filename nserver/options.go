/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 00:27:32
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 21:27:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nserver/options.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nserver

import "github.com/liusuxian/nova/niface"

// TCPServer 的 Option
type TCPOption func(s *TCPServer)

// 只要实现 Packet 接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(pack niface.IDataPack) TCPOption {
	return func(s *TCPServer) {
		// TODO
		// s.SetPacket(pack)
	}
}

// TODO Client 的 Option
// type ClientOption func(c *Client)

// TODO 只要实现 Packet 接口可自由实现数据包解析格式，如果没有则使用默认解析格式
// func WithPacketClient(pack niface.IDataPack) ClientOption {
// 	return func(c *Client) {
// 		c.SetPacket(pack)
// 	}
// }
