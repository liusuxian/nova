/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 17:09:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-20 17:17:01
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/file_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils_test

import (
	"github.com/liusuxian/nova/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils(t *testing.T) {
	assert := assert.New(t)
	assert.True(utils.PathExists("."))
	assert.False(utils.PathExists("config/config.yaml"))
	assert.Equal("yaml", utils.ExtName("config/config.yaml"))
	assert.NotEqual("test", utils.ExtName("config/test"))
	assert.Equal("", utils.ExtName("config/test"))
}
