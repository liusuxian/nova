/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 03:16:46
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-04 10:51:10
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/struct_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

type BBB struct {
	A int
	B float64
	C string
	D []string
	E *BBB
}

type CCC struct {
	Time1 time.Duration
	Time2 time.Time
}

type DDD struct {
	Time1 *time.Duration
	Time2 *time.Time
}

type EEE struct {
	IP    net.IP
	IPNet net.IPNet
	S     []any
}

func TestToStructE(t *testing.T) {
	assert := assert.New(t)
	val1 := &BBB{}
	err := nconv.ToStructE(`{"a":1,"b":1.2,"c":"hello","d":["hello","true"],"e":{"a":1,"b":1.2,"c":"hello","d":["hello","true"]}}`, &val1) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val1)
	}
	val2 := &BBB{}
	err = nconv.ToStructE(map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}, "e": map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}}}, &val2) // map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val2)
	}
	val3 := &BBB{}
	err = nconv.ToStructE(&BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, &val3) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val3)
	}
	val4 := &BBB{}
	err = nconv.ToStructE("hello", &val4) // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(&BBB{}, val4)
	}
	val5 := []*BBB{}
	err = nconv.ToStructE(`[{"a":1,"b":1.2,"c":"hello","d":["hello","true"],"e":{"a":1,"b":1.2,"c":"hello","d":["hello","true"]}}, {"a":1,"b":1.2,"c":"hello","d":["hello","true"],"e":{"a":1,"b":1.2,"c":"hello","d":["hello","true"]}}]`, &val5) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]*BBB{
			{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}},
			{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}},
		}, val5)
	}
	val6 := []*BBB{}
	err = nconv.ToStructE([]map[string]any{
		{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}, "e": map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}}},
		{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}, "e": map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}}},
	}, &val6) // []map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]*BBB{
			{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}},
			{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &BBB{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}},
		}, val6)
	}
	val7 := &CCC{}
	err = nconv.ToStructE(map[string]string{"Time1": "3s", "Time2": "2023-04-18 00:00:00"}, &val7) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&CCC{Time1: 3000000000, Time2: time.Date(2023, 4, 18, 0, 0, 0, 0, time.UTC)}, val7)
	}
	val8 := &DDD{}
	err = nconv.ToStructE(map[string]string{"Time1": "3s", "Time2": "2023-04-18 00:00:00"}, &val8) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		time1 := time.Duration(3000000000)
		time2 := time.Date(2023, 4, 18, 0, 0, 0, 0, time.UTC)
		assert.Equal(&DDD{Time1: &time1, Time2: &time2}, val8)
	}
	val9 := &EEE{}
	err = nconv.ToStructE(map[string]any{
		"IP":    "127.0.0.1",
		"IPNet": map[string]string{"IP": "127.0.0.1", "Mask": "255,255,255,0"},
		"S":     "1,1.2,true,hello",
	}, &val9) // map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&EEE{
			IP:    net.IPv4(127, 0, 0, 1),
			IPNet: net.IPNet{IP: net.IPv4(127, 0, 0, 1), Mask: net.IPv4Mask(255, 255, 255, 0)},
			S:     []any{"1", "1.2", "true", "hello"},
		}, val9)
	}
}
