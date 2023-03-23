/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 21:39:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 22:02:50
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/noverload/overload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package noverload

import "github.com/liusuxian/nova/niface"

// OverLoadMsg 服务器人数超载消息结构
type OverLoadMsg struct {
	makeMsg niface.OverLoadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	msgID   uint16                 // 用户自定义的服务器人数超载消息ID
}

// NewOverLoadMsg 创建服务器人数超载消息
func NewOverLoadMsg() *OverLoadMsg {
	overLoadMsg := &OverLoadMsg{
		makeMsg: makeMsgDefaultFunc,
		msgID:   niface.OverLoadDefaultMsgID,
	}
	return overLoadMsg
}

// SetOverLoadMsgID 设置服务器人数超载消息ID
func (ol *OverLoadMsg) SetOverLoadMsgID(msgID uint16) {
	if msgID != niface.OverLoadDefaultMsgID {
		ol.msgID = msgID
	}
}

// SetOverLoadMsgFunc 设置服务器人数超载消息处理方法
func (ol *OverLoadMsg) SetOverLoadMsgFunc(f niface.OverLoadMsgFunc) {
	if f != nil {
		ol.makeMsg = f
	}
}

// GetMsgID 获取服务器人数超载消息ID
func (ol *OverLoadMsg) GetMsgID() uint16 {
	return ol.msgID
}

// GetMsgData 获取服务器人数超载消息数据
func (ol *OverLoadMsg) GetMsgData() []byte {
	return ol.makeMsg()
}

// makeMsgDefaultFunc 默认的服务器人数超载消息处理方法
func makeMsgDefaultFunc() []byte {
	return []byte("server overload")
}
