/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 14:42:33
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 14:45:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_uint_test.go
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

func TestToUint(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint(1), nconv.ToUint(1))
	assert.Equal(uint(1), nconv.ToUint(float32(1.0)))
	assert.Equal(uint(1), nconv.ToUint(1.0))
	assert.Equal(uint(0), nconv.ToUint("-1.6"))
	assert.Equal(uint(1), nconv.ToUint("1.6"))
}

func TestToUint8(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint8(1), nconv.ToUint8(1))
	assert.Equal(uint8(1), nconv.ToUint8(float32(1.0)))
	assert.Equal(uint8(1), nconv.ToUint8(1.0))
	assert.Equal(uint8(0), nconv.ToUint8("-1.6"))
	assert.Equal(uint8(1), nconv.ToUint8("1.6"))
}

func TestToUint16(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint16(1), nconv.ToUint16(1))
	assert.Equal(uint16(1), nconv.ToUint16(float32(1.0)))
	assert.Equal(uint16(1), nconv.ToUint16(1.0))
	assert.Equal(uint16(0), nconv.ToUint16("-1.6"))
	assert.Equal(uint16(1), nconv.ToUint16("1.6"))
}

func TestToUint32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint32(1), nconv.ToUint32(1))
	assert.Equal(uint32(1), nconv.ToUint32(float32(1.0)))
	assert.Equal(uint32(1), nconv.ToUint32(1.0))
	assert.Equal(uint32(0), nconv.ToUint32("-1.6"))
	assert.Equal(uint32(1), nconv.ToUint32("1.6"))
}

func TestToUint64(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint64(1), nconv.ToUint64(1))
	assert.Equal(uint64(1), nconv.ToUint64(float32(1.0)))
	assert.Equal(uint64(1), nconv.ToUint64(1.0))
	assert.Equal(uint64(0), nconv.ToUint64("-1.6"))
	assert.Equal(uint64(1), nconv.ToUint64("1.6"))
}
