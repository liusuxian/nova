/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 14:03:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-10 19:54:17
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/login/login.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package login

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/utils/nconv"
	"google.golang.org/protobuf/proto"
	"time"
)

// LoginHandler 登录
func LoginHandler(request niface.IRequest) {
	// 获取解析完后的序列化数据
	msgID := request.GetMsgID()
	reqMsg := request.GetResponse().(*pb.LoginRequest)
	nlog.Debug("Receive Login", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", msgID), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
	// 返回
	loginMsg, err := proto.Marshal(&pb.LoginResponse{
		Code: uint32(pb.Code_SUCCESS),
		UserInfo: &pb.UserInfo{
			Uid:         reqMsg.Uid,
			Nickname:    "test_user",
			Sex:         0,
			Avatar:      "",
			Mobile:      "15100000000",
			FromId:      0,
			LastloginAt: time.Now().Format("2006-01-02 15:04:05"),
			RegFromTask: 0,
			IdNo:        "",
			IsRealname:  0,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
	})
	if err != nil {
		nlog.Error("Marshal Login Msg Error", nlog.Err(err))
		return
	}
	if err := request.GetConnection().SendMsg(msgID, loginMsg); err != nil {
		nlog.Error("Send Login Msg Error", nlog.Err(err))
		return
	}
}
