/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-02 21:03:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-06-06 19:59:13
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nmsghandler

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/ninterceptor"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nrouter"
	"github.com/panjf2000/ants/v2"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	workerPool             *ants.Pool                 // Worker 工作池
	workerPoolSize         int                        // Worker 池的最大 Worker 数量
	workerPoolSizeOverflow int                        // 当处理任务超过工作任务池的容量时，增加的 Goroutine 数量
	builder                *ninterceptor.ChainBuilder // 责任链构造器
	router                 *nrouter.Router            // 路由
}

// NewMsgHandle 创建消息处理
func NewMsgHandle(workerPoolSize, workerPoolSizeOverflow int) (handler *MsgHandle) {
	handler = &MsgHandle{
		workerPoolSize:         workerPoolSize,
		workerPoolSizeOverflow: workerPoolSizeOverflow,
		builder:                ninterceptor.NewChainBuilder(),
		router:                 nrouter.NewRouter(),
	}
	// 此处必须把 msghandler 添加到责任链中，并且是责任链最后一环
	handler.builder.Tail(handler)
	return
}

// AddRouter 为消息添加路由处理函数集合
func (mh *MsgHandle) AddRouter(msgID uint16, handlers ...niface.RouterHandler) (router niface.IRouter) {
	mh.router.AddHandler(msgID, handlers...)
	return mh.router
}

// Group 路由组管理
func (mh *MsgHandle) Group(startMsgID, endMsgID uint16, handlers ...niface.RouterHandler) (group niface.IGroupRouter) {
	return nrouter.NewGroupRouter(startMsgID, endMsgID, mh.router, handlers...)
}

// Use 添加全局路由
func (mh *MsgHandle) Use(handlers ...niface.RouterHandler) (router niface.IRouter) {
	mh.router.Use(handlers...)
	return mh.router
}

// PrintRouters 打印所有路由
func (mh *MsgHandle) PrintRouters() {
	mh.router.PrintRouters()
}

// StartWorkerPool 启动 Worker 工作池
func (mh *MsgHandle) StartWorkerPool() {
	if mh.workerPool == nil && mh.workerPoolSize > 0 {
		workerPool, err := ants.NewPool(mh.workerPoolSize, ants.WithNonblocking(true))
		if err != nil {
			nlog.Fatal("StartWorkerPool Fatal", nlog.Err(err))
		}
		mh.workerPool = workerPool
		nlog.Info("StartWorkerPool Succeed", nlog.Int("WorkerPoolSize", mh.workerPoolSize))
	}
}

// StopWorkerPool 停止 Worker 工作池
func (mh *MsgHandle) StopWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Release()
		nlog.Info("StopWorkerPool Succeed")
	}
}

// SubmitToWorkerPool 将请求提交给 Worker 工作池处理
func (mh *MsgHandle) SubmitToWorkerPool(request niface.IRequest) {
	if mh.workerPool != nil && request != nil {
		switch iRequest := request.(type) {
		case niface.IFuncRequest:
			if err := mh.workerPool.Submit(func() { mh.doFuncHandler(iRequest) }); err != nil {
				switch err {
				case ants.ErrPoolOverload:
					nlog.Warn("SubmitToWorkerPool IFuncRequest", nlog.Int("WorkerPoolSize", mh.workerPool.Cap()), nlog.Err(err))
					mh.workerPool.Tune(mh.workerPool.Cap() + mh.workerPoolSizeOverflow)
					mh.SubmitToWorkerPool(request)
				default:
					nlog.Error("SubmitToWorkerPool IFuncRequest Error", nlog.Int("WorkerPoolSize", mh.workerPool.Cap()), nlog.Err(err))
					go mh.doFuncHandler(iRequest)
				}
			}
		case niface.IRequest:
			if err := mh.workerPool.Submit(func() { mh.doMsgHandler(iRequest) }); err != nil {
				switch err {
				case ants.ErrPoolOverload:
					nlog.Warn("SubmitToWorkerPool IRequest", nlog.Int("WorkerPoolSize", mh.workerPool.Cap()), nlog.Err(err))
					mh.workerPool.Tune(mh.workerPool.Cap() + mh.workerPoolSizeOverflow)
					mh.SubmitToWorkerPool(request)
				default:
					nlog.Error("SubmitToWorkerPool IRequest Error", nlog.Int("WorkerPoolSize", mh.workerPool.Cap()), nlog.Err(err))
					go mh.doMsgHandler(iRequest)
				}
			}
		}
	}
}

// AddInterceptor 注册责任链任务入口，每个拦截器处理完后，数据都会传递至下一个拦截器，使得消息可以层层处理层层传递，顺序取决于注册顺序
func (mh *MsgHandle) AddInterceptor(interceptor niface.IInterceptor) {
	if mh.builder != nil {
		mh.builder.AddInterceptor(interceptor)
	}
}

// Execute 执行责任链上的拦截器方法
func (mh *MsgHandle) Execute(request niface.IRequest) {
	mh.builder.Execute(request)
}

// Intercept 默认必经的数据处理拦截器
func (mh *MsgHandle) Intercept(chain niface.IChain) (resp niface.IcResp) {
	request := chain.Request()
	if request != nil {
		switch iRequest := request.(type) {
		case niface.IRequest:
			if mh.workerPool != nil {
				mh.SubmitToWorkerPool(iRequest)
			} else {
				go mh.doMsgHandler(iRequest)
			}
		}
	}

	return chain.Proceed(chain.Request())
}

// doMsgHandler 处理消息
func (mh *MsgHandle) doMsgHandler(request niface.IRequest) {
	defer func() {
		if err := recover(); err != nil {
			nlog.Error("doMsgHandler Panic", nlog.Any("Panic", err))
		}
	}()

	msgID := request.GetMsgID()
	handlers, isExist := mh.router.GetHandlers(msgID)
	if !isExist {
		nlog.Error("Api MsgID Is Not Found !!!", nlog.Uint16("MsgID", msgID))
		return
	}

	request.BindRouter(handlers)
	request.RouterNext()
}

// doFuncHandler 执行函数式请求
func (mh *MsgHandle) doFuncHandler(request niface.IFuncRequest) {
	defer func() {
		if err := recover(); err != nil {
			nlog.Error("doFuncHandler Panic", nlog.Any("Panic", err))
		}
	}()
	// 执行函数式请求
	request.CallFunc()
}
