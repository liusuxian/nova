/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 21:35:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-30 19:05:13
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/request.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "context"

// HandleStep 处理阶段
type HandleStep uint8

// IRequest 请求接口
type IRequest interface {
	SetCtx(ctx context.Context)                   // 设置请求的 Context
	GetCtx() (ctx context.Context)                // 获取请求的 Context
	SetCtxVal(key string, value any)              // 将键值对作为自定义参数设置到请求的 Context 中
	GetCtxVal(key string, def ...any) (value any) // 检索并返回给定键名的值，可选参数 def 指定如果请求的 Context 中不存在给定的 key 时的默认值
	GetConnection() (conn IConnection)            // 获取请求连接信息
	GetMsgID() (msgID uint16)                     // 获取请求的消息 ID
	GetData() (data []byte)                       // 获取请求消息的数据
	GetMessage() (msg IMessage)                   // 获取请求消息的原始数据
	BindRouter(router IRouter)                    // 绑定这次请求由哪个路由处理
	Call()                                        // 转进到下一个处理器开始执行，但是调用此方法的函数会根据先后顺序逆序执行
	Abort()                                       // 终止处理函数的运行，但调用此方法的函数会执行完毕
	Goto(step HandleStep)                         // 指定接下来的 Handle 去执行哪个 Handler 函数（慎用！！！，会导致循环调用）
}
