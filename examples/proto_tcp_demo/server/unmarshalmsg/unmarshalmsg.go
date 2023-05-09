/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 21:35:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-09 21:16:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/proto_tcp_demo/server/unmarshalmsg/unmarshalmsg.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package unmarshalmsg

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/server/router"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"google.golang.org/protobuf/proto"
)

// unmarshalMsg 解析消息拦截器
type unmarshalMsg struct {
}

// Intercept 解析消息
func (um *unmarshalMsg) Intercept(chain niface.IChain) (resp niface.IcResp) {
	iMessage := chain.GetIMessage()
	if iMessage == nil {
		return nil
	}
	request := chain.Request()
	if request == nil {
		return nil
	}

	switch iRequest := request.(type) {
	case niface.IRequest:
		msgID := iRequest.GetMsgID()
		nlog.Debug("Receive MsgID", nlog.Uint16("MsgID", msgID))
		reqMsg := router.GetMessage(msgID)
		if err := proto.Unmarshal(iRequest.GetData(), reqMsg); err != nil {
			nlog.Error("Unmarshal Msg Error", nlog.Uint16("MsgID", msgID), nlog.Err(err))
			return nil
		}
		return chain.ProceedWithIMessage(iMessage, reqMsg)
	}

	return nil
}

// AddInterceptor 添加解析消息拦截器
func AddInterceptor(s niface.IServer) {
	s.AddInterceptor(&unmarshalMsg{})
}
