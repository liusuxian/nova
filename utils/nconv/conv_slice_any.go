/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-06 21:19:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 21:58:43
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_any.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
)

// ToSlice
func ToSlice(val any) (s []any) {
	if val == nil {
		return
	}
	switch value := val.(type) {
	case []any:
		s = value
	case []string:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []int:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []int8:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []int16:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []int32:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []int64:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []uint:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]any, len(value))
			for k, v := range value {
				s[k] = v
			}
		}
	case []uint16:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []uint32:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []uint64:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []bool:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []float32:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	case []float64:
		s = make([]any, len(value))
		for k, v := range value {
			s[k] = v
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iInterfaces); ok {
		return v.Interfaces()
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]any, length)
		for i := 0; i < length; i++ {
			s[i] = originValueAndKind.OriginValue.Index(i).Interface()
		}
		return
	default:
		s = []any{val}
		return
	}
}
