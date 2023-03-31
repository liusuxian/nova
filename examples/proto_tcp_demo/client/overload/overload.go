/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-24 14:45:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 16:28:33
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/overload/overload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package overload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nrouter"
	"google.golang.org/protobuf/proto"
)

// OverLoadRouter 服务器人数超载消息路由
type OverLoadRouter struct {
	nrouter.BaseRouter // 基础路由
}

// Handle 处理服务器人数超载消息
func (olr *OverLoadRouter) Handle(request niface.IRequest) {
	// 收到服务器人数超载消息
	reqMsg := &pb.OverLoad{}
	if err := proto.Unmarshal(request.GetData(), reqMsg); err != nil {
		nlog.Error(request.GetCtx(), "Unmarshal OverLoad Msg Error", nlog.Err(err))
		return
	}
	nlog.Debug(request.GetCtx(), "Receive OverLoadMsg", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
}

// SetOverLoadMsg 设置当前 Client 的服务器人数超载消息
func SetOverLoadMsg(c niface.IClient) {
	c.SetOverLoadMsg(&niface.OverLoadMsgOption{
		MakeMsg: func() []byte {
			msg := &pb.OverLoad{}
			buf, err := proto.Marshal(msg)
			if err != nil {
				nlog.Fatal(c.GetCtx(), "Marshal OverLoad Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID:  uint16(pb.MsgID_SERVER_OVERLOAD),
		Router: &OverLoadRouter{},
	})
}
