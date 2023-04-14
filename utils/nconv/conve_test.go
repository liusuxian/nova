/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-14 13:31:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-15 00:01:41
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
	actualObj, err = nconv.ToBoolE(float64(1.23)) // float64
	errLog(t, err)
	if assert.Error(err) {
		assert.False(actualObj)
	}
}

func TestToFloat64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat64E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.56), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float64(0), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("-1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float64(-1), actualObj)
	}
	actualObj, err = nconv.ToFloat64E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float64(0), actualObj)
	}
}

func TestToFloat32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToFloat32E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.56), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float32(0), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1.23), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("-1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(float32(-1), actualObj)
	}
	actualObj, err = nconv.ToFloat32E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(float32(0), actualObj)
	}
}

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

func TestToUint64E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint64E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(0), actualObj)
	}
	actualObj, err = nconv.ToUint64E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(0), actualObj)
	}
	actualObj, err = nconv.ToUint64E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint64(0), actualObj)
	}
	actualObj, err = nconv.ToUint64E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint64(1), actualObj)
	}
	actualObj, err = nconv.ToUint64E("-1") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint64(0), actualObj)
	}
	actualObj, err = nconv.ToUint64E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint64(0), actualObj)
	}
}

func TestToUint32E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint32E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(0), actualObj)
	}
	actualObj, err = nconv.ToUint32E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(0), actualObj)
	}
	actualObj, err = nconv.ToUint32E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint32(0), actualObj)
	}
	actualObj, err = nconv.ToUint32E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint32(1), actualObj)
	}
	actualObj, err = nconv.ToUint32E("-1") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint32(0), actualObj)
	}
	actualObj, err = nconv.ToUint32E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint32(0), actualObj)
	}
}

func TestToUint16E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint16E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(0), actualObj)
	}
	actualObj, err = nconv.ToUint16E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(0), actualObj)
	}
	actualObj, err = nconv.ToUint16E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint16(0), actualObj)
	}
	actualObj, err = nconv.ToUint16E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint16(1), actualObj)
	}
	actualObj, err = nconv.ToUint16E("-1") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint16(0), actualObj)
	}
	actualObj, err = nconv.ToUint16E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint16(0), actualObj)
	}
}

func TestToUint8E(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUint8E(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(0), actualObj)
	}
	actualObj, err = nconv.ToUint8E(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(0), actualObj)
	}
	actualObj, err = nconv.ToUint8E([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint8(0), actualObj)
	}
	actualObj, err = nconv.ToUint8E("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint8(1), actualObj)
	}
	actualObj, err = nconv.ToUint8E("-1") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint8(0), actualObj)
	}
	actualObj, err = nconv.ToUint8E("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint8(0), actualObj)
	}
}

func TestToUintE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToUintE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(0), actualObj)
	}
	actualObj, err = nconv.ToUintE(int(1)) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE(float64(1.56)) // float64
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE(true) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE(false) // bool
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(0), actualObj)
	}
	actualObj, err = nconv.ToUintE([]byte("1.23")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE([]byte("1.0")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE([]byte("1.")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE([]byte("1")) // []byte
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE([]byte("a")) // []byte
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint(0), actualObj)
	}
	actualObj, err = nconv.ToUintE("1.23") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE("1.0") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE("1.") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE("1") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(uint(1), actualObj)
	}
	actualObj, err = nconv.ToUintE("-1") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint(0), actualObj)
	}
	actualObj, err = nconv.ToUintE("b") // string
	errLog(t, err)
	if assert.Error(err) {
		assert.Equal(uint(0), actualObj)
	}
}
