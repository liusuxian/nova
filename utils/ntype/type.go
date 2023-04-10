/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 12:38:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 12:38:48
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/ntype/type.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
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
