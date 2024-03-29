/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 14:06:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-06-22 12:06:11
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nrequest

import (
	"context"
	"github.com/liusuxian/nova/nerr"
	"github.com/liusuxian/nova/niface"
	"time"
)

// Request 请求结构
type Request struct {
	niface.BaseRequest                        // 基础请求
	ctx                context.Context        // 请求的 Context
	conn               niface.IConnection     // 已经和客户端建立好的连接
	createTime         time.Time              // 请求创建的时间
	msg                niface.IMessage        // 客户端请求的数据
	icResp             niface.IcResp          // 拦截器返回的数据
	handlers           []niface.RouterHandler // 业务处理器集合
	index              int                    // 业务处理器集合索引
}

// RequestFunc 请求函数结构
type RequestFunc struct {
	niface.BaseRequest                    // 基础请求
	conn               niface.IConnection // 已经和客户端建立好的连接
	createTime         time.Time          // 请求创建的时间
	callFunc           func()             // 调用函数
}

// NewRequest 创建请求
func NewRequest(conn niface.IConnection, msg niface.IMessage) (req *Request) {
	req = new(Request)
	req.ctx = context.Background()
	req.conn = conn
	req.createTime = time.Now()
	req.msg = msg
	return
}

// NewRequestFunc 创建请求函数
func NewRequestFunc(conn niface.IConnection, callFunc func()) (request niface.IRequest) {
	req := new(RequestFunc)
	req.conn = conn
	req.createTime = time.Now()
	req.callFunc = callFunc
	return req
}

// GetCtx 获取请求的 Context
func (r *Request) GetCtx() (ctx context.Context) {
	return r.ctx
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() (conn niface.IConnection) {
	return r.conn
}

// GetMsgID 获取请求的消息 ID
func (r *Request) GetMsgID() (msgID uint16) {
	return r.msg.GetMsgID()
}

// GetData 获取请求消息的数据
func (r *Request) GetData() (data []byte) {
	return r.msg.GetData()
}

// GetMessage 获取请求消息的原始数据
func (r *Request) GetMessage() (msg niface.IMessage) {
	return r.msg
}

// GetSerializedData 获取解析完后的序列化数据
func (r *Request) GetSerializedData() (resp niface.IcResp) {
	return r.icResp
}

// SetSerializedData 设置解析完后的序列化数据
func (r *Request) SetSerializedData(resp niface.IcResp) {
	r.icResp = resp
}

// BindRouter 绑定这次请求的业务处理器集合
func (r *Request) BindRouter(handlers []niface.RouterHandler) {
	r.handlers = handlers
	r.index = -1
}

// RouterNext 执行下一个业务处理器
func (r *Request) RouterNext() {
	r.index++
	for r.index < len(r.handlers) {
		r.handlers[r.index](r)
		r.index++
	}
}

// GetHandleTime 获取请求处理的时间
func (r *Request) GetHandleTime() (d time.Duration) {
	return time.Since(r.createTime)
}

// Resp 将数据返回给远程的对端
func (r *Request) Resp(f niface.MsgDataFunc, callback ...niface.SendCallback) (err error) {
	if e := r.conn.Send(f, callback...); e != nil && e != nerr.ErrConnectionClosed {
		return e
	}

	return
}

// RespMsg 将 Message 数据返回给远程的对端（与请求共用消息ID）
func (r *Request) RespMsg(f niface.MsgDataFunc, callback ...niface.SendCallback) (err error) {
	return r.RespMsgWithId(r.msg.GetMsgID(), f, callback...)
}

// RespMsgWithId 将 Message 数据返回给远程的对端（与请求可不共用消息ID）
func (r *Request) RespMsgWithId(msgID uint16, f niface.MsgDataFunc, callback ...niface.SendCallback) (err error) {
	if e := r.conn.SendMsg(msgID, f, callback...); e != nil && e != nerr.ErrConnectionClosed {
		return e
	}

	return
}

// CallFunc 调用函数
func (rf *RequestFunc) CallFunc() {
	if rf.callFunc != nil {
		rf.callFunc()
	}
}

// GetConnection 获取请求连接信息
func (rf *RequestFunc) GetConnection() (conn niface.IConnection) {
	return rf.conn
}

// GetHandleTime 获取函数处理的时间
func (rf *RequestFunc) GetHandleTime() (d time.Duration) {
	return time.Since(rf.createTime)
}
