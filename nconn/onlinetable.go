/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-26 17:48:15
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-26 18:17:33
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconn

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	cmap "github.com/orcaman/concurrent-map/v2"
)

// OnlineTable 在线表结构
type OnlineTable struct {
	connMap cmap.ConcurrentMap[string, niface.IConnection] // 在线表，存放 Connection 的并发安全 Map
}

// NewOnlineTable 创建一个在线表
func NewOnlineTable() (ot *OnlineTable) {
	return &OnlineTable{
		connMap: cmap.New[niface.IConnection](),
	}
}

// Add 添加记录
func (ot *OnlineTable) Add(key string, conn niface.IConnection) {
	ot.connMap.Set(key, conn)
	nlog.Debug("Add To OnlineTable", nlog.String("Key", key), nlog.Int("OnlineCount", ot.OnlineCount()))
}

// Remove 删除记录
func (ot *OnlineTable) Remove(key string) {
	ot.connMap.Remove(key)
	nlog.Debug("Remove From OnlineTable", nlog.String("Key", key), nlog.Int("OnlineCount", ot.OnlineCount()))
}

// Get 获取记录
func (ot *OnlineTable) Get(key string) (conn niface.IConnection, isOnline bool) {
	if c, ok := ot.connMap.Get(key); ok {
		return c, true
	}

	return nil, false
}

// GetAll 获取当前所有记录
func (ot *OnlineTable) GetAll() (connMap map[string]niface.IConnection) {
	return ot.connMap.Items()
}

// OnlineCount 在线数量
func (ot *OnlineTable) OnlineCount() (count int) {
	return ot.connMap.Count()
}
