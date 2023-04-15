/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:18:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-15 13:18:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_int.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"strconv"
)

// ToInt64E 将 any 转换为 int64 类型
func ToInt64E(i any) (iv int64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return val, nil
	case int32:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt64E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return ipv, nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int64(ipf), nil
		}
		return 0, convertError(i, "int64")
	case json.Number:
		return ToInt64E(string(val))
	default:
		return 0, convertError(i, "int64")
	}
}

// ToInt32E 将 any 转换为 int32 类型
func ToInt32E(i any) (iv int32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int32(val), nil
	case int32:
		return val, nil
	case int16:
		return int32(val), nil
	case int8:
		return int32(val), nil
	case uint64:
		return int32(val), nil
	case uint32:
		return int32(val), nil
	case uint16:
		return int32(val), nil
	case uint8:
		return int32(val), nil
	case uint:
		return int32(val), nil
	case float64:
		return int32(val), nil
	case float32:
		return int32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt32E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int32(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int32(ipf), nil
		}
		return 0, convertError(i, "int32")
	case json.Number:
		return ToInt32E(string(val))
	default:
		return 0, convertError(i, "int32")
	}
}

// ToInt16E 将 any 转换为 int16 类型
func ToInt16E(i interface{}) (iv int16, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int16(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int16(val), nil
	case int32:
		return int16(val), nil
	case int16:
		return val, nil
	case int8:
		return int16(val), nil
	case uint64:
		return int16(val), nil
	case uint32:
		return int16(val), nil
	case uint16:
		return int16(val), nil
	case uint8:
		return int16(val), nil
	case uint:
		return int16(val), nil
	case float64:
		return int16(val), nil
	case float32:
		return int16(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt16E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int16(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int16(ipf), nil
		}
		return 0, convertError(i, "int16")
	case json.Number:
		return ToInt16E(string(val))
	default:
		return 0, convertError(i, "int16")
	}
}

// ToInt8E 将 any 转换为 int8 类型
func ToInt8E(i any) (iv int8, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int8(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int8(val), nil
	case int32:
		return int8(val), nil
	case int16:
		return int8(val), nil
	case int8:
		return val, nil
	case uint64:
		return int8(val), nil
	case uint32:
		return int8(val), nil
	case uint16:
		return int8(val), nil
	case uint8:
		return int8(val), nil
	case uint:
		return int8(val), nil
	case float64:
		return int8(val), nil
	case float32:
		return int8(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt8E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int8(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int8(ipf), nil
		}
		return 0, convertError(i, "int8")
	case json.Number:
		return ToInt8E(string(val))
	default:
		return 0, convertError(i, "int8")
	}
}

// ToIntE 将 any 转换为 int 类型
func ToIntE(i interface{}) (iv int, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return intv, nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int(val), nil
	case int32:
		return int(val), nil
	case int16:
		return int(val), nil
	case int8:
		return int(val), nil
	case uint64:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint:
		return int(val), nil
	case float64:
		return int(val), nil
	case float32:
		return int(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToIntE(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int(ipf), nil
		}
		return 0, convertError(i, "int")
	case json.Number:
		return ToIntE(string(val))
	default:
		return 0, convertError(i, "int")
	}
}
