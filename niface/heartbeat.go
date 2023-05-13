/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-30 18:27:46
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-13 19:38:46
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// IHeartBeatChecker 心跳检测器接口
type IHeartBeatChecker interface {
	Start()                                 // 启动心跳检测
	Stop()                                  // 停止心跳检测
	SetHeartBeatMsgFunc(f HeartBeatMsgFunc) // 设置心跳检测消息处理方法
	SetOnRemoteNotAlive(f OnRemoteNotAlive) // 设置远程连接不存活时的处理方法
	BindConn(conn IConnection)              // 绑定连接
	Clone() (checker IHeartBeatChecker)     // 克隆心跳检测器
	SetMsgID(msgID uint16)                  // 设置心跳检测消息ID
}

// HeartBeatMsgFunc 用户自定义的心跳检测消息处理方法
type HeartBeatMsgFunc = MsgDataFunc

// OnRemoteNotAlive 用户自定义的远程连接不存活时的处理方法
type OnRemoteNotAlive func(conn IConnection)

// HeartBeatOption 心跳检测选项
type HeartBeatOption struct {
	MakeMsg          HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	OnRemoteNotAlive OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	MsgID            uint16           // 用户自定义的心跳检测消息ID
}

const (
	HeartBeatDefaultMsgID uint16 = 1 // 默认心跳检测消息ID
)
