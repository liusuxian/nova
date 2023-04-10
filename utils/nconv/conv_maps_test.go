/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 20:10:04
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 20:54:52
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_maps_test.go
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

type TestMaps struct {
	A string    `json:"a" dc:"a"`
	B int       `json:"b" dc:"b"`
	C *TestMaps `json:"c" dc:"c"`
	D bool      `json:"d" dc:"d"`
}

func TestToMaps(t *testing.T) {
	assert := assert.New(t)
	element1 := TestMaps{A: "hello", B: 1, C: &TestMaps{A: "hello", B: 1, C: &TestMaps{A: "hello", B: 1, D: false}, D: false}, D: false}
	element2 := TestMaps{A: "world", B: 2, C: &TestMaps{A: "world", B: 2, C: &TestMaps{A: "world", B: 2, D: true}, D: true}, D: true}
	ms := nconv.ToMaps([]any{element1, element2})
	assert.IsType([]map[string]any{}, ms)
	assert.IsType(&TestMaps{}, ms[0]["c"])
	assert.IsType(&TestMaps{}, ms[1]["c"])
}

func TestToMapsDeep(t *testing.T) {
	assert := assert.New(t)
	element1 := TestMaps{A: "hello", B: 1, C: &TestMaps{A: "hello", B: 1, C: &TestMaps{A: "hello", B: 1, D: false}, D: false}, D: false}
	element2 := TestMaps{A: "world", B: 2, C: &TestMaps{A: "world", B: 2, C: &TestMaps{A: "world", B: 2, D: true}, D: true}, D: true}
	ms := nconv.ToMapsDeep([]any{element1, element2})
	assert.IsType([]map[string]any{}, ms)
	assert.IsType(map[string]any{}, ms[0]["c"])
	assert.IsType(map[string]any{}, ms[1]["c"])
}
