/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:35:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 23:35:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/unmarshalmsg/unmarshalmsg.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package unmarshalmsg

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"google.golang.org/protobuf/proto"
)

// UnmarshalMsg 解析消息拦截器
type UnmarshalMsg struct {
	msgMap map[uint16]func() proto.Message
}

// Intercept 拦截器的拦截处理方法
func (um *UnmarshalMsg) Intercept(chain niface.IChain) (resp niface.IcResp) {
	iMessage := chain.GetIMessage()
	if iMessage == nil {
		return chain.ProceedWithIMessage(iMessage, nil)
	}

	request := chain.Request()
	if request != nil {
		switch iRequest := request.(type) {
		case niface.IRequest:
			msgID := iRequest.GetMsgID()
			nlog.Debug("Receive MsgID", nlog.Uint16("MsgID", msgID))
			reqMsg := um.msgMap[msgID]()
			if err := proto.Unmarshal(iRequest.GetData(), reqMsg); err != nil {
				nlog.Error("Unmarshal Msg Error", nlog.Uint16("MsgID", msgID), nlog.Err(err))
				return chain.ProceedWithIMessage(iMessage, nil)
			}
			return chain.ProceedWithIMessage(iMessage, reqMsg)
		}
	}

	return chain.ProceedWithIMessage(iMessage, nil)
}

// AddInterceptor 添加解析消息拦截器
func AddInterceptor(c niface.IClient) {
	c.AddInterceptor(&UnmarshalMsg{
		msgMap: map[uint16]func() proto.Message{
			uint16(pb.MsgID_SERVER_OVERLOAD): func() proto.Message { return new(pb.ServerOverload) }, // 服务器人数超载
			uint16(pb.MsgID_HEARTBEAT):       func() proto.Message { return new(pb.Heartbeat) },      // 心跳
			uint16(pb.MsgID_LOGIN):           func() proto.Message { return new(pb.LoginRequest) },   // 登录
		},
	})
}
