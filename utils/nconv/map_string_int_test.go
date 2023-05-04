/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-03 16:47:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-04 11:27:57
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_int_test.go
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

type FFF struct {
	A int64
	B int64
	C int64
}

type GGG struct {
	A int64 `json:"a" dc:"a"`
	B int64 `json:"b" dc:"b"`
	C int64 `json:"c" dc:"c"`
}

type HHH struct {
	A int32
	B int32
	C int32
}

type III struct {
	A int32 `json:"a" dc:"a"`
	B int32 `json:"b" dc:"b"`
	C int32 `json:"c" dc:"c"`
}

type JJJ struct {
	A int16
	B int16
	C int16
}

type KKK struct {
	A int16 `json:"a" dc:"a"`
	B int16 `json:"b" dc:"b"`
	C int16 `json:"c" dc:"c"`
}

type LLL struct {
	A int8
	B int8
	C int8
}

type MMM struct {
	A int8 `json:"a" dc:"a"`
	B int8 `json:"b" dc:"b"`
	C int8 `json:"c" dc:"c"`
}

type NNN struct {
	A int
	B int
	C int
}

type OOO struct {
	A int `json:"a" dc:"a"`
	B int `json:"b" dc:"b"`
	C int `json:"c" dc:"c"`
}

func TestToStringMapInt64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt64E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E([]byte(`{"a":1,"b":2,"c":3}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(`{"a":1,"b":2,"c":3}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(map[string]string{"a": "1", "b": "2", "c": "3"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(FFF{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(&FFF{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(GGG{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt64E(&GGG{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int64{"a": 1, "b": 2, "c": 3}, actualObj)
	}
}

func TestToStringMapInt32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt32E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E([]byte(`{"a":1,"b":2,"c":3}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(`{"a":1,"b":2,"c":3}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(map[string]string{"a": "1", "b": "2", "c": "3"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(HHH{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(&HHH{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(III{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt32E(&III{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int32{"a": 1, "b": 2, "c": 3}, actualObj)
	}
}

func TestToStringMapInt16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt16E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E([]byte(`{"a":1,"b":2,"c":3}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(`{"a":1,"b":2,"c":3}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(map[string]string{"a": "1", "b": "2", "c": "3"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(JJJ{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(&JJJ{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(KKK{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt16E(&KKK{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int16{"a": 1, "b": 2, "c": 3}, actualObj)
	}
}

func TestToStringMapInt8E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapInt8E(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E([]byte(`{"a":1,"b":2,"c":3}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(`{"a":1,"b":2,"c":3}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(map[string]string{"a": "1", "b": "2", "c": "3"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(LLL{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(&LLL{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(MMM{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapInt8E(&MMM{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int8{"a": 1, "b": 2, "c": 3}, actualObj)
	}
}

func TestToStringMapIntE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapIntE(map[any]any{"a": "1", "b": 2.6, "c": true}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 1}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE([]byte(`{"a":1,"b":2,"c":3}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(`{"a":1,"b":2,"c":3}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(map[string]string{"a": "1", "b": "2", "c": "3"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(NNN{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(&NNN{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"A": 1, "B": 2, "C": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(OOO{A: 1, B: 2, C: 3}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
	actualObj, err = nconv.ToStringMapIntE(&OOO{A: 1, B: 2, C: 3}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, actualObj)
	}
}
