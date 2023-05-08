/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:33:54
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-08 21:15:45
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/datapack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

// IDataPack 数据包接口
type IDataPack interface {
	GetHeadLen() (headLen int)                           // 获取包头长度(字节数)
	Pack(msg IMessage) (data []byte)                     // 封包
	UnPackHead(headBuf []byte) (msg IMessage, err error) // 拆包头
	UnPackBody(msgBuf []byte, msg IMessage)              // 拆包体
}
