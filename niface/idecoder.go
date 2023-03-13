/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 23:21:33
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 23:55:39
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/idecoder.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "encoding/binary"

// IDecoder 解码器接口
type IDecoder interface {
	Decode(buff []byte) [][]byte
}

// LengthField 长度字段结构
//
//	Order: 大小端排序
//	大端模式: 是指数据的高字节保存在内存的低地址中，而数据的低字节保存在内存的高地址中，地址由小向大增加，而数据从高位往低位放
//	小端模式: 是指数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中，高地址部分权值高，低地址部分权值低，和我们的日常逻辑方法一致
type LengthField struct {
	Order               binary.ByteOrder // 大小端排序
	MaxFrameLength      int64            // 最大帧长度
	LengthFieldOffset   int              // 长度字段偏移量
	LengthFieldLength   int              // 长度字段的字节数
	LengthAdjustment    int              // 长度调整
	InitialBytesToStrip int              // 需要跳过的字节数
}
