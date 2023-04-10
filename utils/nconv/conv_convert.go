/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 21:56:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-11 00:57:54
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_convert.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"reflect"
	"time"
)

// doConvertInput
type doConvertInput struct {
	FromValue              any    // 被转换的值
	ToTypeName             string // 转换值类型名称
	ReferValue             any    // 转换值，ToTypeName 类型
	Extra                  []any  // 用于实现转换的额外值
	alreadySetToReferValue bool   // 表示值已经转换并设置为 ReferValue，调用者可以忽略返回的结果。这是一个内部使用的属性
}

// Convert 将变量 fromValue 转换为类型 toTypeName
//
//	可选参数`extraParams`用于此转换所需的附加参数
//	它支持常见类型的转换，因为其基于类型名称的转换
func Convert(fromValue any, toTypeName string, extraParams ...any) (val any) {
	return doConvert(doConvertInput{
		FromValue:  fromValue,
		ToTypeName: toTypeName,
		ReferValue: nil,
		Extra:      extraParams,
	})
}

// doConvert 通用类型转换
func doConvert(in doConvertInput) (convertedValue any) {
	switch in.ToTypeName {
	case "int":
		return ToInt(in.FromValue)
	case "*int":
		if _, ok := in.FromValue.(*int); ok {
			return in.FromValue
		}
		v := ToInt(in.FromValue)
		return &v
	case "int8":
		return ToInt8(in.FromValue)
	case "*int8":
		if _, ok := in.FromValue.(*int8); ok {
			return in.FromValue
		}
		v := ToInt8(in.FromValue)
		return &v
	case "int16":
		return ToInt16(in.FromValue)
	case "*int16":
		if _, ok := in.FromValue.(*int16); ok {
			return in.FromValue
		}
		v := ToInt16(in.FromValue)
		return &v
	case "int32":
		return ToInt32(in.FromValue)
	case "*int32":
		if _, ok := in.FromValue.(*int32); ok {
			return in.FromValue
		}
		v := ToInt32(in.FromValue)
		return &v
	case "int64":
		return ToInt64(in.FromValue)
	case "*int64":
		if _, ok := in.FromValue.(*int64); ok {
			return in.FromValue
		}
		v := ToInt64(in.FromValue)
		return &v
	case "uint":
		return ToUint(in.FromValue)
	case "*uint":
		if _, ok := in.FromValue.(*uint); ok {
			return in.FromValue
		}
		v := ToUint(in.FromValue)
		return &v
	case "uint8":
		return ToUint8(in.FromValue)
	case "*uint8":
		if _, ok := in.FromValue.(*uint8); ok {
			return in.FromValue
		}
		v := ToUint8(in.FromValue)
		return &v
	case "uint16":
		return ToUint16(in.FromValue)
	case "*uint16":
		if _, ok := in.FromValue.(*uint16); ok {
			return in.FromValue
		}
		v := ToUint16(in.FromValue)
		return &v
	case "uint32":
		return ToUint32(in.FromValue)
	case "*uint32":
		if _, ok := in.FromValue.(*uint32); ok {
			return in.FromValue
		}
		v := ToUint32(in.FromValue)
		return &v
	case "uint64":
		return ToUint64(in.FromValue)
	case "*uint64":
		if _, ok := in.FromValue.(*uint64); ok {
			return in.FromValue
		}
		v := ToUint64(in.FromValue)
		return &v
	case "float32":
		return ToFloat32(in.FromValue)
	case "*float32":
		if _, ok := in.FromValue.(*float32); ok {
			return in.FromValue
		}
		v := ToFloat32(in.FromValue)
		return &v
	case "float64":
		return ToFloat64(in.FromValue)
	case "*float64":
		if _, ok := in.FromValue.(*float64); ok {
			return in.FromValue
		}
		v := ToFloat64(in.FromValue)
		return &v
	case "bool":
		return ToBool(in.FromValue)
	case "*bool":
		if _, ok := in.FromValue.(*bool); ok {
			return in.FromValue
		}
		v := ToBool(in.FromValue)
		return &v
	case "string":
		return ToString(in.FromValue)
	case "*string":
		if _, ok := in.FromValue.(*string); ok {
			return in.FromValue
		}
		v := ToString(in.FromValue)
		return &v
	case "[]byte":
		return ToBytes(in.FromValue)
	case "[]int":
		return ToInts(in.FromValue)
	case "[]int32":
		return ToInt32s(in.FromValue)
	case "[]int64":
		return ToInt64s(in.FromValue)
	case "[]uint":
		return ToUints(in.FromValue)
	case "[]uint8":
		return ToBytes(in.FromValue)
	case "[]uint32":
		return ToUint32s(in.FromValue)
	case "[]uint64":
		return ToUint64s(in.FromValue)
	case "[]float32":
		return ToFloat32s(in.FromValue)
	case "[]float64":
		return ToFloat64s(in.FromValue)
	case "[]string":
		return ToStrings(in.FromValue)
	case "Time", "time.Time":
		return ToTime(in.FromValue)
	case "*time.Time":
		if _, ok := in.FromValue.(*time.Time); ok {
			return in.FromValue
		}
		v := ToTime(in.FromValue)
		return &v
	case "Duration", "time.Duration":
		return ToDuration(in.FromValue)
	case "*time.Duration":
		if _, ok := in.FromValue.(*time.Duration); ok {
			return in.FromValue
		}
		v := ToDuration(in.FromValue)
		return &v
	case "map[string]string":
		return ToMapStrStr(in.FromValue)
	case "map[string]any":
		return ToMap(in.FromValue)
	case "[]map[string]any":
		return ToMaps(in.FromValue)
	case "json.RawMessage":
		return ToBytes(in.FromValue)
	default:
		if in.ReferValue != nil {
			var referReflectValue reflect.Value
			if v, ok := in.ReferValue.(reflect.Value); ok {
				referReflectValue = v
			} else {
				referReflectValue = reflect.ValueOf(in.ReferValue)
			}
			defer func() {
				if recover() != nil {
					if err := bindVarToReflectValue(referReflectValue, in.FromValue, nil); err == nil {
						in.alreadySetToReferValue = true
						convertedValue = referReflectValue.Interface()
					}
				}
			}()
			in.ToTypeName = referReflectValue.Kind().String()
			in.ReferValue = nil
			return reflect.ValueOf(doConvert(in)).Convert(referReflectValue.Type()).Interface()
		}
		return in.FromValue
	}
}

// doConvertWithReflectValueSet
func doConvertWithReflectValueSet(reflectValue reflect.Value, in doConvertInput) {
	convertedValue := doConvert(in)
	if !in.alreadySetToReferValue {
		reflectValue.Set(reflect.ValueOf(convertedValue))
	}
}
