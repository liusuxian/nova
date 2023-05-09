/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:06:13
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-09 20:12:17
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/server/heartbeat/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package heartbeat

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
)

// HeartBeatHandler 心跳消息
func HeartBeatHandler(request niface.IRequest) {
	nlog.Debug("Receive Heartbeat", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("ReqMsg", request.GetData()))
}
