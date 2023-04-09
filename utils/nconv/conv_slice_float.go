/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:16:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 22:45:59
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
func ToFloat32s(val any) (fs []float32) {
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
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []int:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []int8:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []int16:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []int32:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []int64:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []uint:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &fs)
		} else {
			fs = make([]float32, len(value))
			for k, v := range value {
				fs[k] = ToFloat32(v)
			}
		}
	case []uint16:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []uint32:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []uint64:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []bool:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []float32:
		fs = value
	case []float64:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	case []any:
		fs = make([]float32, len(value))
		for k, v := range value {
			fs[k] = ToFloat32(v)
		}
	}
	if fs != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return ToFloat32s(v.Floats())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat32s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &fs) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		fs = make([]float32, length)
		for i := 0; i < length; i++ {
			fs[i] = ToFloat32(originValueAndKind.OriginValue.Index(i).Interface())
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
func ToFloat64s(val any) (fs []float64) {
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
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []int:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []int8:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []int16:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []int32:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []int64:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []uint:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &fs)
		} else {
			fs = make([]float64, len(value))
			for k, v := range value {
				fs[k] = ToFloat64(v)
			}
		}
	case []uint16:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []uint32:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []uint64:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []bool:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []float32:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	case []float64:
		fs = value
	case []any:
		fs = make([]float64, len(value))
		for k, v := range value {
			fs[k] = ToFloat64(v)
		}
	}
	if fs != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return v.Floats()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat64s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &fs) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		fs = make([]float64, length)
		for i := 0; i < length; i++ {
			fs[i] = ToFloat64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{ToFloat64(val)}
	}
}
