/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:22:49
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-15 13:43:42
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_slice_test.go
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

func TestToSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToSliceE("[1, 1.2, true, \"hello\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(1, nconv.ToInt(actualObj[0]))
		assert.Equal(1.2, nconv.ToFloat64(actualObj[1]))
		assert.Equal(true, actualObj[2])
		assert.Equal("hello", actualObj[3])
	}
	actualObj, err = nconv.ToSliceE("[\"h\", \"e\", \"l\", \"l\", \"o\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]any{"h", "e", "l", "l", "o"}, actualObj)
	}
	actualObj, err = nconv.ToSliceE(1) // int
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]any{}, actualObj)
	}
	actualObj, err = nconv.ToSliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]any{map[string]any{"a1": 1, "b1": 2}, map[string]any{"a2": 3, "b2": 4}}, actualObj)
	}
	actualObj, err = nconv.ToSliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]any{[]int{1, 2}, []int{3, 4}}, actualObj)
	}
}

func TestToBoolSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToBoolSliceE([]int{0, 1, 0}) // []int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]bool{false, true, false}, actualObj)
	}
	actualObj, err = nconv.ToBoolSliceE([]string{"true", "false"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]bool{true, false}, actualObj)
	}
	actualObj, err = nconv.ToBoolSliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]bool{true, false}, actualObj)
	}
}
