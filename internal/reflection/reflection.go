/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-04 17:16:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:10:23
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package reflection

import "reflect"

// OriginValueAndKindOutput
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value
	InputKind   reflect.Kind
	OriginValue reflect.Value
	OriginKind  reflect.Kind
}

// OriginValueAndKind 检索并返回原始的反射值和数据类型
func OriginValueAndKind(value any) (out OriginValueAndKindOutput) {
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind
	for out.OriginKind == reflect.Ptr {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

// OriginTypeAndKindOutput
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type
	InputKind  reflect.Kind
	OriginType reflect.Type
	OriginKind reflect.Kind
}

// OriginTypeAndKind 检索并返回原始的反射类型和数据类型
func OriginTypeAndKind(value any) (out OriginTypeAndKindOutput) {
	if value == nil {
		return
	}
	if reflectType, ok := value.(reflect.Type); ok {
		out.InputType = reflectType
	} else {
		if reflectValue, ok := value.(reflect.Value); ok {
			out.InputType = reflectValue.Type()
		} else {
			out.InputType = reflect.TypeOf(value)
		}
	}
	out.InputKind = out.InputType.Kind()
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind
	for out.OriginKind == reflect.Ptr {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// ValueToInterface 将 reflect 值转换为其 any 类型
func ValueToInterface(v reflect.Value) (value any, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	case reflect.Ptr:
		return ValueToInterface(v.Elem())
	case reflect.Interface:
		return ValueToInterface(v.Elem())
	default:
		return nil, false
	}
}
