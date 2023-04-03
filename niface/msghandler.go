/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:49:21
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 21:27:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMsgHandle 消息处理接口
type IMsgHandle interface {
	AddRouter(msgID uint16, router IRouter)  // 为消息添加具体的处理逻辑
	PrintRouters()                           // 打印所有路由
	StartWorkerPool()                        // 启动 Worker 工作池
	StopWorkerPool()                         // 停止 Worker 工作池
	AddInterceptor(interceptor IInterceptor) // 添加拦截器
	Execute(request IRequest)                // 执行当前请求
}
