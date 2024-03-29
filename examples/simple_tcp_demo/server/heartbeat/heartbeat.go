/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:06:13
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-23 01:13:59
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
	nlog.Debug("Receive HeartBeat", nlog.String("From", request.GetConnection().RemoteAddr()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("ReqMsg", request.GetData()))
}
