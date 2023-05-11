/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:38:55
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:40:02
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv_test

import (
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt64E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E(float64(1.56)) // float64
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
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int64(1), actualObj)
	}
	actualObj, err = nconv.ToInt64E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int64(0), actualObj)
	}
}

func TestToInt32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt32E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E(float64(1.56)) // float64
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
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int32(1), actualObj)
	}
	actualObj, err = nconv.ToInt32E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int32(0), actualObj)
	}
}

func TestToInt16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt16E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E(float64(1.56)) // float64
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
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int16(1), actualObj)
	}
	actualObj, err = nconv.ToInt16E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int16(0), actualObj)
	}
}

func TestToInt8E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToInt8E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(0), actualObj)
	}
	actualObj, err = nconv.ToInt8E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(0), actualObj)
	}
	actualObj, err = nconv.ToInt8E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int8(0), actualObj)
	}
	actualObj, err = nconv.ToInt8E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int8(1), actualObj)
	}
	actualObj, err = nconv.ToInt8E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int8(0), actualObj)
	}
}

func TestToIntE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToIntE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(0), actualObj)
	}
	actualObj, err = nconv.ToIntE(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(0), actualObj)
	}
	actualObj, err = nconv.ToIntE([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int(0), actualObj)
	}
	actualObj, err = nconv.ToIntE("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(int(1), actualObj)
	}
	actualObj, err = nconv.ToIntE("b") // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(int(0), actualObj)
	}
}
