/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 13:38:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-07 22:43:21
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/client.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IClient 客户端接口
type IClient interface {
	Start()                                                                            // 启动 Client
	Stop()                                                                             // 停止 Client
	AddRouter(msgID uint16, handlers ...RouterHandler) (router IRouter)                // 添加业务处理器集合
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由分组管理，并且会返回一个组管理器
	Use(handlers ...RouterHandler) (router IRouter)                                    // 添加全局组件
	Conn() (conn IConnection)                                                          // 当前 Client 的连接信息
	SetOnConnStart(f func(conn IConnection))                                           // 设置当前 Client 的连接创建时的 Hook 函数
	SetOnConnStop(f func(conn IConnection))                                            // 设置当前 Client 的连接断开时的 Hook 函数
	GetOnConnStart() (f func(conn IConnection))                                        // 获取当前 Client 的连接创建时的 Hook 函数
	GetOnConnStop() (f func(conn IConnection))                                         // 获取当前 Client 的连接断开时的 Hook 函数
	SetPacket(packet IDataPack)                                                        // 设置当前 Client 绑定的数据协议封包和拆包方式
	GetPacket() (packet IDataPack)                                                     // 获取当前 Client 绑定的数据协议封包和拆包方式
	GetMsgHandler() (handler IMsgHandle)                                               // 获取当前 Client 绑定的消息处理模块
	SetServerOverload(option ...*ServerOverloadOption)                                 // 设置当前 Client 的服务器人数超载检测器
	SetHeartBeat(option ...*HeartBeatOption)                                           // 设置当前 Client 的心跳检测器
	GetHeartBeat() (checker IHeartBeatChecker)                                         // 获取当前 Client 的心跳检测器
	AddInterceptor(interceptor IInterceptor)                                           // 添加拦截器
}
