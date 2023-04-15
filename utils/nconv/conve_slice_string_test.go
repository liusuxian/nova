/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 02:23:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 02:24:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_slice_string_test.go
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

func TestToStringSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToStringSliceE([]int{0, 1, 0}) // []int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"0", "1", "0"}, actualObj)
	}
	actualObj, err = nconv.ToStringSliceE([][]byte{[]byte("1"), []byte("0")}) // [][]byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"1", "0"}, actualObj)
	}
	actualObj, err = nconv.ToStringSliceE([]map[string]any{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"{\"a1\":1,\"b1\":2}", "{\"a2\":3,\"b2\":4}"}, actualObj)
	}
	actualObj, err = nconv.ToStringSliceE([]map[string]int{{"a1": 1, "b1": 2}, {"a2": 3, "b2": 4}}) // []map[string]int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"{\"a1\":1,\"b1\":2}", "{\"a2\":3,\"b2\":4}"}, actualObj)
	}
	actualObj, err = nconv.ToStringSliceE([]map[string]bool{{"a1": true, "b1": false}, {"a2": true, "b2": false}}) // []map[string]bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"{\"a1\":true,\"b1\":false}", "{\"a2\":true,\"b2\":false}"}, actualObj)
	}
	actualObj, err = nconv.ToStringSliceE([][]int{{1, 2}, {3, 4}}) // [][]int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]string{"[1,2]", "[3,4]"}, actualObj)
	}
}
