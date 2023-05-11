/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-07 14:57:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:16:48
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// RouterHandler 路由处理函数
type RouterHandler func(request IRequest)

// IRouter 路由接口
type IRouter interface {
	Use(handlers ...RouterHandler)                                                     // 添加全局组件
	AddHandler(msgID uint16, handlers ...RouterHandler)                                // 添加业务处理器集合
	Group(startMsgID, endMsgID uint16, handlers ...RouterHandler) (group IGroupRouter) // 路由分组管理，并且会返回一个组管理器
	GetHandlers(msgID uint16) (handlers []RouterHandler, isExist bool)                 // 获取 MsgID 注册的处理器集合
	PrintRouters()                                                                     // 打印所有路由
}

// IGroupRouter 路由组接口
type IGroupRouter interface {
	Use(handlers ...RouterHandler)                      // 添加全局组件
	AddHandler(msgID uint16, handlers ...RouterHandler) // 添加业务处理器集合
}
