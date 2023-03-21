/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:45:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-21 20:47:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nmsghandler/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nmsghandler

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	ctx            context.Context           // 当前 Server 的根 Context
	apis           map[uint16]niface.IRouter // 存放每个 MsgID 所对应的处理方法
	workerPool     *ants.Pool                // Worker 工作池
	workerPoolSize int                       // Worker 池的最大 Worker 数量
}

// NewMsgHandle 创建消息处理
func NewMsgHandle(ctx context.Context, workerPoolSize int) *MsgHandle {
	return &MsgHandle{
		ctx:            ctx,
		apis:           make(map[uint16]niface.IRouter),
		workerPoolSize: workerPoolSize,
	}
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request niface.IRequest) {
	handler, ok := mh.apis[request.GetMsgID()]
	if !ok {
		nlog.Error(request.GetCtx(), "DoMsgHandler Api Not Found", zap.Uint16("MsgID", request.GetMsgID()))
		return
	}
	// Request 请求绑定 Router
	request.BindRouter(handler)
	// 执行对应处理方法
	request.Call()
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint16, router niface.IRouter) {
	// 判断当前 msgID 绑定的 API 处理方法是否已经存在
	if _, ok := mh.apis[msgID]; ok {
		nlog.Fatal(mh.ctx, "AddRouter Repeated Api", zap.Uint16("MsgID", msgID))
	}
	// 添加 msgID 与 API 的绑定关系
	mh.apis[msgID] = router
	nlog.Info(mh.ctx, "AddRouter Add Api", zap.Uint16("MsgID", msgID))
}

// StartWorkerPool 启动 Worker 工作池
func (mh *MsgHandle) StartWorkerPool() {
	if mh.workerPool == nil {
		workerPool, err := ants.NewPool(mh.workerPoolSize)
		if err != nil {
			nlog.Fatal(mh.ctx, "StartWorkerPool Fatal", zap.Error(err))
		}
		mh.workerPool = workerPool
		nlog.Info(mh.ctx, "StartWorkerPool Succeed")
	}
}

// StopWorkerPool 停止 Worker 工作池
func (mh *MsgHandle) StopWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Release()
	}
}

// RebootWorkerPool 重启 Worker 工作池
func (mh *MsgHandle) RebootWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Reboot()
	}
}

// SendMsgToWorkerPool 将消息交给 WorkerPool，由 Worker 进行处理
func (mh *MsgHandle) SendMsgToWorkerPool(request niface.IRequest) {
	if mh.workerPool != nil {
		mh.workerPool.Submit(func() {
			mh.DoMsgHandler(request)
		})
	} else {
		go mh.DoMsgHandler(request)
		nlog.Error(request.GetCtx(), "SendMsgToWorkerPool WorkerPool Not Found", zap.Uint16("MsgID", request.GetMsgID()))
	}
}
