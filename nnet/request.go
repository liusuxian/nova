/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:23:33
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 20:41:52
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/request.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nnet

import (
	"github.com/liusuxian/nova/niface"
	"sync"
)

// Request 请求结构
type Request struct {
	conn     niface.IConnection // 已经和客户端建立好的连接
	msg      niface.IMessage    // 客户端请求的数据
	router   niface.IRouter     // 请求处理的函数
	step     niface.HandleStep  // 用来控制路由函数的执行
	stepLock *sync.RWMutex      // 并发互斥
	needNext bool               // 是否需要转进到下一个处理器开始执行
}

const (
	PRE_HANDLE  niface.HandleStep = iota // PreHandle 预处理
	HANDLE                               // Handle 处理
	POST_HANDLE                          // PostHandle 后处理
	HANDLE_OVER                          // HandleOver 处理完成
)

// NewRequest 创建请求
func NewRequest(conn niface.IConnection, msg niface.IMessage) (req *Request) {
	req = new(Request)
	req.conn = conn
	req.msg = msg
	req.step = PRE_HANDLE
	req.stepLock = new(sync.RWMutex)
	req.needNext = true
	return
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() niface.IConnection {
	return r.conn
}

// GetMsgID 获取请求的消息ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// BindRouter 绑定这次请求由哪个路由处理
func (r *Request) BindRouter(router niface.IRouter) {
	r.router = router
}

// Call 转进到下一个处理器开始执行，但是调用此方法的函数会根据先后顺序逆序执行
func (r *Request) Call() {
	if r.router == nil {
		return
	}
	for r.step < HANDLE_OVER {
		switch r.step {
		case PRE_HANDLE:
			r.router.PreHandle(r)
		case HANDLE:
			r.router.Handle(r)
		case POST_HANDLE:
			r.router.PostHandle(r)
		}
		r.next()
	}
}

// Abort 终止处理函数的运行，但调用此方法的函数会执行完毕
func (r *Request) Abort() {
	r.stepLock.Lock()
	r.step = HANDLE_OVER
	r.stepLock.Unlock()
}

// Goto 指定接下来的Handle去执行哪个Handler函数（慎用！！！，会导致循环调用）
func (r *Request) Goto(step niface.HandleStep) {
	r.stepLock.Lock()
	r.step = step
	r.needNext = false
	r.stepLock.Unlock()
}

// 是否需要转进到下一个处理器开始执行
func (r *Request) next() {
	if r.needNext == false {
		r.needNext = true
		return
	}
	r.stepLock.Lock()
	r.step++
	r.stepLock.Unlock()
}
