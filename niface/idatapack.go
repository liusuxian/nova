/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:33:54
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 21:18:20
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/idatapack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IDataPack 包接口
type IDataPack interface {
	GetHeadLen() uint32                // 获取包头长度
	Pack(msg IMessage) ([]byte, error) // 封包
	Unpack([]byte) (IMessage, error)   // 拆包
}

const (
	// 默认封包和拆包方式
	DefaultDataPack string = "default_pack"
	// 自定义封包方式在此添加
)

const (
	// 默认报文协议格式
	DefaultMessage string = "default_message"
)
