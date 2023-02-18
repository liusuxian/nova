/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:35:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:45:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/imessage.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMessage 消息接口
type IMessage interface {
	GetDataLen() uint32 // 获取消息数据段长度
	GetMsgID() uint32   // 获取消息ID
	GetData() []byte    // 获取消息内容
	SetMsgID(uint32)    // 设置消息ID
	SetData([]byte)     // 设置消息内容
	SetDataLen(uint32)  // 设置消息数据段长度
}
