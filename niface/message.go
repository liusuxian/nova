/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 01:01:50
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:15:51
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
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
