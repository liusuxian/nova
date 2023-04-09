/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-09 23:03:24
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 01:21:34
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_slice_str_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToStrings(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]any{1, 2, 3}))
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]int{1, 2, 3}))
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]int8{1, 2, 3}))
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]int16{1, 2, 3}))
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]int32{1, 2, 3}))
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings([]int64{1, 2, 3}))
	f, _ := json.Marshal([]string{"1", "2", "3"})
	assert.Equal([]string{"1", "2", "3"}, nconv.ToStrings(f))
}
