/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 16:09:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 21:34:29
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_float.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

// ToFloat32 将 any 转换为 float32 类型
func ToFloat32(val any) (cVal float32) {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return math.Float32frombits(binary.LittleEndian.Uint32(leFillUpSize(value, 4)))
	default:
		if f, ok := value.(iFloat32); ok {
			return f.Float32()
		}
		v, err := strconv.ParseFloat(ToString(val), 64)
		if err != nil {
			fmt.Printf("ToFloat32 Error: %+v\n", err)
		}
		return float32(v)
	}
}

// ToFloat64 将 any 转换为 float64 类型
func ToFloat64(val any) (cVal float64) {
	if val == nil {
		return 0
	}
	switch value := val.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return math.Float64frombits(binary.LittleEndian.Uint64(leFillUpSize(value, 8)))
	default:
		if f, ok := value.(iFloat64); ok {
			return f.Float64()
		}
		v, err := strconv.ParseFloat(ToString(val), 64)
		if err != nil {
			fmt.Printf("ToFloat64 Error: %+v\n", err)
		}
		return v
	}
}
