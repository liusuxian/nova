/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 17:05:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-27 01:51:06
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// IConnManager 连接管理接口
type IConnManager interface {
	AddConn(conn IConnection)                            // 添加连接
	RemoveConn(connID int)                               // 删除连接
	GetConn(connID int) (conn IConnection, isExist bool) // 通过 ConnID 获取连接
	GetAllConn() (connMap map[string]IConnection)        // 获取当前所有连接
}
