/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 19:28:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-21 14:25:11
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
	"github.com/liusuxian/nova/nrouter"
	"go.uber.org/zap"
)

// HeartbeatChecker 心跳检测器结构
type HeartbeatChecker struct {
	ctx              context.Context         // 当前 Server/Client 的根 Context
	hearbeatMsg      []byte                  // 心跳消息，也可以通过 makeMsgFunc 来动态生成
	makeMsg          niface.HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	onRemoteNotAlive niface.OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	msgID            uint16                  // 用户自定义的心跳检测消息ID
	router           niface.IRouter          // 用户自定义的心跳检测消息业务处理路由
	server           niface.IServer          // 所属服务端
	client           niface.IClient          // 所属客户端
}

// HeartbeatDefaultRouter 收到心跳消息的默认回调路由
type HeartbeatDefaultRouter struct {
	nrouter.BaseRouter
}

// Handle 处理心跳消息
func (hbr *HeartbeatDefaultRouter) Handle(req niface.IRequest) {
	nlog.Debug(req.GetCtx(), "Receive Heartbeat", zap.String("From", req.GetConnection().RemoteAddr().String()), zap.Uint16("MsgID", req.GetMsgID()), zap.ByteString("Data", req.GetData()))
}

// NewHeartbeatCheckerServer Server 创建心跳检测器
func NewHeartbeatCheckerServer(ctx context.Context, server niface.IServer) *HeartbeatChecker {
	heartbeat := &HeartbeatChecker{
		ctx:              ctx,
		makeMsg:          makeMsgDefaultFunc,
		onRemoteNotAlive: onRemoteNotAliveDefaultFunc,
		msgID:            niface.HeartBeatDefaultMsgID,
		router:           &HeartbeatDefaultRouter{},
		server:           server,
		client:           nil,
	}
	return heartbeat
}

// NewHeartbeatCheckerClient Client 创建心跳检测器
func NewHeartbeatCheckerClient(ctx context.Context, client niface.IClient) *HeartbeatChecker {
	heartbeat := &HeartbeatChecker{
		ctx:              ctx,
		makeMsg:          makeMsgDefaultFunc,
		onRemoteNotAlive: onRemoteNotAliveDefaultFunc,
		msgID:            niface.HeartBeatDefaultMsgID,
		router:           &HeartbeatDefaultRouter{},
		server:           nil,
		client:           client,
	}
	return heartbeat
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

// Check 执行心跳检测
func (hbc *HeartbeatChecker) Check() {
	if hbc.server != nil {
		hbc.checkServer()
	} else if hbc.client != nil {
		hbc.checkClient()
	}
}

// checkServer 检测服务端连接
func (hbc *HeartbeatChecker) checkServer() {
	if hbc.server.GetConnManager() != nil {
		allConn := hbc.server.GetConnManager().GetAllConn()
		for _, conn := range allConn {
			if !conn.IsAlive() {
				hbc.onRemoteNotAlive(conn)
			} else {
				msg := hbc.makeMsg(conn)
				if err := conn.SendMsg(hbc.msgID, msg); err != nil {
					nlog.Error(hbc.ctx, "Server Send Heartbeat Msg Error", zap.Uint16("MsgID", hbc.msgID), zap.ByteString("MsgData", msg), zap.Error(err))
				}
			}
		}
	}
}

// checkClient 检测客户端连接
func (hbc *HeartbeatChecker) checkClient() {
	if hbc.client.Conn() != nil {
		if !hbc.client.Conn().IsAlive() {
			hbc.onRemoteNotAlive(hbc.client.Conn())
		} else {
			msg := hbc.makeMsg(hbc.client.Conn())
			if err := hbc.client.Conn().SendMsg(hbc.msgID, msg); err != nil {
				nlog.Error(hbc.ctx, "Client Send Heartbeat Msg Error", zap.Uint16("MsgID", hbc.msgID), zap.ByteString("MsgData", msg), zap.Error(err))
			}
		}
	}
}

// makeMsgDefaultFunc 默认的心跳检测消息处理方法
func makeMsgDefaultFunc(conn niface.IConnection) []byte {
	return []byte("ping")
}

// onRemoteNotAliveDefaultFunc 默认的远程连接不存活时的处理方法
func onRemoteNotAliveDefaultFunc(conn niface.IConnection) {
	conn.Stop()
}
