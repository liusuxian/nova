/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 17:18:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 00:46:07
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
	"github.com/liusuxian/nova/nrouter"
	"google.golang.org/protobuf/proto"
	"time"
)

// HeartBeatRouter 心跳消息路由
type HeartBeatRouter struct {
	nrouter.BaseRouter      // 基础路由
	initiate           bool // 发起心跳
}

// 是否已发送登录消息
var isSendLogin bool

// Handle 处理心跳消息
func (hbr *HeartBeatRouter) Handle(request niface.IRequest) {
	// 收到心跳消息
	reqMsg := &pb.Heartbeat{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		nlog.Error("Unmarshal Heartbeat Msg Error", nlog.Err(err))
		return
	}
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
	// 返回心跳消息
	if !hbr.initiate {
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
	// 登录
	if !isSendLogin {
		loginMsg, err := proto.Marshal(&pb.LoginRequest{
			Mode:  uint32(pb.LoginMode_VISITOR),
			Uid:   7,
			Token: "eyJpZCI6NywidGltZSI6MTY4MzIxMTI1MH0=.008ba9663fa59fb05ce11f80a3df1d8b",
		})
		if err != nil {
			nlog.Error("Marshal Login Msg Error", nlog.Err(err))
			return
		}
		if err := request.GetConnection().SendMsg(uint16(pb.MsgID_LOGIN), loginMsg); err != nil {
			nlog.Error("Send Login Error", nlog.Err(err))
			return
		}
		isSendLogin = true
	}
}

// SetHeartBeat 设置当前 Client 的心跳检测器
func SetHeartBeat(c niface.IClient, initiate bool) {
	c.SetHeartBeat(initiate, &niface.HeartBeatOption{
		MakeMsg: func() []byte {
			buf, err := proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
			if err != nil {
				nlog.Fatal("Marshal Heartbeat Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID:  uint16(pb.MsgID_HEARTBEAT),
		Router: &HeartBeatRouter{initiate: initiate},
	})
}
