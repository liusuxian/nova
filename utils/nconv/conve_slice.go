/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:22:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 01:52:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_slice.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
	"strings"
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
		if json.Valid(val) {
			if err := unmarshalUseNumber(val, &iv); err != nil {
				return []any{}, convertError(i, "[]any")
			}
			return
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
	default:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if checkJsonAndUnmarshalUseNumber(i, &iv) {
			return
		}
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

		return []any{}, convertError(i, "[]any")
	}
}

// ToBoolSliceE  将 any 转换为 []bool 类型
func ToBoolSliceE(i any) (iv []bool, err error) {
	if i == nil {
		return []bool{}, nil
	}

	switch val := i.(type) {
	case []bool:
		return val, nil
	case []any:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int64:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int32:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int16:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int8:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint64:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint32:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint16:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint8:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []string:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case [][]byte:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]bool, length)
			for j := 0; j < length; j++ {
				bl, err := ToBoolE(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []bool{}, convertError(i, "[]bool")
				}
				iv[j] = bl
			}
			return
		}

		return []bool{}, convertError(i, "[]bool")
	}
}

// ToStringSliceE 将 any 转换为 []string 类型
func ToStringSliceE(i any) (iv []string, err error) {
	if i == nil {
		return []string{}, nil
	}

	switch val := i.(type) {
	case []string:
		return val, nil
	case []any:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []int64:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []int32:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []int16:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []int8:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []int:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []uint64:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []uint32:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []uint16:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []uint8:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []uint:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []float64:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []float32:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case []bool:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case [][]byte:
		iv = make([]string, len(val))
		for k, v := range val {
			str, err := ToStringE(v)
			if err != nil {
				return []string{}, convertError(i, "[]string")
			}
			iv[k] = str
		}
		return
	case string:
		return strings.Fields(val), nil
	case []error:
		iv = make([]string, len(val))
		for k, v := range val {
			iv[k] = v.Error()
		}
		return
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]string, length)
			for j := 0; j < length; j++ {
				str, err := ToStringE(originValueAndKind.OriginValue.Index(j).Interface())
				if err != nil {
					return []string{}, convertError(i, "[]string")
				}
				iv[j] = str
			}
			return
		}

		return []string{}, convertError(i, "[]string")
	}
}
