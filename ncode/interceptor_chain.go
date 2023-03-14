/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-14 10:57:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 11:33:50
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/ncode/interceptor_chain.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package ncode

import "github.com/liusuxian/nova/niface"

// InterceptorChain 拦截器链结构
//
//	HTLV+CRC，H头码(1字节)，T功能码(1字节)，L数据长度(1字节)，V数据内容(N字节)，CRC校验(2字节)
type InterceptorChain struct {
	interceptors []niface.Interceptor
	request      niface.Request
}

// NewInterceptorBuilder 创建拦截器构建器
func NewInterceptorBuilder() niface.InterceptorBuilder {
	return &InterceptorChain{
		interceptors: make([]niface.Interceptor, 0),
	}
}

// AddInterceptor 添加拦截器，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
func (ic *InterceptorChain) AddInterceptor(interceptor niface.Interceptor) {
	ic.interceptors = append(ic.interceptors, interceptor)
}

// Execute 将消息丢到责任链，通过责任链里拦截器层层处理层层传递
func (ic *InterceptorChain) Execute(request niface.Request) niface.Response {
	ic.request = request
	chain := NewRealInterceptorChain(ic.interceptors, 0, request)
	return chain.Proceed(ic.request)
}
