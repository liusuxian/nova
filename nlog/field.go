/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 00:46:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 17:26:32
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type logField = zapcore.Field

type ArrayEncoder = zapcore.ArrayEncoder
type ArrayMarshaler interface {
	MarshalLogArray(ArrayEncoder) (err error)
}

type ObjectEncoder = zapcore.ObjectEncoder
type ObjectMarshaler interface {
	MarshalLogObject(ObjectEncoder) (err error)
}

type ObjectMarshalerPtr[T any] interface {
	*T
	ObjectMarshaler
}

// Any
func Any(key string, val any) (nf logField) {
	return zap.Any(key, val)
}

// Array
func Array(key string, val ArrayMarshaler) (nf logField) {
	return zap.Array(key, val)
}

// Binary
func Binary(key string, val []byte) (nf logField) {
	return zap.Binary(key, val)
}

// Bool
func Bool(key string, val bool) (nf logField) {
	return zap.Bool(key, val)
}

// Boolp
func Boolp(key string, val *bool) (nf logField) {
	return zap.Boolp(key, val)
}

// Bools
func Bools(key string, bs []bool) (nf logField) {
	return zap.Bools(key, bs)
}

// ByteString
func ByteString(key string, val []byte) (nf logField) {
	return zap.ByteString(key, val)
}

// ByteStrings
func ByteStrings(key string, bss [][]byte) (nf logField) {
	return zap.ByteStrings(key, bss)
}

// Complex128
func Complex128(key string, val complex128) (nf logField) {
	return zap.Complex128(key, val)
}

// Complex128p
func Complex128p(key string, val *complex128) (nf logField) {
	return zap.Complex128p(key, val)
}

// Complex128s
func Complex128s(key string, nums []complex128) (nf logField) {
	return zap.Complex128s(key, nums)
}

// Complex64
func Complex64(key string, val complex64) (nf logField) {
	return zap.Complex64(key, val)
}

// Complex64p
func Complex64p(key string, val *complex64) (nf logField) {
	return zap.Complex64p(key, val)
}

// Complex64s
func Complex64s(key string, nums []complex64) (nf logField) {
	return zap.Complex64s(key, nums)
}

// Duration
func Duration(key string, val time.Duration) (nf logField) {
	return zap.Duration(key, val)
}

// Durationp
func Durationp(key string, val *time.Duration) (nf logField) {
	return zap.Durationp(key, val)
}

// Durations
func Durations(key string, ds []time.Duration) (nf logField) {
	return zap.Durations(key, ds)
}

// Err
func Err(err error) (nf logField) {
	return zap.Error(err)
}

// Errs
func Errs(key string, errs []error) (nf logField) {
	return zap.Errors(key, errs)
}

// Float32
func Float32(key string, val float32) (nf logField) {
	return zap.Float32(key, val)
}

// Float32p
func Float32p(key string, val *float32) (nf logField) {
	return zap.Float32p(key, val)
}

// Float32s
func Float32s(key string, nums []float32) (nf logField) {
	return zap.Float32s(key, nums)
}

// Float64
func Float64(key string, val float64) (nf logField) {
	return zap.Float64(key, val)
}

// Float64p
func Float64p(key string, val *float64) (nf logField) {
	return zap.Float64p(key, val)
}

// Float64s
func Float64s(key string, nums []float64) (nf logField) {
	return zap.Float64s(key, nums)
}

// Inline
func Inline(val ObjectMarshaler) (nf logField) {
	return zap.Inline(val)
}

// Int
func Int(key string, val int) (nf logField) {
	return zap.Int(key, val)
}

// Int16
func Int16(key string, val int16) (nf logField) {
	return zap.Int16(key, val)
}

// Int16p
func Int16p(key string, val *int16) (nf logField) {
	return zap.Int16p(key, val)
}

// Int16s
func Int16s(key string, nums []int16) (nf logField) {
	return zap.Int16s(key, nums)
}

// Int32
func Int32(key string, val int32) (nf logField) {
	return zap.Int32(key, val)
}

// Int32p
func Int32p(key string, val *int32) (nf logField) {
	return zap.Int32p(key, val)
}

// Int32s
func Int32s(key string, nums []int32) (nf logField) {
	return zap.Int32s(key, nums)
}

// Int64
func Int64(key string, val int64) (nf logField) {
	return zap.Int64(key, val)
}

// Int64p
func Int64p(key string, val *int64) (nf logField) {
	return zap.Int64p(key, val)
}

// Int64s
func Int64s(key string, nums []int64) (nf logField) {
	return zap.Int64s(key, nums)
}

// Int8
func Int8(key string, val int8) (nf logField) {
	return zap.Int8(key, val)
}

// Int8p
func Int8p(key string, val *int8) (nf logField) {
	return zap.Int8p(key, val)
}

// Int8s
func Int8s(key string, nums []int8) (nf logField) {
	return zap.Int8s(key, nums)
}

// Intp
func Intp(key string, val *int) (nf logField) {
	return zap.Intp(key, val)
}

// Ints
func Ints(key string, nums []int) (nf logField) {
	return zap.Ints(key, nums)
}

// NamedError
func NamedError(key string, err error) (nf logField) {
	return zap.NamedError(key, err)
}

// Namespace
func Namespace(key string) (nf logField) {
	return zap.Namespace(key)
}

// Object
func Object(key string, val ObjectMarshaler) (nf logField) {
	return zap.Object(key, val)
}

// ObjectValues
func ObjectValues[T any, P ObjectMarshalerPtr[T]](key string, vals []T) (nf logField) {
	return zap.ObjectValues[T, P](key, vals)
}

// Objects
func Objects[T ObjectMarshaler](key string, vals []T) (nf logField) {
	return zap.Objects(key, vals)
}

// Reflect
func Reflect(key string, val any) (nf logField) {
	return zap.Reflect(key, val)
}

// Skip
func Skip() (nf logField) {
	return zap.Skip()
}

// Stack
func Stack(key string) (nf logField) {
	return zap.Stack(key)
}

// StackSkip
func StackSkip(key string, skip int) (nf logField) {
	return zap.StackSkip(key, skip)
}

// String
func String(key string, val string) (nf logField) {
	return zap.String(key, val)
}

// Stringer
func Stringer(key string, val fmt.Stringer) (nf logField) {
	return zap.Stringer(key, val)
}

// Stringers
func Stringers[T fmt.Stringer](key string, vals []T) (nf logField) {
	return zap.Stringers(key, vals)
}

// Stringp
func Stringp(key string, val *string) (nf logField) {
	return zap.Stringp(key, val)
}

// Strings
func Strings(key string, ss []string) (nf logField) {
	return zap.Strings(key, ss)
}

// Time
func Time(key string, val time.Time) (nf logField) {
	return zap.Time(key, val)
}

// Timep
func Timep(key string, val *time.Time) (nf logField) {
	return zap.Timep(key, val)
}

// Times
func Times(key string, ts []time.Time) (nf logField) {
	return zap.Times(key, ts)
}

// Uint
func Uint(key string, val uint) (nf logField) {
	return zap.Uint(key, val)
}

// Uint16
func Uint16(key string, val uint16) (nf logField) {
	return zap.Uint16(key, val)
}

// Uint16p
func Uint16p(key string, val *uint16) (nf logField) {
	return zap.Uint16p(key, val)
}

// Uint16s
func Uint16s(key string, nums []uint16) (nf logField) {
	return zap.Uint16s(key, nums)
}

// Uint32
func Uint32(key string, val uint32) (nf logField) {
	return zap.Uint32(key, val)
}

// Uint32p
func Uint32p(key string, val *uint32) (nf logField) {
	return zap.Uint32p(key, val)
}

// Uint32s
func Uint32s(key string, nums []uint32) (nf logField) {
	return zap.Uint32s(key, nums)
}

// Uint64
func Uint64(key string, val uint64) (nf logField) {
	return zap.Uint64(key, val)
}

// Uint64p
func Uint64p(key string, val *uint64) (nf logField) {
	return zap.Uint64p(key, val)
}

// Uint64s
func Uint64s(key string, nums []uint64) (nf logField) {
	return zap.Uint64s(key, nums)
}

// Uint8
func Uint8(key string, val uint8) (nf logField) {
	return zap.Uint8(key, val)
}

// Uint8p
func Uint8p(key string, val *uint8) (nf logField) {
	return zap.Uint8p(key, val)
}

// Uint8s
func Uint8s(key string, nums []uint8) (nf logField) {
	return zap.Uint8s(key, nums)
}

// Uintp
func Uintp(key string, val *uint) (nf logField) {
	return zap.Uintp(key, val)
}

// Uintptr
func Uintptr(key string, val uintptr) (nf logField) {
	return zap.Uintptr(key, val)
}

// Uintptrp
func Uintptrp(key string, val *uintptr) (nf logField) {
	return zap.Uintptrp(key, val)
}

// Uintptrs
func Uintptrs(key string, us []uintptr) (nf logField) {
	return zap.Uintptrs(key, us)
}

// Uints
func Uints(key string, nums []uint) (nf logField) {
	return zap.Uints(key, nums)
}
