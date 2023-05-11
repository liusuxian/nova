/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:58:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:03:34
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package router

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/login"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/serveroverload"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"google.golang.org/protobuf/proto"
)

type Router struct {
	c      niface.IClient
	msgMap map[uint16]func() proto.Message
}

// StartRouter 启动路由
func StartRouter(c niface.IClient) (r *Router) {
	r = &Router{
		c:      c,
		msgMap: make(map[uint16]func() proto.Message),
	}
	// 添加业务处理器集合
	r.addRouter(pb.MsgID_SERVER_OVERLOAD, func() proto.Message { return new(pb.ServerOverload) }, serveroverload.ServerOverloadHandler)                                     // 服务器人数超载
	r.addRouter(pb.MsgID_HEARTBEAT, func() proto.Message { return new(pb.Heartbeat) }, heartbeat.HeartBeatHandler, heartbeat.ReplyHeartBeatHandler, heartbeat.LoginHandler) // 心跳
	r.addRouter(pb.MsgID_LOGIN, func() proto.Message { return new(pb.LoginResponse) }, login.LoginHandler)                                                                  // 登录
	return
}

// GetMessage 获取 proto 消息实例
func (r *Router) GetMessage(msgID uint16) (msg proto.Message) {
	return r.msgMap[msgID]()
}

// addRouter 添加业务处理器集合
func (r *Router) addRouter(msgID pb.MsgID, msgFun func() proto.Message, handlers ...niface.RouterHandler) {
	id := uint16(msgID)
	r.msgMap[id] = msgFun
	r.c.AddRouter(id, handlers...)
}
