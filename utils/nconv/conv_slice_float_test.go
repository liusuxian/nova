/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 22:16:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 01:11:15
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_float_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"encoding/json"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToFloat32s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]any{1, 2, 3}))
	assert.Equal([]float32{1.0, 2.0, 3.0}, nconv.ToFloat32s([]any{1, 2, 3}))
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]int{1, 2, 3}))
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]int8{1, 2, 3}))
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]int16{1, 2, 3}))
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]int32{1, 2, 3}))
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]float32{1, 2, 3})
	assert.Equal([]float32{1, 2, 3}, nconv.ToFloat32s(f))
}

func TestToFloat64s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]any{1, 2, 3}))
	assert.Equal([]float64{1.0, 2.0, 3.0}, nconv.ToFloat64s([]any{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]int{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]int8{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]int16{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]int32{1, 2, 3}))
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]float64{1, 2, 3})
	assert.Equal([]float64{1, 2, 3}, nconv.ToFloat64s(f))
}
