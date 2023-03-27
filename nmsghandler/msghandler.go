/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 20:45:01
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-27 11:20:00
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nmsghandler/msghandler.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nmsghandler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/olekukonko/tablewriter"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// MsgHandle 消息处理回调结构
type MsgHandle struct {
	ctx            context.Context           // 消息处理的 Context
	apis           map[uint16]niface.IRouter // 存放每个 MsgID 所对应的处理方法
	workerPool     *ants.Pool                // Worker 工作池
	workerPoolSize int                       // Worker 池的最大 Worker 数量
}

// NewMsgHandle 创建消息处理
func NewMsgHandle(workerPoolSize int) *MsgHandle {
	return &MsgHandle{
		ctx:            context.Background(),
		apis:           make(map[uint16]niface.IRouter),
		workerPoolSize: workerPoolSize,
	}
}

// HandleRequest 处理请求消息
func (mh *MsgHandle) HandleRequest(request niface.IRequest) {
	if mh.workerPool != nil {
		mh.workerPool.Submit(func() {
			mh.doRequest(request)
		})
	} else {
		go mh.doRequest(request)
	}
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint16, router niface.IRouter) {
	// 判断当前 msgID 绑定的 API 处理方法是否已经存在
	if _, ok := mh.apis[msgID]; ok {
		nlog.Fatal(mh.ctx, "AddRouter Repeated Api", zap.Uint16("MsgID", msgID))
	}
	// 添加 msgID 与 API 的绑定关系
	mh.apis[msgID] = router
}

// PrintRouters 打印所有路由
func (mh *MsgHandle) PrintRouters() {
	routerNum := len(mh.apis)
	if routerNum == 0 {
		return
	}
	// 组装打印数据
	printData := make([][]string, 0, routerNum)
	for msgID, router := range mh.apis {
		rowData := make([]string, 0, 3)
		// msgID
		rowData = append(rowData, strconv.FormatInt(int64(msgID), 10))
		// 获取 Router 的类型信息
		t := reflect.TypeOf(router)
		// Router
		rowData = append(rowData, t.Elem().String())
		// handler
		handlerData := strings.Builder{}
		for i := t.NumMethod() - 1; i >= 0; i-- {
			handlerData.WriteString(fmt.Sprintf("%s.(*%s).%s\n", t.Elem().PkgPath(), t.Elem().Name(), t.Method(i).Name))
		}
		rowData = append(rowData, handlerData.String())
		printData = append(printData, rowData)
	}
	// 打印数据
	tableData := &bytes.Buffer{}
	table := tablewriter.NewWriter(io.MultiWriter(os.Stdout, tableData))
	table.SetHeader([]string{"MSGID", "ROUTER", "HANDLER"})
	table.SetCaption(true, time.Now().Local().String())
	for _, v := range printData {
		table.Append(v)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetRowLine(true)
	}
	table.Render()
	os.Stdout.Sync()
	// 输出到日志文件
	// 打开文件，如果不存在则创建
	logFileName := fmt.Sprintf("%s/router-%s.log", nlog.GetLoggerPath(), time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(errors.Wrapf(err, "PrintRouters OpenFile[%s] Error", logFileName))
	}
	defer f.Close()
	// 写入文件
	if _, err := io.WriteString(f, tableData.String()); err != nil {
		panic(errors.Wrapf(err, "PrintRouters WriteFile[%s] Error", logFileName))
	}
}

// StartWorkerPool 启动 Worker 工作池
func (mh *MsgHandle) StartWorkerPool() {
	if mh.workerPool == nil && mh.workerPoolSize > 0 {
		workerPool, err := ants.NewPool(mh.workerPoolSize)
		if err != nil {
			nlog.Fatal(mh.ctx, "StartWorkerPool Fatal", zap.Error(err))
		}
		mh.workerPool = workerPool
		nlog.Info(mh.ctx, "StartWorkerPool Succeed", zap.Int("WorkerPoolSize", mh.workerPoolSize))
	}
}

// StopWorkerPool 停止 Worker 工作池
func (mh *MsgHandle) StopWorkerPool() {
	if mh.workerPool != nil {
		mh.workerPool.Release()
		nlog.Info(mh.ctx, "StopWorkerPool Succeed")
	}
}

// doRequest 处理请求
func (mh *MsgHandle) doRequest(request niface.IRequest) {
	handler, ok := mh.apis[request.GetMsgID()]
	if !ok {
		nlog.Error(request.GetCtx(), "HandlerMsg Api Not Found", zap.Uint16("MsgID", request.GetMsgID()))
		return
	}
	// Request 请求绑定 Router
	request.BindRouter(handler)
	// 执行对应处理方法
	request.Call()
}
