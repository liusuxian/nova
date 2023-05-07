/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 17:25:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-07 22:26:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/serveroverload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IServerOverloadChecker 服务器人数超载检测器接口
type IServerOverloadChecker interface {
	Check(server IServer, maxConn int) (isOverload bool) // 服务器人数超载检测
	SetServerOverloadMsgFunc(f ServerOverloadMsgFunc)    // 设置服务器人数超载消息处理方法
	BindRouter(msgID uint16, handlers ...RouterHandler)  // 绑定服务器人数超载消息的业务处理器集合
	GetMsgID() (msgID uint16)                            // 获取服务器人数超载消息ID
	GetMessage() (msg IMessage)                          // 获取服务器人数超载消息
	GetHandlers() (handlers []RouterHandler)             // 获取服务器人数超载消息的业务处理器集合
}

// ServerOverloadMsgFunc 用户自定义的服务器人数超载消息处理方法
type ServerOverloadMsgFunc func() (buf []byte)

// ServerOverloadOption 服务器人数超载检测选项
type ServerOverloadOption struct {
	MakeMsg        ServerOverloadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	MsgID          uint16                // 用户自定义的服务器人数超载消息ID
	RouterHandlers []RouterHandler       // 用户自定义的服务器人数超载消息的业务处理器集合
}

const (
	ServerOverloadDefaultMsgID uint16 = 0 // 默认服务器人数超载消息ID
)
