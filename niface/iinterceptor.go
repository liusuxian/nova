/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 23:28:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 15:16:26
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iinterceptor.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// 请求接口
//
//	定义空接口，用于扩展支持任意类型
type Request interface {
}

// 回复接口
//
//	定义空接口，用于扩展支持任意类型
type Response interface {
}

// Interceptor
type Interceptor interface {
	Intercept(Chain) Response
}

// Chain
type Chain interface {
	Request() Request
	Proceed(Request) Response
}

// InterceptorBuilder
type InterceptorBuilder interface {
	AddInterceptor(interceptor Interceptor)
	Execute(request Request) Response
}
