/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 13:31:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-14 16:55:10
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conve_test.go
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

func errLog(t *testing.T, err error) {
	if err != nil {
		t.Logf("Error: %+v\n", err.Error())
	}
}

func TestToBoolE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToBoolE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int64(0)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int64(-1)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int64(1)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int32(0)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int32(-1)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int32(1)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int16(0)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int16(-1)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int16(1)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int8(0)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int8(-1)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(int8(1)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(byte('a')) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(0) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(-1) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(1) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint64(0)) // uint64
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint64(1)) // uint64
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint32(0)) // uint32
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint32(1)) // uint32
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint16(0)) // uint16
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint16(1)) // uint16
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint8(0)) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint8(1)) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(byte('A')) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint(0)) // uint
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(uint(1)) // uint
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte("true")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte("")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE([]byte{}) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("true") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.True(actualObj)
	}
	actualObj, err = nconv.ToBoolE("false") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE(" ") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.False(actualObj)
	}
	actualObj, err = nconv.ToBoolE("hello") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.False(actualObj)
	}
}

func TestToInt64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt64E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int64(1)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int32(1)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int16(1)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int8(1)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(uint64(1)) // uint64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(uint32(1)) // uint32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(uint16(1)) // uint16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(uint8(1)) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(uint(1)) // uint
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(float64(1.23)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(float32(1.23)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(float32(1.56)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.23") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
}

func TestToInt32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt32E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int64(1)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int32(1)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int16(1)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int8(1)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(uint64(1)) // uint64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(uint32(1)) // uint32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(uint16(1)) // uint16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(uint8(1)) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(uint(1)) // uint
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(float64(1.23)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(float32(1.23)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(float32(1.56)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.23") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
}

func TestToInt16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt16E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int64(1)) // int64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int32(1)) // int32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int16(1)) // int16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int8(1)) // int8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(uint64(1)) // uint64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(uint32(1)) // uint32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(uint16(1)) // uint16
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(uint8(1)) // uint8
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(uint(1)) // uint
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(float64(1.23)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(float32(1.23)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(float32(1.56)) // float32
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.23") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
}
