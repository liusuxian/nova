/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 17:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:16:18
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// INotify 通知接口
type INotify interface {
	HasIdConn(id uint64) (has bool)                                    // 是否存在这个通知 ID
	ConnNums() (nums uint32)                                           // 通知连接的数量
	AddNotify(id uint64, conn IConnection)                             // 添加通知连接
	GetNotify(id uint64) (conn IConnection, err error)                 // 获取通知连接
	DelNotify(id uint64)                                               // 删除通知连接
	NotifyToConn(id uint64, msgID uint16, data []byte) (err error)     // 通知某个连接
	NotifyAll(msgID uint16, data []byte) (err error)                   // 通知所有连接
	NotifyBuffToConn(id uint64, msgID uint16, data []byte) (err error) // 通过缓冲队列通知某个连接
	NotifyBuffAll(msgID uint16, data []byte) (err error)               // 通过缓冲队列通知所有连接
}
