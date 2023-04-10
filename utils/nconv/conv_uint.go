/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 14:42:33
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 15:43:00
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_uint.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/binary"
	"math"
	"strconv"
)

// ToUint 将 any 转换为 uint 类型
func ToUint(val any) (cVal uint) {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint); ok {
		return v
	}
	return uint(ToUint64(val))
}

// ToUint8 将 any 转换为 uint8 类型
func ToUint8(val any) (cVal uint8) {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint8); ok {
		return v
	}
	return uint8(ToUint64(val))
}

// ToUint16 将 any 转换为 uint16 类型
func ToUint16(val any) (cVal uint16) {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint16); ok {
		return v
	}
	return uint16(ToUint64(val))
}

// ToUint32 将 any 转换为 uint32 类型
func ToUint32(val any) (cVal uint32) {
	if val == nil {
		return 0
	}
	if v, ok := val.(uint32); ok {
		return v
	}
	return uint32(ToUint64(val))
}

// ToUint64 将 any 转换为 uint64 类型
func ToUint64(val any) (cVal uint64) {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return binary.LittleEndian.Uint64(leFillUpSize(value, 8))
	default:
		if f, ok := value.(iUint64); ok {
			return f.Uint64()
		}
		s := ToString(value)
		// 十六进制
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// 十进制
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		if valueFloat64 := ToFloat64(value); math.IsNaN(valueFloat64) {
			return 0
		} else {
			return uint64(valueFloat64)
		}
	}
}
