/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 03:16:46
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 16:51:26
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/struct_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

type B struct {
	A int      `json:"a" dc:"a"`
	B float64  `json:"b" dc:"b"`
	C string   `json:"c" dc:"c"`
	D []string `json:"d" dc:"d"`
	E *B       `json:"e" dc:"e"`
}

func TestToStructE(t *testing.T) {
	assert := assert.New(t)
	val1 := &B{}
	err := nconv.ToStructE(`{"a":1,"b":1.2,"c":"hello","d":["hello","true"],"e":{"a":1,"b":1.2,"c":"hello","d":["hello","true"]}}`, &val1) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val1)
	}
	val2 := &B{}
	err = nconv.ToStructE(map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}, "e": map[string]any{"a": 1, "b": 1.2, "c": "hello", "d": []string{"hello", "true"}}}, &val2) // map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val2)
	}
	val3 := &B{}
	err = nconv.ToStructE(&B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, &val3) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(&B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}, E: &B{A: 1, B: 1.2, C: "hello", D: []string{"hello", "true"}}}, val3)
	}
	val4 := &B{}
	err = nconv.ToStructE("hello", &val4) // string
	errLog(t, err)
	if assert.Error(err) {
		t.Logf("val4: %+v\n", val4)
	}
}
