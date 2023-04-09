/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:17:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 23:30:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nstr/str.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nstr

import "strings"

// DefaultTrimChars 是默认情况下被 Trim* 函数剥离的字符
var DefaultTrimChars = string([]byte{
	'\t', // 制表符
	'\v', // 垂直制表符
	'\n', // 换行符
	'\r', // 回车符
	'\f', // 换页符
	' ',  // 普通空格
	0x00, // NULL字符
	0x85, // 删除符
	0xA0, // 不换行空格
})

// TrimAll 删除字符串 str 中的所有字符
func TrimAll(str string, characterMask ...string) (newStr string) {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	var filtered bool
	var slice = make([]rune, 0, len(str))
	for _, char := range str {
		filtered = false
		for _, trimChar := range trimChars {
			if char == trimChar {
				filtered = true
				break
			}
		}
		if !filtered {
			slice = append(slice, char)
		}
	}
	return string(slice)
}

// Split 用 sep 将 str 拆分为 []string
func Split(str, sep string) (list []string) {
	str = TrimAll(str)
	if str == "" {
		return []string{}
	}
	list = strings.Split(str, sep)
	return
}

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
