/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 14:24:41
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:25:13
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
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
	assert.ElementsMatch([]string{"a", "log"}, nstr.Split("a.log", "."))
	assert.ElementsMatch([]string{"a", "log"}, nstr.Split(" a . log ", "."))
	assert.ElementsMatch([]string{"a", "b", "log"}, nstr.Split("a.b.log", "."))
	assert.ElementsMatch([]string{"a", "b", "log"}, nstr.Split(" a . b . log", "."))
	assert.ElementsMatch([]string{}, nstr.Split("", "."))
	assert.ElementsMatch([]string{}, nstr.Split(" ", "."))
	assert.ElementsMatch([]string{}, nstr.Split("   ", "."))
}
