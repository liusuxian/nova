/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 01:39:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-10 01:39:32
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/client/login/login.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package login

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
)

// LoginHandler 登录
func LoginHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	reqMsg := request.GetResponse().(*pb.LoginResponse)
	nlog.Debug("Receive Login", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.Reflect("ReqMsg", reqMsg))
}
