/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 15:25:32
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:43:08
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

type FFFFFF struct {
	A any
	B any
	C any
}

type GGGGGG struct {
	A any `json:"a" dc:"a"`
	B any `json:"b" dc:"b"`
	C any `json:"c" dc:"c"`
}

func TestToStringMapFloat64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapFloat64E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1, "b": 2.6, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1.6, "b": 2.7, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1.6, "b": 2.7, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1.6, "b": 2.7, "c": 3.1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(FFFFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"A": 1.6, "B": 0, "C": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(&FFFFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"A": 1.6, "B": 0, "C": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(GGGGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1.6, "b": 0, "c": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat64E(&GGGGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float64{"a": 1.6, "b": 0, "c": 2.7}, actualObj)
	}
}

func TestToStringMapFloat32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapFloat32E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1, "b": 2.6, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1.6, "b": 2.7, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1.6, "b": 2.7, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1.6, "b": 2.7, "c": 3.1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(FFFFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"A": 1.6, "B": 0, "C": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(&FFFFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"A": 1.6, "B": 0, "C": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(GGGGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1.6, "b": 0, "c": 2.7}, actualObj)
	}
	actualObj, err = nconv.ToStringMapFloat32E(&GGGGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]float32{"a": 1.6, "b": 0, "c": 2.7}, actualObj)
	}
}
