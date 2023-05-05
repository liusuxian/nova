/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:17:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:24:33
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_any.go
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

// ToSliceE 将 any 转换为 []any 类型
func ToSliceE(i any) (iv []any, err error) {
	if i == nil {
		return []any{}, nil
	}

	switch val := i.(type) {
	case []any:
		return val, nil
	case []int64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int16:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int8:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint16:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint8:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if json.Valid(val) {
			if e := json.Unmarshal(val, &iv); e == nil {
				return
			}
		}
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []float64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []float32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []bool:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []string:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case [][]byte:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []map[string]any:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case string:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if e := json.Unmarshal(anyBytes, &iv); e == nil {
				return
			}
		}
		return []any{val}, nil
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]any, length)
			for i := 0; i < length; i++ {
				iv[i] = originValueAndKind.OriginValue.Index(i).Interface()
			}
			return
		}

		return []any{val}, nil
	}
}
