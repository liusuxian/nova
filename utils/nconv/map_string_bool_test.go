/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 14:29:29
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:42:06
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

type FFFFF struct {
	A any
	B any
	C any
}

type GGGGG struct {
	A any `json:"a" dc:"a"`
	B any `json:"b" dc:"b"`
	C any `json:"c" dc:"c"`
}

func TestToStringMapBoolE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapBoolE(map[any]any{"a": "1", "b": 2.6, "c": -1}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": true, "c": false}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE([]byte(`{"a": 1.6, "b": "1", "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": true, "c": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(`{"a": 1.6, "b": "1", "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": true, "c": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(map[string]string{"a": "1", "b": "0", "c": "1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": false, "c": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(FFFFF{A: 1.6, B: false, C: "1"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"A": true, "B": false, "C": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(&FFFFF{A: 1.6, B: false, C: "1"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"A": true, "B": false, "C": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(GGGGG{A: 1.6, B: false, C: "1"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": false, "c": true}, actualObj)
	}
	actualObj, err = nconv.ToStringMapBoolE(&GGGGG{A: 1.6, B: false, C: "1"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]bool{"a": true, "b": false, "c": true}, actualObj)
	}
}
