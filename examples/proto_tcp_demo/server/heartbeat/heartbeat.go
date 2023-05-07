/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-23 17:18:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 01:20:53
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
	"google.golang.org/protobuf/proto"
	"time"
)

// HeartBeatHandler 心跳消息
func HeartBeatHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetResponse().(*pb.Heartbeat)
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
}

// SetHeartBeat 设置当前 Server 的心跳检测器
func SetHeartBeat(s niface.IServer, initiate bool) {
	s.SetHeartBeat(initiate, &niface.HeartBeatOption{
		MakeMsg: func() []byte {
			msg := &pb.Heartbeat{Timestamp: time.Now().Unix()}
			buf, err := proto.Marshal(msg)
			if err != nil {
				nlog.Fatal("Marshal Heartbeat Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID:          uint16(pb.MsgID_HEARTBEAT),
		RouterHandlers: []niface.RouterHandler{HeartBeatHandler},
	})
}
