/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 01:33:54
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 20:17:52
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/datapack.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "github.com/panjf2000/gnet/v2"

// IDataPack 包接口
type IDataPack interface {
	GetHeadLen() uint8                       // 获取包头长度
	Pack(msg IMessage) ([]byte, error)       // 封包
	Unpack(conn gnet.Conn) (IMessage, error) // 拆包
}
