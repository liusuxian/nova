/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 22:24:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 22:33:31
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_ptr.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

// PtrAny 创建并返回指向该值的 any 指针变量
func PtrAny(val any) (ptr *any) {
	return &val
}

// PtrString 创建并返回指向该值的 string 指针变量
func PtrString(val any) (ptr *string) {
	v := ToString(val)
	return &v
}

// PtrBool 创建并返回指向该值的 bool 指针变量
func PtrBool(val any) (ptr *bool) {
	v := ToBool(val)
	return &v
}

// PtrInt 创建并返回指向该值的 int 指针变量
func PtrInt(val any) (ptr *int) {
	v := ToInt(val)
	return &v
}

// PtrInt8 创建并返回指向该值的 int8 指针变量
func PtrInt8(val any) (ptr *int8) {
	v := ToInt8(val)
	return &v
}

// PtrInt16 创建并返回指向该值的 int16 指针变量
func PtrInt16(val any) (ptr *int16) {
	v := ToInt16(val)
	return &v
}

// PtrInt32 创建并返回指向该值的 int32 指针变量
func PtrInt32(val any) (ptr *int32) {
	v := ToInt32(val)
	return &v
}

// PtrInt64 创建并返回指向该值的 int64 指针变量
func PtrInt64(val any) (ptr *int64) {
	v := ToInt64(val)
	return &v
}

// PtrUint 创建并返回指向该值的 uint 指针变量
func PtrUint(val any) (ptr *uint) {
	v := ToUint(val)
	return &v
}

// PtrUint8 创建并返回指向该值的 uint8 指针变量
func PtrUint8(val any) (ptr *uint8) {
	v := ToUint8(val)
	return &v
}

// PtrUint16 创建并返回指向该值的 uint16 指针变量
func PtrUint16(val any) (ptr *uint16) {
	v := ToUint16(val)
	return &v
}

// PtrUint32 创建并返回指向该值的 uint32 指针变量
func PtrUint32(val any) (ptr *uint32) {
	v := ToUint32(val)
	return &v
}

// PtrUint64 创建并返回指向该值的 uint64 指针变量
func PtrUint64(val any) (ptr *uint64) {
	v := ToUint64(val)
	return &v
}

// PtrFloat32 创建并返回指向该值的 float32 指针变量
func PtrFloat32(val any) (ptr *float32) {
	v := ToFloat32(val)
	return &v
}

// PtrFloat64 创建并返回指向该值的 float64 指针变量
func PtrFloat64(val any) (ptr *float64) {
	v := ToFloat64(val)
	return &v
}
