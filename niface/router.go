/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:27:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-07 14:58:41
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
	Use(handlers ...RouterHandler)                                                     // 添加全局组件
	AddHandler(msgID uint16, handlers ...RouterHandler)                                // 添加业务处理器集合
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由分组管理，并且会返回一个组管理器
	GetHandlers(msgID uint16) (handlers []RouterHandler, isExist bool)                 // 获取当前所有注册在 MsgID 的处理器集合
	PrintRouters()                                                                     // 打印所有路由
}

// IGroupRouter 路由组接口
type IGroupRouter interface {
	Use(handlers ...RouterHandler)                      // 添加全局组件
	AddHandler(msgID uint16, handlers ...RouterHandler) // 添加业务处理器集合
}
