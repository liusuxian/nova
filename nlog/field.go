/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-31 15:14:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 16:17:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/field.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type nField struct {
	field zapcore.Field
}

// Any
func Any(key string, val any) nField {
	return nField{field: zap.Any(key, val)}
}

// Array
func Array(key string, val zapcore.ArrayMarshaler) nField {
	return nField{field: zap.Array(key, val)}
}

// Binary
func Binary(key string, val []byte) nField {
	return nField{field: zap.Binary(key, val)}
}

// Bool
func Bool(key string, val bool) nField {
	return nField{field: zap.Bool(key, val)}
}

// Boolp
func Boolp(key string, val *bool) nField {
	return nField{field: zap.Boolp(key, val)}
}

// Bools
func Bools(key string, bs []bool) nField {
	return nField{field: zap.Bools(key, bs)}
}

// ByteString
func ByteString(key string, val []byte) nField {
	return nField{field: zap.ByteString(key, val)}
}

// ByteStrings
func ByteStrings(key string, bss [][]byte) nField {
	return nField{field: zap.ByteStrings(key, bss)}
}

// Complex128
func Complex128(key string, val complex128) nField {
	return nField{field: zap.Complex128(key, val)}
}

// Complex128p
func Complex128p(key string, val *complex128) nField {
	return nField{field: zap.Complex128p(key, val)}
}

// Complex128s
func Complex128s(key string, nums []complex128) nField {
	return nField{field: zap.Complex128s(key, nums)}
}

// Complex64
func Complex64(key string, val complex64) nField {
	return nField{field: zap.Complex64(key, val)}
}

// Complex64p
func Complex64p(key string, val *complex64) nField {
	return nField{field: zap.Complex64p(key, val)}
}

// Complex64s
func Complex64s(key string, nums []complex64) nField {
	return nField{field: zap.Complex64s(key, nums)}
}

// Duration
func Duration(key string, val time.Duration) nField {
	return nField{field: zap.Duration(key, val)}
}

// Durationp
func Durationp(key string, val *time.Duration) nField {
	return nField{field: zap.Durationp(key, val)}
}

// Durations
func Durations(key string, ds []time.Duration) nField {
	return nField{field: zap.Durations(key, ds)}
}

// Err
func Err(err error) nField {
	return nField{field: zap.Error(err)}
}

// Errs
func Errs(key string, errs []error) nField {
	return nField{field: zap.Errors(key, errs)}
}

// Float32
func Float32(key string, val float32) nField {
	return nField{field: zap.Float32(key, val)}
}

// Float32p
func Float32p(key string, val *float32) nField {
	return nField{field: zap.Float32p(key, val)}
}

// Float32s
func Float32s(key string, nums []float32) nField {
	return nField{field: zap.Float32s(key, nums)}
}

// Float64
func Float64(key string, val float64) nField {
	return nField{field: zap.Float64(key, val)}
}

// Float64p
func Float64p(key string, val *float64) nField {
	return nField{field: zap.Float64p(key, val)}
}

// Float64s
func Float64s(key string, nums []float64) nField {
	return nField{field: zap.Float64s(key, nums)}
}

// Inline
func Inline(val zapcore.ObjectMarshaler) nField {
	return nField{field: zap.Inline(val)}
}

// Int
func Int(key string, val int) nField {
	return nField{field: zap.Int(key, val)}
}

// Int16
func Int16(key string, val int16) nField {
	return nField{field: zap.Int16(key, val)}
}

// Int16p
func Int16p(key string, val *int16) nField {
	return nField{field: zap.Int16p(key, val)}
}

// Int16s
func Int16s(key string, nums []int16) nField {
	return nField{field: zap.Int16s(key, nums)}
}

// Int32
func Int32(key string, val int32) nField {
	return nField{field: zap.Int32(key, val)}
}

// Int32p
func Int32p(key string, val *int32) nField {
	return nField{field: zap.Int32p(key, val)}
}

// Int32s
func Int32s(key string, nums []int32) nField {
	return nField{field: zap.Int32s(key, nums)}
}

// Int64
func Int64(key string, val int64) nField {
	return nField{field: zap.Int64(key, val)}
}

// Int64p
func Int64p(key string, val *int64) nField {
	return nField{field: zap.Int64p(key, val)}
}

// Int64s
func Int64s(key string, nums []int64) nField {
	return nField{field: zap.Int64s(key, nums)}
}

// Int8
func Int8(key string, val int8) nField {
	return nField{field: zap.Int8(key, val)}
}

// Int8p
func Int8p(key string, val *int8) nField {
	return nField{field: zap.Int8p(key, val)}
}

// Int8s
func Int8s(key string, nums []int8) nField {
	return nField{field: zap.Int8s(key, nums)}
}

// Intp
func Intp(key string, val *int) nField {
	return nField{field: zap.Intp(key, val)}
}

// Ints
func Ints(key string, nums []int) nField {
	return nField{field: zap.Ints(key, nums)}
}

// NamedError
func NamedError(key string, err error) nField {
	return nField{field: zap.NamedError(key, err)}
}

// Namespace
func Namespace(key string) nField {
	return nField{field: zap.Namespace(key)}
}

// Object
func Object(key string, val zapcore.ObjectMarshaler) nField {
	return nField{field: zap.Object(key, val)}
}

// ObjectValues
func ObjectValues[T any, P zap.ObjectMarshalerPtr[T]](key string, vals []T) nField {
	return nField{field: zap.ObjectValues[T, P](key, vals)}
}

// Objects
func Objects[T zapcore.ObjectMarshaler](key string, vals []T) nField {
	return nField{field: zap.Objects(key, vals)}
}

// Reflect
func Reflect(key string, val any) nField {
	return nField{field: zap.Reflect(key, val)}
}

// Skip
func Skip() nField {
	return nField{field: zap.Skip()}
}

// Stack
func Stack(key string) nField {
	return nField{field: zap.Stack(key)}
}

// StackSkip
func StackSkip(key string, skip int) nField {
	return nField{field: zap.StackSkip(key, skip)}
}

// String
func String(key string, val string) nField {
	return nField{field: zap.String(key, val)}
}

// Stringer
func Stringer(key string, val fmt.Stringer) nField {
	return nField{field: zap.Stringer(key, val)}
}

// Stringers
func Stringers[T fmt.Stringer](key string, vals []T) nField {
	return nField{field: zap.Stringers(key, vals)}
}

// Stringp
func Stringp(key string, val *string) nField {
	return nField{field: zap.Stringp(key, val)}
}

// Strings
func Strings(key string, ss []string) nField {
	return nField{field: zap.Strings(key, ss)}
}

// Time
func Time(key string, val time.Time) nField {
	return nField{field: zap.Time(key, val)}
}

// Timep
func Timep(key string, val *time.Time) nField {
	return nField{field: zap.Timep(key, val)}
}

// Times
func Times(key string, ts []time.Time) nField {
	return nField{field: zap.Times(key, ts)}
}

// Uint
func Uint(key string, val uint) nField {
	return nField{field: zap.Uint(key, val)}
}

// Uint16
func Uint16(key string, val uint16) nField {
	return nField{field: zap.Uint16(key, val)}
}

// Uint16p
func Uint16p(key string, val *uint16) nField {
	return nField{field: zap.Uint16p(key, val)}
}

// Uint16s
func Uint16s(key string, nums []uint16) nField {
	return nField{field: zap.Uint16s(key, nums)}
}

// Uint32
func Uint32(key string, val uint32) nField {
	return nField{field: zap.Uint32(key, val)}
}

// Uint32p
func Uint32p(key string, val *uint32) nField {
	return nField{field: zap.Uint32p(key, val)}
}

// Uint32s
func Uint32s(key string, nums []uint32) nField {
	return nField{field: zap.Uint32s(key, nums)}
}

// Uint64
func Uint64(key string, val uint64) nField {
	return nField{field: zap.Uint64(key, val)}
}

// Uint64p
func Uint64p(key string, val *uint64) nField {
	return nField{field: zap.Uint64p(key, val)}
}

// Uint64s
func Uint64s(key string, nums []uint64) nField {
	return nField{field: zap.Uint64s(key, nums)}
}

// Uint8
func Uint8(key string, val uint8) nField {
	return nField{field: zap.Uint8(key, val)}
}

// Uint8p
func Uint8p(key string, val *uint8) nField {
	return nField{field: zap.Uint8p(key, val)}
}

// Uint8s
func Uint8s(key string, nums []uint8) nField {
	return nField{field: zap.Uint8s(key, nums)}
}

// Uintp
func Uintp(key string, val *uint) nField {
	return nField{field: zap.Uintp(key, val)}
}

// Uintptr
func Uintptr(key string, val uintptr) nField {
	return nField{field: zap.Uintptr(key, val)}
}

// Uintptrp
func Uintptrp(key string, val *uintptr) nField {
	return nField{field: zap.Uintptrp(key, val)}
}

// Uintptrs
func Uintptrs(key string, us []uintptr) nField {
	return nField{field: zap.Uintptrs(key, us)}
}

// Uints
func Uints(key string, nums []uint) nField {
	return nField{field: zap.Uints(key, nums)}
}
