/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 13:38:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 19:42:28
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iclient.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IClient 客户端接口
type IClient interface {
	Start() // 启动 Client
	Stop()  // 停止 Client
	// AddRouter(msgID uint32, router IRouter)                   // 给当前 Client 添加路由
	// Conn() IConnection                                        // 当前 Client 的连接信息
	// SetOnConnStart(func(IConnection))                         // 设置当前 Client 的连接创建时Hook函数
	// SetOnConnStop(func(IConnection))                          // 设置当前 Client 的连接断开时的Hook函数
	// GetOnConnStart() func(IConnection)                        // 获取当前 Client 的连接创建时Hook函数
	// GetOnConnStop() func(IConnection)                         // 获取当前 Client 的连接断开时的Hook函数
	// SetPacket(IDataPack)                                      // 设置当前 Client 绑定的数据协议封包方式
	// GetPacket() IDataPack                                     // 获取当前 Client 绑定的数据协议封包方式
	// GetMsgHandler() IMsgHandle                                // 获取当前 Client 绑定的消息处理模块
	// StartHeartBeat(time.Duration)                             // 启动心跳检测
	// StartHeartBeatWithOption(time.Duration, *HeartBeatOption) // 启动心跳检测(自定义回调)
	// AddInterceptor(interceptor Interceptor)                   // 添加拦截器，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
	// GetLengthField() LengthField                              // 获取长度字段解码器，用于解码二进制数据流中表示长度字段的工具
	// SetLengthField(LengthField)                               // 设置长度字段解码器
}
