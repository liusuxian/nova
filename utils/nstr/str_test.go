/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:35:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-02 02:34:12
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

func TestStr(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("alog", nstr.TrimAll("a.log", "."))
	assert.Equal("alog", nstr.TrimAll(" a . log ", "."))
	assert.Equal("ablog", nstr.TrimAll("a.b.log", "."))
	assert.Equal("ablog", nstr.TrimAll(" a . b . log", "."))
	assert.Equal("", nstr.TrimAll("", "."))
	assert.Equal("", nstr.TrimAll(" ", "."))
	assert.Equal("", nstr.TrimAll("   ", "."))

	assert.ElementsMatch([]string{"a", "log"}, nstr.Split("a.log", "."))
	assert.ElementsMatch([]string{"a", "log"}, nstr.Split(" a . log ", "."))
	assert.ElementsMatch([]string{"a", "b", "log"}, nstr.Split("a.b.log", "."))
	assert.ElementsMatch([]string{"a", "b", "log"}, nstr.Split(" a . b . log", "."))
	assert.ElementsMatch([]string{}, nstr.Split("", "."))
	assert.ElementsMatch([]string{}, nstr.Split(" ", "."))
	assert.ElementsMatch([]string{}, nstr.Split("   ", "."))
}
