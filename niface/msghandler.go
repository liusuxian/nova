/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:49:21
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-07 23:04:38
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMsgHandle 消息处理接口
type IMsgHandle interface {
	AddRouter(msgID uint16, handlers ...RouterHandler) (router IRouter)                // 为消息添加路由处理函数集合
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由组管理
	Use(handlers ...RouterHandler) (router IRouter)                                    // 添加全局路由
	PrintRouters()                                                                     // 打印所有路由
	StartWorkerPool()                                                                  // 启动 Worker 工作池
	StopWorkerPool()                                                                   // 停止 Worker 工作池
	SubmitToWorkerPool(request IRequest)                                               // 将请求提交给 Worker 工作池处理
	AddInterceptor(interceptor IInterceptor)                                           // 注册责任链任务入口，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
	Execute(request IRequest)                                                          // 执行责任链上的拦截器方法
}
