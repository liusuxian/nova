/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 00:49:32
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 21:27:19
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconn/connmanager.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconn

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ConnManager 连接管理结构
type ConnManager struct {
	ctx     context.Context                                // 当前 Server 的根 Context
	connMap cmap.ConcurrentMap[string, niface.IConnection] // 存放 Connection 的并发安全 Map
}

// NewConnManager 创建一个连接管理
func NewConnManager(ctx context.Context) *ConnManager {
	return &ConnManager{
		ctx:     ctx,
		connMap: cmap.New[niface.IConnection](),
	}
}

// AddConn 添加连接
func (cm *ConnManager) AddConn(conn niface.IConnection) {
	cm.connMap.Set(conn.GetConnID(), conn)
	nlog.Debug(cm.ctx, "Connection Add To ConnManager Success", zap.Int("Count", cm.connMap.Count()))
}

// RemoveConn 删除连接
func (cm *ConnManager) RemoveConn(conn niface.IConnection) {
	cm.connMap.Remove(conn.GetConnID())
	nlog.Debug(cm.ctx, "Connection Remove From ConnManager Success", zap.String("ConnID", conn.GetConnID()), zap.Int("Count", cm.connMap.Count()))
}

// GetConn 通过 ConnID 获取连接
func (cm *ConnManager) GetConn(connID string) (niface.IConnection, error) {
	if conn, ok := cm.connMap.Get(connID); ok {
		return conn, nil
	}
	return nil, errors.New("Connection Not Found")
}

// ClearAllConn 清除并停止当前所有连接
func (cm *ConnManager) ClearAllConn() {
	for item := range cm.connMap.IterBuffered() {
		item.Val.Stop()             // 停止当前连接
		cm.connMap.Remove(item.Key) // 清除当前连接
	}
	nlog.Debug(cm.ctx, "Clear All Connection From ConnManager Success", zap.Int("Count", cm.connMap.Count()))
}

// GetAllConnID 获取当前所有 ConnID
func (cm *ConnManager) GetAllConnID() []string {
	ids := make([]string, 0, cm.connMap.Count())
	for item := range cm.connMap.IterBuffered() {
		ids = append(ids, item.Val.GetConnID())
	}
	return ids
}
