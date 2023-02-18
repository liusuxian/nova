/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:30:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:32:07
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/iconnmanager.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IConnManager 连接管理接口
type IConnManager interface {
	Add(conn IConnection)                   // 添加连接
	Remove(conn IConnection)                // 删除连接
	Get(connID uint32) (IConnection, error) // 利用ConnID获取连接
	Len() int                               // 获取当前连接
	ClearConn()                             // 删除并停止所有连接
}
