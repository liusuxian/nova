/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:40:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 14:52:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/Inotify.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// INotify 通知接口
type INotify interface {
	HasIdConn(id uint64) bool                                    // 是否存在这个通知 ID
	ConnNums() uint32                                            // 通知连接的数量
	AddNotify(id uint64, conn IConnection)                       // 添加通知连接
	GetNotify(id uint64) (IConnection, error)                    // 获取通知连接
	DelNotify(id uint64)                                         // 删除通知连接
	NotifyToConn(id uint64, msgID uint32, data []byte) error     // 通知某个连接
	NotifyAll(msgID uint32, data []byte) error                   // 通知所有连接
	NotifyBuffToConn(id uint64, msgID uint32, data []byte) error // 通过缓冲队列通知某个连接
	NotifyBuffAll(msgID uint32, data []byte) error               // 通过缓冲队列通知所有连接
}
