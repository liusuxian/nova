/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 14:59:29
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iserver.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "time"

// IServer 服务器接口
type IServer interface {
	Start()                                                   // 启动服务器
	Stop()                                                    // 停止服务器
	Server()                                                  // 开启业务服务
	AddRouter(msgID uint32, router IRouter)                   // 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
	GetConnMgr() IConnManager                                 // 得到连接管理
	SetOnConnStart(func(IConnection))                         // 设置当前 Server 的连接创建时Hook函数
	SetOnConnStop(func(IConnection))                          // 设置当前 Server 的连接断开时的Hook函数
	GetOnConnStart() func(IConnection)                        // 获取当前 Server 的连接创建时Hook函数
	GetOnConnStop() func(IConnection)                         // 获取当前 Server 的连接断开时的Hook函数
	SetPacket(IDataPack)                                      // 设置当前 Server 绑定的数据协议封包方式
	GetPacket() IDataPack                                     // 获取当前 Server 绑定的数据协议封包方式
	GetMsgHandler() IMsgHandle                                // 获取当前 Server 绑定的消息处理模块
	StartHeartBeat(time.Duration)                             // 启动心跳检测
	StartHeartBeatWithOption(time.Duration, *HeartBeatOption) // 启动心跳检测(自定义回调)
}
