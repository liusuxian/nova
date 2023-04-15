/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:21:13
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-15 13:21:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_float.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"strconv"
)

// ToFloat64E 将 any 转换为 float64 类型
func ToFloat64E(i any) (iv float64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToFloat64E(string(val))
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return v, nil
		}
		return 0, convertError(i, "float64")
	case json.Number:
		v, err := val.Float64()
		if err == nil {
			return v, nil
		}
		return 0, convertError(i, "float64")
	default:
		return 0, convertError(i, "float64")
	}
}

// ToFloat32E 将 any 转换为 float32 类型
func ToFloat32E(i any) (iv float32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case float64:
		return float32(val), nil
	case float32:
		return val, nil
	case int64:
		return float32(val), nil
	case int32:
		return float32(val), nil
	case int16:
		return float32(val), nil
	case int8:
		return float32(val), nil
	case uint64:
		return float32(val), nil
	case uint32:
		return float32(val), nil
	case uint16:
		return float32(val), nil
	case uint8:
		return float32(val), nil
	case uint:
		return float32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToFloat32E(string(val))
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return float32(v), nil
		}
		return 0, convertError(i, "float32")
	case json.Number:
		v, err := val.Float64()
		if err == nil {
			return float32(v), nil
		}
		return 0, convertError(i, "float32")
	default:
		return 0, convertError(i, "float32")
	}
}
