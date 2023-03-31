/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:35:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-30 19:04:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/message.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMessage 消息接口
type IMessage interface {
	SetMsgID(msgID uint16)     // 设置消息 ID
	SetDataLen(dataLen int)    // 设置消息体长度
	SetData(data []byte)       // 设置消息内容
	GetMsgID() (msgID uint16)  // 获取消息 ID
	GetDataLen() (dataLen int) // 获取消息体长度
	GetData() (data []byte)    // 获取消息内容
}
