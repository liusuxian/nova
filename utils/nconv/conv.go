/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 13:11:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 18:42:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"google.golang.org/protobuf/proto"
	"time"
)

// ToBool 将 any 转换为 bool 类型
func ToBool(i any) (v bool) {
	v, _ = ToBoolE(i)
	return
}

// ToTime 将 any 转换为 time.Time 类型
func ToTime(i any) (v time.Time) {
	v, _ = ToTimeE(i)
	return
}

// ToTimeInDefaultLocation 将 any 转换为 time.Time 类型
func ToTimeInDefaultLocation(i any, location *time.Location) (v time.Time) {
	v, _ = ToTimeInDefaultLocationE(i, location)
	return
}

// ToDuration 将 any 转换为 time.Duration 类型
func ToDuration(i any) (v time.Duration) {
	v, _ = ToDurationE(i)
	return
}

// ToFloat64 将 any 转换为 float64 类型
func ToFloat64(i any) (v float64) {
	v, _ = ToFloat64E(i)
	return
}

// ToFloat32 将 any 转换为 float32 类型
func ToFloat32(i any) (v float32) {
	v, _ = ToFloat32E(i)
	return
}

// ToInt64 将 any 转换为 int64 类型
func ToInt64(i any) (v int64) {
	v, _ = ToInt64E(i)
	return
}

// ToInt32 将 any 转换为 int32 类型
func ToInt32(i any) (v int32) {
	v, _ = ToInt32E(i)
	return
}

// ToInt16 将 any 转换为 int16 类型
func ToInt16(i any) (v int16) {
	v, _ = ToInt16E(i)
	return
}

// ToInt8 将 any 转换为 int8 类型
func ToInt8(i any) (v int8) {
	v, _ = ToInt8E(i)
	return
}

// ToInt 将 any 转换为 int 类型
func ToInt(i any) (v int) {
	v, _ = ToIntE(i)
	return
}

// ToUint64 将 any 转换为 uint64 类型
func ToUint64(i any) (v uint64) {
	v, _ = ToUint64E(i)
	return
}

// ToUint32 将 any 转换为 uint32 类型
func ToUint32(i any) (v uint32) {
	v, _ = ToUint32E(i)
	return
}

// ToUint16 将 any 转换为 uint16 类型
func ToUint16(i any) (v uint16) {
	v, _ = ToUint16E(i)
	return
}

// ToUint8 将 any 转换为 uint8 类型
func ToUint8(i any) (v uint8) {
	v, _ = ToUint8E(i)
	return
}

// ToUint 将 any 转换为 uint 类型
func ToUint(i any) (v uint) {
	v, _ = ToUintE(i)
	return
}

// ToString 将 any 转换为 string 类型
func ToString(i any) (v string) {
	v, _ = ToStringE(i)
	return
}

// ToStringMapString 将 any 转换为 map[string]string 类型
func ToStringMapString(i any) (v map[string]string) {
	v, _ = ToStringMapStringE(i)
	return
}

// ToStringMapStringSlice 将 any 转换为 map[string][]string 类型
func ToStringMapStringSlice(i any) (v map[string][]string) {
	v, _ = ToStringMapStringSliceE(i)
	return
}

// ToStringMapBool 将 any 转换为 map[string]bool 类型
func ToStringMapBool(i any) (v map[string]bool) {
	v, _ = ToStringMapBoolE(i)
	return
}

// ToStringMapInt64 将 any 转换为 map[string]int64 类型
func ToStringMapInt64(i any) (v map[string]int64) {
	v, _ = ToStringMapInt64E(i)
	return
}

// ToStringMapInt32 将 any 转换为 map[string]int32 类型
func ToStringMapInt32(i any) (v map[string]int32) {
	v, _ = ToStringMapInt32E(i)
	return
}

// ToStringMapInt16 将 any 转换为 map[string]int16 类型
func ToStringMapInt16(i any) (v map[string]int16) {
	v, _ = ToStringMapInt16E(i)
	return
}

// ToStringMapInt8 将 any 转换为 map[string]int8 类型
func ToStringMapInt8(i any) (v map[string]int8) {
	v, _ = ToStringMapInt8E(i)
	return
}

// ToStringMapInt 将 any 转换为 map[string]int 类型
func ToStringMapInt(i any) (v map[string]int) {
	v, _ = ToStringMapIntE(i)
	return
}

// ToStringMapUint64 将 any 转换为 map[string]uint64 类型
func ToStringMapUint64(i any) (v map[string]uint64) {
	v, _ = ToStringMapUint64E(i)
	return
}

// ToStringMapUint32 将 any 转换为 map[string]uint32 类型
func ToStringMapUint32(i any) (v map[string]uint32) {
	v, _ = ToStringMapUint32E(i)
	return
}

// ToStringMapUint16 将 any 转换为 map[string]uint16 类型
func ToStringMapUint16(i any) (v map[string]uint16) {
	v, _ = ToStringMapUint16E(i)
	return
}

// ToStringMapUint8 将 any 转换为 map[string]uint8 类型
func ToStringMapUint8(i any) (v map[string]uint8) {
	v, _ = ToStringMapUint8E(i)
	return
}

// ToStringMapUint 将 any 转换为 map[string]uint 类型
func ToStringMapUint(i any) (v map[string]uint) {
	v, _ = ToStringMapUintE(i)
	return
}

// ToStringMapFloat64 将 any 转换为 map[string]float64 类型
func ToStringMapFloat64(i any) (v map[string]float64) {
	v, _ = ToStringMapFloat64E(i)
	return
}

// ToStringMapFloat32 将 any 转换为 map[string]float32 类型
func ToStringMapFloat32(i any) (v map[string]float32) {
	v, _ = ToStringMapFloat32E(i)
	return
}

// ToStringMap 将 any 转换为 map[string]any 类型
func ToStringMap(i any, opts ...DecoderConfigOption) (v map[string]any) {
	v, _ = ToStringMapE(i, opts...)
	return
}

// ToSlice 将 any 转换为 []any 类型
func ToSlice(i any) (v []any) {
	v, _ = ToSliceE(i)
	return
}

// ToBoolSlice 将 any 转换为 []bool 类型
func ToBoolSlice(i any) (v []bool) {
	v, _ = ToBoolSliceE(i)
	return
}

// ToStringSlice 将 any 转换为 []string 类型
func ToStringSlice(i any) (v []string) {
	v, _ = ToStringSliceE(i)
	return
}

// ToInt64Slice 将 any 转换为 []int64 类型
func ToInt64Slice(i any) (v []int64) {
	v, _ = ToInt64SliceE(i)
	return
}

// ToInt32Slice 将 any 转换为 []int32 类型
func ToInt32Slice(i any) (v []int32) {
	v, _ = ToInt32SliceE(i)
	return
}

// ToInt16Slice 将 any 转换为 []int16 类型
func ToInt16Slice(i any) (v []int16) {
	v, _ = ToInt16SliceE(i)
	return
}

// ToInt8Slice 将 any 转换为 []int8 类型
func ToInt8Slice(i any) (v []int8) {
	v, _ = ToInt8SliceE(i)
	return
}

// ToIntSlice 将 any 转换为 []int 类型
func ToIntSlice(i any) (v []int) {
	v, _ = ToIntSliceE(i)
	return
}

// ToUint64Slice 将 any 转换为 []uint64 类型
func ToUint64Slice(i any) (v []uint64) {
	v, _ = ToUint64SliceE(i)
	return
}

// ToUint32Slice 将 any 转换为 []uint32 类型
func ToUint32Slice(i any) (v []uint32) {
	v, _ = ToUint32SliceE(i)
	return
}

// ToUint16Slice 将 any 转换为 []uint16 类型
func ToUint16Slice(i any) (v []uint16) {
	v, _ = ToUint16SliceE(i)
	return
}

// ToUint8Slice 将 any 转换为 []uint8 类型
func ToUint8Slice(i any) (v []uint8) {
	v, _ = ToUint8SliceE(i)
	return
}

// ToUintSlice 将 any 转换为 []uint 类型
func ToUintSlice(i any) (v []uint) {
	v, _ = ToUintSliceE(i)
	return
}

// ToFloat64Slice 将 any 转换为 []float64 类型
func ToFloat64Slice(i any) (v []float64) {
	v, _ = ToFloat64SliceE(i)
	return
}

// ToFloat32Slice 将 any 转换为 []float32 类型
func ToFloat32Slice(i any) (v []float32) {
	v, _ = ToFloat32SliceE(i)
	return
}

// ToDurationSlice 将 any 转换为 []time.Duration 类型
func ToDurationSlice(i any) (v []time.Duration) {
	v, _ = ToDurationSliceE(i)
	return
}

// ToStruct 将 any 转换为 struct/[]struct 类型
func ToStruct(input, output any, opts ...DecoderConfigOption) {
	_ = ToStructE(input, output, opts...)
}

// ProtoMsgToMap 将 protobuf 消息转换为 Map 类型
func ProtoMsgToMap(msg proto.Message) (m map[string]any) {
	m, _ = ProtoMsgToMapE(msg)
	return
}
