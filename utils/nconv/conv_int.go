/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 17:18:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 21:35:04
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_int.go
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

// ToInt
func ToInt(val any) (cVal int) {
	if val == nil {
		return 0
	}
	if v, ok := val.(int); ok {
		return v
	}
	return int(ToInt64(val))
}

// ToInt8
func ToInt8(val any) (cVal int8) {
	if val == nil {
		return 0
	}
	if v, ok := val.(int8); ok {
		return v
	}
	return int8(ToInt64(val))
}

// ToInt16
func ToInt16(val any) (cVal int16) {
	if val == nil {
		return 0
	}
	if v, ok := val.(int16); ok {
		return v
	}
	return int16(ToInt64(val))
}

// ToInt32
func ToInt32(val any) (cVal int32) {
	if val == nil {
		return 0
	}
	if v, ok := val.(int32); ok {
		return v
	}
	return int32(ToInt64(val))
}

// ToInt64
func ToInt64(val any) (cVal int64) {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return int64(binary.LittleEndian.Uint64(leFillUpSize(value, 8)))
	default:
		if f, ok := value.(iInt64); ok {
			return f.Int64()
		}
		var s = ToString(value)
		var isMinus = false
		if len(s) > 0 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// 十六进制
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// 十进制
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			if isMinus {
				return -v
			}
			return v
		}
		// Float64
		if valueInt64 := ToFloat64(value); math.IsNaN(valueInt64) {
			return 0
		} else {
			return int64(valueInt64)
		}
	}
}
