/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 14:55:53
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:46:10
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
)

// ToFloat64SliceE 将 any 转换为 []float64 类型
func ToFloat64SliceE(i any) (iv []float64, err error) {
	if i == nil {
		return []float64{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []int32:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []int16:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []int8:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []int:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []uint64:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []uint32:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []uint16:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]float64, len(anyV))
				for k, v := range anyV {
					floatv, err := ToFloat64E(v)
					if err != nil {
						return []float64{}, convertError(i, "[]float64")
					}
					iv[k] = floatv
				}
				return
			}
		}
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []uint:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []float64:
		return val, nil
	case []float32:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []bool:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case [][]byte:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []string:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case []any:
		iv = make([]float64, len(val))
		for k, v := range val {
			floatv, err := ToFloat64E(v)
			if err != nil {
				return []float64{}, convertError(i, "[]float64")
			}
			iv[k] = floatv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]float64, len(anyV))
				for k, v := range anyV {
					floatv, err := ToFloat64E(v)
					if err != nil {
						return []float64{}, convertError(i, "[]float64")
					}
					iv[k] = floatv
				}
				return
			}
		}
		return []float64{}, convertError(i, "[]float64")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]float64, length)
			for j := 0; j < length; j++ {
				floatv, err := ToFloat64E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []float64{}, convertError(i, "[]float64")
				}
				iv[j] = floatv
			}
			return
		}

		return []float64{}, convertError(i, "[]float64")
	}
}

// ToFloat32SliceE 将 any 转换为 []float32 类型
func ToFloat32SliceE(i any) (iv []float32, err error) {
	if i == nil {
		return []float32{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []int32:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []int16:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []int8:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []int:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []uint64:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []uint32:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []uint16:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]float32, len(anyV))
				for k, v := range anyV {
					floatv, err := ToFloat32E(v)
					if err != nil {
						return []float32{}, convertError(i, "[]float32")
					}
					iv[k] = floatv
				}
				return
			}
		}
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []uint:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []float64:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []float32:
		return val, nil
	case []bool:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case [][]byte:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []string:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case []any:
		iv = make([]float32, len(val))
		for k, v := range val {
			floatv, err := ToFloat32E(v)
			if err != nil {
				return []float32{}, convertError(i, "[]float32")
			}
			iv[k] = floatv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]float32, len(anyV))
				for k, v := range anyV {
					floatv, err := ToFloat32E(v)
					if err != nil {
						return []float32{}, convertError(i, "[]float32")
					}
					iv[k] = floatv
				}
				return
			}
		}
		return []float32{}, convertError(i, "[]float32")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]float32, length)
			for j := 0; j < length; j++ {
				floatv, err := ToFloat32E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []float32{}, convertError(i, "[]float32")
				}
				iv[j] = floatv
			}
			return
		}

		return []float32{}, convertError(i, "[]float32")
	}
}
