/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:59:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 00:11:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/utils/utils_str.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils

// IsNumeric 检查给定字符串 s 是否为数字
// 注意，类似于 "123.456" 的浮点数字符串也是数字
func IsNumeric(s string) (isNumeric bool) {
	dotCount := 0
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] == '-' && i == 0 {
			continue
		}
		if s[i] == '.' {
			dotCount++
			if i > 0 && i < length-1 {
				continue
			} else {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return dotCount <= 1
}
