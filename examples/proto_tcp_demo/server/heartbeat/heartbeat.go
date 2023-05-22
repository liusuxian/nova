/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-08 02:22:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-23 01:13:41
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
	nlog.Debug("Receive HeartBeat", nlog.String("From", request.GetConnection().RemoteAddr()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
}

// SetHeartBeat 设置当前 Server 的心跳检测器
func SetHeartBeat(s niface.IServer, initiate bool) {
	s.SetHeartBeat(initiate, &niface.HeartBeatOption{
		MakeMsg: func() (buf []byte, err error) {
			return proto.Marshal(&pb.Heartbeat{Timestamp: time.Now().Unix()})
		},
		MsgID: uint16(pb.MsgID_HEARTBEAT),
	})
}
