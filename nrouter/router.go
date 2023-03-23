/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 15:22:54
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 21:27:50
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nrouter/router.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nrouter

import "github.com/liusuxian/nova/niface"

// BaseRouter 基础路由结构
// 实现 Router 时，先嵌入这个结构，然后根据需要对这个结构的方法进行重写
type BaseRouter struct {
}

// 之所以 BaseRouter 的方法都为空，是因为有的 Router 不希望有 PreHandle 或 PostHandle
// 所以 Router 全部复合 BaseRouter 的好处是，不需要实现 PreHandle 和 PostHandle 也可以实例化

// PreHandle 在处理 Connection 业务之前的钩子
func (br *BaseRouter) PreHandle(request niface.IRequest) {

}

// Handle 处理 Connection 业务
func (br *BaseRouter) Handle(request niface.IRequest) {

}

// PostHandle 处理 Connection 业务之后的钩子
func (br *BaseRouter) PostHandle(request niface.IRequest) {

}
