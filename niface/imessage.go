/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:35:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 18:24:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/imessage.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMessage 消息接口
type IMessage interface {
	GetMsgID() uint32   // 获取消息ID
	GetDataLen() uint32 // 获取消息体长度
	GetData() []byte    // 获取消息内容
	SetMsgID(uint32)    // 设置消息ID
	SetDataLen(uint32)  // 设置消息体长度
	SetData([]byte)     // 设置消息内容
}
