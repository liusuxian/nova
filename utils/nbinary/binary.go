/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 15:54:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-07 17:24:30
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nbinary/binary.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nbinary

// DecodeToInt64
func DecodeToInt64(b []byte) int64 {
	return LeDecodeToInt64(b)
}

// DecodeToUint64
func DecodeToUint64(b []byte) (val uint64) {
	return LeDecodeToUint64(b)
}

// DecodeToFloat32
func DecodeToFloat32(b []byte) (val float32) {
	return LeDecodeToFloat32(b)
}

// DecodeToFloat64
func DecodeToFloat64(b []byte) (val float64) {
	return LeDecodeToFloat64(b)
}
