/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-03 17:10:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:43:36
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

type FFF struct {
	A any
	B any
	C any
}

type GGG struct {
	A any `json:"a" dc:"a"`
	B any `json:"b" dc:"b"`
	C any `json:"c" dc:"c"`
}

func TestToStringMapInt64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt64E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(map[string]string{}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(FFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(&FFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(GGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(&GGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapInt32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt32E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(map[string]string{}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(FFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(&FFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(GGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(&GGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapInt16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt16E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(map[string]string{}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(FFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(&FFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(GGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(&GGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapInt8E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt8E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(map[string]string{}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(FFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(&FFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(GGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(&GGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapIntE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapIntE(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(map[string]string{}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(FFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(&FFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(GGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(&GGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}
