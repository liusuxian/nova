/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:07:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:09:10
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package heartbeat

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
)

// HeartBeatHandler 心跳消息
func HeartBeatHandler(request niface.IRequest) {
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("ReqMsg", request.GetData()))
	if err := request.GetConnection().SendMsg(request.GetMsgID(), request.GetData()); err != nil {
		nlog.Error("Send Heartbeat Error", nlog.Err(err))
		return
	}
}
