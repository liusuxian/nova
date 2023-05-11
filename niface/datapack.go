/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-08 18:49:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:15:10
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

// IDataPack 数据包接口
type IDataPack interface {
	GetHeadLen() (headLen int)                           // 获取包头长度(字节数)
	Pack(msg IMessage) (data []byte)                     // 封包
	UnPackHead(headBuf []byte) (msg IMessage, err error) // 拆包头
	UnPackBody(msgBuf []byte, msg IMessage)              // 拆包体
}
