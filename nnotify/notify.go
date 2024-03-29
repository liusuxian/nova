/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-27 01:28:36
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-06-01 13:12:09
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nnotify

import (
	"context"
	"github.com/liusuxian/nova/nerr"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	cmap "github.com/orcaman/concurrent-map/v2"
	"strconv"
)

// Notify 通知管理结构
type Notify struct {
	connMap cmap.ConcurrentMap[string, niface.IConnection] // 存放 Connection 的并发安全 Map
}

// NewNotify 创建通知管理
func NewNotify() (n niface.INotify) {
	return &Notify{
		connMap: cmap.New[niface.IConnection](),
	}
}

// HasID 是否存在这个 ID
func (n *Notify) HasID(id uint32) (has bool) {
	key := strconv.FormatInt(int64(id), 10)
	return n.connMap.Has(key)
}

// ConnCount 连接的数量
func (n *Notify) ConnCount() (count int) {
	return n.connMap.Count()
}

// AddNotifyID 添加连接
func (n *Notify) AddNotifyID(id uint32, conn niface.IConnection) {
	key := strconv.FormatInt(int64(id), 10)
	n.connMap.Set(key, conn)
	nlog.Debug("Add To Notify", nlog.Uint32("ID", id), nlog.Int("ConnCount", n.ConnCount()))
}

// RemoveNotifyID 删除连接
func (n *Notify) RemoveNotifyID(id uint32) {
	key := strconv.FormatInt(int64(id), 10)
	n.connMap.Remove(key)
	nlog.Debug("Remove From Notify", nlog.Uint32("ID", id), nlog.Int("ConnCount", n.ConnCount()))
}

// GetNotifyConn 获取连接
func (n *Notify) GetNotifyConn(id uint32) (conn niface.IConnection, isExist bool) {
	key := strconv.FormatInt(int64(id), 10)
	if c, ok := n.connMap.Get(key); ok {
		if c == nil {
			return nil, false
		}
		return c, true
	}
	return nil, false
}

// Notify 通知连接
func (n *Notify) Notify(id uint32, msgID uint16, f niface.MsgDataFunc, callback ...niface.SendCallback) {
	// 获取连接
	conn, isExist := n.GetNotifyConn(id)
	if !isExist {
		return
	}
	// 将 Message 数据发送给远程的对端
	if err := conn.SendMsg(msgID, f, callback...); err != nil && err != nerr.ErrConnectionClosed {
		nlog.Error("Notify Error", nlog.Uint32("Id", id), nlog.Uint16("MsgID", msgID), nlog.Err(err))
		return
	}
}

// NotifyAll 通知所有连接
func (n *Notify) NotifyAll(msgID uint16, f niface.MsgDataFunc, callback ...niface.SendCallback) {
	items := n.connMap.Items()
	// 循环发送
	for _, conn := range items {
		if conn == nil {
			continue
		}
		// 将 Message 数据发送给远程的对端
		if err := conn.SendMsg(msgID, f, callback...); err != nil && err != nerr.ErrConnectionClosed {
			nlog.Error("NotifyAll Error", nlog.Uint16("MsgID", msgID), nlog.Err(err))
		}
	}
}

// NotifySaveOfflineMsg 通知连接，当连接不存在或处于关闭状态时保存离线消息
func (n *Notify) NotifySaveOfflineMsg(ctx context.Context, id uint32, msgID uint16, f niface.MsgDataFunc, offlineMsg niface.IOfflineMsg, callback ...niface.SendCallback) {
	// 获取连接
	conn, isExist := n.GetNotifyConn(id)
	if !isExist {
		if offlineMsg != nil {
			// 保存离线消息
			if err := offlineMsg.Save(ctx, id, msgID, f); err != nil {
				nlog.Error("Save Offline Msg Error", nlog.Err(err))
			}
		}
		return
	}
	// 将 Message 数据发送给远程的对端
	if err := conn.SendMsg(msgID, f, callback...); err != nil {
		if err == nerr.ErrConnectionClosed {
			if offlineMsg != nil {
				// 保存离线消息
				if err := offlineMsg.Save(ctx, id, msgID, f); err != nil {
					nlog.Error("Save Offline Msg Error", nlog.Err(err))
				}
			}
		} else {
			nlog.Error("NotifySaveOfflineMsg Error", nlog.Uint32("Id", id), nlog.Uint16("MsgID", msgID), nlog.Err(err))
		}
		return
	}
}
