/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:45:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-14 15:03:03
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
	"go.uber.org/zap"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	apis           map[uint32]niface.IRouter // 存放每个 MsgID 所对应的处理方法
	workerPoolSize uint32                    // 业务工作 Worker 池的数量
	taskQueue      []chan niface.IRequest    // Worker 负责取任务的消息队列
	builder        niface.InterceptorBuilder // 责任链构造器
}

// NewMsgHandle 创建消息处理
func NewMsgHandle() *MsgHandle {
	workerPoolSize := nconf.WorkerPoolSize()
	return &MsgHandle{
		apis:           make(map[uint32]niface.IRouter),
		workerPoolSize: workerPoolSize,
		taskQueue:      make([]chan niface.IRequest, workerPoolSize), // 一个 Worker 对应一个 Queue
		builder:        ncode.NewInterceptorBuilder(),
	}
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request niface.IRequest) {
	handler, ok := mh.apis[request.GetMsgID()]
	if !ok {
		nlog.Error(request.GetCtx(), "DoMsgHandler Api Not Found", zap.Uint32("MsgID", request.GetMsgID()))
		return
	}
	// Request 请求绑定 Router
	request.BindRouter(handler)
	// 执行对应处理方法
	request.Call()
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
	// 此处必须把 Msghandler 添加到责任链中，并且是责任链最后一环，在 Msghandler 中进行解码后由 Router 做数据分发
	mh.AddInterceptor(mh)
	// 遍历需要启动 Worker 的数量，依此启动
	for i := 0; i < int(mh.workerPoolSize); i++ {
		// 给当前 Worker 对应的任务队列开辟空间
		mh.taskQueue[i] = make(chan niface.IRequest, nconf.MaxWorkerTaskLen())
		// 启动当前 Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.taskQueue[i])
	}
}

// StartOneWorker 启动一个 Worker
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan niface.IRequest) {
	nlog.Info(nil, "Worker Is Started", zap.Int("WorkerID", workerID))
	for {
		select {
		// 有消息则取出队列的 Request，并执行绑定的业务方法
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

// SendMsgToTaskQueue 将消息交给 TaskQueue，由 Worker 进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request niface.IRequest) {
	// 根据 ConnID 来分配当前的连接应该由哪个 Worker 负责处理
	// 轮询的平均分配法则，得到需要处理当前连接的 WorkerID
	workerID := request.GetConnection().GetConnID() % uint64(mh.workerPoolSize)
	// 将请求消息发送给任务队列
	mh.taskQueue[workerID] <- request
	nlog.Debug(request.GetCtx(), "SendMsgToTaskQueue", zap.Uint64("WorkerID", workerID), zap.String("Data", hex.EncodeToString(request.GetData())))
}

// Intercept 拦截并处理
func (mh *MsgHandle) Intercept(chain niface.Chain) niface.Response {
	request := chain.Request()
	if request != nil {
		switch request.(type) {
		case niface.IRequest:
			iRequest := request.(niface.IRequest)
			if mh.workerPoolSize > 0 {
				// 已经启动工作池机制，将消息交给 Worker 处理
				mh.SendMsgToTaskQueue(iRequest)
			} else {
				// 从绑定好的消息和对应的处理方法中执行对应的 Handle 方法
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
