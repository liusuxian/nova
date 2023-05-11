/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:34:47
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:37:56
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func errLog(t *testing.T, err error) {
	if err != nil {
		t.Logf("Error: %+v\n", err.Error())
	}
}

func TestToBoolE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToBoolE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(byte('a')) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(0) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(-1) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(1) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte("true")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte("")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte{}) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("true") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE("false") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(" ") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("hello") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(float64(1.23)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
}
