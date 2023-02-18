/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:24:32
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iserver.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IServer 服务器接口
type IServer interface {
	Start()                                 // 启动服务器
	Stop()                                  // 停止服务器
	Server()                                // 开启业务服务
	AddRouter(msgID uint32, router IRouter) // 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
	GetConnMgr() IConnManager               // 得到连接管理
	SetOnConnStart(func(IConnection))       // 设置该Server的连接创建时Hook函数
	SetOnConnStop(func(IConnection))        // 设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn IConnection)       // 调用连接OnConnStart Hook函数
	CallOnConnStop(conn IConnection)        // 调用连接OnConnStop Hook函数
	Packet() IDataPack
}
