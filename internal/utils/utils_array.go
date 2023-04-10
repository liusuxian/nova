/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 17:43:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 17:43:31
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/utils/utils_array.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils

import "reflect"

// IsArray 检查给定的值是否为数组/切片
//
//	注意，它使用 reflect 内部实现此功能
func IsArray(value any) (isArray bool) {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}
