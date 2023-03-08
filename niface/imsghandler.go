/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:49:21
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 14:33:10
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/imsghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IMsgHandle 消息处理接口
type IMsgHandle interface {
	DoMsgHandler(request IRequest)          // 马上以非阻塞方式处理消息
	AddRouter(msgID uint32, router IRouter) // 为消息添加具体的处理逻辑
	StartWorkerPool()                       // 启动worker工作池
	SendMsgToTaskQueue(request IRequest)    // 将消息交给TaskQueue，由worker进行处理
}
