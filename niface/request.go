/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 21:35:02
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 18:29:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/request.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IFuncRequest 函数消息接口
type IFuncRequest interface {
	CallFunc() // 调用函数
}

// IRequest 请求接口
type IRequest interface {
	GetConnection() (conn IConnection)   // 获取请求连接信息
	GetMsgID() (msgID uint16)            // 获取请求的消息 ID
	GetData() (data []byte)              // 获取请求消息的数据
	GetMessage() (msg IMessage)          // 获取请求消息的原始数据
	GetResponse() (resp IcResp)          // 获取解析完后的序列化数据
	SetResponse(resp IcResp)             // 设置解析完后的序列化数据
	BindRouter(handlers []RouterHandler) // 绑定这次请求的路由处理函数集合
	RouterNext()                         // 执行下一个路由处理函数
}

// BaseRequest 基础请求结构
type BaseRequest struct{}

// GetConnection 获取请求连接信息
func (br *BaseRequest) GetConnection() (conn IConnection) { return nil }

// GetMsgID 获取请求的消息 ID
func (br *BaseRequest) GetMsgID() (msgID uint16) { return 0 }

// GetData 获取请求消息的数据
func (br *BaseRequest) GetData() (data []byte) { return nil }

// GetMessage 获取请求消息的原始数据
func (br *BaseRequest) GetMessage() (msg IMessage) { return nil }

// GetResponse 获取解析完后的序列化数据
func (br *BaseRequest) GetResponse() (resp IcResp) { return nil }

// SetResponse 设置解析完后的序列化数据
func (br *BaseRequest) SetResponse(resp IcResp) {}

// BindRouter 绑定这次请求的路由处理函数集合
func (br *BaseRequest) BindRouter(handlers []RouterHandler) {}

// RouterNext 执行下一个路由处理函数
func (br *BaseRequest) RouterNext() {}
