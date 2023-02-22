/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 21:13:41
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 19:02:09
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/packfactory.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"sync"
)

// 初始化单例资源
var packOnce sync.Once

// 封包拆包方式结构
type packFactory struct {
}

// 封包拆包方式实例
var packFactoryInstance *packFactory

// 生成不同封包拆包的方式
func Factory() *packFactory {
	packOnce.Do(func() {
		packFactoryInstance = new(packFactory)
	})
	return packFactoryInstance
}

// NewPack 创建一个具体的封包拆包对象
func (pf *packFactory) NewPack() (dataPack niface.IDataPack) {
	packetMethod := nconf.GetUint8("server.packetMethod")
	switch packetMethod {
	case 1:
		// 默认封包拆包方式
		// 消息ID-消息体长度-消息内容
		dataPack = NewDefaultPack()
	default:
		dataPack = NewDefaultPack()
	}
	return
}
