/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 17:44:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 20:01:08
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nheartbeat/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nheartbeat

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrouter"
	"time"
)

// HeartbeatChecker 心跳检测器结构
type HeartbeatChecker struct {
	ctx              context.Context         // 心跳检测器的 Context
	interval         time.Duration           // 心跳检测时间间隔
	quitChan         chan bool               // 退出信号
	makeMsg          niface.HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	onRemoteNotAlive niface.OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	msgID            uint16                  // 用户自定义的心跳检测消息ID
	router           niface.IRouter          // 用户自定义的心跳检测消息业务处理路由
	conn             niface.IConnection      // 绑定的连接
	initiate         bool                    // 发起心跳
}

// HeartbeatDefaultRouter 收到心跳消息的默认回调路由
type HeartbeatDefaultRouter struct {
	nrouter.BaseRouter                   // 基础路由
	heartbeatChecker   *HeartbeatChecker // 所属心跳检测器
}

// Handle 处理心跳消息
func (hbr *HeartbeatDefaultRouter) Handle(request niface.IRequest) {
	nlog.Debug(request.GetCtx(), "Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("Data", request.GetData()))
	// 回复心跳消息
	hbr.heartbeatChecker.replyHeartBeatMsg()
}

// NewHeartbeatChecker 创建心跳检测器
func NewHeartbeatChecker(interval time.Duration, initiate bool) (checker niface.IHeartBeatChecker) {
	heartbeat := &HeartbeatChecker{
		ctx:              context.Background(),
		interval:         interval,
		quitChan:         make(chan bool),
		makeMsg:          makeMsgDefaultFunc,
		onRemoteNotAlive: onRemoteNotAliveDefaultFunc,
		msgID:            niface.HeartBeatDefaultMsgID,
		conn:             nil,
		initiate:         initiate,
	}
	heartbeat.router = &HeartbeatDefaultRouter{heartbeatChecker: heartbeat}
	return heartbeat
}

// Start 启动心跳检测
func (hbc *HeartbeatChecker) Start() {
	go hbc.start()
}

// Stop 停止心跳检测
func (hbc *HeartbeatChecker) Stop() {
	hbc.quitChan <- true
}

// SetHeartBeatMsgFunc 设置心跳检测消息处理方法
func (hbc *HeartbeatChecker) SetHeartBeatMsgFunc(f niface.HeartBeatMsgFunc) {
	if f != nil {
		hbc.makeMsg = f
	}
}

// SetOnRemoteNotAlive 设置远程连接不存活时的处理方法
func (hbc *HeartbeatChecker) SetOnRemoteNotAlive(f niface.OnRemoteNotAlive) {
	if f != nil {
		hbc.onRemoteNotAlive = f
	}
}

// BindRouter 绑定心跳检测消息业务处理路由
func (hbc *HeartbeatChecker) BindRouter(msgID uint16, router niface.IRouter) {
	if router != nil && msgID != niface.HeartBeatDefaultMsgID {
		hbc.msgID = msgID
		hbc.router = router
	}
}

// BindConn 绑定连接
func (hbc *HeartbeatChecker) BindConn(conn niface.IConnection) {
	hbc.conn = conn
	conn.SetHeartBeat(hbc)
}

// Clone 克隆心跳检测器
func (hbc *HeartbeatChecker) Clone() (checker niface.IHeartBeatChecker) {
	heartbeat := &HeartbeatChecker{
		ctx:              context.Background(),
		interval:         hbc.interval,
		quitChan:         make(chan bool),
		makeMsg:          hbc.makeMsg,
		onRemoteNotAlive: hbc.onRemoteNotAlive,
		msgID:            hbc.msgID,
		router:           hbc.router,
		conn:             nil,
		initiate:         hbc.initiate,
	}
	return heartbeat
}

// GetMsgID 获取心跳检测消息ID
func (hbc *HeartbeatChecker) GetMsgID() (msgID uint16) {
	return hbc.msgID
}

// GetMessage 获取心跳检测消息
func (hbc *HeartbeatChecker) GetMessage() (msg niface.IMessage) {
	return npack.NewMsgPackage(hbc.msgID, hbc.makeMsg())
}

// GetRouter 获取心跳检测消息业务处理路由
func (hbc *HeartbeatChecker) GetRouter() (router niface.IRouter) {
	return hbc.router
}

// start 启动心跳检测
func (hbc *HeartbeatChecker) start() {
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
func (hbc *HeartbeatChecker) check() {
	if hbc.conn == nil {
		return
	}
	if !hbc.conn.IsAlive() {
		hbc.onRemoteNotAlive(hbc.conn)
	} else {
		hbc.sendHeartBeatMsg()
	}
}

// sendHeartBeatMsg 发送心跳消息
func (hbc *HeartbeatChecker) sendHeartBeatMsg() {
	if hbc.initiate {
		msg := hbc.makeMsg()
		if err := hbc.conn.SendMsg(hbc.msgID, msg); err != nil {
			nlog.Error(hbc.ctx, "Send HeartBeatMsg Error", nlog.Uint16("MsgID", hbc.msgID), nlog.Err(err))
		}
	}
}

// replyHeartBeatMsg 回复心跳消息
func (hbc *HeartbeatChecker) replyHeartBeatMsg() {
	if !hbc.initiate {
		msg := hbc.makeMsg()
		if err := hbc.conn.SendMsg(hbc.msgID, msg); err != nil {
			nlog.Error(hbc.ctx, "Reply HeartBeatMsg Error", nlog.Uint16("MsgID", hbc.msgID), nlog.Err(err))
		}
	}
}

// makeMsgDefaultFunc 默认的心跳检测消息处理方法
func makeMsgDefaultFunc() []byte {
	return []byte("ping")
}

// onRemoteNotAliveDefaultFunc 默认的远程连接不存活时的处理方法
func onRemoteNotAliveDefaultFunc(conn niface.IConnection) {
	conn.Stop()
}
