/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 17:44:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 00:36:41
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nheartbeat/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
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
	handlers         []niface.RouterHandler  // 用户自定义的心跳检测消息的业务处理器集合
	conn             niface.IConnection      // 绑定的连接
	initiate         bool                    // 发起心跳
}

// HeartBeatDefaultHandler 默认的心跳消息业务处理器
func HeartBeatDefaultHandler(request niface.IRequest) {
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("Data", request.GetData()))
}

// NewHeartBeatChecker 创建心跳检测器
func NewHeartBeatChecker(interval time.Duration, initiate bool) (checker niface.IHeartBeatChecker) {
	heartbeat := &HeartBeatChecker{
		interval:         interval,
		quitChan:         make(chan bool),
		makeMsg:          makeMsgDefaultFunc,
		onRemoteNotAlive: onRemoteNotAliveDefaultFunc,
		msgID:            niface.HeartBeatDefaultMsgID,
		handlers:         []niface.RouterHandler{HeartBeatDefaultHandler},
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

// BindRouter 绑定心跳检测消息的业务处理器集合
func (hbc *HeartBeatChecker) BindRouter(msgID uint16, handlers ...niface.RouterHandler) {
	if msgID != niface.HeartBeatDefaultMsgID && len(handlers) > 0 {
		hbc.msgID = msgID
		hbc.handlers = handlers
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
		handlers:         hbc.handlers,
		conn:             nil,
		initiate:         hbc.initiate,
	}
	return heartbeat
}

// GetMsgID 获取心跳检测消息ID
func (hbc *HeartBeatChecker) GetMsgID() (msgID uint16) {
	return hbc.msgID
}

// GetMessage 获取心跳检测消息
func (hbc *HeartBeatChecker) GetMessage() (msg niface.IMessage) {
	return npack.NewMsgPackage(hbc.msgID, hbc.makeMsg())
}

// GetHandlers 获取心跳检测消息的业务处理器集合
func (hbc *HeartBeatChecker) GetHandlers() (handlers []niface.RouterHandler) {
	return hbc.handlers
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
