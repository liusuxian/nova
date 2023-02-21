/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 21:13:41
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 22:02:27
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/packfactory.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack

import (
	"github.com/liusuxian/nova/niface"
	"sync"
)

var packOnce sync.Once

type packFactory struct{}

var factoryInstance *packFactory

// 生成不同封包解包的方式，单例
func Factory() *packFactory {
	packOnce.Do(func() {
		factoryInstance = new(packFactory)
	})
	return factoryInstance
}

// NewPack 创建一个具体的拆包解包对象
func (pf *packFactory) NewPack(kind string) niface.IDataPack {
	var dataPack niface.IDataPack
	switch kind {
	// 标准默认封包拆包方式
	case niface.DefaultDataPack:
		dataPack = NewDataPack()
		// case 自定义封包拆包方式case
	default:
		dataPack = NewDataPack()
	}
	return dataPack
}
