/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 21:08:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 18:30:04
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/message.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

// Message 消息结构
type Message struct {
	ID      uint32 // 消息的ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		ID:      ID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// Init 初始化Message消息包
func (msg *Message) InitMsgPackage(ID uint32, data []byte) {
	msg.ID = ID
	msg.DataLen = uint32(len(data))
	msg.Data = data
}

// GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

// GetDataLen 获取消息体长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetMsgID 设置消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

// SetDataLen 设置消息体长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// SetData 设置消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
