/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 17:18:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-24 14:47:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/heartbeat/heartbeat.go
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
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

// HeartBeatRouter 心跳消息路由
type HeartBeatRouter struct {
	nrouter.BaseRouter      // 基础路由
	initiate           bool // 发起心跳
}

// Handle 处理心跳消息
func (hbr *HeartBeatRouter) Handle(request niface.IRequest) {
	// 收到心跳消息
	reqMsg := &pb.Heartbeat{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		nlog.Error(request.GetCtx(), "Unmarshal Heartbeat Msg Error", zap.Error(err))
		return
	}
	nlog.Debug(request.GetCtx(), "Receive Heartbeat", zap.String("From", request.GetConnection().RemoteAddr().String()), zap.Uint16("MsgID", request.GetMsgID()), zap.Reflect("ReqMsg", reqMsg))
	// 返回心跳消息
	if !hbr.initiate {
		resMsg, err := proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
		if err != nil {
			nlog.Error(request.GetCtx(), "Marshal Heartbeat Msg Error", zap.Error(err))
			return
		}
		if err := request.GetConnection().SendMsg(request.GetMsgID(), resMsg, nil); err != nil {
			nlog.Error(request.GetCtx(), "Send Heartbeat Error", zap.Error(err))
			return
		}
	}
}

// SetHeartBeat 设置当前 Server 的心跳检测
func SetHeartBeat(s niface.IServer, initiate bool) {
	s.SetHeartBeat(&niface.HeartBeatOption{
		MakeMsg: func() []byte {
			msg := &pb.Heartbeat{Timestamp: time.Now().Unix()}
			buf, err := proto.Marshal(msg)
			if err != nil {
				nlog.Fatal(s.GetCtx(), "Marshal Heartbeat Msg Error", zap.Error(err))
			}
			return buf
		},
		MsgID:  uint16(pb.MsgID_HEARTBEAT),
		Router: &HeartBeatRouter{initiate: initiate},
	}, initiate)
}
