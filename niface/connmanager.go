/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:30:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 20:41:22
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/connmanager.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IConnManager 连接管理接口
type IConnManager interface {
	AddConn(conn IConnection)                   // 添加连接
	RemoveConn(conn IConnection)                // 删除连接
	GetConn(connID string) (IConnection, error) // 通过 ConnID 获取连接
	ClearAllConn()                              // 清除并停止当前所有连接
	GetAllConnID() []string                     // 获取当前所有 ConnID
}
