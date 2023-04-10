/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 22:47:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-11 00:50:04
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_time.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"time"
)

// timeFormatType 时间格式类型
type timeFormatType int

const (
	timeFormatNoTimezone timeFormatType = iota
	timeFormatNamedTimezone
	timeFormatNumericTimezone
	timeFormatNumericAndNamedTimezone
	timeFormatTimeOnly
)

// timeFormat 时间格式结构
type timeFormat struct {
	format string
	typ    timeFormatType
}

// timeFormats 时间格式
var timeFormats = []timeFormat{
	{time.RFC3339, timeFormatNumericTimezone},
	{"2006-01-02T15:04:05", timeFormatNoTimezone}, // iso8601 without timezone
	{time.RFC1123Z, timeFormatNumericTimezone},
	{time.RFC1123, timeFormatNamedTimezone},
	{time.RFC822Z, timeFormatNumericTimezone},
	{time.RFC822, timeFormatNamedTimezone},
	{time.RFC850, timeFormatNamedTimezone},
	{"2006-01-02 15:04:05.999999999 -0700 MST", timeFormatNumericAndNamedTimezone}, // Time.String()
	{"2006-01-02T15:04:05-0700", timeFormatNumericTimezone},                        // RFC3339 without timezone hh:mm colon
	{"2006-01-02 15:04:05Z0700", timeFormatNumericTimezone},                        // RFC3339 without T or timezone hh:mm colon
	{"2006-01-02 15:04:05", timeFormatNoTimezone},
	{time.ANSIC, timeFormatNoTimezone},
	{time.UnixDate, timeFormatNamedTimezone},
	{time.RubyDate, timeFormatNumericTimezone},
	{"2006-01-02 15:04:05Z07:00", timeFormatNumericTimezone},
	{"2006-01-02", timeFormatNoTimezone},
	{"02 Jan 2006", timeFormatNoTimezone},
	{"2006-01-02 15:04:05 -07:00", timeFormatNumericTimezone},
	{"2006-01-02 15:04:05 -0700", timeFormatNumericTimezone},
	{time.Kitchen, timeFormatTimeOnly},
	{time.Stamp, timeFormatTimeOnly},
	{time.StampMilli, timeFormatTimeOnly},
	{time.StampMicro, timeFormatTimeOnly},
	{time.StampNano, timeFormatTimeOnly},
}

// ToTime 将 any 转换为 time.Time 类型
func ToTime(val any) (t time.Time) {
	return ToTimeInDefaultLocation(val, time.UTC)
}

// ToTimeInDefaultLocation 将 any 转换为 time.Time 类型
func ToTimeInDefaultLocation(val any, location *time.Location) (t time.Time) {
	val = indirect(val)
	switch v := val.(type) {
	case time.Time:
		return v
	case string:
		return StringToDateInDefaultLocation(v, location)
	case json.Number:
		return time.Unix(ToInt64(v), 0)
	case int:
		return time.Unix(int64(v), 0)
	case int64:
		return time.Unix(v, 0)
	case int32:
		return time.Unix(int64(v), 0)
	case uint:
		return time.Unix(int64(v), 0)
	case uint64:
		return time.Unix(int64(v), 0)
	case uint32:
		return time.Unix(int64(v), 0)
	default:
		return time.Time{}
	}
}

// StringToDateInDefaultLocation 将 string 转换为 time.Time 类型
func StringToDateInDefaultLocation(s string, location *time.Location) (t time.Time) {
	d, err := parseDateWith(s, location, timeFormats)
	if err != nil {
		fmt.Printf("StringToDateInDefaultLocation Error: %+v\n", err)
	}
	return d
}

// StringToDate 将 string 转换为 time.Time 类型
func StringToDate(s string) (t time.Time) {
	d, err := parseDateWith(s, time.UTC, timeFormats)
	if err != nil {
		fmt.Printf("StringToDate Error: %+v\n", err)
	}
	return d
}

// ToDuration 将 any 转换为 time.Duration 类型
func ToDuration(val any) (d time.Duration) {
	var err error
	val = indirect(val)
	switch s := val.(type) {
	case time.Duration:
		return s
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		d = time.Duration(ToInt64(s))
		return
	case float32, float64:
		d = time.Duration(ToFloat64(s))
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		if err != nil {
			fmt.Printf("ToDuration Error: %+v\n", err)
		}
		return
	case json.Number:
		var v float64
		v, err = s.Float64()
		if err != nil {
			fmt.Printf("ToDuration Error: %+v\n", err)
		}
		d = time.Duration(v)
		return
	default:
		fmt.Printf("Unable To %#v Of Type %T ToDuration Error\n", val, val)
		return
	}
}

// ToDurations 将 any 转换为 []time.Duration 类型
func ToDurations(val any) (ds []time.Duration) {
	if val == nil {
		return []time.Duration{}
	}
	switch v := val.(type) {
	case []time.Duration:
		return v
	}
	kind := reflect.TypeOf(val).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(val)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			a[j] = ToDuration(s.Index(j).Interface())
		}
		return a
	default:
		return []time.Duration{}
	}
}

// indirect 将 value 解引用，直到达到基本类型（或nil）并返回
func indirect(a any) (val any) {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// 如果不是指针类型，避免创建反射（reflect）值
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// parseDateWith
func parseDateWith(s string, location *time.Location, formats []timeFormat) (d time.Time, e error) {
	for _, format := range formats {
		if d, e = time.Parse(format.format, s); e == nil {
			if format.typ <= timeFormatNamedTimezone {
				if location == nil {
					location = time.Local
				}
				year, month, day := d.Date()
				hour, min, sec := d.Clock()
				d = time.Date(year, month, day, hour, min, sec, d.Nanosecond(), location)
			}
			return
		}
	}
	return d, errors.Errorf("Unable To Parse Date: %s", s)
}
