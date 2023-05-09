/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 17:18:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-10 01:48:33
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/heartbeat/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package heartbeat

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"google.golang.org/protobuf/proto"
	"time"
)

// HeartBeatHandler 心跳消息
func HeartBeatHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetResponse().(*pb.Heartbeat)
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
}

// ReplyHeartBeatHandler 心跳消息回复
func ReplyHeartBeatHandler(request niface.IRequest) {
	resMsg, err := proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
	if err != nil {
		nlog.Error("Marshal Heartbeat Msg Error", nlog.Err(err))
		return
	}
	if err := request.GetConnection().SendMsg(request.GetMsgID(), resMsg); err != nil {
		nlog.Error("Send Heartbeat Error", nlog.Err(err))
		return
	}
}

// LoginHandler 登录
func LoginHandler(request niface.IRequest) {
	loginMsg, err := proto.Marshal(&pb.LoginRequest{
		Mode:  uint32(pb.LoginMode_VISITOR),
		Uid:   7,
		Token: "eyJpZCI6NywidGltZSI6MTY4MzQ2OTk4N30=.4adee61bd4c47b6c73542d25a354b784",
	})
	if err != nil {
		nlog.Error("Marshal Login Msg Error", nlog.Err(err))
		return
	}
	if err := request.GetConnection().SendMsg(uint16(pb.MsgID_LOGIN), loginMsg); err != nil {
		nlog.Error("Send Login Error", nlog.Err(err))
		return
	}
}

// SetHeartBeat 设置当前 Client 的心跳检测器
func SetHeartBeat(c niface.IClient, initiate bool) {
	c.SetHeartBeat(initiate, &niface.HeartBeatOption{
		MakeMsg: func() []byte {
			buf, err := proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
			if err != nil {
				nlog.Error("Marshal Heartbeat Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID: uint16(pb.MsgID_HEARTBEAT),
	})
}
