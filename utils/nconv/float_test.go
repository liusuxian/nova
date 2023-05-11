/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:37:29
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:39:37
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

func TestToFloat64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat64E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.56), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("-1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(-1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float64(0), actualObj)
	}
}

func TestToFloat32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat32E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.56), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("-1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(-1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float32(0), actualObj)
	}
}
