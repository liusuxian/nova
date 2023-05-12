/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 14:03:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 14:20:12
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package login

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/proto/pb"
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/redisdb"
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
	reqMsg := request.GetSerializedData().(*pb.LoginRequest)
	nlog.Debug("Receive Login", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", msgID), nlog.Any("ReqMsg", nconv.ProtoMsgToMap(reqMsg)))
	// 测试读取 redis
	_, err := redisdb.Instance().Do(request.GetCtx(), "GET", "aaa")
	if err != nil {
		nlog.Error("Redis Error", nlog.Err(err))
	}
	// 返回
	if err := request.RespMsg(func() (buf []byte, err error) {
		return proto.Marshal(&pb.LoginResponse{
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
	}); err != nil {
		nlog.Error("Response Login Msg Error", nlog.Err(err))
		return
	}
}
