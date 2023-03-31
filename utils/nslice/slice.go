/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-27 22:38:10
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-27 23:09:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nslice/slice.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nslice

import "sort"

// ContainsInt
func ContainsInt(slice []int, target int) bool {
	if len(slice) == 0 {
		return false
	}
	// 先将切片排序
	sort.Ints(slice)
	// 查找
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})
	return index < len(slice) && slice[index] == target
}