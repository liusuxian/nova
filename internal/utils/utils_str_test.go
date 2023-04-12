/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:59:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-12 20:53:37
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/utils/utils_str_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils_test

import (
	"github.com/liusuxian/nova/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLetterUpper(t *testing.T) {
	assert := assert.New(t)
	assert.False(utils.IsLetterUpper(101))
	assert.False(utils.IsLetterUpper('a'))
	assert.True(utils.IsLetterUpper(65))
	assert.True(utils.IsLetterUpper('A'))
}

func TestIsLetterLower(t *testing.T) {
	assert := assert.New(t)
	assert.False(utils.IsLetterLower(65))
	assert.False(utils.IsLetterLower('A'))
	assert.True(utils.IsLetterLower(101))
	assert.True(utils.IsLetterLower('a'))
}

func TestIsLetter(t *testing.T) {
	assert := assert.New(t)
	assert.False(utils.IsLetter(200))
	assert.True(utils.IsLetter(65))
	assert.True(utils.IsLetter('A'))
	assert.True(utils.IsLetter(101))
	assert.True(utils.IsLetter('a'))
}

func TestIsNumeric(t *testing.T) {
	assert := assert.New(t)
	assert.True(utils.IsNumeric("-123456"))
	assert.True(utils.IsNumeric("123456"))
	assert.True(utils.IsNumeric("-123.456"))
	assert.True(utils.IsNumeric("123.456"))
	assert.False(utils.IsNumeric("-123.456.789"))
	assert.False(utils.IsNumeric("123.456.789"))
	assert.False(utils.IsNumeric("-123-456"))
	assert.False(utils.IsNumeric("123-456"))
}

func TestRemoveSymbols(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abc123", utils.RemoveSymbols("abc&123"))
	assert.Equal("abc123", utils.RemoveSymbols("abc.123"))
	assert.Equal("abc123", utils.RemoveSymbols("abc.123!@#$%"))
}

func TestTrim(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abc", utils.Trim(" abc"))
	assert.Equal("abc", utils.Trim("abc "))
	assert.Equal("abc", utils.Trim(" abc "))
}

func TestSplitAndTrim(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"a", "b", "c"}, utils.SplitAndTrim("a,b,c", ","))
	assert.Equal([]string{"a", "b", "c"}, utils.SplitAndTrim(",a,b,c", ","))
	assert.Equal([]string{"a", "b", "c"}, utils.SplitAndTrim("a,b,c,", ","))
}
