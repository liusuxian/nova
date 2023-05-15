/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-08 01:20:29
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 14:07:03
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

import "context"

// IFuncRequest 函数消息接口
type IFuncRequest interface {
	CallFunc() // 调用函数
}

// IRequest 请求接口
type IRequest interface {
	GetCtx() (ctx context.Context)                                                   // 获取请求的 Context
	GetConnection() (conn IConnection)                                               // 获取请求连接信息
	GetMsgID() (msgID uint16)                                                        // 获取请求的消息 ID
	GetData() (data []byte)                                                          // 获取请求消息的数据
	GetMessage() (msg IMessage)                                                      // 获取请求消息的原始数据
	GetSerializedData() (resp IcResp)                                                // 获取解析完后的序列化数据
	SetSerializedData(resp IcResp)                                                   // 设置解析完后的序列化数据
	BindRouter(handlers []RouterHandler)                                             // 绑定这次请求的业务处理器集合
	RouterNext()                                                                     // 执行下一个业务处理器
	Resp(f MsgDataFunc, callback ...SendCallback) (err error)                        // 将数据返回给远程的对端
	RespMsg(f MsgDataFunc, callback ...SendCallback) (err error)                     // 将 Message 数据返回给远程的对端（与请求共用消息ID）
	RespMsgWithId(msgID uint16, f MsgDataFunc, callback ...SendCallback) (err error) // 将 Message 数据返回给远程的对端（与请求可不共用消息ID）
}

// BaseRequest 基础请求结构
type BaseRequest struct{}

// GetConnection 获取请求连接信息
func (br *BaseRequest) GetConnection() (conn IConnection) {
	return nil
}

// GetMsgID 获取请求的消息 ID
func (br *BaseRequest) GetMsgID() (msgID uint16) {
	return 0
}

// GetData 获取请求消息的数据
func (br *BaseRequest) GetData() (data []byte) {
	return nil
}

// GetMessage 获取请求消息的原始数据
func (br *BaseRequest) GetMessage() (msg IMessage) {
	return nil
}

// GetSerializedData 获取解析完后的序列化数据
func (br *BaseRequest) GetSerializedData() (resp IcResp) {
	return nil
}

// SetSerializedData 设置解析完后的序列化数据
func (br *BaseRequest) SetSerializedData(resp IcResp) {

}

// BindRouter 绑定这次请求的业务处理器集合
func (br *BaseRequest) BindRouter(handlers []RouterHandler) {

}

// RouterNext 执行下一个业务处理器
func (br *BaseRequest) RouterNext() {
}

// Resp 将数据返回给远程的对端
func (br *BaseRequest) Resp(f MsgDataFunc, callback ...SendCallback) (err error) {
	return nil
}

// RespMsg 将 Message 数据返回给远程的对端（与请求共用消息ID）
func (br *BaseRequest) RespMsg(f MsgDataFunc, callback ...SendCallback) (err error) {
	return nil
}

// RespMsgWithId 将 Message 数据返回给远程的对端（与请求可不共用消息ID）
func (br *BaseRequest) RespMsgWithId(msgID uint16, f MsgDataFunc, callback ...SendCallback) (err error) {
	return nil
}
