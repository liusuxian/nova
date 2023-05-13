/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-01 17:52:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-12 12:41:36
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nserveroverload

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/npack"
)

// ServerOverloadChecker 服务器人数超载检测器结构
type ServerOverloadChecker struct {
	makeMsg niface.ServerOverloadMsgFunc // 用户自定义的服务器人数超载消息处理方法
	msgID   uint16                       // 用户自定义的服务器人数超载消息ID
}

// NewServerOverloadChecker 创建服务器人数超载检测器
func NewServerOverloadChecker() (checker niface.IServerOverloadChecker) {
	serverOverload := &ServerOverloadChecker{
		makeMsg: makeMsgDefaultFunc,
		msgID:   niface.ServerOverloadDefaultMsgID,
	}
	return serverOverload
}

// Check 服务器人数超载检测
func (soc *ServerOverloadChecker) Check(server niface.IServer, maxConn int) (isOverload bool) {
	return server.GetConnections() > maxConn
}

// SetServerOverloadMsgFunc 设置服务器人数超载消息处理方法
func (soc *ServerOverloadChecker) SetServerOverloadMsgFunc(f niface.ServerOverloadMsgFunc) {
	if f != nil {
		soc.makeMsg = f
	}
}

// SetMsgID 设置服务器人数超载消息ID
func (soc *ServerOverloadChecker) SetMsgID(msgID uint16) {
	if msgID != niface.ServerOverloadDefaultMsgID {
		soc.msgID = msgID
	}
}

// GetMessage 获取服务器人数超载消息
func (soc *ServerOverloadChecker) GetMessage() (msg niface.IMessage, err error) {
	var buf []byte
	if buf, err = soc.makeMsg(); err != nil {
		return
	}

	msg = npack.NewMsgPackage(soc.msgID, buf)
	return
}

// makeMsgDefaultFunc 默认的服务器人数超载消息处理方法
func makeMsgDefaultFunc() (buf []byte, err error) {
	return []byte("server overload"), nil
}
