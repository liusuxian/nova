/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:45:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 01:29:37
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nmsghandler/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nmsghandler

import (
	"encoding/hex"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	apis       map[uint32]niface.IRouter // 存放每个 MsgID 所对应的处理方法
	workerPool *ants.Pool                // Worker 工作池
	builder    niface.InterceptorBuilder // 责任链构造器
}

// NewMsgHandle 创建消息处理
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		apis:    make(map[uint32]niface.IRouter),
		builder: ncode.NewInterceptorBuilder(),
	}
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(req niface.IRequest) {
	var handler niface.IRouter
	var ok bool
	if handler, ok = mh.apis[req.GetMsgID()]; !ok {
		nlog.Error(req.GetCtx(), "DoMsgHandler Api Not Found", zap.Uint32("MsgID", req.GetMsgID()))
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
		nlog.Fatal(nil, "AddRouter Repeated Api", zap.Uint32("MsgID", msgID))
	}
	// 添加 msgID 与 API 的绑定关系
	mh.apis[msgID] = router
	nlog.Info(nil, "AddRouter Add Api", zap.Uint32("MsgID", msgID))
}

// StartWorkerPool 启动 Worker 工作池
func (mh *MsgHandle) StartWorkerPool() {
	if mh.workerPool == nil {
		// 此处必须把 msghandler 添加到责任链中，并且是责任链最后一环，在 msghandler 中进行解码后由 router 做数据分发
		mh.AddInterceptor(mh)
		var workerPool *ants.Pool
		var err error
		if workerPool, err = ants.NewPool(int(nconf.WorkerPoolSize())); err != nil {
			nlog.Fatal(nil, "StartWorkerPool Fatal", zap.Error(err))
		}
		mh.workerPool = workerPool
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
func (mh *MsgHandle) SendMsgToWorkerPool(req niface.IRequest) {
	nlog.Debug(req.GetCtx(), "SendMsgToWorkerPool", zap.String("Data", hex.EncodeToString(req.GetData())))
	if mh.workerPool != nil {
		mh.workerPool.Submit(func() {
			mh.DoMsgHandler(req)
		})
	} else {
		go mh.DoMsgHandler(req)
		nlog.Error(req.GetCtx(), "SendMsgToWorkerPool WorkerPool Not Found", zap.Uint32("MsgID", req.GetMsgID()))
	}
}

// Intercept 拦截器
func (mh *MsgHandle) Intercept(chain niface.Chain) niface.Response {
	request := chain.Request()
	if request != nil {
		switch request.(type) {
		case niface.IRequest:
			iRequest := request.(niface.IRequest)
			if mh.workerPool != nil {
				mh.workerPool.Submit(func() {
					mh.DoMsgHandler(iRequest)
				})
			} else {
				go mh.DoMsgHandler(iRequest)
			}
		}
	}
	return chain.Proceed(chain.Request())
}

// Decode 解码
func (mh *MsgHandle) Decode(request niface.IRequest) {
	mh.builder.Execute(request) // 将消息丢到责任链，通过责任链里拦截器层层处理层层传递
}

// AddInterceptor 添加拦截器，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
func (mh *MsgHandle) AddInterceptor(interceptor niface.Interceptor) {
	if mh.builder != nil {
		mh.builder.AddInterceptor(interceptor)
	}
}
