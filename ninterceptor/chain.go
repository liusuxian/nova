/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:03:53
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:17:28
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package ninterceptor

import "github.com/liusuxian/nova/niface"

// Chain 责任链结构
type Chain struct {
	req          niface.IcReq          // 当前请求
	position     int                   // 当前执行到责任链的哪一个拦截器
	interceptors []niface.IInterceptor // 拦截器链
}

// NewChain 创建拦截器责任链
func NewChain(list []niface.IInterceptor, pos int, req niface.IcReq) (chain niface.IChain) {
	return &Chain{
		req:          req,
		position:     pos,
		interceptors: list,
	}
}

// Request 获取当前责任链中的请求数据(当前拦截器)
func (c *Chain) Request() (req niface.IcReq) {
	return c.req
}

// GetIMessage 从当前责任链中获取 IMessage
func (c *Chain) GetIMessage() (msg niface.IMessage) {
	req := c.Request()
	if req == nil {
		return nil
	}

	iRequest := c.ShouldIRequest(req)
	if iRequest == nil {
		return nil
	}

	return iRequest.GetMessage()
}

// Proceed 进入并执行下一个拦截器，且将请求数据传递给下一个拦截器
func (c *Chain) Proceed(req niface.IcReq) (resp niface.IcResp) {
	if c.position < len(c.interceptors) {
		chain := NewChain(c.interceptors, c.position+1, req)
		interceptor := c.interceptors[c.position]
		resp = interceptor.Intercept(chain)
		return
	}
	return req
}

// ProceedWithIMessage 通过 IMessage 和解码后的数据进入下一个拦截器
// msg 为解码后的 IMessage
// response 为解码后的数据
func (c *Chain) ProceedWithIMessage(msg niface.IMessage, response niface.IcReq) (resp niface.IcResp) {
	if msg == nil || response == nil {
		return c.Proceed(c.Request())
	}

	req := c.Request()
	if req == nil {
		return c.Proceed(c.Request())
	}

	iRequest := c.ShouldIRequest(req)
	if iRequest == nil {
		return c.Proceed(c.Request())
	}

	// 设置解析完后的序列化数据
	iRequest.SetResponse(response)

	return c.Proceed(iRequest)
}

// ShouldIRequest 判断是否是 IRequest
func (c *Chain) ShouldIRequest(icReq niface.IcReq) (request niface.IRequest) {
	if icReq == nil {
		return nil
	}

	switch iRequest := icReq.(type) {
	case niface.IRequest:
		return iRequest
	default:
		return nil
	}
}
