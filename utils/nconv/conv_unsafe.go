/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 22:52:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 22:57:37
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_unsafe.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "unsafe"

// UnsafeStrToBytes 将 string 转换为 []byte，无需内存复制
//
//	注意，如果你完全确定在未来永远不会使用`s`变量，则可以使用此不安全的函数以实现高性能的类型转换
func UnsafeStrToBytes(s string) (bs []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeBytesToStr 将 []byte 转换为 string，无需内存复制
//
//	注意，如果你完全确定在未来永远不会使用`bs`变量，则可以使用此不安全的函数以实现高性能的类型转换
func UnsafeBytesToStr(bs []byte) (s string) {
	return *(*string)(unsafe.Pointer(&bs))
}
