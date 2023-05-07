/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:35:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 01:39:58
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
	"google.golang.org/protobuf/reflect/protoreflect"
)

// UnmarshalMsg 解析消息拦截器
type UnmarshalMsg struct {
}

// Intercept 拦截器的拦截处理方法
func (um *UnmarshalMsg) Intercept(chain niface.IChain) (resp niface.IcResp) {
	request := chain.Request()
	if request != nil {
		switch iRequest := request.(type) {
		case niface.IRequest:
			msgID := iRequest.GetMsgID()
			nlog.Debug("Receive MsgID", nlog.Uint16("MsgID", msgID))
			switch msgID {
			case uint16(pb.MsgID_SERVER_OVERLOAD):
				// 服务器人数超载
				if err := unmarshalMsg(iRequest, &pb.ServerOverload{}); err != nil {
					return
				}
			case uint16(pb.MsgID_HEARTBEAT):
				// 心跳
				if err := unmarshalMsg(iRequest, &pb.Heartbeat{}); err != nil {
					return
				}
			case uint16(pb.MsgID_LOGIN):
				// 登录
				if err := unmarshalMsg(iRequest, &pb.LoginResponse{}); err != nil {
					return
				}
			default:
				nlog.Error("Receive Unknown MsgID", nlog.Uint16("MsgID", msgID))
				return
			}
		}
	}
	return chain.Proceed(chain.Request())
}

// unmarshalMsg 解析消息
func unmarshalMsg(request niface.IRequest, msg protoreflect.ProtoMessage) (err error) {
	if err = proto.Unmarshal(request.GetData(), msg); err != nil {
		nlog.Error("Unmarshal Msg Error", nlog.Uint16("MsgID", request.GetMsgID()), nlog.Err(err))
		return
	}
	request.SetResponse(msg)
	return
}
