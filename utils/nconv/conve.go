/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 13:31:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-14 17:30:39
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"github.com/pkg/errors"
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
		return false, errors.Errorf("unable to convert %#v of type %T to bool", i, i)
	case json.Number:
		iv, err := ToInt64E(val)
		if err == nil {
			return iv > 0, nil
		}
		return false, errors.Errorf("unable to convert %#v of type %T to bool", i, i)
	default:
		return false, errors.Errorf("unable to convert %#v of type %T to bool", i, i)
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
		return 0, errors.Errorf("unable to convert %#v of type %T to int64", i, i)
	case json.Number:
		return ToInt64E(string(val))
	default:
		return 0, errors.Errorf("unable to convert %#v of type %T to int64", i, i)
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
		return 0, errors.Errorf("unable to convert %#v of type %T to int32", i, i)
	case json.Number:
		return ToInt32E(string(val))
	default:
		return 0, errors.Errorf("unable to convert %#v of type %T to int32", i, i)
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
		return 0, errors.Errorf("unable to convert %#v of type %T to int16", i, i)
	case json.Number:
		return ToInt16E(string(val))
	default:
		return 0, errors.Errorf("unable to convert %#v of type %T to int16", i, i)
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
		return 0, errors.Errorf("unable to convert %#v of type %T to int8", i, i)
	case json.Number:
		return ToInt8E(string(val))
	default:
		return 0, errors.Errorf("unable to convert %#v of type %T to int8", i, i)
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
		return 0, errors.Errorf("unable to convert %#v of type %T to int", i, i)
	case json.Number:
		return ToIntE(string(val))
	default:
		return 0, errors.Errorf("unable to convert %#v of type %T to int", i, i)
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
