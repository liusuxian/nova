/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 23:08:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-01 23:09:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/serveroverload/serveroverload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package serveroverload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"google.golang.org/protobuf/proto"
)

// SetServerOverload 设置当前 Server 的服务器人数超载检测器
func SetServerOverload(s niface.IServer) {
	s.SetServerOverload(&niface.ServerOverloadOption{
		MakeMsg: func() []byte {
			buf, err := proto.Marshal(&pb.ServerOverload{})
			if err != nil {
				nlog.Fatal(s.GetCtx(), "Marshal ServerOverload Msg Error", nlog.Err(err))
			}
			return buf
		},
		MsgID: uint16(pb.MsgID_SERVER_OVERLOAD),
	})
}
