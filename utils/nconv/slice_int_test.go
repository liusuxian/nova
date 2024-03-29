/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:29:36
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:46:22
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

func TestToInt64SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt64SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int64{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int64{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int64{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int64{}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int64{}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int64{}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int64{}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int64{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt64SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int64{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToInt32SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt32SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int32{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int32{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int32{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int32{}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int32{}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int32{}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int32{}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int32{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt32SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int32{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToInt16SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt16SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int16{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int16{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int16{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int16{}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int16{}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int16{}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int16{}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int16{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt16SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int16{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToInt8SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt8SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int8{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int8{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int8{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int8{}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int8{}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int8{}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int8{}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int8{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToInt8SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int8{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToIntSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToIntSliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int{}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int{}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int{}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]int{}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToIntSliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]int{1, 2, 1, 0, 1}, actualObj)
	}
}
