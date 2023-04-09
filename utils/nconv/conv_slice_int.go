/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:47:12
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 23:06:38
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_int.go
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

// ToInts
func ToInts(val any) (s []int) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case []string:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []int:
		s = value
	case []int8:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []int16:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []int32:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []int64:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []uint:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]int, len(value))
			for k, v := range value {
				s[k] = ToInt(v)
			}
		}
	case []uint16:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []uint32:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []uint64:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []bool:
		s = make([]int, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []float64:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case []any:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	case [][]byte:
		s = make([]int, len(value))
		for k, v := range value {
			s[k] = ToInt(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iInts); ok {
		return v.Ints()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToInts(v.Interfaces())
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
		s = make([]int, length)
		for i := 0; i < length; i++ {
			s[i] = ToInt(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return

	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int{}
		}
		return []int{ToInt(val)}
	}
}

// ToInt32s
func ToInt32s(val any) (s []int32) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case []string:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []int:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []int8:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []int16:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []int32:
		s = value
	case []int64:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []uint:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]int32, len(value))
			for k, v := range value {
				s[k] = ToInt32(v)
			}
		}
	case []uint16:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []uint32:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []uint64:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []bool:
		s = make([]int32, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []float64:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case []any:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	case [][]byte:
		s = make([]int32, len(value))
		for k, v := range value {
			s[k] = ToInt32(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iInts); ok {
		return ToInt32s(v.Ints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToInt32s(v.Interfaces())
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
		s = make([]int32, length)
		for i := 0; i < length; i++ {
			s[i] = ToInt32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int32{}
		}
		return []int32{ToInt32(val)}
	}
}

// ToInt64s
func ToInt64s(val any) (s []int64) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case []string:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []int:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []int8:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []int16:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []int32:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []int64:
		s = value
	case []uint:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]int64, len(value))
			for k, v := range value {
				s[k] = ToInt64(v)
			}
		}
	case []uint16:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []uint32:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []uint64:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []bool:
		s = make([]int64, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []float64:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case []interface{}:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	case [][]byte:
		s = make([]int64, len(value))
		for k, v := range value {
			s[k] = ToInt64(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iInts); ok {
		return ToInt64s(v.Ints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToInt64s(v.Interfaces())
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
		s = make([]int64, length)
		for i := 0; i < length; i++ {
			s[i] = ToInt64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []int64{}
		}
		return []int64{ToInt64(val)}
	}
}
