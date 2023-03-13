/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:35:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 21:38:50
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
	assert.Equal(2, len(nstr.Split("a.log", ".")))
	assert.Equal(2, len(nstr.Split(" a . log ", ".")))
	assert.Equal(3, len(nstr.Split("a.b.log", ".")))
	assert.Equal(3, len(nstr.Split(" a . b . log", ".")))
	assert.Equal(0, len(nstr.Split("", ".")))
	assert.Equal(0, len(nstr.Split(" ", ".")))
	assert.Equal(0, len(nstr.Split("   ", ".")))
}
