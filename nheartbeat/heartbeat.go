/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 19:28:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 21:34:48
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nheartbeat/heartbeat.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nheartbeat

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nrouter"
	"time"
)

type HeartbeatChecker struct {
	interval         time.Duration           // 心跳检测时间间隔
	quitChan         chan bool               // 退出信号
	hearbeatMsg      []byte                  // 心跳消息，也可以通过 makeMsgFunc 来动态生成
	makeMsg          niface.HeartBeatMsgFunc // 用户自定义的心跳检测消息处理方法
	onRemoteNotAlive niface.OnRemoteNotAlive // 用户自定义的远程连接不存活时的处理方法
	msgID            uint32                  // 用户自定义的心跳检测消息ID
	router           niface.IRouter          // 用户自定义的心跳检测消息业务处理路由
	server           niface.IServer          // 所属服务端
	client           niface.IClient          // 所属客户端
}

type HeartbeatDefaultRouter struct {
	nrouter.BaseRouter
}
