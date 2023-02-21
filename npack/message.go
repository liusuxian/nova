/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 21:08:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 21:10:04
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/message.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

// Message 消息结构
type Message struct {
	DataLen uint32 // 消息的长度
	ID      uint32 // 消息的ID
	Data    []byte // 消息的内容
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		ID:      ID,
		Data:    data,
	}
}

// Init 初始化消息
func (msg *Message) Init(ID uint32, data []byte) {
	msg.ID = ID
	msg.Data = data
	msg.DataLen = uint32(len(data))
}

// GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetDataLen 设置消息的长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// SetMsgID 设置消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

// SetData 设置消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
