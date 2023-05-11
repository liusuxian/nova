/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:19:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:51:19
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"strconv"
)

// ToUint64E 将 any 转换为 uint64 类型
func ToUint64E(i any) (iv uint64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case uint64:
		return val, nil
	case uint32:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case uint:
		return uint64(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint64E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return ipv, nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint64")
			}
			return uint64(ipf), nil
		}
		return 0, convertError(i, "uint64")
	case json.Number:
		return ToUint64E(string(val))
	default:
		return 0, convertError(i, "uint64")
	}
}

// ToUint32E 将 any 转换为 uint32 类型
func ToUint32E(i any) (iv uint32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case uint64:
		return uint32(val), nil
	case uint32:
		return val, nil
	case uint16:
		return uint32(val), nil
	case uint8:
		return uint32(val), nil
	case uint:
		return uint32(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint32E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint32(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint32")
			}
			return uint32(ipf), nil
		}
		return 0, convertError(i, "uint32")
	case json.Number:
		return ToUint32E(string(val))
	default:
		return 0, convertError(i, "uint32")
	}
}

// ToUint16E 将 any 转换为 uint16 类型
func ToUint16E(i any) (iv uint16, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case uint64:
		return uint16(val), nil
	case uint32:
		return uint16(val), nil
	case uint16:
		return val, nil
	case uint8:
		return uint16(val), nil
	case uint:
		return uint16(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint16E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint16(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint16")
			}
			return uint16(ipf), nil
		}
		return 0, convertError(i, "uint16")
	case json.Number:
		return ToUint16E(string(val))
	default:
		return 0, convertError(i, "uint16")
	}
}

// ToUint8E 将 any 转换为 uint8 类型
func ToUint8E(i any) (iv uint8, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case uint64:
		return uint8(val), nil
	case uint32:
		return uint8(val), nil
	case uint16:
		return uint8(val), nil
	case uint8:
		return val, nil
	case uint:
		return uint8(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint8E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint8(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint8")
			}
			return uint8(ipf), nil
		}
		return 0, convertError(i, "uint8")
	case json.Number:
		return ToUint8E(string(val))
	default:
		return 0, convertError(i, "uint8")
	}
}

// ToUintE 将 any 转换为 uint 类型
func ToUintE(i any) (iv uint, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case uint64:
		return uint(val), nil
	case uint32:
		return uint(val), nil
	case uint16:
		return uint(val), nil
	case uint8:
		return uint(val), nil
	case uint:
		return val, nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUintE(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint")
			}
			return uint(ipf), nil
		}
		return 0, convertError(i, "uint")
	case json.Number:
		return ToUintE(string(val))
	default:
		return 0, convertError(i, "uint")
	}
}
