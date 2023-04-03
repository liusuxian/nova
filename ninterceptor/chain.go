/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:03:53
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 21:11:10
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/ninterceptor/chain.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package ninterceptor

import "github.com/liusuxian/nova/niface"

// Chain 责任链结构
type Chain struct {
	req          niface.IcReq          // 当前请求
	position     int                   // 当前执行到责任链的哪一个拦截器
	interceptors []niface.IInterceptor // 拦截器链
}

// NewChain 创建责任链
func NewChain(list []niface.IInterceptor, pos int, req niface.IcReq) (chain niface.IChain) {
	return &Chain{
		req:          req,
		position:     pos,
		interceptors: list,
	}
}

// Request 获取当前请求信息
func (c *Chain) Request() (req niface.IcReq) {
	return c.req
}

// Proceed 继续执行当前请求
func (c *Chain) Proceed(req niface.IcReq) (resp niface.IcResp) {
	if c.position < len(c.interceptors) {
		chain := NewChain(c.interceptors, c.position+1, req)
		interceptor := c.interceptors[c.position]
		resp = interceptor.Intercept(chain)
		return
	}
	return req
}
