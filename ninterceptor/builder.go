/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 20:45:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 18:01:07
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/ninterceptor/builder.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package ninterceptor

import "github.com/liusuxian/nova/niface"

// chainBuilder 责任链构造器结构
type chainBuilder struct {
	head niface.IInterceptor   // 责任链头
	tail niface.IInterceptor   // 责任链尾
	body []niface.IInterceptor // 责任链主体
}

// newChainBuilder 创建责任链构造器实例
func newChainBuilder() (builder *chainBuilder) {
	return &chainBuilder{
		body: make([]niface.IInterceptor, 0),
	}
}

// Head 将拦截器添加到责任链的开头
func (ic *chainBuilder) Head(interceptor niface.IInterceptor) {
	ic.head = interceptor
}

// Tail 将拦截器添加到责任链的尾部
func (ic *chainBuilder) Tail(interceptor niface.IInterceptor) {
	ic.tail = interceptor
}

// AddInterceptor 将拦截器添加到责任链的主体中
func (ic *chainBuilder) AddInterceptor(interceptor niface.IInterceptor) {
	ic.body = append(ic.body, interceptor)
}

// Execute 按顺序执行当前责任链中的所有拦截器
func (ic *chainBuilder) Execute(req niface.IcReq) (resp niface.IcResp) {
	// 将所有拦截器放入责任链构造器中
	var interceptors []niface.IInterceptor
	if ic.head != nil {
		interceptors = append(interceptors, ic.head)
	}
	if len(ic.body) > 0 {
		interceptors = append(interceptors, ic.body...)
	}
	if ic.tail != nil {
		interceptors = append(interceptors, ic.tail)
	}
	// 创建一个新的拦截器责任链，并执行每个拦截器
	chain := NewChain(interceptors, 0, req)
	// 进入并执行下一个拦截器，且将请求数据传递给下一个拦截器
	return chain.Proceed(req)
}
