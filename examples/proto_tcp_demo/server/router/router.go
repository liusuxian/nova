/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:43:47
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:08:08
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package router

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/heartbeat"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/login"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"google.golang.org/protobuf/proto"
)

type sRouter struct {
	s      niface.IServer
	msgMap map[uint16]func() proto.Message
}

var instance *sRouter

func init() {
	instance = &sRouter{
		msgMap: make(map[uint16]func() proto.Message),
	}
}

// StartRouter 启动路由
func StartRouter(s niface.IServer) {
	instance.s = s
	// 添加业务处理器集合
	addRouter(pb.MsgID_HEARTBEAT, func() proto.Message { return new(pb.Heartbeat) }, heartbeat.HeartBeatHandler) // 心跳
	addRouter(pb.MsgID_LOGIN, func() proto.Message { return new(pb.LoginRequest) }, login.LoginHandler)          // 登录
}

// GetMessage 获取 proto 消息实例
func GetMessage(msgID uint16) (msg proto.Message) {
	return instance.msgMap[msgID]()
}

// addRouter 添加业务处理器集合
func addRouter(msgID pb.MsgID, msgFun func() proto.Message, handlers ...niface.RouterHandler) {
	id := uint16(msgID)
	instance.msgMap[id] = msgFun
	instance.s.AddRouter(id, handlers...)
}
