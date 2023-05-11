/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 12:27:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:25:00
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nslice

import (
	"github.com/liusuxian/nova/utils/ntype"
	"sort"
)

// IsContains 判断切片 s 中是否包含 target 元素
func IsContains[T ntype.Number | ntype.String | byte](s []T, target T) (isContains bool) {
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
