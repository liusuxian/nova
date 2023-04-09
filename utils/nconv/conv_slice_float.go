/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:16:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 22:42:05
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
func ToFloat32s(val any) (cVals []float32) {
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
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []int:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []int8:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []int16:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []int32:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []int64:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []uint:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &cVals)
		} else {
			cVals = make([]float32, len(value))
			for k, v := range value {
				cVals[k] = ToFloat32(v)
			}
		}
	case []uint16:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []uint32:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []uint64:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []bool:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []float32:
		cVals = value
	case []float64:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	case []any:
		cVals = make([]float32, len(value))
		for k, v := range value {
			cVals[k] = ToFloat32(v)
		}
	}
	if cVals != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return ToFloat32s(v.Floats())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat32s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &cVals) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		cVals = make([]float32, length)
		for i := 0; i < length; i++ {
			cVals[i] = ToFloat32(originValueAndKind.OriginValue.Index(i).Interface())
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
func ToFloat64s(val any) (cVals []float64) {
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
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []int:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []int8:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []int16:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []int32:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []int64:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []uint:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &cVals)
		} else {
			cVals = make([]float64, len(value))
			for k, v := range value {
				cVals[k] = ToFloat64(v)
			}
		}
	case []uint16:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []uint32:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []uint64:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []bool:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []float32:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	case []float64:
		cVals = value
	case []any:
		cVals = make([]float64, len(value))
		for k, v := range value {
			cVals[k] = ToFloat64(v)
		}
	}
	if cVals != nil {
		return
	}
	if v, ok := val.(iFloats); ok {
		return v.Floats()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToFloat64s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &cVals) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		cVals = make([]float64, length)
		for i := 0; i < length; i++ {
			cVals[i] = ToFloat64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []float64{}
		}
		return []float64{ToFloat64(val)}
	}
}
