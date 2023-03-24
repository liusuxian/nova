/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-24 14:45:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-24 14:54:38
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/overload/overload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package overload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// SetOverLoadMsg 设置当前 Client 的服务器人数超载消息
func SetOverLoadMsg(s niface.IServer) {
	s.SetOverLoadMsg(&niface.OverLoadMsgOption{
		MakeMsg: func() []byte {
			msg := &pb.OverLoad{}
			buf, err := proto.Marshal(msg)
			if err != nil {
				nlog.Fatal(s.GetCtx(), "Marshal OverLoad Msg Error", zap.Error(err))
			}
			return buf
		},
		MsgID: uint16(pb.MsgID_SERVER_OVERLOAD),
	})
}
