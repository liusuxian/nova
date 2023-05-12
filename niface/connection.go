/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 00:52:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 13:08:40
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

import (
	"context"
	"github.com/panjf2000/gnet/v2"
	"net"
)

// IConnection 连接接口
type IConnection interface {
	Start()                                                                          // 启动连接
	Stop()                                                                           // 停止连接
	GetCancelCtx() (ctx context.Context)                                             // 返回 Cancel Context，用于用户自定义的 Goroutine 获取连接退出状态
	GetConnection() (conn gnet.Conn)                                                 // 从当前连接获取原始的 gnet.Conn
	GetConnID() (connID int)                                                         // 获取当前 ConnID
	RemoteAddr() (addr net.Addr)                                                     // 获取当前连接远程地址信息
	LocalAddr() (addr net.Addr)                                                      // 获取当前连接本地地址信息
	Send(f MsgDataFunc, callback ...gnet.AsyncCallback) (err error)                  // 将数据发送给远程的对端
	SendMsg(msgID uint16, f MsgDataFunc, callback ...gnet.AsyncCallback) (err error) // 将 Message 数据发送给远程的对端
	SetProperty(key string, value any)                                               // 设置当前连接属性
	GetProperty(key string) (value any, err error)                                   // 获取当前连接属性
	RemoveProperty(key string)                                                       // 移除当前连接属性
	IsAlive() (isAlive bool)                                                         // 判断当前连接是否存活
	UpdateActivity()                                                                 // 更新连接活动时间
	SetHeartBeat(checker IHeartBeatChecker)                                          // 设置心跳检测器
}

// MsgDataFunc 消息数据的处理方法
type MsgDataFunc func() (buf []byte, err error)
