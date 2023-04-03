/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:35:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 21:42:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/interceptor/interceptor.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package interceptor

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
)

// Interceptor 自定义拦截器
type Interceptor struct {
}

// Intercept 拦截
func (ic *Interceptor) Intercept(chain niface.IChain) (resp niface.IcResp) {
	request := chain.Request()
	if request != nil {
		switch iRequest := request.(type) {
		case niface.IRequest:
			nlog.Debug("Interceptor Receive Data", nlog.Binary("Data", iRequest.GetData()))
		}
	}
	return chain.Proceed(chain.Request())
}
