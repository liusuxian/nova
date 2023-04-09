/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 12:48:18
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 22:30:06
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/internal/reflection"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var emptyStringMap = map[string]struct{}{
	"":      {},
	"0":     {},
	"no":    {},
	"off":   {},
	"false": {},
}

// ToByte
func ToByte(val any) (b byte) {
	if v, ok := val.(byte); ok {
		return v
	}
	return ToUint8(val)
}

// ToBytes
func ToBytes(val any) (bs []byte) {
	if val == nil {
		return nil
	}
	switch value := val.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		if f, ok := value.(iBytes); ok {
			return f.Bytes()
		}
		originValueAndKind := reflection.OriginValueAndKind(val)
		switch originValueAndKind.OriginKind {
		case reflect.Map:
			byteList, err := json.Marshal(val)
			if err != nil {
				fmt.Printf("ToBytes Error: %+v\n", err)
			}
			return byteList
		case reflect.Array, reflect.Slice:
			ok := true
			byteList := make([]byte, originValueAndKind.OriginValue.Len())
			for i := range byteList {
				int32Value := ToInt32(originValueAndKind.OriginValue.Index(i).Interface())
				if int32Value < 0 || int32Value > math.MaxUint8 {
					ok = false
					break
				}
				byteList[i] = byte(int32Value)
			}
			if ok {
				return byteList
			}
		}
		return leEncode(val)
	}
}

// ToRune
func ToRune(val any) (cVal rune) {
	if v, ok := val.(rune); ok {
		return v
	}
	return ToInt32(val)
}

// ToRunes
func ToRunes(val any) (cVals []rune) {
	if v, ok := val.([]rune); ok {
		return v
	}
	return []rune(ToString(val))
}

// ToString
func ToString(val any) (cVal string) {
	if val == nil {
		return ""
	}
	switch value := val.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// 空值检查
		if value == nil {
			return ""
		}
		if f, ok := value.(iString); ok {
			return f.String()
		}
		if f, ok := value.(iError); ok {
			return f.Error()
		}
		// 反射检查
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return ToString(rv.Elem().Interface())
		}
		// 最后，使用 json.Marshal 函数进行转换
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

// ToBool
func ToBool(val any) (cVal bool) {
	if val == nil {
		return false
	}
	switch value := val.(type) {
	case bool:
		return value
	case []byte:
		if _, ok := emptyStringMap[strings.ToLower(string(value))]; ok {
			return false
		}
		return true
	case string:
		if _, ok := emptyStringMap[strings.ToLower(value)]; ok {
			return false
		}
		return true
	default:
		if f, ok := value.(iBool); ok {
			return f.Bool()
		}
		rv := reflect.ValueOf(val)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			s := strings.ToLower(ToString(val))
			if _, ok := emptyStringMap[s]; ok {
				return false
			}
			return true
		}
	}
}

// checkJsonAndUnmarshalUseNumber 检查给定的 val 是否为 JSON 格式的字符串值，并使用 json.UnmarshalUseNumber 进行转换
func checkJsonAndUnmarshalUseNumber(val, target any) (isJson bool) {
	switch value := val.(type) {
	case []byte:
		if json.Valid(value) {
			if err := json.UnmarshalUseNumber(value, &target); err != nil {
				return false
			}
			return true
		}
	case string:
		anyAsBytes := []byte(value)
		if json.Valid(anyAsBytes) {
			if err := json.UnmarshalUseNumber(anyAsBytes, &target); err != nil {
				return false
			}
			return true
		}
	}
	return false
}

// leFillUpSize 使用 LittleEndian 填充字节切片 b 到给定的长度 l
func leFillUpSize(b []byte, l int) (bs []byte) {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}

// leEncode 将一个或多个 values 使用 LittleEndian 编码为字节
func leEncode(values ...any) (bs []byte) {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}
		switch value := values[i].(type) {
		case int:
			buf.Write(leEncodeInt(value))
		case int8:
			buf.Write(leEncodeInt8(value))
		case int16:
			buf.Write(leEncodeInt16(value))
		case int32:
			buf.Write(leEncodeInt32(value))
		case int64:
			buf.Write(leEncodeInt64(value))
		case uint:
			buf.Write(leEncodeUint(value))
		case uint8:
			buf.Write(leEncodeUint8(value))
		case uint16:
			buf.Write(leEncodeUint16(value))
		case uint32:
			buf.Write(leEncodeUint32(value))
		case uint64:
			buf.Write(leEncodeUint64(value))
		case bool:
			buf.Write(leEncodeBool(value))
		case string:
			buf.Write(leEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(leEncodeFloat32(value))
		case float64:
			buf.Write(leEncodeFloat64(value))
		default:
			if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
				fmt.Printf("LeEncode Error: %+v\n", err)
				buf.Write(leEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

// leEncodeInt
func leEncodeInt(i int) (bs []byte) {
	if i <= math.MaxInt8 {
		return leEncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return leEncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return leEncodeInt32(int32(i))
	} else {
		return leEncodeInt64(int64(i))
	}
}

// leEncodeUint
func leEncodeUint(i uint) (bs []byte) {
	if i <= math.MaxUint8 {
		return leEncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return leEncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return leEncodeUint32(uint32(i))
	} else {
		return leEncodeUint64(uint64(i))
	}
}

// leEncodeInt8
func leEncodeInt8(i int8) (bs []byte) {
	return []byte{byte(i)}
}

// leEncodeUint8
func leEncodeUint8(i uint8) (bs []byte) {
	return []byte{i}
}

// leEncodeInt16
func leEncodeInt16(i int16) (bs []byte) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

// leEncodeUint16
func leEncodeUint16(i uint16) (bs []byte) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

// leEncodeInt32
func leEncodeInt32(i int32) (bs []byte) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

// leEncodeUint32
func leEncodeUint32(i uint32) (bs []byte) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

// leEncodeInt64
func leEncodeInt64(i int64) (bs []byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

// leEncodeUint64
func leEncodeUint64(i uint64) (bs []byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

// leEncodeFloat32
func leEncodeFloat32(f float32) (bs []byte) {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

// leEncodeFloat64
func leEncodeFloat64(f float64) (bs []byte) {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

// leEncodeBool
func leEncodeBool(b bool) (bs []byte) {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// leEncodeString
func leEncodeString(s string) (bs []byte) {
	return []byte(s)
}
