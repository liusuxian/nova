/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:03:24
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 15:41:54
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_str.go
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

// ToStrings 将 any 转换为 []string 类型
func ToStrings(val any) (s []string) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case []int:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []int8:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []int16:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []int32:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []int64:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []uint:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]string, len(value))
			for k, v := range value {
				s[k] = ToString(v)
			}
		}
	case []uint16:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []uint32:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []uint64:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []bool:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []float32:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []float64:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []any:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	case []string:
		s = value
	case [][]byte:
		s = make([]string, len(value))
		for k, v := range value {
			s[k] = ToString(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iStrings); ok {
		return v.Strings()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToStrings(v.Interfaces())
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
		s = make([]string, length)
		for i := 0; i < length; i++ {
			s[i] = ToString(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []string{}
		}
		return []string{ToString(val)}
	}
}
