/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 17:56:55
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 19:54:38
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/time_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv_test

import (
	"encoding/json"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToTimeE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToTimeE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Time{}, actualObj)
	}
	actualObj, err = nconv.ToTimeE("2023-05-05 18:00:00") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2023, time.May, 5, 18, 0, 0, 0, time.UTC), actualObj)
	}
	actualObj, err = nconv.ToTimeE(json.Number("1594477475")) // json.Number
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2020, time.July, 11, 22, 24, 35, 0, time.Local), actualObj)
	}
	actualObj, err = nconv.ToTimeE(1594477475) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2020, time.July, 11, 22, 24, 35, 0, time.Local), actualObj)
	}
}

func TestToTimeInDefaultLocationE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToTimeInDefaultLocationE(nil, time.UTC) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Time{}, actualObj)
	}
	actualObj, err = nconv.ToTimeInDefaultLocationE("2023-05-05 18:00:00", time.UTC) // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2023, time.May, 5, 18, 0, 0, 0, time.UTC), actualObj)
	}
	actualObj, err = nconv.ToTimeInDefaultLocationE(json.Number("1594477475"), time.UTC) // json.Number
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2020, time.July, 11, 22, 24, 35, 0, time.Local), actualObj)
	}
	actualObj, err = nconv.ToTimeInDefaultLocationE(1594477475, time.UTC) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Date(2020, time.July, 11, 22, 24, 35, 0, time.Local), actualObj)
	}
}

func TestToDurationE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToDurationE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Duration(0), actualObj)
	}
	actualObj, err = nconv.ToDurationE("5m30s") // string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Duration(time.Minute*5+time.Second*30), actualObj)
	}
	actualObj, err = nconv.ToDurationE(json.Number("1594477475")) // json.Number
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Duration(1594477475), actualObj)
	}
	actualObj, err = nconv.ToDurationE(1594477475) // int
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal(time.Duration(1594477475), actualObj)
	}
}

func TestToDurationSliceE(t *testing.T) {
	assert := assert.New(t)
	actualObj, err := nconv.ToDurationSliceE(nil) // nil
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]time.Duration{}, actualObj)
	}
	actualObj, err = nconv.ToDurationSliceE([]string{"5m30s", "6m30s"}) // []string
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]time.Duration{time.Duration(time.Minute*5 + time.Second*30), time.Duration(time.Minute*6 + time.Second*30)}, actualObj)
	}
	actualObj, err = nconv.ToDurationSliceE([]any{json.Number("1594477475"), 1594477475}) // []any
	errLog(t, err)
	if assert.NoError(err) {
		assert.Equal([]time.Duration{time.Duration(1594477475), time.Duration(1594477475)}, actualObj)
	}
}
