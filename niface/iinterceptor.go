/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 23:28:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 10:45:26
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

// Interceptor 拦截器接口
type Interceptor interface {
	Intercept(Chain) Response // 拦截并处理
}

// Chain 责任链模式接口
type Chain interface {
	Request() Request         // 请求
	Proceed(Request) Response // 继续处理
}

// InterceptorBuilder 拦截器构建器接口
type InterceptorBuilder interface {
	AddInterceptor(interceptor Interceptor) // 添加拦截器，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
	Execute(request Request) Response       // 将消息丢到责任链，通过责任链里拦截器层层处理层层传递
}
