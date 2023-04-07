/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 15:55:48
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-07 17:26:22
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nbinary/binary_le.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nbinary

import (
	"encoding/binary"
	"math"
)

// LeDecodeToInt64
func LeDecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

// LeDecodeToUint64
func LeDecodeToUint64(b []byte) (val uint64) {
	return binary.LittleEndian.Uint64(LeFillUpSize(b, 8))
}

// LeDecodeToFloat32
func LeDecodeToFloat32(b []byte) (val float32) {
	return math.Float32frombits(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

// LeDecodeToFloat64
func LeDecodeToFloat64(b []byte) (val float64) {
	return math.Float64frombits(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

// LeFillUpSize 使用 LittleEndian 填充字节切片 b 到给定的长度 l
func LeFillUpSize(b []byte, l int) (buf []byte) {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
