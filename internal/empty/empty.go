/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 18:03:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 18:19:58
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/empty/empty.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package empty

import (
	"github.com/liusuxian/nova/internal/reflection"
	"reflect"
	"time"
)

// iString 使用 String() 方法的类型断言 API
type iString interface {
	String() (val string)
}

// iInterfaces 使用 Interfaces() 方法的类型断言 API
type iInterfaces interface {
	Interfaces() (vals []any)
}

// iMapStrAny 将 struct 参数转换为 map
type iMapStrAny interface {
	MapStrAny() (m map[string]any)
}

// iTime 是用于 time.Time 转换的接口
type iTime interface {
	Date() (year int, month time.Month, day int)
	IsZero() (isZero bool)
}

// IsEmpty 判断给定的 value 是否为空
//
//	如果`value`值为`0, nil, false, "", len(slice/map/chan) == 0`，则返回`true`，否则返回`false`
func IsEmpty(value any) (isEmpty bool) {
	if value == nil {
		return true
	}
	// 首先会使用断言来检查变量是否是常见的类型以提高性能，然后再使用反射
	switch result := value.(type) {
	case int:
		return result == 0
	case int8:
		return result == 0
	case int16:
		return result == 0
	case int32:
		return result == 0
	case int64:
		return result == 0
	case uint:
		return result == 0
	case uint8:
		return result == 0
	case uint16:
		return result == 0
	case uint32:
		return result == 0
	case uint64:
		return result == 0
	case float32:
		return result == 0
	case float64:
		return result == 0
	case bool:
		return !result
	case string:
		return result == ""
	case []byte:
		return len(result) == 0
	case []rune:
		return len(result) == 0
	case []int:
		return len(result) == 0
	case []string:
		return len(result) == 0
	case []float32:
		return len(result) == 0
	case []float64:
		return len(result) == 0
	case map[string]any:
		return len(result) == 0
	default:
		if f, ok := value.(iTime); ok {
			if f == (*time.Time)(nil) {
				return true
			}
			return f.IsZero()
		}
		if f, ok := value.(iString); ok {
			if f == nil {
				return true
			}
			return f.String() == ""
		}
		if f, ok := value.(iInterfaces); ok {
			if f == nil {
				return true
			}
			return len(f.Interfaces()) == 0
		}
		if f, ok := value.(iMapStrAny); ok {
			if f == nil {
				return true
			}
			return len(f.MapStrAny()) == 0
		}
		// 使用反射
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			rv = reflect.ValueOf(value)
		}
		switch rv.Kind() {
		case reflect.Bool:
			return !rv.Bool()
		case
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			return rv.Int() == 0
		case
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			return rv.Uint() == 0
		case
			reflect.Float32,
			reflect.Float64:
			return rv.Float() == 0
		case reflect.String:
			return rv.Len() == 0
		case reflect.Struct:
			var fieldValueInterface any
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflection.ValueToInterface(rv.Field(i))
				if !IsEmpty(fieldValueInterface) {
					return false
				}
			}
			return true
		case
			reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0
		case
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return true
			}
		}
	}
	return false
}

// 检查给定的 value 是否为 nil
//
//	参数`traceSource`用于跟踪源变量，如果给定的`value`是指向指针的指针类型，则返回`nil`，
//	当`traceSource`为`true`时，如果源为`nil`，则返回`nil`，
//	注意，这可能使用反射功能，会稍微影响性能
func IsNil(value any, traceSource ...bool) (isNil bool) {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Interface,
		reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}
	}
	return false
}
