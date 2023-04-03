/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 20:45:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 21:09:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/ninterceptor/builder.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package ninterceptor

import "github.com/liusuxian/nova/niface"

// Builder 责任链构造器结构
type Builder struct {
	head niface.IInterceptor   // 责任链头
	tail niface.IInterceptor   // 责任链尾
	list []niface.IInterceptor // 拦截器链
	req  niface.IcReq          // 当前请求
}

// NewBuilder 创建责任链构造器
func NewBuilder() (builder niface.IBuilder) {
	return &Builder{
		list: make([]niface.IInterceptor, 0),
	}
}

// Head 设置责任链头
func (ic *Builder) Head(interceptor niface.IInterceptor) {
	ic.head = interceptor
}

// Tail 设置责任链尾
func (ic *Builder) Tail(interceptor niface.IInterceptor) {
	ic.tail = interceptor
}

// AddInterceptor 添加拦截器
func (ic *Builder) AddInterceptor(interceptor niface.IInterceptor) {
	ic.list = append(ic.list, interceptor)
}

// Execute 执行当前请求
func (ic *Builder) Execute(req niface.IcReq) (resp niface.IcResp) {
	ic.req = req
	// 将全部拦截器放入 Builder 中
	var interceptors []niface.IInterceptor
	if ic.head != nil {
		interceptors = append(interceptors, ic.head)
	}
	if len(ic.list) > 0 {
		interceptors = append(interceptors, ic.list...)
	}
	if ic.tail != nil {
		interceptors = append(interceptors, ic.tail)
	}
	// 创建一个拦截器责任链，执行每一个拦截器
	chain := NewChain(interceptors, 0, req)
	// 进入责任链执行
	return chain.Proceed(ic.req)
}
