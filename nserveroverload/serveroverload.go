/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 17:52:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-01 21:49:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nserveroverload/serveroverload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nserveroverload

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/npack"
	"github.com/liusuxian/nova/nrouter"
)

// ServerOverloadChecker 服务器人数超载检测器结构
type ServerOverloadChecker struct {
	makeMsg    niface.ServerOverloadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	msgID      uint16                       // 用户自定义的服务器人数超载消息ID
	router     niface.IRouter               // 用户自定义的服务器人数超载消息业务处理路由
	clientCall bool                         // 是否是客户端调用
}

// ServerOverloadDefaultRouter 收到服务器人数超载消息的默认回调路由
type ServerOverloadDefaultRouter struct {
	nrouter.BaseRouter
}

// Handle 处理服务器人数超载消息
func (sor *ServerOverloadDefaultRouter) Handle(request niface.IRequest) {
	nlog.Debug(request.GetCtx(), "Receive Server Overload Msg", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("Data", request.GetData()))
}

// NewServerOverloadChecker 创建服务器人数超载检测器
func NewServerOverloadChecker(clientCall bool) (checker niface.IServerOverloadChecker) {
	serverOverload := &ServerOverloadChecker{
		makeMsg:    makeMsgDefaultFunc,
		msgID:      niface.ServerOverloadDefaultMsgID,
		router:     nil,
		clientCall: clientCall,
	}
	if clientCall {
		serverOverload.router = &ServerOverloadDefaultRouter{}
	}
	return serverOverload
}

// Check 服务器人数超载检测
func (soc *ServerOverloadChecker) Check(server niface.IServer, maxConn int) (isOverload bool) {
	return server.GetConnections() > maxConn
}

// SetServerOverloadMsgFunc 设置服务器人数超载消息处理方法
func (soc *ServerOverloadChecker) SetServerOverloadMsgFunc(f niface.ServerOverloadMsgFunc) {
	if f != nil {
		soc.makeMsg = f
	}
}

// BindRouter 绑定服务器人数超载消息业务处理路由
func (soc *ServerOverloadChecker) BindRouter(msgID uint16, router niface.IRouter) {
	if soc.clientCall {
		if router != nil && msgID != niface.ServerOverloadDefaultMsgID {
			soc.msgID = msgID
			soc.router = router
		}
	} else {
		if msgID != niface.ServerOverloadDefaultMsgID {
			soc.msgID = msgID
		}
	}
}

// GetMsgID 获取服务器人数超载消息ID
func (soc *ServerOverloadChecker) GetMsgID() (msgID uint16) {
	return soc.msgID
}

// GetMessage 获取服务器人数超载消息
func (soc *ServerOverloadChecker) GetMessage() (msg niface.IMessage) {
	return npack.NewMsgPackage(soc.msgID, soc.makeMsg())
}

// GetRouter 获取服务器人数超载消息业务处理路由
func (soc *ServerOverloadChecker) GetRouter() (router niface.IRouter) {
	return soc.router
}

// makeMsgDefaultFunc 默认的服务器人数超载消息处理方法
func makeMsgDefaultFunc() []byte {
	return []byte("server overload")
}
