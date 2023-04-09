/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:16:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 23:01:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_float.go
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

// ToFloat32s
func ToFloat32s(val any) (s []float32) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		if value == "" {
			return []float32{}
		}
		return []float32{ToFloat32(value)}
	case []string:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []int:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []int8:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []int16:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []int32:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []int64:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []uint:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]float32, len(value))
			for k, v := range value {
				s[k] = ToFloat32(v)
			}
		}
	case []uint16:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []uint32:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []uint64:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []bool:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []float32:
		s = value
	case []float64:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	case []any:
		s = make([]float32, len(value))
		for k, v := range value {
			s[k] = ToFloat32(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return ToFloat32s(v.Floats())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat32s(v.Interfaces())
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
		s = make([]float32, length)
		for i := 0; i < length; i++ {
			s[i] = ToFloat32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float32{}
		}
		return []float32{ToFloat32(val)}
	}
}

// ToFloat64s
func ToFloat64s(val any) (s []float64) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		if value == "" {
			return []float64{}
		}
		return []float64{ToFloat64(value)}
	case []string:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []int:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []int8:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []int16:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []int32:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []int64:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []uint:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]float64, len(value))
			for k, v := range value {
				s[k] = ToFloat64(v)
			}
		}
	case []uint16:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []uint32:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []uint64:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []bool:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []float32:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	case []float64:
		s = value
	case []any:
		s = make([]float64, len(value))
		for k, v := range value {
			s[k] = ToFloat64(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return v.Floats()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat64s(v.Interfaces())
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
		s = make([]float64, length)
		for i := 0; i < length; i++ {
			s[i] = ToFloat64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{ToFloat64(val)}
	}
}
