/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 23:08:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 14:21:14
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package serveroverload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"google.golang.org/protobuf/proto"
)

// SetServerOverload 设置当前 Server 的服务器人数超载检测器
func SetServerOverload(s niface.IServer) {
	s.SetServerOverload(&niface.ServerOverloadOption{
		MakeMsg: func() (buf []byte, err error) {
			return proto.Marshal(&pb.ServerOverload{})
		},
		MsgID: uint16(pb.MsgID_SERVER_OVERLOAD),
	})
}
