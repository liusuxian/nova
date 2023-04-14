/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 13:31:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-15 03:04:24
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liusuxian/nova/internal/reflection"
	"github.com/pkg/errors"
	"html/template"
	"reflect"
	"strconv"
	"time"
)

// ToBoolE 将 any 转换为 bool 类型
func ToBoolE(i any) (bl bool, err error) {
	i = indirect(i)

	switch val := i.(type) {
	case nil:
		return false, nil
	case bool:
		return val, nil
	case int64:
		return val > 0, nil
	case int32:
		return val > 0, nil
	case int16:
		return val > 0, nil
	case int8:
		return val > 0, nil
	case int:
		return val > 0, nil
	case uint64:
		return val > 0, nil
	case uint32:
		return val > 0, nil
	case uint16:
		return val > 0, nil
	case uint8:
		return val > 0, nil
	case uint:
		return val > 0, nil
	case []byte:
		return ToBoolE(string(val))
	case string:
		if val == "" {
			return false, nil
		}
		iv, err := strconv.ParseBool(val)
		if err == nil {
			return iv, nil
		}
		return false, convertError(i, "bool")
	case json.Number:
		iv, err := ToInt64E(val)
		if err == nil {
			return iv > 0, nil
		}
		return false, convertError(i, "bool")
	default:
		return false, convertError(i, "bool")
	}
}

// ToFloat64E 将 any 转换为 float64 类型
func ToFloat64E(i any) (iv float64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToFloat64E(string(val))
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return v, nil
		}
		return 0, convertError(i, "float64")
	case json.Number:
		v, err := val.Float64()
		if err == nil {
			return v, nil
		}
		return 0, convertError(i, "float64")
	default:
		return 0, convertError(i, "float64")
	}
}

// ToFloat32E 将 any 转换为 float32 类型
func ToFloat32E(i any) (iv float32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case float64:
		return float32(val), nil
	case float32:
		return val, nil
	case int64:
		return float32(val), nil
	case int32:
		return float32(val), nil
	case int16:
		return float32(val), nil
	case int8:
		return float32(val), nil
	case uint64:
		return float32(val), nil
	case uint32:
		return float32(val), nil
	case uint16:
		return float32(val), nil
	case uint8:
		return float32(val), nil
	case uint:
		return float32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToFloat32E(string(val))
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return float32(v), nil
		}
		return 0, convertError(i, "float32")
	case json.Number:
		v, err := val.Float64()
		if err == nil {
			return float32(v), nil
		}
		return 0, convertError(i, "float32")
	default:
		return 0, convertError(i, "float32")
	}
}

// ToInt64E 将 any 转换为 int64 类型
func ToInt64E(i any) (iv int64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return val, nil
	case int32:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt64E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return ipv, nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int64(ipf), nil
		}
		return 0, convertError(i, "int64")
	case json.Number:
		return ToInt64E(string(val))
	default:
		return 0, convertError(i, "int64")
	}
}

// ToInt32E 将 any 转换为 int32 类型
func ToInt32E(i any) (iv int32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int32(val), nil
	case int32:
		return val, nil
	case int16:
		return int32(val), nil
	case int8:
		return int32(val), nil
	case uint64:
		return int32(val), nil
	case uint32:
		return int32(val), nil
	case uint16:
		return int32(val), nil
	case uint8:
		return int32(val), nil
	case uint:
		return int32(val), nil
	case float64:
		return int32(val), nil
	case float32:
		return int32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt32E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int32(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int32(ipf), nil
		}
		return 0, convertError(i, "int32")
	case json.Number:
		return ToInt32E(string(val))
	default:
		return 0, convertError(i, "int32")
	}
}

// ToInt16E 将 any 转换为 int16 类型
func ToInt16E(i interface{}) (iv int16, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int16(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int16(val), nil
	case int32:
		return int16(val), nil
	case int16:
		return val, nil
	case int8:
		return int16(val), nil
	case uint64:
		return int16(val), nil
	case uint32:
		return int16(val), nil
	case uint16:
		return int16(val), nil
	case uint8:
		return int16(val), nil
	case uint:
		return int16(val), nil
	case float64:
		return int16(val), nil
	case float32:
		return int16(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt16E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int16(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int16(ipf), nil
		}
		return 0, convertError(i, "int16")
	case json.Number:
		return ToInt16E(string(val))
	default:
		return 0, convertError(i, "int16")
	}
}

// ToInt8E 将 any 转换为 int8 类型
func ToInt8E(i any) (iv int8, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int8(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int8(val), nil
	case int32:
		return int8(val), nil
	case int16:
		return int8(val), nil
	case int8:
		return val, nil
	case uint64:
		return int8(val), nil
	case uint32:
		return int8(val), nil
	case uint16:
		return int8(val), nil
	case uint8:
		return int8(val), nil
	case uint:
		return int8(val), nil
	case float64:
		return int8(val), nil
	case float32:
		return int8(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToInt8E(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int8(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int8(ipf), nil
		}
		return 0, convertError(i, "int8")
	case json.Number:
		return ToInt8E(string(val))
	default:
		return 0, convertError(i, "int8")
	}
}

// ToIntE 将 any 转换为 int 类型
func ToIntE(i interface{}) (iv int, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return intv, nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		return int(val), nil
	case int32:
		return int(val), nil
	case int16:
		return int(val), nil
	case int8:
		return int(val), nil
	case uint64:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint:
		return int(val), nil
	case float64:
		return int(val), nil
	case float32:
		return int(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToIntE(string(val))
	case string:
		ipv, err := strconv.ParseInt(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return int(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return int(ipf), nil
		}
		return 0, convertError(i, "int")
	case json.Number:
		return ToIntE(string(val))
	default:
		return 0, convertError(i, "int")
	}
}

// ToUint64E 将 any 转换为 uint64 类型
func ToUint64E(i any) (iv uint64, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case uint64:
		return val, nil
	case uint32:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case uint:
		return uint64(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint64")
		}
		return uint64(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint64E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return ipv, nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint64")
			}
			return uint64(ipf), nil
		}
		return 0, convertError(i, "uint64")
	case json.Number:
		return ToUint64E(string(val))
	default:
		return 0, convertError(i, "uint64")
	}
}

// ToUint32E 将 any 转换为 uint32 类型
func ToUint32E(i any) (iv uint32, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case uint64:
		return uint32(val), nil
	case uint32:
		return val, nil
	case uint16:
		return uint32(val), nil
	case uint8:
		return uint32(val), nil
	case uint:
		return uint32(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint32")
		}
		return uint32(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint32E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint32(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint32")
			}
			return uint32(ipf), nil
		}
		return 0, convertError(i, "uint32")
	case json.Number:
		return ToUint32E(string(val))
	default:
		return 0, convertError(i, "uint32")
	}
}

// ToUint16E 将 any 转换为 uint16 类型
func ToUint16E(i any) (iv uint16, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case uint64:
		return uint16(val), nil
	case uint32:
		return uint16(val), nil
	case uint16:
		return val, nil
	case uint8:
		return uint16(val), nil
	case uint:
		return uint16(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint16")
		}
		return uint16(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint16E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint16(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint16")
			}
			return uint16(ipf), nil
		}
		return 0, convertError(i, "uint16")
	case json.Number:
		return ToUint16E(string(val))
	default:
		return 0, convertError(i, "uint16")
	}
}

// ToUint8E 将 any 转换为 uint8 类型
func ToUint8E(i any) (iv uint8, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case uint64:
		return uint8(val), nil
	case uint32:
		return uint8(val), nil
	case uint16:
		return uint8(val), nil
	case uint8:
		return val, nil
	case uint:
		return uint8(val), nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint8")
		}
		return uint8(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUint8E(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint8(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint8")
			}
			return uint8(ipf), nil
		}
		return 0, convertError(i, "uint8")
	case json.Number:
		return ToUint8E(string(val))
	default:
		return 0, convertError(i, "uint8")
	}
}

// ToUintE 将 any 转换为 uint 类型
func ToUintE(i any) (iv uint, err error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(intv), nil
	}

	switch val := i.(type) {
	case nil:
		return 0, nil
	case int64:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int32:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int16:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case int8:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case uint64:
		return uint(val), nil
	case uint32:
		return uint(val), nil
	case uint16:
		return uint(val), nil
	case uint8:
		return uint(val), nil
	case uint:
		return val, nil
	case float64:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case float32:
		if val < 0 {
			return 0, convertError(i, "uint")
		}
		return uint(val), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		return ToUintE(string(val))
	case string:
		ipv, err := strconv.ParseUint(trimZeroDecimal(val), 0, 0)
		if err == nil {
			return uint(ipv), nil
		}
		ipf, err := strconv.ParseFloat(val, 64)
		if err == nil {
			if ipf < 0 {
				return 0, convertError(i, "uint")
			}
			return uint(ipf), nil
		}
		return 0, convertError(i, "uint")
	case json.Number:
		return ToUintE(string(val))
	default:
		return 0, convertError(i, "uint")
	}
}

// ToStringE 将 any 转换为 string 类型
func ToStringE(i any) (iv string, err error) {
	i = indirectToStringerOrError(i)

	switch val := i.(type) {
	case nil:
		return "", nil
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case bool:
		return strconv.FormatBool(val), nil
	case json.Number:
		return val.String(), nil
	case template.HTML:
		return string(val), nil
	case template.URL:
		return string(val), nil
	case template.JS:
		return string(val), nil
	case template.CSS:
		return string(val), nil
	case template.HTMLAttr:
		return string(val), nil
	case fmt.Stringer:
		return val.String(), nil
	case error:
		return val.Error(), nil
	default:
		// 使用 json.Marshal 函数进行转换
		jsonContent, err := json.Marshal(val)
		if err == nil {
			return string(jsonContent), nil
		}
		return "", convertError(i, "string")
	}
}

// ToSliceE 将 any 转换为 []any 类型
func ToSliceE(i any) (iv []any, err error) {
	if i == nil {
		return []any{}, convertError(i, "[]any")
	}

	switch val := i.(type) {
	case []any:
		return val, nil
	case []int64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int16:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int8:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []int:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint16:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint8:
		if json.Valid(val) {
			if err := unmarshalUseNumber(val, &iv); err != nil {
				return []any{}, convertError(i, "[]any")
			}
			return
		}
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []uint:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []float64:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []float32:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []bool:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []string:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case [][]byte:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	case []map[string]any:
		iv = make([]any, len(val))
		for k, v := range val {
			iv[k] = v
		}
		return
	default:
		// 检查给定的 i 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
		if checkJsonAndUnmarshalUseNumber(i, &iv) {
			return
		}
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]any, length)
			for i := 0; i < length; i++ {
				iv[i] = originValueAndKind.OriginValue.Index(i).Interface()
			}
			return
		}

		return []any{}, convertError(i, "[]any")
	}
}

// ToBoolSliceE  将 any 转换为 []bool 类型
func ToBoolSliceE(i any) (iv []bool, err error) {
	if i == nil {
		return []bool{}, convertError(i, "[]bool")
	}

	switch val := i.(type) {
	case []bool:
		return val, nil
	case []any:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int64:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int32:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int16:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int8:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []int:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint64:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint32:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint16:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint8:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []uint:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case []string:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	case [][]byte:
		iv = make([]bool, len(val))
		for k, v := range val {
			bl, err := ToBoolE(v)
			if err != nil {
				return []bool{}, convertError(i, "[]bool")
			}
			iv[k] = bl
		}
		return
	default:
		// 使用反射进行转换
		originValueAndKind := reflection.OriginValueAndKind(i)
		originKind := originValueAndKind.OriginKind
		if originKind == reflect.Slice || originKind == reflect.Array {
			length := originValueAndKind.OriginValue.Len()
			iv = make([]bool, length)
			for i := 0; i < length; i++ {
				bl, err := ToBoolE(originValueAndKind.OriginValue.Index(i).Interface())
				if err != nil {
					return []bool{}, convertError(i, "[]bool")
				}
				iv[i] = bl
			}
			return
		}

		return []bool{}, convertError(i, "[]bool")
	}
}

// indirect 对给定的值进行多次解引用以达到基本类型（或 nil）
func indirect(i any) (iv any) {
	if i == nil {
		return nil
	}
	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		// 如果不是指针类型，避免创建 reflect.Value
		return i
	}
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// indirectToStringerOrError 通过解引用直到达到基本类型（或 nil）或实现了 fmt.Stringer 或 error 接口的对象
func indirectToStringerOrError(i any) (iv any) {
	if i == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(i)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// toInt 如果 i 或 i 的底层类型是 int，则返回 i 的 int 值
func toInt(i any) (iv int, bl bool) {
	switch i := i.(type) {
	case int:
		return i, true
	case time.Weekday:
		return int(i), true
	case time.Month:
		return int(i), true
	default:
		return 0, false
	}
}

// trimZeroDecimal 删除字符串中末尾的零和小数点
func trimZeroDecimal(s string) (v string) {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}

// unmarshalUseNumber 使用 number 选项将 JSON 数据字节解码为目标接口
func unmarshalUseNumber(data []byte, v any) (err error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err = decoder.Decode(v); err != nil {
		err = errors.Wrap(err, `json.UnmarshalUseNumber failed`)
	}
	return
}

// checkJsonAndUnmarshalUseNumber 检查给定的 i 是否为 JSON 格式的字符串值，并使用 unmarshalUseNumber 进行转换
func checkJsonAndUnmarshalUseNumber(i, iv any) (isJson bool) {
	switch val := i.(type) {
	case []byte:
		if json.Valid(val) {
			if err := unmarshalUseNumber(val, &iv); err != nil {
				return false
			}
			return true
		}
	case string:
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := unmarshalUseNumber(anyBytes, &iv); err != nil {
				return false
			}
			return true
		}
	}
	return false
}

// convertError 转换错误
func convertError(i any, typ string) (err error) {
	return errors.Errorf("unable to convert %#v of type %T to %s", i, i, typ)
}
