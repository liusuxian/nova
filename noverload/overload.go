/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 21:39:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-23 22:40:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/noverload/overload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package noverload

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nrouter"
	"go.uber.org/zap"
)

// OverLoadMsg 服务器人数超载消息结构
type OverLoadMsg struct {
	makeMsg niface.OverLoadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	msgID   uint16                 // 用户自定义的服务器人数超载消息ID
	router  niface.IRouter         // 用户自定义的服务器人数超载消息业务处理路由
	server  niface.IServer         // 所属服务端
	client  niface.IClient         // 所属客户端
}

// OverLoadMsgDefaultRouter 服务器人数超载消息的默认回调路由
type OverLoadMsgDefaultRouter struct {
	nrouter.BaseRouter
}

// Handle 处理服务器人数超载消息
func (olr *OverLoadMsgDefaultRouter) Handle(request niface.IRequest) {
	nlog.Debug(request.GetCtx(), "Receive OverLoadMsg", zap.String("From", request.GetConnection().RemoteAddr().String()), zap.Uint16("MsgID", request.GetMsgID()), zap.ByteString("Data", request.GetData()))
}

// NewOverLoadMsgServer Server 创建服务器人数超载消息
func NewOverLoadMsgServer(server niface.IServer) *OverLoadMsg {
	overLoadMsg := &OverLoadMsg{
		makeMsg: makeMsgDefaultFunc,
		msgID:   niface.OverLoadDefaultMsgID,
		router:  nil,
		server:  server,
		client:  nil,
	}
	return overLoadMsg
}

// NewOverLoadMsgClient Client 创建服务器人数超载消息
func NewOverLoadMsgClient(client niface.IClient) *OverLoadMsg {
	overLoadMsg := &OverLoadMsg{
		makeMsg: makeMsgDefaultFunc,
		msgID:   niface.OverLoadDefaultMsgID,
		router:  &OverLoadMsgDefaultRouter{},
		server:  nil,
		client:  client,
	}
	return overLoadMsg
}

// SetOverLoadMsgFunc 设置服务器人数超载消息处理方法
func (ol *OverLoadMsg) SetOverLoadMsgFunc(f niface.OverLoadMsgFunc) {
	if f != nil {
		ol.makeMsg = f
	}
}

// BindRouter 绑定服务器人数超载消息业务处理路由
func (ol *OverLoadMsg) BindRouter(msgID uint16, router niface.IRouter) {
	if ol.server != nil {
		if msgID != niface.OverLoadDefaultMsgID {
			ol.msgID = msgID
		}
	} else if ol.client != nil {
		if router != nil && msgID != niface.OverLoadDefaultMsgID {
			ol.msgID = msgID
			ol.router = router
		}
	}
}

// GetMsgID 获取服务器人数超载消息ID
func (ol *OverLoadMsg) GetMsgID() uint16 {
	return ol.msgID
}

// GetMsgData 获取服务器人数超载消息数据
func (ol *OverLoadMsg) GetMsgData() []byte {
	return ol.makeMsg()
}

// GetRouter 获取心跳检测消息业务处理路由
func (ol *OverLoadMsg) GetRouter() niface.IRouter {
	return ol.router
}

// makeMsgDefaultFunc 默认的服务器人数超载消息处理方法
func makeMsgDefaultFunc() []byte {
	return []byte("server overload")
}