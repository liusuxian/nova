/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:59:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-12 20:54:00
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/utils/utils_str.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils

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

// IsLetterUpper 检查给定的字节 b 是否为大写字母
func IsLetterUpper(b byte) (isLetterUpper bool) {
	return b >= byte('A') && b <= byte('Z')
}

// IsLetterUpper 检查给定的字节 b 是否为小写字母
func IsLetterLower(b byte) (isLetterLower bool) {
	return b >= byte('a') && b <= byte('z')
}

// IsLetter 检查给定的字节 b 是否是一个字母
func IsLetter(b byte) (isLetter bool) {
	return IsLetterUpper(b) || IsLetterLower(b)
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

// RemoveSymbols 从字符串中删除所有符号，只留下数字和字母
func RemoveSymbols(str string) (newStr string) {
	b := make([]rune, 0, len(str))
	for _, c := range str {
		if c > 127 {
			b = append(b, c)
		} else if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			b = append(b, c)
		}
	}
	return string(b)
}

// Trim 用于去除字符串开头和结尾的空白字符(或其他字符)，可选参数 characterMask 指定了需要去除的其他字符
func Trim(str string, characterMask ...string) (newStr string) {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// SplitAndTrim 将字符串 str 按照分隔符 delimiter 分割成一个字符串数组，并对数组中的每个元素进行 Trim 操作。在 Trim 操作后，忽略那些变成空字符串的元素
func SplitAndTrim(str, delimiter string, characterMask ...string) (strs []string) {
	strs = make([]string, 0)
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			strs = append(strs, v)
		}
	}
	return
}
