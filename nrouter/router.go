/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 15:22:54
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 17:19:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nrouter/router.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nrouter

import (
	"github.com/liusuxian/nova/niface"
	"strconv"
	"sync"
)

// Router 路由结构
type Router struct {
	apis         map[uint16][]niface.RouterHandler // 存放每个 MsgID 所对应的路由处理函数集合
	handlers     []niface.RouterHandler            // 存放全局路由处理函数集合
	sync.RWMutex                                   // 并发读写锁
}

// GroupRouter 路由组结构
type GroupRouter struct {
	startMsgID uint16                 // 起始 MsgID
	endMsgID   uint16                 // 终止 MsgID
	handlers   []niface.RouterHandler // 存放全局路由处理函数集合
	router     niface.IRouter         // 路由
}

// NewRouter 创建路由
func NewRouter() (r *Router) {
	return &Router{
		apis:     make(map[uint16][]niface.RouterHandler, 10),
		handlers: make([]niface.RouterHandler, 0, 6),
	}
}

// NewGroupRouter 创建路由组
func NewGroupRouter(startMsgID, endMsgID uint16, router *Router, handlers ...niface.RouterHandler) (group *GroupRouter) {
	group = &GroupRouter{
		startMsgID: startMsgID,
		endMsgID:   endMsgID,
		handlers:   make([]niface.RouterHandler, 0, len(handlers)),
		router:     router,
	}
	group.handlers = append(group.handlers, handlers...)
	return
}

// Use 添加全局路由
func (r *Router) Use(handlers ...niface.RouterHandler) {
	r.handlers = append(r.handlers, handlers...)
}

// AddHandler 添加路由
func (r *Router) AddHandler(msgID uint16, handlers ...niface.RouterHandler) {
	if _, ok := r.apis[msgID]; ok {
		panic("Repeated Api MsgID = " + strconv.Itoa(int(msgID)))
	}

	length := len(r.handlers) + len(handlers)
	newHandlers := make([]niface.RouterHandler, length)
	copy(newHandlers, r.handlers)
	copy(newHandlers[len(r.handlers):], handlers)
	r.apis[msgID] = append(r.apis[msgID], newHandlers...)
}

// Group 路由组管理
func (r *Router) Group(startMsgID uint16, endMsgID uint16, handlers ...niface.RouterHandler) (group niface.IGroupRouter) {
	return NewGroupRouter(startMsgID, endMsgID, r, handlers...)
}

// GetHandlers 获取路由处理函数集合
func (r *Router) GetHandlers(msgID uint16) (handlers []niface.RouterHandler, isExist bool) {
	r.RLock()
	defer r.RUnlock()
	handlers, isExist = r.apis[msgID]
	return
}

// Use 添加全局路由
func (gr *GroupRouter) Use(handlers ...niface.RouterHandler) {
	gr.handlers = append(gr.handlers, handlers...)
}

// AddHandler 添加路由
func (gr *GroupRouter) AddHandler(msgID uint16, handlers ...niface.RouterHandler) {
	if msgID < gr.startMsgID || msgID > gr.endMsgID {
		panic("Add Router To Group Error In MsgID: " + strconv.Itoa(int(msgID)))
	}

	length := len(gr.handlers) + len(handlers)
	newHandlers := make([]niface.RouterHandler, length)
	copy(newHandlers, gr.handlers)
	copy(newHandlers[len(gr.handlers):], handlers)
	gr.router.AddHandler(msgID, newHandlers...)
}
