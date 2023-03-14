/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 00:18:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iserver.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IServer 服务器接口
type IServer interface {
	Start()                                 // 启动服务器
	Stop()                                  // 停止服务器
	AddRouter(msgID uint32, router IRouter) // 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
	GetConnManager() IConnManager           // 获取连接管理
	SetOnConnStart(func(IConnection))       // 设置当前 Server 的连接创建时的 Hook 函数
	SetOnConnStop(func(IConnection))        // 设置当前 Server 的连接断开时的 Hook 函数
	GetOnConnStart() func(IConnection)      // 获取当前 Server 的连接创建时的 Hook 函数
	GetOnConnStop() func(IConnection)       // 获取当前 Server 的连接断开时的 Hook 函数
	SetPacket(IDataPack)                    // 设置当前 Server 绑定的数据协议封包方式
	GetPacket() IDataPack                   // 获取当前 Server 绑定的数据协议封包方式
	GetMsgHandler() IMsgHandle              // 获取当前 Server 绑定的消息处理模块
	SetHeartBeat(*HeartBeatOption)          // 设置心跳检测
}
