/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:20:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 02:21:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_slice_bool.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
)

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
