/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-18 00:54:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-03 16:36:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_any_test.go
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

type AddressNoTag struct {
	Street string
	City   string
}

type PersonNoTag struct {
	Name    string
	Age     int
	Address AddressNoTag
	sex     int
}

type AddressTag struct {
	Street string `json:"street" dc:"street"`
	City   string `json:"city" dc:"city"`
}

type PersonTag struct {
	Name    string      `json:"name" dc:"name"`
	Age     int         `json:"age" dc:"age"`
	Address *AddressTag `json:"address" dc:"address"`
	sex     int         `json:"sex" dc:"sex"`
}

func TestToStringMapE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringMapE(map[any]any{"a": "hello", "b": []any{"hello", "true"}, "c": map[string]any{"a": "hello", "b": []any{"hello", "true"}}}) // map[any]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"a": "hello", "b": []any{"hello", "true"}, "c": map[string]any{"a": "hello", "b": []any{"hello", "true"}}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE([]byte(`{"a":"hello","b":["hello","true"],"c":{"a":"hello","b":["hello","true"]}}`)) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"a": "hello", "b": []any{"hello", "true"}, "c": map[string]any{"a": "hello", "b": []any{"hello", "true"}}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(`{"a":"hello","b":["hello","true"],"c":{"a":"hello","b":["hello","true"]}}`) // json
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"a": "hello", "b": []any{"hello", "true"}, "c": map[string]any{"a": "hello", "b": []any{"hello", "true"}}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(map[string]string{"a": "hello", "b": "world", "c": "true"}) // map[string]string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"a": "hello", "b": "world", "c": "true"}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(PersonNoTag{Name: "lsx", Age: 18, Address: AddressNoTag{Street: "hz-123", City: "hz"}, sex: 1}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"Name": "lsx", "Age": 18, "Address": map[string]any{"City": "hz", "Street": "hz-123"}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(&PersonNoTag{Name: "lsx", Age: 18, Address: AddressNoTag{Street: "hz-123", City: "hz"}, sex: 1}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"Name": "lsx", "Age": 18, "Address": map[string]any{"City": "hz", "Street": "hz-123"}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(PersonTag{Name: "lsx", Age: 18, Address: &AddressTag{Street: "hz-123", City: "hz"}, sex: 1}) // struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"name": "lsx", "age": 18, "address": map[string]any{"city": "hz", "street": "hz-123"}}, actualObj)
	}
	actualObj, err = nconv.ToStringMapE(&PersonTag{Name: "lsx", Age: 18, Address: &AddressTag{Street: "hz-123", City: "hz"}, sex: 1}) // *struct
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(map[string]any{"name": "lsx", "age": 18, "address": map[string]any{"city": "hz", "street": "hz-123"}}, actualObj)
	}
}
