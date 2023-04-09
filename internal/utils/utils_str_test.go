/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:59:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 00:11:05
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
