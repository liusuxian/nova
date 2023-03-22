/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-21 22:19:14
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-22 22:02:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nserver"
	"google.golang.org/protobuf/proto"
	"time"
)

func main() {
	// 创建 Server
	s := nserver.NewServer(
		nserver.WithMulticore(true),
		nserver.WithReuseAddr(true),
		nserver.WithReusePort(true),
		nserver.WithLockOSThread(true),
		nserver.WithTicker(true),
	)
	// 设置当前 Server 的心跳检测
	s.SetHeartBeat(&niface.HeartBeatOption{
		MakeMsg: func() []byte {
			msg := &pb.Heartbeat{Timestamp: time.Now().Unix()}
			buf, _ := proto.Marshal(msg)
			if len(buf) == 0 {
				return []byte("ping")
			}
			return buf
		},
		MsgID: uint16(pb.MsgID_HEARTBEAT),
	})
	// 启动服务器
	s.Start()
}
