/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 20:13:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-09 20:14:00
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/simple_tcp_demo/client/serveroverload/serveroverload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package serveroverload

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
)

// ServerOverloadHandler 服务器人数超载消息
func ServerOverloadHandler(request niface.IRequest) {
	nlog.Debug("Receive ServerOverload", nlog.String("From", request.GetConnection().RemoteAddr().String()), nlog.Uint16("MsgID", request.GetMsgID()), nlog.ByteString("ReqMsg", request.GetData()))
}
