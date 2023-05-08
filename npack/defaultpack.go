/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 13:49:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-09 02:16:28
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/defaultpack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"encoding/binary"
	"github.com/liusuxian/nova/niface"
	"github.com/pkg/errors"
)

// defaultPack 默认封包拆包结构，消息ID(2字节)-消息体长度(4字节)-消息内容
type defaultPack struct {
	endian        int // 字节存储次序，1: 小端 2: 大端，默认 1
	maxPacketSize int // 数据包的最大值（单位:字节），默认 4096
}

const (
	msgIdLen   int = 2 // 消息ID长度 uint16(2字节)
	msgBodyLen int = 4 // 消息体长度 uint32(4字节)
)

// newDefaultPack 创建默认封包拆包实例
func newDefaultPack(endian, maxPacketSize int) (packet niface.IDataPack) {
	return &defaultPack{
		endian:        endian,
		maxPacketSize: maxPacketSize,
	}
}

// GetHeadLen 获取包头长度(字节数)
func (p *defaultPack) GetHeadLen() (headLen int) {
	return msgIdLen + msgBodyLen
}

// Pack 封包
func (p *defaultPack) Pack(msg niface.IMessage) (data []byte) {
	// 创建消息包的缓冲区
	headLen := p.GetHeadLen()
	msgLen := headLen + msg.GetDataLen()
	data = make([]byte, msgLen)
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	switch p.endian {
	case LittleEndian:
		// 小端
		endianOrder = binary.LittleEndian
	case BigEndian:
		// 大端
		endianOrder = binary.BigEndian
	default:
		// 默认小端
		endianOrder = binary.LittleEndian
	}
	// 写消息 ID
	msgIdBytes := make([]byte, msgIdLen)
	endianOrder.PutUint16(msgIdBytes, msg.GetMsgID())
	copy(data, msgIdBytes)
	// 写消息体长度
	endianOrder.PutUint32(data[msgIdLen:headLen], uint32(msg.GetDataLen()))
	// 写消息内容
	copy(data[headLen:], msg.GetData())
	return
}

// 拆包头
func (p *defaultPack) UnPackHead(headBuf []byte) (msg niface.IMessage, err error) {
	headLen := p.GetHeadLen()
	if len(headBuf) < headLen {
		err = ErrIncompletePacket
		return
	}
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	switch p.endian {
	case LittleEndian:
		// 小端
		endianOrder = binary.LittleEndian
	case BigEndian:
		// 大端
		endianOrder = binary.BigEndian
	default:
		// 默认小端
		endianOrder = binary.LittleEndian
	}
	// 读消息 ID
	msgID := endianOrder.Uint16(headBuf[:msgIdLen])
	// 读消息体长度
	bodyLen := int(endianOrder.Uint32(headBuf[msgIdLen:]))
	// 创建 Message 消息包
	msg = NewMsgPackage(msgID, nil)
	// 设置消息体长度
	msg.SetDataLen(bodyLen)
	// 判断消息体长度是否超出我们允许的最大包长度
	if p.maxPacketSize > 0 && bodyLen > p.maxPacketSize {
		err = errors.New("too large msg data received")
		return
	}
	return
}

// UnPackBody 拆包体
func (p *defaultPack) UnPackBody(msgBuf []byte, msg niface.IMessage) {
	// 读消息内容
	msgData := make([]byte, msg.GetDataLen())
	copy(msgData, msgBuf[p.GetHeadLen():])
	// 设置消息内容
	msg.SetData(msgData)
}
