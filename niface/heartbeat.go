/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-30 18:27:46
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 19:55:36
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IHeartBeatChecker 心跳检测器接口
type IHeartBeatChecker interface {
	Start()                                  // 启动心跳检测
	Stop()                                   // 停止心跳检测
	SetHeartBeatMsgFunc(f HeartBeatMsgFunc)  // 设置心跳检测消息处理方法
	SetOnRemoteNotAlive(f OnRemoteNotAlive)  // 设置远程连接不存活时的处理方法
	BindRouter(msgID uint16, router IRouter) // 绑定心跳检测消息业务处理路由
	BindConn(conn IConnection)               // 绑定连接
	Clone() (checker IHeartBeatChecker)      // 克隆心跳检测器
	GetMsgID() (msgID uint16)                // 获取心跳检测消息ID
	GetMessage() (msg IMessage)              // 获取心跳检测消息
	GetRouter() (router IRouter)             // 获取心跳检测消息业务处理路由
}

// 用户自定义的心跳检测消息处理方法
type HeartBeatMsgFunc func() []byte

// 用户自定义的远程连接不存活时的处理方法
type OnRemoteNotAlive func(IConnection)

// HeartBeatOption 心跳检测
type HeartBeatOption struct {
	MakeMsg          HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	OnRemoteNotAlive OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	MsgID            uint16           // 用户自定义的心跳检测消息ID
	Router           IRouter          // 用户自定义的心跳检测消息业务处理路由
}

const (
	HeartBeatDefaultMsgID uint16 = 1 // 默认心跳检测消息ID
)
