/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-09 19:59:56
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IServer 服务器接口
type IServer interface {
	Start()                                                                            // 启动 Server
	Stop()                                                                             // 停止 Server
	AddRouter(msgID uint16, handlers ...RouterHandler) (router IRouter)                // 添加业务处理器集合
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由分组管理，并且会返回一个组管理器
	Use(handlers ...RouterHandler) (router IRouter)                                    // 添加全局组件
	GetConnManager() (connMgr IConnManager)                                            // 获取当前 Server 的连接管理
	GetConnections() (nums int)                                                        // 获取当前 Server 的活跃连接数
	SetOnConnStart(f func(conn IConnection))                                           // 设置当前 Server 的连接创建时的 Hook 函数
	SetOnConnStop(f func(conn IConnection))                                            // 设置当前 Server 的连接断开时的 Hook 函数
	GetOnConnStart() (f func(conn IConnection))                                        // 获取当前 Server 的连接创建时的 Hook 函数
	GetOnConnStop() (f func(conn IConnection))                                         // 获取当前 Server 的连接断开时的 Hook 函数
	SetOnStart(f func(s IServer))                                                      // 设置当前 Server 启动时的 Hook 函数
	SetOnStop(f func(s IServer))                                                       // 设置当前 Server 停止时的 Hook 函数
	SetPacket(packet IDataPack)                                                        // 设置当前 Server 绑定的数据协议封包和拆包方式
	GetPacket() (packet IDataPack)                                                     // 获取当前 Server 绑定的数据协议封包和拆包方式
	GetMsgHandler() (handler IMsgHandle)                                               // 获取当前 Server 绑定的消息处理模块
	SetServerOverload(option ...*ServerOverloadOption)                                 // 设置当前 Server 的服务器人数超载检测器
	GetServerOverload() (checker IServerOverloadChecker)                               // 获取当前 Server 的服务器人数超载检测器
	SetHeartBeat(initiate bool, option ...*HeartBeatOption)                            // 设置当前 Server 的心跳检测器
	GetHeartBeat() (checker IHeartBeatChecker)                                         // 获取当前 Server 的心跳检测器
	AddInterceptor(interceptor IInterceptor)                                           // 添加拦截器
}
