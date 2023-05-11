/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 16:19:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:04:03
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package unmarshalmsg

import (
	"github.com/liusuxian/nova/examples/proto_tcp_demo/client/router"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"google.golang.org/protobuf/proto"
)

// unmarshalMsg 解析消息拦截器
type unmarshalMsg struct {
	r *router.Router
}

// Intercept 拦截器的拦截处理方法
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
		reqMsg := um.r.GetMessage(msgID)
		if err := proto.Unmarshal(iRequest.GetData(), reqMsg); err != nil {
			nlog.Error("Unmarshal Msg Error", nlog.Uint16("MsgID", msgID), nlog.Err(err))
			return nil
		}
		return chain.ProceedWithIMessage(iMessage, reqMsg)
	}

	return nil
}

// AddInterceptor 添加解析消息拦截器
func AddInterceptor(c niface.IClient, r *router.Router) {
	c.AddInterceptor(&unmarshalMsg{
		r: r,
	})
}
