/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:28:27
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:48:33
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/irequest.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IRequest 请求接口
type IRequest interface {
	GetConnection() IConnection // 获取请求连接信息
	GetData() []byte            // 获取请求消息的数据
	GetMsgID() uint32           // 获取请求的消息ID
	BindRouter(router IRouter)  // 绑定这次请求由哪个路由处理
	Next()                      // 转进到下一个处理器开始执行 但是调用此方法的函数会根据先后顺序逆序执行
	Abort()                     // 终止处理函数的运行 但调用此方法的函数会执行完毕
}
