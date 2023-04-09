/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:19:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 23:54:56
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_uint.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/internal/reflection"
	"github.com/liusuxian/nova/utils/nstr"
	"reflect"
	"strings"
)

// ToUints
func ToUints(val any) (s []uint) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint{}
		}
		if nstr.IsNumeric(value) {
			return []uint{ToUint(value)}
		}
	case []string:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []int8:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []int16:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []int32:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []int64:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []uint:
		s = value
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]uint, len(value))
			for k, v := range value {
				s[k] = ToUint(v)
			}
		}
	case []uint16:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []uint32:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []uint64:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []bool:
		s = make([]uint, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []float64:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case []any:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	case [][]byte:
		s = make([]uint, len(value))
		for k, v := range value {
			s[k] = ToUint(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iUints); ok {
		return v.Uints()
	}
	if v, ok := val.(iInterfaces); ok {
		return ToUints(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]uint, length)
		for i := 0; i < length; i++ {
			s[i] = ToUint(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint{}
		}
		return []uint{ToUint(val)}
	}
}

// ToUint8s
func ToUint8s(val any) (s []uint8) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint8{}
		}
		if nstr.IsNumeric(value) {
			return []uint8{ToUint8(value)}
		}
	case []string:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []int8:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []int16:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []int32:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []int64:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []uint:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = value
		}
	case []uint16:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []uint32:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []uint64:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []bool:
		s = make([]uint8, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []float64:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case []any:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	case [][]byte:
		s = make([]uint8, len(value))
		for k, v := range value {
			s[k] = ToUint8(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iUints); ok {
		return ToUint8s(v.Uints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToUint8s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]uint8, length)
		for i := 0; i < length; i++ {
			s[i] = ToUint8(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint8{}
		}
		return []uint8{ToUint8(val)}
	}
}

// ToUint16s
func ToUint16s(val any) (s []uint16) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint16{}
		}
		if nstr.IsNumeric(value) {
			return []uint16{ToUint16(value)}
		}
	case []string:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []int8:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []int16:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []int32:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []int64:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []uint:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]uint16, len(value))
			for k, v := range value {
				s[k] = ToUint16(v)
			}
		}
	case []uint16:
		s = value
	case []uint32:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []uint64:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []bool:
		s = make([]uint16, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []float64:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case []any:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	case [][]byte:
		s = make([]uint16, len(value))
		for k, v := range value {
			s[k] = ToUint16(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iUints); ok {
		return ToUint16s(v.Uints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToUint16s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]uint16, length)
		for i := 0; i < length; i++ {
			s[i] = ToUint16(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint16{}
		}
		return []uint16{ToUint16(val)}
	}
}

// ToUint32s
func ToUint32s(val any) (s []uint32) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint32{}
		}
		if nstr.IsNumeric(value) {
			return []uint32{ToUint32(value)}
		}
	case []string:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []int8:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []int16:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []int32:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []int64:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []uint:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]uint32, len(value))
			for k, v := range value {
				s[k] = ToUint32(v)
			}
		}
	case []uint16:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []uint32:
		s = value
	case []uint64:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []bool:
		s = make([]uint32, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []float64:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case []any:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	case [][]byte:
		s = make([]uint32, len(value))
		for k, v := range value {
			s[k] = ToUint32(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iUints); ok {
		return ToUint32s(v.Uints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToUint32s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]uint32, length)
		for i := 0; i < length; i++ {
			s[i] = ToUint32(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint32{}
		}
		return []uint32{ToUint32(val)}
	}
}

// ToUint64s
func ToUint64s(val any) (s []uint64) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return []uint64{}
		}
		if nstr.IsNumeric(value) {
			return []uint64{ToUint64(value)}
		}
	case []string:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []int8:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []int16:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []int32:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []int64:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []uint:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &s)
		} else {
			s = make([]uint64, len(value))
			for k, v := range value {
				s[k] = ToUint64(v)
			}
		}
	case []uint16:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []uint32:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []uint64:
		s = value
	case []bool:
		s = make([]uint64, len(value))
		for k, v := range value {
			if v {
				s[k] = 1
			} else {
				s[k] = 0
			}
		}
	case []float32:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []float64:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case []any:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	case [][]byte:
		s = make([]uint64, len(value))
		for k, v := range value {
			s[k] = ToUint64(v)
		}
	}
	if s != nil {
		return
	}
	if v, ok := val.(iUints); ok {
		return ToUint64s(v.Uints())
	}
	if v, ok := val.(iInterfaces); ok {
		return ToUint64s(v.Interfaces())
	}
	// 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
	if checkJsonAndUnmarshalUseNumber(val, &s) {
		return
	}
	// 传入的参数不是常见的类型，则会使用反射进行转换
	originValueAndKind := reflection.OriginValueAndKind(val)
	switch originValueAndKind.OriginKind {
	case reflect.Slice, reflect.Array:
		length := originValueAndKind.OriginValue.Len()
		s = make([]uint64, length)
		for i := 0; i < length; i++ {
			s[i] = ToUint64(originValueAndKind.OriginValue.Index(i).Interface())
		}
		return
	default:
		if originValueAndKind.OriginValue.IsZero() {
			return []uint64{}
		}
		return []uint64{ToUint64(val)}
	}
}
