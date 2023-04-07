/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 16:09:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-07 17:10:22
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_float.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"fmt"
	"github.com/liusuxian/nova/utils/nbinary"
	"strconv"
)

// ToFloat32
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
		return nbinary.DecodeToFloat32(value)
	default:
		if f, ok := value.(iFloat32); ok {
			return f.Float32()
		}
		v, err := strconv.ParseFloat(String(val), 64)
		if err != nil {
			fmt.Printf("ToFloat32 Error: %+v\n", err)
		}
		return float32(v)
	}
}

// ToFloat64
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
		return nbinary.DecodeToFloat64(value)
	default:
		if f, ok := value.(iFloat64); ok {
			return f.Float64()
		}
		v, err := strconv.ParseFloat(String(val), 64)
		if err != nil {
			fmt.Printf("ToFloat64 Error: %+v\n", err)
		}
		return v
	}
}
