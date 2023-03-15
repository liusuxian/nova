/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 18:49:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 18:16:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/defaultpack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"encoding/binary"
	"github.com/liusuxian/nova/niface"
	"github.com/panjf2000/gnet/v2"
	"github.com/pkg/errors"
)

// DefaultPack 默认封包拆包结构，消息ID(2字节)-消息体长度(4字节)-消息内容
type DefaultPack struct {
	endian        uint8  // 字节存储次序，默认小端，1: 小端 2: 大端（单位:字节 uint8）
	maxPacketSize uint32 // 数据包的最大值，默认 4096（单位:字节 uint32）
}

// 自定义错误，不完整的包
var ErrIncompletePacket = errors.New("incomplete packet")

const (
	msgIdSize   = 2 // 消息ID长度 uint16(2字节)
	msgBodySize = 4 // 消息体长度 uint32(4字节)
)

// NewDefaultPack 创建默认封包拆包实例
func NewDefaultPack(endian uint8, maxPacketSize uint32) niface.IDataPack {
	return &DefaultPack{
		endian:        endian,
		maxPacketSize: maxPacketSize,
	}
}

// GetHeadLen 获取包头长度
func (p *DefaultPack) GetHeadLen() uint8 {
	return msgIdSize + msgBodySize
}

// Pack 封包
func (p *DefaultPack) Pack(msg niface.IMessage) (data []byte, err error) {
	// 创建消息包的缓冲区
	bodyOffset := msgIdSize + msgBodySize
	msgLen := bodyOffset + int(msg.GetDataLen())
	data = make([]byte, msgLen)
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	switch p.endian {
	case 1:
		// 小端
		endianOrder = binary.LittleEndian
	case 2:
		// 大端
		endianOrder = binary.BigEndian
	default:
		// 默认小端
		endianOrder = binary.LittleEndian
	}
	// 写消息ID
	msgIdBytes := make([]byte, msgIdSize)
	endianOrder.PutUint16(msgIdBytes, msg.GetMsgID())
	copy(data, msgIdBytes)
	// 写消息体长度
	endianOrder.PutUint32(data[msgIdSize:bodyOffset], msg.GetDataLen())
	// 写消息内容
	copy(data[bodyOffset:msgLen], msg.GetData())
	return
}

// Unpack 拆包
func (p *DefaultPack) Unpack(conn gnet.Conn) (data niface.IMessage, err error) {
	// 读消息头
	var buf []byte
	bodyOffset := msgIdSize + msgBodySize
	buf, _ = conn.Peek(bodyOffset)
	if len(buf) < bodyOffset {
		err = ErrIncompletePacket
		return
	}
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	switch p.endian {
	case 1:
		// 小端
		endianOrder = binary.LittleEndian
	case 2:
		// 大端
		endianOrder = binary.BigEndian
	default:
		// 默认小端
		endianOrder = binary.LittleEndian
	}
	// 读取并判断消息体长度是否超出我们允许的最大包长度
	msgBodyLen := endianOrder.Uint32(buf[msgIdSize:bodyOffset])
	if p.maxPacketSize > 0 && msgBodyLen > p.maxPacketSize {
		err = errors.New("Too Large Msg Data Received")
		return
	}
	// 读取整个消息数据
	msgLen := bodyOffset + int(msgBodyLen)
	if conn.InboundBuffered() < msgLen {
		err = ErrIncompletePacket
		return
	}
	buf, _ = conn.Peek(msgLen)
	_, _ = conn.Discard(msgLen)
	// 创建 Message 消息包
	msgID := endianOrder.Uint16(buf[:msgIdSize])
	msgData := buf[bodyOffset:msgLen]
	data = NewMsgPackage(msgID, msgData)
	return
}
