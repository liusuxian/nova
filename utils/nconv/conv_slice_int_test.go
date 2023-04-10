/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:47:12
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 14:38:25
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_int_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInts(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]any{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]int{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]int8{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]int16{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]int32{1, 2, 3}))
	assert.Equal([]int{1, 2, 3}, nconv.ToInts([]int64{1, 2, 3}))
	f, _ := json.Marshal([]int{1, 2, 3})
	assert.Equal([]int{1, 2, 3}, nconv.ToInts(f))
}

func TestToInt8s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]any{1, 2, 3}))
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]int{1, 2, 3}))
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]int8{1, 2, 3}))
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]int16{1, 2, 3}))
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]int32{1, 2, 3}))
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]int8{1, 2, 3})
	assert.Equal([]int8{1, 2, 3}, nconv.ToInt8s(f))
}

func TestToInt16s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]any{1, 2, 3}))
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]int{1, 2, 3}))
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]int8{1, 2, 3}))
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]int16{1, 2, 3}))
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]int32{1, 2, 3}))
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]int16{1, 2, 3})
	assert.Equal([]int16{1, 2, 3}, nconv.ToInt16s(f))
}

func TestToInt32s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]any{1, 2, 3}))
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]int{1, 2, 3}))
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]int8{1, 2, 3}))
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]int16{1, 2, 3}))
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]int32{1, 2, 3}))
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]int32{1, 2, 3})
	assert.Equal([]int32{1, 2, 3}, nconv.ToInt32s(f))
}

func TestToInt64s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]any{1, 2, 3}))
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]int{1, 2, 3}))
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]int8{1, 2, 3}))
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]int16{1, 2, 3}))
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]int32{1, 2, 3}))
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]int64{1, 2, 3})
	assert.Equal([]int64{1, 2, 3}, nconv.ToInt64s(f))
}
