/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 00:16:21
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:23:38
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nfile_test

import (
	"github.com/liusuxian/nova/utils/nfile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathExists(t *testing.T) {
	assert := assert.New(t)
	assert.True(nfile.PathExists("."))
	assert.False(nfile.PathExists("config/config.yaml"))
}

func TestExtName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("yaml", nfile.ExtName("config/config.yaml"))
	assert.NotEqual("test", nfile.ExtName("config/test"))
	assert.Equal("", nfile.ExtName("config/test"))
}

func TestGetContents(t *testing.T) {
	assert := assert.New(t)
	assert.NotEmpty(nfile.GetContents("file_test.go"))
}

func TestGetBytes(t *testing.T) {
	assert := assert.New(t)
	assert.NotEmpty(nfile.GetBytes("file_test.go"))
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("file", nfile.Name("/var/www/file.js"))
	assert.Equal("file", nfile.Name("file.js"))
}
