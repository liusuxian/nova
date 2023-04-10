/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-27 22:38:10
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 10:49:18
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nslice/slice.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nslice

import "sort"

// NumberT 自定义数字类型
type NumberT interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// StringT 自定义字符串类型
type StringT interface {
	string | rune
}

// IsContains
func IsContains[T NumberT | StringT | byte](s []T, target T) (isContains bool) {
	if len(s) == 0 {
		return false
	}
	// 先将切片排序
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	// 使用二分查找算法查询
	index := sort.Search(len(s), func(i int) bool {
		return s[i] >= target
	})
	return index < len(s) && s[index] == target
}
