/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 20:58:27
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 22:00:04
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/datapack.go
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

var defaultHeaderLen uint32 = 8

// DataPack 封包拆包类实例，暂时不需要成员
type DataPack struct{}

// NewDataPack 封包拆包实例初始化方法
func NewDataPack() niface.IDataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// ID uint32(4字节) + DataLen uint32(4字节)
	return defaultHeaderLen
}

// Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg niface.IMessage) (data []byte, err error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	// 写dataLen
	if err = binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return
	}
	// 写msgID
	if err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return
	}
	// 写data数据
	if err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return
	}
	data = dataBuff.Bytes()
	return
}

// Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (data niface.IMessage, err error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	// 只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	// 读dataLen
	if err = binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return
	}
	// 读msgID
	if err = binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return
	}
	// 判断dataLen的长度是否超出我们允许的最大包长度
	maxPacketSize := nconf.GetUint32("server.maxPacketSize")
	if maxPacketSize > 0 && msg.DataLen > maxPacketSize {
		err = errors.New("Too Large Msg Data Received")
		return
	}
	// 这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	data = msg
	return
}
