/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 15:15:07
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 20:07:29
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_map_test.go
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

type TestMap struct {
	A string   `json:"a" dc:"a"`
	B int      `json:"b" dc:"b"`
	C *TestMap `json:"c" dc:"c"`
	D bool     `json:"d" dc:"d"`
	e string   `json:"e" dc:"e"`
}

func TestToMap(t *testing.T) {
	assert := assert.New(t)
	m1 := nconv.ToMap(`{"a": 1, "b": 2.2, "c": "hello", "d": true, "e": {"a": "1", "b": "2"}}`)
	assert.Equal(1, nconv.ToInt(m1["a"]))
	assert.Equal(2.2, nconv.ToFloat64(m1["b"]))
	assert.Equal("hello", m1["c"])
	assert.True(nconv.ToBool(m1["d"]))
	assert.Equal(map[string]any{"a": "1", "b": "2"}, m1["e"])
	assert.Equal(map[string]any{"1": 2, "3": 4}, nconv.ToMap([]any{1, 2, 3, 4}))

	m2 := nconv.ToMap(TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, D: true}, D: true}, D: true})
	assert.Equal("hello", m2["a"])
	assert.Equal(1, nconv.ToInt(m2["b"]))
	assert.IsType(&TestMap{}, m2["c"])
	assert.True(nconv.ToBool(m2["d"]))
	assert.Nil(m2["e"])
}

func TestToMapDeep(t *testing.T) {
	assert := assert.New(t)
	m := nconv.ToMapDeep(TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, D: true}, D: true}, D: true})
	assert.Equal("hello", m["a"])
	assert.Equal(1, nconv.ToInt(m["b"]))
	assert.IsType(map[string]any{}, m["c"])
	assert.True(nconv.ToBool(m["d"]))
	assert.Nil(m["e"])
}

func TestToMapStrStr(t *testing.T) {
	assert := assert.New(t)
	m1 := nconv.ToMapStrStr(`{"a": 1, "b": 2.2, "c": "hello", "d": true, "e": {"a": "1", "b": "2"}}`)
	assert.Equal(1, nconv.ToInt(m1["a"]))
	assert.Equal(2.2, nconv.ToFloat64(m1["b"]))
	assert.Equal("hello", m1["c"])
	assert.True(nconv.ToBool(m1["d"]))
	assert.Equal(map[string]string{"a": "1", "b": "2"}, nconv.ToMapStrStr(m1["e"]))
	assert.Equal(map[string]string{"1": "2", "3": "4"}, nconv.ToMapStrStr([]any{1, 2, 3, 4}))

	m2 := nconv.ToMapStrStr(TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, D: true}, D: true}, D: true})
	assert.Equal("hello", m2["a"])
	assert.Equal(1, nconv.ToInt(m2["b"]))
	assert.IsType("", m2["c"])
	assert.True(nconv.ToBool(m2["d"]))
	assert.Empty(m2["e"])
}

func TestMapStrStrDeep(t *testing.T) {
	assert := assert.New(t)
	m := nconv.ToMapStrStrDeep(TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, C: &TestMap{A: "hello", B: 1, D: true}, D: true}, D: true})
	assert.Equal("hello", m["a"])
	assert.Equal(1, nconv.ToInt(m["b"]))
	assert.IsType(map[string]string{}, nconv.ToMapStrStrDeep(m["c"]))
	assert.True(nconv.ToBool(m["d"]))
	assert.Empty(m["e"])
}
