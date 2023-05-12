/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 00:16:13
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 14:51:39
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package heartbeat

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/utils/nconv"
	"google.golang.org/protobuf/proto"
	"time"
)

// HeartBeatHandler 心跳消息
func HeartBeatHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetSerializedData().(*pb.Heartbeat)
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
}

// ReplyHeartBeatHandler 心跳消息回复
func ReplyHeartBeatHandler(request niface.IRequest) {
	if err := request.RespMsg(func() (buf []byte, err error) {
		return proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
	}); err != nil {
		nlog.Error("Reply Heartbeat Msg Error", nlog.Err(err))
		return
	}
}

// LoginHandler 登录
func LoginHandler(request niface.IRequest) {
	if err := request.RespMsgWithId(uint16(pb.MsgID_LOGIN), func() (buf []byte, err error) {
		return proto.Marshal(&pb.LoginRequest{
			Mode:  uint32(pb.LoginMode_VISITOR),
			Uid:   8,
			Token: "eyJpZCI6OCwidGltZSI6MTY4MzY4OTAwM30=.6fb749f16570e415b7936c636e0c759a",
		})
	}); err != nil {
		nlog.Error("Send Login Msg Error", nlog.Err(err))
		return
	}
}

// SetHeartBeat 设置当前 Client 的心跳检测器
func SetHeartBeat(c niface.IClient, initiate bool) {
	c.SetHeartBeat(initiate, &niface.HeartBeatOption{
		MakeMsg: func() (buf []byte, err error) {
			return proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
		},
		MsgID: uint16(pb.MsgID_HEARTBEAT),
	})
}
