/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 13:41:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-03 16:42:39
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconn/connmanager.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconn

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"strconv"
)

// ConnManager 连接管理结构
type ConnManager struct {
	connMap cmap.ConcurrentMap[string, niface.IConnection] // 存放 Connection 的并发安全 Map
}

// NewConnManager 创建一个连接管理
func NewConnManager() (connMgr *ConnManager) {
	return &ConnManager{
		connMap: cmap.New[niface.IConnection](),
	}
}

// AddConn 添加连接
func (cm *ConnManager) AddConn(conn niface.IConnection) {
	key := strconv.FormatInt(int64(conn.GetConnID()), 10)
	cm.connMap.Set(key, conn)
	nlog.Debug("Connection Add To ConnManager", nlog.Int("ConnID", conn.GetConnID()), nlog.Int("ConnCount", cm.connMap.Count()))
}

// RemoveConn 删除连接
func (cm *ConnManager) RemoveConn(connID int) {
	key := strconv.FormatInt(int64(connID), 10)
	cm.connMap.Remove(key)
	nlog.Debug("Connection Remove From ConnManager", nlog.Int("ConnID", connID), nlog.Int("ConnCount", cm.connMap.Count()))
}

// GetConn 通过 ConnID 获取连接
func (cm *ConnManager) GetConn(connID int) (conn niface.IConnection, err error) {
	key := strconv.FormatInt(int64(connID), 10)
	if c, ok := cm.connMap.Get(key); ok {
		return c, nil
	}
	return nil, errors.New("Connection Not Found")
}

// GetAllConn 获取当前所有连接
func (cm *ConnManager) GetAllConn() (connMap map[string]niface.IConnection) {
	return cm.connMap.Items()
}
