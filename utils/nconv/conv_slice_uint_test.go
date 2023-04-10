/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:19:09
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 14:41:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_uint_test.go
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

func TestToUints(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]any{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]int{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]int8{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]int16{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]int32{1, 2, 3}))
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints([]int64{1, 2, 3}))
	f, _ := json.Marshal([]uint{1, 2, 3})
	assert.Equal([]uint{1, 2, 3}, nconv.ToUints(f))
}

func TestToUint8s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]any{1, 2, 3}))
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]int{1, 2, 3}))
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]int8{1, 2, 3}))
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]int16{1, 2, 3}))
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]int32{1, 2, 3}))
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]uint8{1, 2, 3})
	assert.Equal([]uint8{1, 2, 3}, nconv.ToUint8s(f))
}

func TestToUint16s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]any{1, 2, 3}))
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]int{1, 2, 3}))
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]int8{1, 2, 3}))
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]int16{1, 2, 3}))
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]int32{1, 2, 3}))
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]uint16{1, 2, 3})
	assert.Equal([]uint16{1, 2, 3}, nconv.ToUint16s(f))
}

func TestToUint32s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]any{1, 2, 3}))
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]int{1, 2, 3}))
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]int8{1, 2, 3}))
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]int16{1, 2, 3}))
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]int32{1, 2, 3}))
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]uint32{1, 2, 3})
	assert.Equal([]uint32{1, 2, 3}, nconv.ToUint32s(f))
}

func TestToUint64s(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]any{1, 2, 3}))
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]int{1, 2, 3}))
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]int8{1, 2, 3}))
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]int16{1, 2, 3}))
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]int32{1, 2, 3}))
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s([]int64{1, 2, 3}))
	f, _ := json.Marshal([]uint64{1, 2, 3})
	assert.Equal([]uint64{1, 2, 3}, nconv.ToUint64s(f))
}
