/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 12:48:18
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 15:05:46
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_test.go
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

func TestToByte(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(byte('h'), nconv.ToByte('h'))
	assert.Equal(byte('h'), nconv.ToByte([]byte{'h', 'l'}))
}

func TestToBytes(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{'h', 'e', 'l', 'l', 'o'}, nconv.ToBytes([]any{'h', 'e', 'l', 'l', 'o'}))
	assert.Equal([]byte{'h', 'e', 'l', 'l', 'o'}, nconv.ToBytes("hello"))
}

func TestToRune(t *testing.T) {
	assert := assert.New(t)
	assert.Equal('中', nconv.ToRune('中'))
	assert.Equal('国', nconv.ToRune('国'))
}

func TestToRunes(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]rune{'中', '国'}, nconv.ToRunes("中国"))
}

func TestToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", nconv.ToString([]byte{'h', 'e', 'l', 'l', 'o'}))
}

func TestToBool(t *testing.T) {
	assert := assert.New(t)
	assert.False(nconv.ToBool(0))
	assert.True(nconv.ToBool(1.0))
	assert.True(nconv.ToBool(1))
	assert.True(nconv.ToBool("ok"))
}
