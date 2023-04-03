/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 20:28:07
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 20:48:59
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/interceptor.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// 请求父类，定义空接口，用于扩展支持任意类型
type IcReq interface{}

// 回复父类，定义空接口，用于扩展支持任意类型
type IcResp interface{}

// 拦截器接口
type IInterceptor interface {
	Intercept(chain IChain) (resp IcResp) // 拦截
}

// 责任链接口
type IChain interface {
	Request() (req IcReq)            // 获取当前请求信息
	Proceed(req IcReq) (resp IcResp) // 继续执行当前请求
}

// IBuilder 责任链构造器接口
type IBuilder interface {
	Head(interceptor IInterceptor)           // 设置责任链头
	Tail(interceptor IInterceptor)           // 设置责任链尾
	AddInterceptor(interceptor IInterceptor) // 添加拦截器
	Execute(req IcReq) (resp IcResp)         // 执行当前请求
}
