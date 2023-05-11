/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:39:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:51:05
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
