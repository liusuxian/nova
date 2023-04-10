/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 18:03:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 18:41:01
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/empty/empty_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package empty_test

import (
	"github.com/liusuxian/nova/internal/empty"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestInt int

type TestString string

type TestPerson interface {
	Say() string
}

type TestWoman struct {
}

func (woman TestWoman) Say() string {
	return "nice"
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)
	tmpT1 := "0"
	var tmpT2 *int
	tmpT3 := make(chan int)
	var (
		tmpT4 TestPerson  = nil
		tmpT5 *TestPerson = nil
		tmpT6 TestPerson  = TestWoman{}
		tmpT7 TestInt     = 0
		tmpT8 TestString  = ""
	)
	tmpF1 := "1"
	tmpF2 := func(a string) string { return "1" }
	tmpF3 := make(chan int, 1)
	tmpF3 <- 1
	var (
		tmpF4 TestPerson = &TestWoman{}
		tmpF5 TestInt    = 1
		tmpF6 TestString = "1"
	)

	assert.True(empty.IsEmpty(nil))
	assert.True(empty.IsEmpty(0))
	assert.True(empty.IsEmpty(nconv.ToInt(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToInt8(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToInt16(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToInt32(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToInt64(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToUint64(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToUint(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToUint16(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToUint32(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToUint64(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToFloat32(tmpT1)))
	assert.True(empty.IsEmpty(nconv.ToFloat64(tmpT1)))
	assert.True(empty.IsEmpty(false))
	assert.True(empty.IsEmpty([]byte("")))
	assert.True(empty.IsEmpty(""))
	assert.True(empty.IsEmpty(map[string]any{}))
	assert.True(empty.IsEmpty([]any{}))
	assert.True(empty.IsEmpty(tmpT2))
	assert.True(empty.IsEmpty(tmpT3))
	assert.True(empty.IsEmpty(tmpT3))
	assert.True(empty.IsEmpty(tmpT4))
	assert.True(empty.IsEmpty(tmpT5))
	assert.True(empty.IsEmpty(tmpT6))
	assert.True(empty.IsEmpty(tmpT7))
	assert.True(empty.IsEmpty(tmpT8))

	assert.False(empty.IsEmpty(nconv.ToInt(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToInt8(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToInt16(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToInt32(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToInt64(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToUint(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToUint8(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToUint16(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToUint32(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToUint64(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToFloat32(tmpF1)))
	assert.False(empty.IsEmpty(nconv.ToFloat64(tmpF1)))
	assert.False(empty.IsEmpty(true))
	assert.False(empty.IsEmpty(tmpT1))
	assert.False(empty.IsEmpty([]byte("1")))
	assert.False(empty.IsEmpty(map[string]any{"a": 1}))
	assert.False(empty.IsEmpty([]any{"1"}))
	assert.False(empty.IsEmpty(tmpF2))
	assert.False(empty.IsEmpty(tmpF3))
	assert.False(empty.IsEmpty(tmpF4))
	assert.False(empty.IsEmpty(tmpF5))
	assert.False(empty.IsEmpty(tmpF6))
}

func TestIsNil(t *testing.T) {
	assert := assert.New(t)
	assert.True(empty.IsNil(nil))
	var i int
	assert.False(empty.IsNil(i))
	var pi1 *int
	assert.True(empty.IsNil(pi1))
	var pi2 *int
	assert.False(empty.IsNil(&pi2))
	assert.True(empty.IsNil(&pi2, true))
}
