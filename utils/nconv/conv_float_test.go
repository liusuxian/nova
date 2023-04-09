/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 16:09:39
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 00:30:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_float_test.go
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

func TestToFloat32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(float32(1.0), nconv.ToFloat32(1.0))
	assert.Equal(float32(1.0), nconv.ToFloat32(1))
	assert.Equal(float32(1.0), nconv.ToFloat32("1"))
	assert.Equal(float32(-1.23), nconv.ToFloat32("-1.23"))
	assert.Equal(float32(1.23), nconv.ToFloat32("1.23"))
}

func TestToFloat64(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1.0, nconv.ToFloat64(1.0))
	assert.Equal(1.0, nconv.ToFloat64(1))
	assert.Equal(1.0, nconv.ToFloat64("1"))
	assert.Equal(-1.23, nconv.ToFloat64("-1.23"))
	assert.Equal(1.23, nconv.ToFloat64("1.23"))
}
