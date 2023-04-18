/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:29:45
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-18 14:21:21
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/utils.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

// indirect 对给定的值进行多次解引用以达到基本类型（或 nil）
func indirect(i any) (iv any) {
	if i == nil {
		return nil
	}
	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		// 如果不是指针类型，避免创建 reflect.Value
		return i
	}
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// indirectToStringerOrError 通过解引用直到达到基本类型（或 nil）或实现了 fmt.Stringer 或 error 接口的对象
func indirectToStringerOrError(i any) (iv any) {
	if i == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(i)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// toInt 如果 i 或 i 的底层类型是 int，则返回 i 的 int 值
func toInt(i any) (iv int, bl bool) {
	switch i := i.(type) {
	case int:
		return i, true
	case time.Weekday:
		return int(i), true
	case time.Month:
		return int(i), true
	default:
		return 0, false
	}
}

// trimZeroDecimal 删除字符串中末尾的零和小数点
func trimZeroDecimal(s string) (v string) {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}

// convertError 转换错误
func convertError(i any, typ string) (err error) {
	return errors.Errorf("unable to convert %#v of type %T to %s", i, i, typ)
}
