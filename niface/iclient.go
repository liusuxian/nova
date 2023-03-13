/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 13:38:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 23:26:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iclient.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "time"

// IClient 客户端接口
type IClient interface {
	Start()                                                   // 启动 Client
	Stop()                                                    // 停止 Client
	AddRouter(msgID uint32, router IRouter)                   // 给当前 Client 添加路由
	Conn() IConnection                                        // 当前 Client 的连接信息
	SetOnConnStart(func(IConnection))                         // 设置当前 Client 的连接创建时Hook函数
	SetOnConnStop(func(IConnection))                          // 设置当前 Client 的连接断开时的Hook函数
	GetOnConnStart() func(IConnection)                        // 获取当前 Client 的连接创建时Hook函数
	GetOnConnStop() func(IConnection)                         // 获取当前 Client 的连接断开时的Hook函数
	SetPacket(IDataPack)                                      // 设置当前 Client 绑定的数据协议封包方式
	GetPacket() IDataPack                                     // 获取当前 Client 绑定的数据协议封包方式
	GetMsgHandler() IMsgHandle                                // 获取当前 Client 绑定的消息处理模块
	StartHeartBeat(time.Duration)                             // 启动心跳检测
	StartHeartBeatWithOption(time.Duration, *HeartBeatOption) // 启动心跳检测(自定义回调)
	AddInterceptor(interceptor Interceptor)                   // 添加协议解析拦截器
	GetLengthField() LengthField                              // 获取长度字段解码器，用于解码二进制数据流中表示长度字段的工具
}
