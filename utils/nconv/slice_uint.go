/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-19 11:58:36
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:13:18
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_uint.go
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

// ToUint64SliceE 将 any 转换为 []uint64 类型
func ToUint64SliceE(i any) (iv []uint64, err error) {
	if i == nil {
		return []uint64{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		return val, nil
	case []uint32:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]uint64, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint64E(v)
					if err != nil {
						return []uint64{}, convertError(i, "[]uint64")
					}
					iv[k] = intv
				}
				return
			}
		}
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]uint64, len(val))
		for k, v := range val {
			intv, err := ToUint64E(v)
			if err != nil {
				return []uint64{}, convertError(i, "[]uint64")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]uint64, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint64E(v)
					if err != nil {
						return []uint64{}, convertError(i, "[]uint64")
					}
					iv[k] = intv
				}
				return
			}
		}
		return []uint64{}, convertError(i, "[]uint64")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]uint64, length)
			for j := 0; j < length; j++ {
				intv, err := ToUint64E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []uint64{}, convertError(i, "[]uint64")
				}
				iv[j] = intv
			}
			return
		}

		return []uint64{}, convertError(i, "[]uint64")
	}
}

// ToUint32SliceE 将 any 转换为 []uint32 类型
func ToUint32SliceE(i any) (iv []uint32, err error) {
	if i == nil {
		return []uint32{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		return val, nil
	case []uint16:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]uint32, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint32E(v)
					if err != nil {
						return []uint32{}, convertError(i, "[]uint32")
					}
					iv[k] = intv
				}
				return
			}
		}
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]uint32, len(val))
		for k, v := range val {
			intv, err := ToUint32E(v)
			if err != nil {
				return []uint32{}, convertError(i, "[]uint32")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]uint32, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint32E(v)
					if err != nil {
						return []uint32{}, convertError(i, "[]uint32")
					}
					iv[k] = intv
				}
				return
			}
		}
		return []uint32{}, convertError(i, "[]uint32")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]uint32, length)
			for j := 0; j < length; j++ {
				intv, err := ToUint32E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []uint32{}, convertError(i, "[]uint32")
				}
				iv[j] = intv
			}
			return
		}

		return []uint32{}, convertError(i, "[]uint32")
	}
}

// ToUint16SliceE 将 any 转换为 []uint16 类型
func ToUint16SliceE(i any) (iv []uint16, err error) {
	if i == nil {
		return []uint16{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		return val, nil
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]uint16, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint16E(v)
					if err != nil {
						return []uint16{}, convertError(i, "[]uint16")
					}
					iv[k] = intv
				}
				return
			}
		}
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []uint:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]uint16, len(val))
		for k, v := range val {
			intv, err := ToUint16E(v)
			if err != nil {
				return []uint16{}, convertError(i, "[]uint16")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]uint16, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint16E(v)
					if err != nil {
						return []uint16{}, convertError(i, "[]uint16")
					}
					iv[k] = intv
				}
				return
			}
		}
		return []uint16{}, convertError(i, "[]uint16")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]uint16, length)
			for j := 0; j < length; j++ {
				intv, err := ToUint16E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []uint16{}, convertError(i, "[]uint16")
				}
				iv[j] = intv
			}
			return
		}

		return []uint16{}, convertError(i, "[]uint16")
	}
}

// ToUint8SliceE 将 any 转换为 []uint8 类型
func ToUint8SliceE(i any) (iv []uint8, err error) {
	if i == nil {
		return []uint8{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]uint8, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint8E(v)
					if err != nil {
						return []uint8{}, convertError(i, "[]uint8")
					}
					iv[k] = intv
				}
				return
			}
		}
		return val, nil
	case []uint:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []float64:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]uint8, len(val))
		for k, v := range val {
			intv, err := ToUint8E(v)
			if err != nil {
				return []uint8{}, convertError(i, "[]uint8")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]uint8, len(anyV))
				for k, v := range anyV {
					intv, err := ToUint8E(v)
					if err != nil {
						return []uint8{}, convertError(i, "[]uint8")
					}
					iv[k] = intv
				}
				return
			}
		}
		return []uint8{}, convertError(i, "[]uint8")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]uint8, length)
			for j := 0; j < length; j++ {
				intv, err := ToUint8E(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []uint8{}, convertError(i, "[]uint8")
				}
				iv[j] = intv
			}
			return
		}

		return []uint8{}, convertError(i, "[]uint8")
	}
}

// ToUintSliceE 将 any 转换为 []uint 类型
func ToUintSliceE(i any) (iv []uint, err error) {
	if i == nil {
		return []uint{}, nil
	}

	switch val := i.(type) {
	case []int64:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []int32:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []int16:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []int8:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []int:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []uint64:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []uint32:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []uint16:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(val, &anyV); e == nil {
				iv = make([]uint, len(anyV))
				for k, v := range anyV {
					intv, err := ToUintE(v)
					if err != nil {
						return []uint{}, convertError(i, "[]uint")
					}
					iv[k] = intv
				}
				return
			}
		}
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []uint:
		return val, nil
	case []float64:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []float32:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []bool:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case [][]byte:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []string:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case []any:
		iv = make([]uint, len(val))
		for k, v := range val {
			intv, err := ToUintE(v)
			if err != nil {
				return []uint{}, convertError(i, "[]uint")
			}
			iv[k] = intv
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			anyV := make([]any, len(val))
			if e := json.Unmarshal(anyBytes, &anyV); e == nil {
				iv = make([]uint, len(anyV))
				for k, v := range anyV {
					intv, err := ToUintE(v)
					if err != nil {
						return []uint{}, convertError(i, "[]uint")
					}
					iv[k] = intv
				}
				return
			}
		}
		return []uint{}, convertError(i, "[]uint")
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]uint, length)
			for j := 0; j < length; j++ {
				intv, err := ToUintE(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []uint{}, convertError(i, "[]uint")
				}
				iv[j] = intv
			}
			return
		}

		return []uint{}, convertError(i, "[]uint")
	}
}
