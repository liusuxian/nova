/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-06 13:50:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:23:13
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nenv_test

import (
	"github.com/liusuxian/nova/utils/nenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual([]string{}, nenv.All())
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(map[string]string{}, nenv.Map())
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", nenv.Get("a123"))
	assert.Equal("321", nenv.Get("a123", "321"))
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	err := nenv.Set("a123", "321")
	if assert.NoError(err) {
		assert.Equal("321", nenv.Get("a123"))
		nenv.Remove("a123")
	}
}

func TestSetMap(t *testing.T) {
	assert := assert.New(t)
	err := nenv.SetMap(map[string]string{
		"a123": "321",
		"b123": "321",
		"c123": "321",
	})
	if assert.NoError(err) {
		assert.Equal("321", nenv.Get("a123"))
		assert.Equal("321", nenv.Get("b123"))
		assert.Equal("321", nenv.Get("c123"))
		nenv.Remove("a123", "b123", "c123")
	}
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	assert.False(nenv.Contains("a123"))
	assert.False(nenv.Contains("b123"))
	assert.False(nenv.Contains("c123"))
	assert.True(nenv.Contains("GOROOT"))
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	err := nenv.Remove("a123", "b123", "c123")
	if assert.NoError(err) {
		assert.False(nenv.Contains("a123"))
		assert.False(nenv.Contains("b123"))
		assert.False(nenv.Contains("c123"))
	}
}

func TestMapFromEnv(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(map[string]string{"a123": "321", "b123": "321"}, nenv.MapFromEnv([]string{"a123=321", "b123=321"}))
}

func TestMapToEnv(t *testing.T) {
	assert := assert.New(t)
	assert.ElementsMatch([]string{"a123=321", "b123=321"}, nenv.MapToEnv(map[string]string{"a123": "321", "b123": "321"}))
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	assert.ElementsMatch([]string{"a123=321", "b123=321"}, nenv.Filter([]string{"a123=321", "b123=321", "a123=321"}))
}
