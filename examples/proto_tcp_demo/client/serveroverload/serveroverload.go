/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-02 16:20:06
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 14:31:59
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package serveroverload

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/utils/nconv"
	"google.golang.org/protobuf/proto"
)

// ServerOverloadHandler 服务器人数超载消息
func ServerOverloadHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetSerializedData().(*pb.ServerOverload)
	nlog.Debug("Receive ServerOverload", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
}

// SetServerOverload 设置当前 Client 的服务器人数超载检测器
func SetServerOverload(c niface.IClient) {
	c.SetServerOverload(&niface.ServerOverloadOption{
		MakeMsg: func() (buf []byte, err error) {
			return proto.Marshal(&pb.ServerOverload{})
		},
		MsgID: uint16(pb.MsgID_SERVER_OVERLOAD),
	})
}
