/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:23:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 03:25:37
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_string.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
	"strings"
)

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