/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 01:01:50
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:12:35
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nheartbeat

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"time"
)

// HeartBeatChecker 心跳检测器结构
type HeartBeatChecker struct {
	interval         time.Duration           // 心跳检测时间间隔
	quitChan         chan bool               // 退出信号
	makeMsg          niface.HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	onRemoteNotAlive niface.OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	msgID            uint16                  // 用户自定义的心跳检测消息ID
	conn             niface.IConnection      // 绑定的连接
	initiate         bool                    // 发起心跳
}

// NewHeartBeatChecker 创建心跳检测器
func NewHeartBeatChecker(interval time.Duration, initiate bool) (checker niface.IHeartBeatChecker) {
	heartbeat := &HeartBeatChecker{
		interval:         interval,
		quitChan:         make(chan bool),
		makeMsg:          makeMsgDefaultFunc,
		onRemoteNotAlive: onRemoteNotAliveDefaultFunc,
		msgID:            niface.HeartBeatDefaultMsgID,
		conn:             nil,
		initiate:         initiate,
	}
	return heartbeat
}

// Start 启动心跳检测
func (hbc *HeartBeatChecker) Start() {
	go hbc.start()
}

// Stop 停止心跳检测
func (hbc *HeartBeatChecker) Stop() {
	hbc.quitChan <- true
}

// SetHeartBeatMsgFunc 设置心跳检测消息处理方法
func (hbc *HeartBeatChecker) SetHeartBeatMsgFunc(f niface.HeartBeatMsgFunc) {
	if f != nil {
		hbc.makeMsg = f
	}
}

// SetOnRemoteNotAlive 设置远程连接不存活时的处理方法
func (hbc *HeartBeatChecker) SetOnRemoteNotAlive(f niface.OnRemoteNotAlive) {
	if f != nil {
		hbc.onRemoteNotAlive = f
	}
}

// BindConn 绑定连接
func (hbc *HeartBeatChecker) BindConn(conn niface.IConnection) {
	hbc.conn = conn
	// 设置心跳检测器
	conn.SetHeartBeat(hbc)
}

// Clone 克隆心跳检测器
func (hbc *HeartBeatChecker) Clone() (checker niface.IHeartBeatChecker) {
	heartbeat := &HeartBeatChecker{
		interval:         hbc.interval,
		quitChan:         make(chan bool),
		makeMsg:          hbc.makeMsg,
		onRemoteNotAlive: hbc.onRemoteNotAlive,
		msgID:            hbc.msgID,
		conn:             nil,
		initiate:         hbc.initiate,
	}
	return heartbeat
}

// SetMsgID 设置心跳检测消息ID
func (hbc *HeartBeatChecker) SetMsgID(msgID uint16) {
	if msgID != niface.HeartBeatDefaultMsgID {
		hbc.msgID = msgID
	}
}

// GetMsgID 获取心跳检测消息ID
func (hbc *HeartBeatChecker) GetMsgID() (msgID uint16) {
	return hbc.msgID
}

// GetMessage 获取心跳检测消息
func (hbc *HeartBeatChecker) GetMessage() (msg niface.IMessage) {
	return npack.NewMsgPackage(hbc.msgID, hbc.makeMsg())
}

// start 启动心跳检测
func (hbc *HeartBeatChecker) start() {
	ticker := time.NewTicker(hbc.interval)
	for {
		select {
		case <-ticker.C:
			hbc.check()
		case <-hbc.quitChan:
			ticker.Stop()
			return
		}
	}
}

// check 执行心跳检测
func (hbc *HeartBeatChecker) check() {
	if hbc.conn == nil {
		return
	}
	if !hbc.conn.IsAlive() {
		hbc.onRemoteNotAlive(hbc.conn)
	} else {
		hbc.sendHeartBeatMsg(hbc.conn)
	}
}

// sendHeartBeatMsg 发送心跳消息
func (hbc *HeartBeatChecker) sendHeartBeatMsg(conn niface.IConnection) {
	if hbc.initiate {
		if err := conn.SendMsg(hbc.msgID, hbc.makeMsg()); err != nil {
			nlog.Error("Send HeartBeatMsg Error", nlog.Uint16("MsgID", hbc.msgID), nlog.Err(err))
		}
	}
}

// makeMsgDefaultFunc 默认的心跳检测消息处理方法
func makeMsgDefaultFunc() (buf []byte) {
	return []byte("ping")
}

// onRemoteNotAliveDefaultFunc 默认的远程连接不存活时的处理方法
func onRemoteNotAliveDefaultFunc(conn niface.IConnection) {
	conn.Stop()
}
