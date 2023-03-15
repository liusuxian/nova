/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 20:17:18
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 20:26:20
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/connection.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"net"
)

// IConnection 连接接口
type IConnection interface {
	Start()                                      // 启动连接
	Stop()                                       // 停止连接
	Context() context.Context                    // 返回 Context，用于用户自定义的 Goroutine 获取连接退出状态
	GetConnection() gnet.Conn                    // 从当前连接获取原始的 Socket Conn
	GetConnID() string                           // 获取当前 ConnID
	RemoteAddr() net.Addr                        // 获取当前连接远程地址信息
	LocalAddr() net.Addr                         // 获取当前连接本地地址信息
	SendMsg(msgID uint16, data []byte) error     // 直接将 Message 数据发送给远程的客户端(无缓冲)
	SendBuffMsg(msgID uint16, data []byte) error // 直接将 Message 数据发送给远程的客户端(有缓冲)
	SetProperty(key string, value any)           // 设置当前连接属性
	GetProperty(key string) (any, error)         // 获取当前连接属性
	RemoveProperty(key string)                   // 移除当前连接属性
	IsAlive() bool                               // 判断当前连接是否存活
}

// 用户自定义的心跳检测消息处理方法
type HeartBeatMsgFunc func(IConnection) []byte

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
	HeartBeatDefaultMsgID uint16 = 0 // 默认心跳检测消息ID
)
