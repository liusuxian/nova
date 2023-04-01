/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 22:52:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-01 23:10:06
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/serveroverload/serveroverload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package serveroverload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nrouter"
	"google.golang.org/protobuf/proto"
)

// ServerOverloadRouter 服务器人数超载消息路由
type ServerOverloadRouter struct {
	nrouter.BaseRouter // 基础路由
}

// Handle 处理服务器人数超载消息
func (sor *ServerOverloadRouter) Handle(request niface.IRequest) {
	// 收到服务器人数超载消息
	reqMsg := &pb.ServerOverload{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		nlog.Error(request.GetCtx(), "Unmarshal ServerOverload Msg Error", nlog.Err(err))
		return
	}
	nlog.Debug(request.GetCtx(), "Receive ServerOverload", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
}

// SetServerOverload 设置当前 Client 的服务器人数超载检测器
func SetServerOverload(c niface.IClient) {
	c.SetServerOverload(&niface.ServerOverloadOption{
		MakeMsg: func() []byte {
			buf, err := proto.Marshal(&pb.ServerOverload{})
			if err != nil {
				nlog.Fatal(c.GetCtx(), "Marshal ServerOverload Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID:  uint16(pb.MsgID_SERVER_OVERLOAD),
		Router: &ServerOverloadRouter{},
	})
}
