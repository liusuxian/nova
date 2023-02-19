/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 00:58:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 14:44:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iconnection.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import (
	"context"
	"net"
)

// IConnection 连接接口
type IConnection interface {
	Start()                                      // 启动连接
	Stop()                                       // 停止连接
	Context() context.Context                    // 返回ctx，用于用户自定义的go程获取连接退出状态
	GetConnection() net.Conn                     // 从当前连接获取原始的socket Conn
	GetConnID() uint32                           // 获取当前连接ID
	RemoteAddr() net.Addr                        // 获取远程客户端地址信息
	SendMsg(msgID uint32, data []byte) error     // 直接将数据发送给远程的客户端(无缓冲)
	SendBuffMsg(msgID uint32, data []byte) error // 直接将数据发送给远程的客户端(有缓冲)
	SetProperty(key string, value interface{})   // 设置连接属性
	GetProperty(key string) (interface{}, error) // 获取连接属性
	RemoveProperty(key string)                   // 移除连接属性
}
