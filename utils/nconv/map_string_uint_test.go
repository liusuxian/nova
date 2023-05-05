/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-04 14:02:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 14:34:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_uint_test.go
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

type FFFF struct {
	A any
	B any
	C any
}

type GGGG struct {
	A any `json:"a" dc:"a"`
	B any `json:"b" dc:"b"`
	C any `json:"c" dc:"c"`
}

func TestToStringMapUint64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapUint64E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(FFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(&FFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(GGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint64E(&GGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint64{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapUint32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapUint32E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(FFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(&FFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(GGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint32E(&GGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint32{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapUint16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapUint16E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(FFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(&FFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(GGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint16E(&GGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint16{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapUint8E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapUint8E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(FFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(&FFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(GGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUint8E(&GGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint8{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}

func TestToStringMapUintE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapUintE(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE([]byte(`{"a": "1.6", "b": 2.7, "c": true}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(`{"a": "1.6", "b": 2.7, "c": true}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(map[string]string{"a": "1.6", "b": "2.7", "c": "3.1"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(FFFF{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(&FFFF{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"A": 1, "B": 0, "C": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(GGGG{A: 1.6, B: false, C: "2.7"}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 0, "c": 2}, actualObj)
	}
	actualObj, err = nconv.ToStringMapUintE(&GGGG{A: 1.6, B: false, C: "2.7"}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]uint{"a": 1, "b": 0, "c": 2}, actualObj)
	}
}
