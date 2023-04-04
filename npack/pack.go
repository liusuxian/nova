/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 15:42:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-04 10:43:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/pack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"github.com/liusuxian/nova/niface"
	"github.com/pkg/errors"
)

const (
	DefaultPacketMethod int = 1 // 默认封包拆包方式
)
const (
	LittleEndian int = iota + 1 // 字节存储次序，小端
	BigEndian                   // 字节存储次序，大端
)

// 自定义错误，不完整的包
var ErrIncompletePacket = errors.New("incomplete packet")

// NewPack 创建一个具体的封包拆包对象
func NewPack(packetMethod, endian, maxPacketSize int) (dataPack niface.IDataPack) {
	switch packetMethod {
	case DefaultPacketMethod:
		// 默认封包拆包方式
		// 消息ID(2字节)-消息体长度(4字节)-消息内容
		dataPack = newDefaultPack(endian, maxPacketSize)
	default:
		// 未知的封包拆包方式，则使用默认封包拆包方式
		dataPack = newDefaultPack(endian, maxPacketSize)
	}
	return
}
