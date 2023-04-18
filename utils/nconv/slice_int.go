/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:26:46
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-18 14:26:19
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_int.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
)

// ToInt64SliceE 将 any 转换为 []int64 类型
func ToInt64SliceE(i any) (iv []int64, err error) {
	if i == nil {
		return []int64{}, nil
	}

	switch val := i.(type) {
	case []int64:
		return val, nil
	case []int32:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(val, &anyV); err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv = make([]int64, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt64E(v)
				if err != nil {
					return []int64{}, convertError(i, "[]int64")
				}
				iv[k] = intv
			}
			return
		}
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]int64, len(val))
		for k, v := range val {
			intv, err := ToInt64E(v)
			if err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(anyBytes, &anyV); err != nil {
				return []int64{}, convertError(i, "[]int64")
			}
			iv = make([]int64, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt64E(v)
				if err != nil {
					return []int64{}, convertError(i, "[]int64")
				}
				iv[k] = intv
			}
			return
		}
		return []int64{}, convertError(i, "[]int64")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]int64, length)
			for j := 0; j < length; j++ {
				intv, err := ToInt64E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []int64{}, convertError(i, "[]int64")
				}
				iv[j] = intv
			}
			return
		}

		return []int64{}, convertError(i, "[]int64")
	}
}

// ToInt32SliceE 将 any 转换为 []int32 类型
func ToInt32SliceE(i any) (iv []int32, err error) {
	if i == nil {
		return []int32{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []int32:
		return val, nil
	case []int16:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(val, &anyV); err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv = make([]int32, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt32E(v)
				if err != nil {
					return []int32{}, convertError(i, "[]int32")
				}
				iv[k] = intv
			}
			return
		}
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]int32, len(val))
		for k, v := range val {
			intv, err := ToInt32E(v)
			if err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(anyBytes, &anyV); err != nil {
				return []int32{}, convertError(i, "[]int32")
			}
			iv = make([]int32, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt32E(v)
				if err != nil {
					return []int32{}, convertError(i, "[]int32")
				}
				iv[k] = intv
			}
			return
		}
		return []int32{}, convertError(i, "[]int32")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]int32, length)
			for j := 0; j < length; j++ {
				intv, err := ToInt32E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []int32{}, convertError(i, "[]int32")
				}
				iv[j] = intv
			}
			return
		}

		return []int32{}, convertError(i, "[]int32")
	}
}

// ToInt16SliceE 将 any 转换为 []int16 类型
func ToInt16SliceE(i any) (iv []int16, err error) {
	if i == nil {
		return []int16{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []int16:
		return val, nil
	case []int8:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(val, &anyV); err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv = make([]int16, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt16E(v)
				if err != nil {
					return []int16{}, convertError(i, "[]int16")
				}
				iv[k] = intv
			}
			return
		}
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]int16, len(val))
		for k, v := range val {
			intv, err := ToInt16E(v)
			if err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(anyBytes, &anyV); err != nil {
				return []int16{}, convertError(i, "[]int16")
			}
			iv = make([]int16, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt16E(v)
				if err != nil {
					return []int16{}, convertError(i, "[]int16")
				}
				iv[k] = intv
			}
			return
		}
		return []int16{}, convertError(i, "[]int16")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]int16, length)
			for j := 0; j < length; j++ {
				intv, err := ToInt16E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []int16{}, convertError(i, "[]int16")
				}
				iv[j] = intv
			}
			return
		}

		return []int16{}, convertError(i, "[]int16")
	}
}

// ToInt8SliceE 将 any 转换为 []int8 类型
func ToInt8SliceE(i any) (iv []int8, err error) {
	if i == nil {
		return []int8{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []int8:
		return val, nil
	case []int:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(val, &anyV); err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv = make([]int8, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt8E(v)
				if err != nil {
					return []int8{}, convertError(i, "[]int8")
				}
				iv[k] = intv
			}
			return
		}
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]int8, len(val))
		for k, v := range val {
			intv, err := ToInt8E(v)
			if err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(anyBytes, &anyV); err != nil {
				return []int8{}, convertError(i, "[]int8")
			}
			iv = make([]int8, len(anyV))
			for k, v := range anyV {
				intv, err := ToInt8E(v)
				if err != nil {
					return []int8{}, convertError(i, "[]int8")
				}
				iv[k] = intv
			}
			return
		}
		return []int8{}, convertError(i, "[]int8")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]int8, length)
			for j := 0; j < length; j++ {
				intv, err := ToInt8E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []int8{}, convertError(i, "[]int8")
				}
				iv[j] = intv
			}
			return
		}

		return []int8{}, convertError(i, "[]int8")
	}
}

// ToIntSliceE 将 any 转换为 []int 类型
func ToIntSliceE(i any) (iv []int, err error) {
	if i == nil {
		return []int{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []int:
		return val, nil
	case []uint64:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(val, &anyV); err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv = make([]int, len(anyV))
			for k, v := range anyV {
				intv, err := ToIntE(v)
				if err != nil {
					return []int{}, convertError(i, "[]int")
				}
				iv[k] = intv
			}
			return
		}
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]int, len(val))
		for k, v := range val {
			intv, err := ToIntE(v)
			if err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if err := json.Unmarshal(anyBytes, &anyV); err != nil {
				return []int{}, convertError(i, "[]int")
			}
			iv = make([]int, len(anyV))
			for k, v := range anyV {
				intv, err := ToIntE(v)
				if err != nil {
					return []int{}, convertError(i, "[]int")
				}
				iv[k] = intv
			}
			return
		}
		return []int{}, convertError(i, "[]int")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]int, length)
			for j := 0; j < length; j++ {
				intv, err := ToIntE(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []int{}, convertError(i, "[]int")
				}
				iv[j] = intv
			}
			return
		}

		return []int{}, convertError(i, "[]int")
	}
}
