/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-27 22:38:10
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 11:01:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nslice/slice_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nslice_test

import (
	"github.com/liusuxian/nova/utils/nslice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsContains(t *testing.T) {
	assert := assert.New(t)
	assert.False(nslice.IsContains([]int{}, 0))
	assert.True(nslice.IsContains([]int{0, 1, 2}, 0))
	assert.True(nslice.IsContains([]int{0, 1, 2}, 1))
	assert.True(nslice.IsContains([]int{0, 1, 2}, 2))
	assert.False(nslice.IsContains([]int{0, 1, 2}, 3))
	assert.True(nslice.IsContains([]int{10, 9, 9, 1, 6, 6, 5, 5, 4, 4, 4, 3, 3, 2}, 1))
	assert.True(nslice.IsContains([]int{10, 9, 9, 1, 6, 6, 5, 5, 4, 4, 4, 3, 3, 2}, 2))
	assert.True(nslice.IsContains([]int{10, 9, 9, 1, 6, 6, 5, 5, 4, 4, 4, 3, 3, 2}, 3))
	assert.False(nslice.IsContains([]int{10, 9, 9, 1, 6, 6, 5, 5, 4, 4, 4, 3, 3, 2}, 0))

	assert.False(nslice.IsContains([]float32{}, 1.0))
	assert.True(nslice.IsContains([]float32{1.0, 1.1, 1.2}, 1.0))
	assert.True(nslice.IsContains([]float32{1.0, 1.1, 1.2}, 1.1))
	assert.True(nslice.IsContains([]float32{1.0, 1.1, 1.2}, 1.2))
	assert.False(nslice.IsContains([]float32{1.0, 1.1, 1.2}, 1.3))
	assert.True(nslice.IsContains([]float32{10.1, 9.1, 9.1, 1.1, 6.1, 6.1, 5.1, 5.1, 4.1, 4.1, 4.1, 3.1, 3.1, 2.1}, 1.1))
	assert.True(nslice.IsContains([]float32{10.1, 9.1, 9.1, 1.1, 6.1, 6.1, 5.1, 5.1, 4.1, 4.1, 4.1, 3.1, 3.1, 2.1}, 2.1))
	assert.True(nslice.IsContains([]float32{10.1, 9.1, 9.1, 1.1, 6.1, 6.1, 5.1, 5.1, 4.1, 4.1, 4.1, 3.1, 3.1, 2.1}, 3.1))
	assert.False(nslice.IsContains([]float32{10.1, 9.1, 9.1, 1.1, 6.1, 6.1, 5.1, 5.1, 4.1, 4.1, 4.1, 3.1, 3.1, 2.1}, 1.0))

	assert.False(nslice.IsContains([]string{"hello", "jack", "hello", "world", "tom", "hay", "tom"}, "lsx"))
	assert.False(nslice.IsContains([]string{"hello", "jack", "hello", "world", "tom", "hay", "tom"}, "hell"))
	assert.True(nslice.IsContains([]string{"hello", "jack", "hello", "world", "tom", "hay", "tom"}, "hay"))

	assert.False(nslice.IsContains([]string{"我", "是", "中", "国", "人"}, "hello"))
	assert.False(nslice.IsContains([]string{"我", "是", "中", "国", "人"}, "哈"))
	assert.True(nslice.IsContains([]string{"我", "是", "中", "国", "人"}, "中"))

	assert.False(nslice.IsContains([]rune("我是中国人"), []rune("hello")[0]))
	assert.False(nslice.IsContains([]rune("我是中国人"), []rune("哈罗")[0]))
	assert.True(nslice.IsContains([]rune("我是中国人"), []rune("中国")[0]))
	assert.True(nslice.IsContains([]rune("hello我是中国人"), []rune("hello")[0]))

	assert.False(nslice.IsContains([]byte{4, 3, 1, 5, 6, 9, 8, 2, 1, 7}, 0))
	assert.True(nslice.IsContains([]byte{4, 3, 1, 5, 6, 9, 8, 2, 1, 7}, 2))
}
