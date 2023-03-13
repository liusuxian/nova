/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 18:49:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 18:01:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/defaultpack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"bytes"
	"encoding/binary"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/pkg/errors"
)

// DefaultPack 默认封包拆包结构，消息ID(4字节)-消息体长度(4字节)-消息内容
type DefaultPack struct {
}

// 默认包头长度，消息ID uint32(4字节) + 消息体 uint32(4字节)
var defaultHeadLen uint8 = 8

// NewDefaultPack 创建默认封包拆包实例
func NewDefaultPack() niface.IDataPack {
	return &DefaultPack{}
}

// GetHeadLen 获取包头长度
func (p *DefaultPack) GetHeadLen() uint8 {
	return defaultHeadLen
}

// Pack 封包
func (p *DefaultPack) Pack(msg niface.IMessage) (data []byte, err error) {
	// 创建一个存放字节切片的缓冲区
	dataBuf := bytes.NewBuffer([]byte{})
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	endian := nconf.Endian()
	switch endian {
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
	if err = binary.Write(dataBuf, endianOrder, msg.GetMsgID()); err != nil {
		return
	}
	// 写消息体长度
	if err = binary.Write(dataBuf, endianOrder, msg.GetDataLen()); err != nil {
		return
	}
	// 写消息内容
	if err = binary.Write(dataBuf, endianOrder, msg.GetData()); err != nil {
		return
	}
	data = dataBuf.Bytes()
	return
}

// Unpack 拆包
func (p *DefaultPack) Unpack(binaryData []byte) (data niface.IMessage, err error) {
	// 创建一个存放二进制数据的ioReader
	dataBuf := bytes.NewReader(binaryData)
	// 只拆包包头的信息，得到消息ID和消息体长度
	msg := &Message{}
	// 获取字节存储次序
	var endianOrder binary.ByteOrder
	endian := nconf.Endian()
	switch endian {
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
	// 读消息ID
	if err = binary.Read(dataBuf, endianOrder, &msg.ID); err != nil {
		return
	}
	// 读消息体长度
	if err = binary.Read(dataBuf, endianOrder, &msg.DataLen); err != nil {
		return
	}
	// 判断消息体长度是否超出我们允许的最大包长度
	maxPacketSize := nconf.MaxPacketSize()
	if maxPacketSize > 0 && msg.DataLen > maxPacketSize {
		err = errors.New("Too Large Msg Data Received")
		return
	}
	// 这里只需要把包头的数据拆包出来就可以了，然后再通过包头中的消息体长度，再从连接中读取一次数据
	data = msg
	return
}
