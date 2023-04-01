/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-01 17:49:25
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "context"

// IServer 服务器接口
type IServer interface {
	Start()                                                 // 启动 Server
	Stop()                                                  // 停止 Server
	AddRouter(msgID uint16, router IRouter)                 // 给当前 Server 添加路由
	GetCtx() (ctx context.Context)                          // 获取当前 Server 的 Context
	GetConnManager() (connMgr IConnManager)                 // 获取当前 Server 的连接管理
	GetConnections() (nums int)                             // 获取当前 Server 的活跃连接数
	SetOnConnStart(f func(conn IConnection))                // 设置当前 Server 的连接创建时的 Hook 函数
	SetOnConnStop(f func(conn IConnection))                 // 设置当前 Server 的连接断开时的 Hook 函数
	GetOnConnStart() (f func(conn IConnection))             // 获取当前 Server 的连接创建时的 Hook 函数
	GetOnConnStop() (f func(conn IConnection))              // 获取当前 Server 的连接断开时的 Hook 函数
	SetPacket(packet IDataPack)                             // 设置当前 Server 绑定的数据协议封包和拆包方式
	GetPacket() (packet IDataPack)                          // 获取当前 Server 绑定的数据协议封包和拆包方式
	GetMsgHandler() (handler IMsgHandle)                    // 获取当前 Server 绑定的消息处理模块
	SetServerOverload(option ...*ServerOverloadOption)      // 设置当前 Server 的服务器人数超载检测器
	SetHeartBeat(initiate bool, option ...*HeartBeatOption) // 设置当前 Server 的心跳检测器
	GetHeartBeat() (checker IHeartBeatChecker)              // 获取当前 Server 的心跳检测器
}
