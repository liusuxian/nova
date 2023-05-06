/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 14:06:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 20:13:58
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nrequest/request.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nrequest

import "github.com/liusuxian/nova/niface"

// Request 请求结构
type Request struct {
	niface.BaseRequest                        // 基础请求
	conn               niface.IConnection     // 已经和客户端建立好的连接
	msg                niface.IMessage        // 客户端请求的数据
	icResp             niface.IcResp          // 拦截器返回的数据
	handlers           []niface.RouterHandler // 路由处理函数集合
	index              int                    // 路由处理函数集合索引
}

// NewRequest 创建请求
func NewRequest(conn niface.IConnection, msg niface.IMessage) (req *Request) {
	req = new(Request)
	req.conn = conn
	req.msg = msg
	return
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

// GetResponse 获取解析完后的序列化数据
func (r *Request) GetResponse() (resp niface.IcResp) {
	return r.icResp
}

// SetResponse 设置解析完后的序列化数据
func (r *Request) SetResponse(resp niface.IcResp) {
	r.icResp = resp
}

// BindRouter 绑定这次请求的路由处理函数集合
func (r *Request) BindRouter(handlers []niface.RouterHandler) {
	r.handlers = handlers
	r.index = -1
}

// RouterNext 执行下一个路由处理函数
func (r *Request) RouterNext() {
	r.index++
	for r.index < len(r.handlers) {
		r.handlers[r.index](r)
		r.index++
	}
}
