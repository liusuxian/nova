/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 17:25:31
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-31 01:39:24
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// IOfflineMsg 离线消息接口
type IOfflineMsg interface {
	Save() (err error) // 保存离线消息
}

// INotify 通知接口
type INotify interface {
	HasID(id uint32) (has bool)                                                                      // 是否存在这个 ID
	ConnCount() (count int)                                                                          // 连接的数量
	AddNotifyID(id uint32, conn IConnection)                                                         // 添加连接
	RemoveNotifyID(id uint32)                                                                        // 删除连接
	GetNotifyConn(id uint32) (conn IConnection, isExist bool)                                        // 获取连接
	Notify(id uint32, msgID uint16, f MsgDataFunc, offlineMsg IOfflineMsg, callback ...SendCallback) // 通知连接
	NotifyAll(msgID uint16, f MsgDataFunc, callback ...SendCallback)                                 // 通知所有连接
}
