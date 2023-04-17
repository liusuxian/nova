/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:21:26
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-18 00:11:40
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/slice_bool_test.go
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
	actualObj, err = nconv.ToBoolSliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal([]bool{}, actualObj)
	}
	actualObj, err = nconv.ToBoolSliceE([]byte("[1, 0, true, false, \"true\", \"false\"]")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]bool{true, false, true, false, true, false}, actualObj)
	}
	actualObj, err = nconv.ToBoolSliceE("[1, 0, true, false, \"true\", \"false\"]") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]bool{true, false, true, false, true, false}, actualObj)
	}
}
