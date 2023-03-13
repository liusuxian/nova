/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 17:09:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 13:53:26
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nfile/nfile_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nfile_test

import (
	"github.com/liusuxian/nova/utils/nfile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile(t *testing.T) {
	assert := assert.New(t)
	assert.True(nfile.PathExists("."))
	assert.False(nfile.PathExists("config/config.yaml"))
	assert.Equal("yaml", nfile.ExtName("config/config.yaml"))
	assert.NotEqual("test", nfile.ExtName("config/test"))
	assert.Equal("", nfile.ExtName("config/test"))
}
