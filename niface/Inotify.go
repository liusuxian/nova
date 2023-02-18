/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:40:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:48:11
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/Inotify.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// Inotify 通知接口
type Inotify interface {
	HasIdConn(id uint64) bool                                        // 是否有这个id
	ConnNums() int                                                   // 存储的map长度
	SetNotifyID(id uint64, conn IConnection)                         // 添加连接
	GetNotifyByID(id uint64) (IConnection, error)                    // 得到某个连接
	DelNotifyByID(id uint64)                                         // 删除某个连接
	NotifyToConnByID(id uint64, msgID uint32, data []byte) error     // 通知某个id
	NotifyAll(msgID uint32, data []byte) error                       // 通知所有人
	NotifyBuffToConnByID(id uint64, msgID uint32, data []byte) error // 通过缓冲队列通知某个id
	NotifyBuffAll(msgID uint32, data []byte) error                   // 缓冲队列通知所有人
}
