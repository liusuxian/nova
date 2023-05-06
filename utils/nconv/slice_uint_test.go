/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-19 11:58:36
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-06 14:17:15
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_uint_test.go
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

func TestToUint64SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint64SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint64{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint64{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint64{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint64{}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint64{}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint64{}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint64{}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint64{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint64SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint64{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToUint32SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint32SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint32{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint32{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint32{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint32{}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint32{}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint32{}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint32{}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint32{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint32SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint32{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToUint16SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint16SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint16{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint16{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint16{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint16{}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint16{}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint16{}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint16{}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint16{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint16SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint16{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToUint8SliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint8SliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint8{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint8{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint8{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint8{}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint8{}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint8{}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint8{}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint8{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToUint8SliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint8{1, 2, 1, 0, 1}, actualObj)
	}
}

func TestToUintSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUintSliceE([]any{1, 1.2, "1.56"}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint{1, 1, 1}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint{1, 0}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([]string{"1.57", "2", "0.1"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint{1, 2, 0}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint{}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint{}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint{}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.Error(err) {
		assert.ElementsMatch([]uint{}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE([]byte("[1, 2, true, \"0\", \"1.2\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint{1, 2, 1, 0, 1}, actualObj)
	}
	actualObj, err = nconv.ToUintSliceE("[1, 2, true, \"0\", \"1.2\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.ElementsMatch([]uint{1, 2, 1, 0, 1}, actualObj)
	}
}
