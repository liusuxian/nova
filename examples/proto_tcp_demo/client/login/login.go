/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 22:42:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-23 01:13:04
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package login

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/utils/nconv"
)

// LoginHandler 登录
func LoginHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetSerializedData().(*pb.LoginResponse)
	nlog.Debug("Receive Login", nlog.String("From", request.GetConnection().RemoteAddr()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
}
