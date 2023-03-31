/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 13:54:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 20:39:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/message.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

// Message 消息结构
type Message struct {
	ID      uint16 // 消息的 ID
	DataLen int    // 消息的长度
	Data    []byte // 消息的内容
}

// NewMsgPackage 创建 Message 消息包
func NewMsgPackage(ID uint16, data []byte) (msg *Message) {
	return &Message{
		ID:      ID,
		DataLen: len(data),
		Data:    data,
	}
}

// SetMsgID 设置消息 ID
func (msg *Message) SetMsgID(msgID uint16) {
	msg.ID = msgID
}

// SetDataLen 设置消息体长度
func (msg *Message) SetDataLen(dataLen int) {
	msg.DataLen = dataLen
}

// SetData 设置消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

// GetMsgID 获取消息 ID
func (msg *Message) GetMsgID() (msgID uint16) {
	return msg.ID
}

// GetDataLen 获取消息体长度
func (msg *Message) GetDataLen() (dataLen int) {
	return msg.DataLen
}

// GetData 获取消息内容
func (msg *Message) GetData() (data []byte) {
	return msg.Data
}
