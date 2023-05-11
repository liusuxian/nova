/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 17:56:55
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:50:46
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
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

// ToTimeE 将 any 转换为 time.Time 类型
func ToTimeE(i any) (iv time.Time, err error) {
	return ToTimeInDefaultLocationE(i, time.UTC)
}

// ToTimeInDefaultLocationE 将 any 转换为 time.Time 类型
func ToTimeInDefaultLocationE(i any, location *time.Location) (iv time.Time, err error) {
	i = indirect(i)

	switch val := i.(type) {
	case nil:
		return time.Time{}, nil
	case time.Time:
		return val, nil
	case string:
		return StringToDateInDefaultLocation(val, location)
	case json.Number:
		s, e := ToInt64E(val)
		if e != nil {
			return time.Time{}, convertError(i, "Time")
		}
		return time.Unix(s, 0), nil
	case int:
		return time.Unix(int64(val), 0), nil
	case int64:
		return time.Unix(val, 0), nil
	case int32:
		return time.Unix(int64(val), 0), nil
	case uint:
		return time.Unix(int64(val), 0), nil
	case uint64:
		return time.Unix(int64(val), 0), nil
	case uint32:
		return time.Unix(int64(val), 0), nil
	default:
		return time.Time{}, convertError(i, "Time")
	}
}

// ToDurationE 将 any 转换为 time.Duration 类型
func ToDurationE(i any) (iv time.Duration, err error) {
	i = indirect(i)

	switch val := i.(type) {
	case nil:
		return time.Duration(0), nil
	case time.Duration:
		return val, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return time.Duration(ToInt64(val)), nil
	case float32, float64:
		return time.Duration(ToFloat64(val)), nil
	case string:
		var d time.Duration
		var e error
		if strings.ContainsAny(val, "nsuµmh") {
			d, e = time.ParseDuration(val)
		} else {
			d, e = time.ParseDuration(val + "ns")
		}
		if e != nil {
			return time.Duration(0), convertError(i, "Duration")
		}
		return d, nil
	case json.Number:
		f, e := val.Float64()
		if e != nil {
			return time.Duration(0), convertError(i, "Duration")
		}
		return time.Duration(f), nil
	default:
		return time.Duration(0), convertError(i, "Duration")
	}
}

// ToDurationSliceE 将 any 转换为 []time.Duration 类型
func ToDurationSliceE(i any) ([]time.Duration, error) {
	if i == nil {
		return []time.Duration{}, nil
	}

	switch val := i.(type) {
	case []time.Duration:
		return val, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			iVal, err := ToDurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, convertError(i, "[]time.Duration")
			}
			a[j] = iVal
		}
		return a, nil
	default:
		return []time.Duration{}, convertError(i, "[]time.Duration")
	}
}

// StringToDateInDefaultLocation 将 string 转换为 time.Time 类型
func StringToDateInDefaultLocation(s string, location *time.Location) (iv time.Time, err error) {
	return parseDateWith(s, location, timeFormats)
}

// parseDateWith 解析时间
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
	return d, errors.Errorf("unable to parse date: %s", s)
}
