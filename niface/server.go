/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 22:18:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "context"

// IServer 服务器接口
type IServer interface {
	Start()                                 // 启动 Server
	Stop()                                  // 停止 Server
	AddRouter(msgID uint16, router IRouter) // 给当前 Server 添加路由
	GetCtx() context.Context                // 获取当前 Server 的 Context
	GetConnManager() IConnManager           // 获取当前 Server 的连接管理
	GetConnections() int                    // 获取当前 Server 的活跃连接数
	SetOnConnStart(func(IConnection))       // 设置当前 Server 的连接创建时的 Hook 函数
	SetOnConnStop(func(IConnection))        // 设置当前 Server 的连接断开时的 Hook 函数
	GetOnConnStart() func(IConnection)      // 获取当前 Server 的连接创建时的 Hook 函数
	GetOnConnStop() func(IConnection)       // 获取当前 Server 的连接断开时的 Hook 函数
	SetPacket(IDataPack)                    // 设置当前 Server 绑定的数据协议封包和拆包方式
	GetPacket() IDataPack                   // 获取当前 Server 绑定的数据协议封包和拆包方式
	GetMsgHandler() IMsgHandle              // 获取当前 Server 绑定的消息处理模块
	SetOverLoadMsg(*OverLoadMsgOption)      // 设置当前 Server 的服务器人数超载消息
	SetHeartBeat(*HeartBeatOption)          // 设置当前 Server 的心跳检测
}

// 用户自定义的服务器人数超载消息处理方法
type OverLoadMsgFunc func() []byte

// OverLoadMsgOption 服务器人数超载消息选项
type OverLoadMsgOption struct {
	MakeMsg OverLoadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	MsgID   uint16          // 用户自定义的服务器人数超载消息ID
	Router  IRouter         // 用户自定义的服务器人数超载消息业务处理路由
}

const (
	OverLoadDefaultMsgID uint16 = 0 // 默认服务器人数超载消息ID
)
