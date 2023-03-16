/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 00:49:32
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-16 19:44:59
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
	"strconv"
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
	key := strconv.FormatInt(int64(conn.GetConnID()), 10)
	cm.connMap.Set(key, conn)
	nlog.Debug(cm.ctx, "Connection Add To ConnManager Success", zap.Int("ConnID", conn.GetConnID()), zap.Int("ConnCount", cm.connMap.Count()))
}

// RemoveConn 删除连接
func (cm *ConnManager) RemoveConn(connID int) {
	key := strconv.FormatInt(int64(connID), 10)
	cm.connMap.Remove(key)
	nlog.Debug(cm.ctx, "Connection Remove From ConnManager Success", zap.Int("ConnID", connID), zap.Int("ConnCount", cm.connMap.Count()))
}

// GetConn 通过 ConnID 获取连接
func (cm *ConnManager) GetConn(connID int) (niface.IConnection, error) {
	key := strconv.FormatInt(int64(connID), 10)
	if conn, ok := cm.connMap.Get(key); ok {
		return conn, nil
	}
	return nil, errors.New("Connection Not Found")
}

// GetAllConn 获取当前所有连接
func (cm *ConnManager) GetAllConn() map[string]niface.IConnection {
	return cm.connMap.Items()
}
