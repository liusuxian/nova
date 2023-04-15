/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:27:35
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 01:11:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_string_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type A struct {
	A int
	B float64
	C string
	D []any
}

func TestToStringE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("", actualObj)
	}
	actualObj, err = nconv.ToStringE(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("1", actualObj)
	}
	actualObj, err = nconv.ToStringE(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("1.56", actualObj)
	}
	actualObj, err = nconv.ToStringE(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("true", actualObj)
	}
	actualObj, err = nconv.ToStringE(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("false", actualObj)
	}
	actualObj, err = nconv.ToStringE([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("1.23", actualObj)
	}
	actualObj, err = nconv.ToStringE(time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC)) // time.Time
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("2023-04-15 00:00:00 +0000 UTC", actualObj)
	}
	actualObj, err = nconv.ToStringE([]any{1, 1.2, "hello", true}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("[1,1.2,\"hello\",true]", actualObj)
	}
	actualObj, err = nconv.ToStringE(map[string]any{"a": 1, "b": 1}) // map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("{\"a\":1,\"b\":1}", actualObj)
	}
	actualObj, err = nconv.ToStringE(&A{A: 1, B: 1.2, C: "hello", D: []any{1, 1.2, "hello", true}}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal("{\"A\":1,\"B\":1.2,\"C\":\"hello\",\"D\":[1,1.2,\"hello\",true]}", actualObj)
	}
}
