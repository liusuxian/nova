/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:27:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:27:13
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/irouter.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IRouter 路由接口
type IRouter interface {
	PreHandle(request IRequest)  // 在处理 Connection 业务之前的钩子
	Handle(request IRequest)     // 处理 Connection 业务
	PostHandle(request IRequest) // 处理 Connection 业务之后的钩子
}
