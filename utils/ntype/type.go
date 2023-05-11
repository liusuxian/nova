/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 12:30:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:25:59
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package ntype

// Number 自定义数字类型
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// String 自定义字符串类型
type String interface {
	string | rune
}
