/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 17:18:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 00:39:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_int_test.go
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

func TestToInt(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, nconv.ToInt(1))
	assert.Equal(1, nconv.ToInt(float32(1.0)))
	assert.Equal(1, nconv.ToInt(1.0))
	assert.Equal(-1, nconv.ToInt("-1.6"))
	assert.Equal(1, nconv.ToInt("1.6"))
}

func TestToInt8(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int8(1), nconv.ToInt8(1))
	assert.Equal(int8(1), nconv.ToInt8(float32(1.0)))
	assert.Equal(int8(1), nconv.ToInt8(1.0))
	assert.Equal(int8(-1), nconv.ToInt8("-1.6"))
	assert.Equal(int8(1), nconv.ToInt8("1.6"))
}

func TestToInt16(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int16(1), nconv.ToInt16(1))
	assert.Equal(int16(1), nconv.ToInt16(float32(1.0)))
	assert.Equal(int16(1), nconv.ToInt16(1.0))
	assert.Equal(int16(-1), nconv.ToInt16("-1.6"))
	assert.Equal(int16(1), nconv.ToInt16("1.6"))
}

func TestToInt32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int32(1), nconv.ToInt32(1))
	assert.Equal(int32(1), nconv.ToInt32(float32(1.0)))
	assert.Equal(int32(1), nconv.ToInt32(1.0))
	assert.Equal(int32(-1), nconv.ToInt32("-1.6"))
	assert.Equal(int32(1), nconv.ToInt32("1.6"))
}

func TestToInt64(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(1), nconv.ToInt64(1))
	assert.Equal(int64(1), nconv.ToInt64(float32(1.0)))
	assert.Equal(int64(1), nconv.ToInt64(1.0))
	assert.Equal(int64(-1), nconv.ToInt64("-1.6"))
	assert.Equal(int64(1), nconv.ToInt64("1.6"))
}
