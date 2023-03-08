/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:23:33
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 21:51:32
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/request.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package request

import (
	"context"
	"github.com/liusuxian/nova/niface"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// Request 请求结构
type Request struct {
	context  context.Context    // 请求的上下文信息
	ctxKeys  []string           // context keys
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

const (
	ctxKeyForRequest = "NovaRequestObject" // 请求的 Request 对象
	requestTraceID   = "RequestTraceID"    // 请求的 Trace ID
)

// NewRequest 创建请求
func NewRequest(conn niface.IConnection, msg niface.IMessage) (req *Request) {
	req = new(Request)
	req.conn = conn
	req.msg = msg
	req.step = PRE_HANDLE
	req.stepLock = new(sync.RWMutex)
	req.needNext = true
	req.SetCtxVal(requestTraceID, uuid.NewV4().String())
	return
}

// RequestFromCtx 从上下文中检索并返回 Request 对象
func RequestFromCtx(ctx context.Context) *Request {
	if v := ctx.Value(ctxKeyForRequest); v != nil {
		return v.(*Request)
	}
	return nil
}

// Context 是函数 GetCtx 的别名
func (r *Request) Context() context.Context {
	if r.context == nil {
		r.context = context.Background()
	}
	// 将 Request 对象注入到上下文中
	if RequestFromCtx(r.context) == nil {
		r.context = context.WithValue(r.context, ctxKeyForRequest, r)
	}
	return r.context
}

// SetCtx 设置请求的上下文信息
func (r *Request) SetCtx(ctx context.Context) {
	r.context = ctx
}

// GetCtx 获取请求的上下文信息
func (r *Request) GetCtx() context.Context {
	return r.Context()
}

// SetCtxVal 将键值对作为自定义参数设置到请求的上下文信息中
func (r *Request) SetCtxVal(key string, value any) {
	r.context = context.WithValue(r.Context(), key, value)
	for _, k := range r.ctxKeys {
		if k == key {
			return
		}
	}
	r.ctxKeys = append(r.ctxKeys, key)
}

// GetCtxVal 检索并返回给定键名的值，可选参数 def 指定如果请求的上下文信息中不存在给定的 key 时的默认值
func (r *Request) GetCtxVal(key string, def ...any) any {
	val := r.Context().Value(key)
	if val == nil && len(def) > 0 {
		val = def[0]
	}
	return val
}

// GetCtxKeys 获取所有的 context key
func (r *Request) GetCtxKeys() []string {
	return r.ctxKeys
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() niface.IConnection {
	return r.conn
}

// GetMsgID 获取请求的消息 ID
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

// Goto 指定接下来的 Handle 去执行哪个 Handler 函数（慎用！！！，会导致循环调用）
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
