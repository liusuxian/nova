/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:27:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 19:09:38
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/router.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// RouterHandler 路由处理函数
type RouterHandler func(request IRequest)

// IRouter 路由接口
type IRouter interface {
	Use(handlers ...RouterHandler)                                                     // 添加全局路由
	AddHandler(msgID uint16, handlers ...RouterHandler)                                // 添加路由
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由组管理
	GetHandlers(msgID uint16) (handlers []RouterHandler, isExist bool)                 // 获取路由处理函数集合
	PrintRouters()                                                                     // 打印所有路由
}

// IGroupRouter 路由组接口
type IGroupRouter interface {
	Use(handlers ...RouterHandler)                      // 添加全局路由
	AddHandler(msgID uint16, handlers ...RouterHandler) // 添加路由
}
