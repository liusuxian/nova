/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:17:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-09 23:36:34
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nstr/str_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nstr_test

import (
	"github.com/liusuxian/nova/utils/nstr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimAll(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("alog", nstr.TrimAll("a.log", "."))
	assert.Equal("alog", nstr.TrimAll(" a . log ", "."))
	assert.Equal("ablog", nstr.TrimAll("a.b.log", "."))
	assert.Equal("ablog", nstr.TrimAll(" a . b . log", "."))
	assert.Equal("", nstr.TrimAll("", "."))
	assert.Equal("", nstr.TrimAll(" ", "."))
	assert.Equal("", nstr.TrimAll("   ", "."))
}

func TestSplit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"a", "log"}, nstr.Split("a.log", "."))
	assert.Equal([]string{"a", "log"}, nstr.Split(" a . log ", "."))
	assert.Equal([]string{"a", "b", "log"}, nstr.Split("a.b.log", "."))
	assert.Equal([]string{"a", "b", "log"}, nstr.Split(" a . b . log", "."))
	assert.Equal([]string{}, nstr.Split("", "."))
	assert.Equal([]string{}, nstr.Split(" ", "."))
	assert.Equal([]string{}, nstr.Split("   ", "."))
}

func TestIsNumeric(t *testing.T) {
	assert := assert.New(t)
	assert.True(nstr.IsNumeric("-123456"))
	assert.True(nstr.IsNumeric("123456"))
	assert.True(nstr.IsNumeric("-123.456"))
	assert.True(nstr.IsNumeric("123.456"))
	assert.False(nstr.IsNumeric("-123.456.789"))
	assert.False(nstr.IsNumeric("123.456.789"))
	assert.False(nstr.IsNumeric("-123-456"))
	assert.False(nstr.IsNumeric("123-456"))
}
