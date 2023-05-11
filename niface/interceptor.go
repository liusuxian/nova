/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 20:28:07
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:15:30
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// IcReq 拦截器输入数据
type IcReq interface{}

// IcResp 拦截器输出数据
type IcResp interface{}

// 拦截器接口
type IInterceptor interface {
	Intercept(chain IChain) (resp IcResp) // 拦截器的拦截处理方法，由开发者定义
}

// 责任链接口
type IChain interface {
	Request() (req IcReq)                                      // 获取当前责任链中的请求数据(当前拦截器)
	GetIMessage() (msg IMessage)                               // 从当前责任链中获取 IMessage
	Proceed(req IcReq) (resp IcResp)                           // 进入并执行下一个拦截器，且将请求数据传递给下一个拦截器
	ProceedWithIMessage(msg IMessage, req IcReq) (resp IcResp) // 通过 IMessage 和解码后的数据进入下一个拦截器
}
