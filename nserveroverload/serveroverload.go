/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 17:52:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-07 22:34:15
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
)

// ServerOverloadChecker 服务器人数超载检测器结构
type ServerOverloadChecker struct {
	makeMsg    niface.ServerOverloadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	msgID      uint16                       // 用户自定义的服务器人数超载消息ID
	handlers   []niface.RouterHandler       // 用户自定义的服务器人数超载消息的业务处理器集合
	clientCall bool                         // 是否是客户端调用
}

// ServerOverloadDefaultHandlers 默认的处理服务器人数超载消息业务处理器
func ServerOverloadDefaultHandlers(request niface.IRequest) {
	nlog.Debug("Receive Server Overload Msg", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("Data", request.GetData()))
}

// NewServerOverloadChecker 创建服务器人数超载检测器
func NewServerOverloadChecker(clientCall bool) (checker niface.IServerOverloadChecker) {
	serverOverload := &ServerOverloadChecker{
		makeMsg:    makeMsgDefaultFunc,
		msgID:      niface.ServerOverloadDefaultMsgID,
		handlers:   nil,
		clientCall: clientCall,
	}
	if clientCall {
		serverOverload.handlers = []niface.RouterHandler{ServerOverloadDefaultHandlers}
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
func (soc *ServerOverloadChecker) BindRouter(msgID uint16, handlers ...niface.RouterHandler) {
	if soc.clientCall {
		if msgID != niface.ServerOverloadDefaultMsgID && len(handlers) > 0 {
			soc.msgID = msgID
			soc.handlers = handlers
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

// GetHandlers 获取服务器人数超载消息的业务处理器集合
func (soc *ServerOverloadChecker) GetHandlers() (handlers []niface.RouterHandler) {
	return soc.handlers
}

// makeMsgDefaultFunc 默认的服务器人数超载消息处理方法
func makeMsgDefaultFunc() (buf []byte) {
	return []byte("server overload")
}
