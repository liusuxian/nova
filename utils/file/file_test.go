/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 17:09:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 18:12:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/file/file_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package file_test

import (
	"github.com/liusuxian/nova/utils/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils(t *testing.T) {
	assert := assert.New(t)
	assert.True(file.PathExists("."))
	assert.False(file.PathExists("config/config.yaml"))
	assert.Equal("yaml", file.ExtName("config/config.yaml"))
	assert.NotEqual("test", file.ExtName("config/test"))
	assert.Equal("", file.ExtName("config/test"))
}
