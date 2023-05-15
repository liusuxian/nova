/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:07:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 15:51:05
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
	nlog.Debug("Receive HeartBeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("ReqMsg", request.GetData()))
	if err := request.RespMsg(func() (buf []byte, err error) {
		return []byte("pong"), nil
	}, func(c niface.Conn, e error) (err error) {
		if e != nil {
			nlog.Error("Send HeartBeat Callback Error", nlog.Err(err))
			return e
		}

		nlog.Debug("Send HeartBeat Callback")
		return nil
	}); err != nil {
		nlog.Error("Send HeartBeat Error", nlog.Err(err))
		return
	}
}
