/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:45:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 22:59:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nnet

import (
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	apis       map[uint32]niface.IRouter // 存放每个 MsgID 所对应的处理方法
	workerPool *ants.Pool                // worker 工作池
}

// NewMsgHandle 创建消息处理
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		apis: make(map[uint32]niface.IRouter),
	}
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(req niface.IRequest) {
	var handler niface.IRouter
	var ok bool
	if handler, ok = mh.apis[req.GetMsgID()]; !ok {
		nlog.Error(req.GetCtx(), "DoMsgHandler Api Not Found", zap.Uint32("msgID", req.GetMsgID()))
		return
	}
	// Request 请求绑定 Router
	req.BindRouter(handler)
	// 执行对应处理方法
	req.Call()
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router niface.IRouter) {
	// 判断当前 msgID 绑定的 API 处理方法是否已经存在
	if _, ok := mh.apis[msgID]; ok {
		nlog.Fatal(nil, "AddRouter Repeated Api", zap.Uint32("msgID", msgID))
	}
	// 添加 msgID 与 API 的绑定关系
	mh.apis[msgID] = router
	nlog.Info(nil, "AddRouter Add Api", zap.Uint32("msgID", msgID))
}

// StartWorkerPool 启动 worker 工作池
func (mh *MsgHandle) StartWorkerPool() {
	if mh.workerPool == nil {
		var workerPool *ants.Pool
		var err error
		if workerPool, err = ants.NewPool(nconf.GetInt("server.workerPoolSize")); err != nil {
			nlog.Fatal(nil, "StartWorkerPool Fatal", zap.Error(err))
		}
		mh.workerPool = workerPool
	}
}

// StopWorkerPool 停止 worker 工作池
func (mh *MsgHandle) StopWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Release()
	}
}

// RebootWorkerPool 重启 worker 工作池
func (mh *MsgHandle) RebootWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Reboot()
	}
}

// SendMsgToWorkerPool 将消息交给 WorkerPool，由 worker 进行处理
func (mh *MsgHandle) SendMsgToWorkerPool(req niface.IRequest) {
	if mh.workerPool != nil {
		mh.workerPool.Submit(func() {
			mh.DoMsgHandler(req)
		})
	} else {
		mh.DoMsgHandler(req)
		nlog.Error(req.GetCtx(), "SendMsgToWorkerPool WorkerPool Not Found", zap.Uint32("msgID", req.GetMsgID()))
	}
}
