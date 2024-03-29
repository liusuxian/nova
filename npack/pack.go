/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 15:42:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:19:30
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package npack

import "github.com/liusuxian/nova/niface"

const (
	DefaultPacketMethod int = 1 // 默认封包拆包方式
)
const (
	LittleEndian int = iota + 1 // 字节存储次序，小端
	BigEndian                   // 字节存储次序，大端
)

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
