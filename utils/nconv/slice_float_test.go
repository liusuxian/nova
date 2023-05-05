/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 14:55:53
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 15:20:03
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_float_test.go
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

func TestToFloat64SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat64SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float64{1, 1.2, 1.56}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float64{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float64{1.57, 2, 0.1}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float64{}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float64{}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float64{}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float64{}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float64{1, 2, 1, 0, 1.2}, actualObj)
	}
	actualObj, err = nconv.ToFloat64SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float64{1, 2, 1, 0, 1.2}, actualObj)
	}
}

func TestToFloat32SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat32SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float32{1, 1.2, 1.56}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float32{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float32{1.57, 2, 0.1}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float32{}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float32{}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float32{}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]float32{}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float32{1, 2, 1, 0, 1.2}, actualObj)
	}
	actualObj, err = nconv.ToFloat32SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]float32{1, 2, 1, 0, 1.2}, actualObj)
	}
}
