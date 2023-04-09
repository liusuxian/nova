/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-06 21:19:37
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 00:56:17
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_any_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSlice(t *testing.T) {
	assert := assert.New(t)
	s1 := nconv.ToSlice("[1, 1.2, true, \"hello\"]")
	assert.Equal(1, nconv.ToInt(s1[0]))
	assert.Equal(1.2, nconv.ToFloat64(s1[1]))
	assert.Equal(true, s1[2])
	assert.Equal("hello", s1[3])
	assert.Equal([]any{"h", "e", "l", "l", "o"}, nconv.ToSlice("[\"h\", \"e\", \"l\", \"l\", \"o\"]"))
	assert.Equal([]any{1, 2, 3}, nconv.ToSlice([]any{1, 2, 3}))
	assert.Equal([]any{1, 2, 3}, nconv.ToSlice([]int{1, 2, 3}))
	assert.Equal([]any{1}, nconv.ToSlice(1))
	assert.Equal([]any{1.2}, nconv.ToSlice(1.2))
	assert.Equal([]any{"2"}, nconv.ToSlice("2"))
}
