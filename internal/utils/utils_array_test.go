/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 17:43:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 17:46:15
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/utils/utils_array_test.go
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

func TestIsArray(t *testing.T) {
	assert := assert.New(t)
	assert.False(utils.IsArray(map[int]int{}))
	assert.True(utils.IsArray([]any{}))
}
