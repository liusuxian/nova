/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 00:49:32
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 13:38:25
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/connmanager/connmanager.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package connmanager

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
)

// ConnManager 连接管理结构
type ConnManager struct {
	connMap cmap.ConcurrentMap[string, niface.IConnection] // 存放 connection 的并发安全 map
}

// NewConnManager 创建一个连接管理
func NewConnManager() *ConnManager {
	return &ConnManager{connMap: cmap.New[niface.IConnection]()}
}

// AddConn 添加连接
func (cm *ConnManager) AddConn(conn niface.IConnection) {
	key := strconv.FormatUint(conn.GetConnID(), 10)
	cm.connMap.Set(key, conn)
	nlog.Debug(nil, "Connection Add To ConnManager Success", zap.Int("connected", cm.connMap.Count()))
}

// RemoveConn 删除连接
func (cm *ConnManager) RemoveConn(conn niface.IConnection) {
	key := strconv.FormatUint(conn.GetConnID(), 10)
	cm.connMap.Remove(key)
	nlog.Debug(nil, "Connection Remove From ConnManager Success", zap.Uint64("connID", conn.GetConnID()), zap.Int("connected", cm.connMap.Count()))
}

// GetConn 通过 ConnID 获取连接
func (cm *ConnManager) GetConn(connID uint64) (niface.IConnection, error) {
	key := strconv.FormatUint(connID, 10)
	if conn, ok := cm.connMap.Get(key); ok {
		return conn, nil
	}
	return nil, errors.New("Connection Not Found")
}

// GetConnected 获取当前连接的数量
func (cm *ConnManager) GetConnected() uint32 {
	return uint32(cm.connMap.Count())
}

// ClearAllConn 清除并停止当前所有连接
func (cm *ConnManager) ClearAllConn() {
	for item := range cm.connMap.IterBuffered() {
		item.Val.Stop()             // 停止当前连接
		cm.connMap.Remove(item.Key) // 清除当前连接
	}
	nlog.Debug(nil, "Clear All Connection From ConnManager Success", zap.Int("connected", cm.connMap.Count()))
}

// GetAllConnID 获取当前所有 ConnID
func (cm *ConnManager) GetAllConnID() []uint64 {
	ids := make([]uint64, 0, cm.connMap.Count())
	for item := range cm.connMap.IterBuffered() {
		ids = append(ids, item.Val.GetConnID())
	}
	return ids
}
