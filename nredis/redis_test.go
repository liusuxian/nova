/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-04 12:14:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-10-10 12:00:22
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nredis_test

import (
	"context"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nredis"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

type A struct {
	A int
	B float64
	C string
	D []any
}

func TestRedis(t *testing.T) {
	client := nredis.NewClient(func(cc *nredis.ClientConfig) {
		cc.Addr = "127.0.0.1:6379"
		cc.Password = ""
		cc.DB = 2
	})
	defer client.Close()

	ctx := context.Background()
	assert := assert.New(t)
	actualObj, err := client.Do(ctx, "FLUSHDB")
	if assert.NoError(err) {
		assert.Equal("OK", actualObj)
	}
	actualObj, err = client.Do(ctx, "SET", "aaa", 1)
	if assert.NoError(err) {
		assert.Equal("OK", actualObj)
	}
	actualObj, err = client.Do(ctx, "GET", "aaa")
	if assert.NoError(err) {
		assert.Equal(1, nconv.ToInt(actualObj))
	}
	var actualPipelineObj []*niface.PipelineResult
	actualPipelineObj, err = client.Pipeline(ctx, []any{"SET", "bbb", 2}, []any{"SADD", "ccc", 3})
	if assert.NoError(err) {
		for k, v := range actualPipelineObj {
			assert.IsType(&niface.PipelineResult{}, v)
			assert.NoError(v.Err)
			if k == 0 {
				assert.Equal("OK", v.Val)
			} else {
				assert.Equal(int64(1), v.Val)
			}
		}
	}
	actualPipelineObj, err = client.Pipeline(ctx, []any{"GET", "bbb"}, []any{"SMEMBERS", "ccc"})
	if assert.NoError(err) {
		for k, v := range actualPipelineObj {
			assert.IsType(&niface.PipelineResult{}, v)
			assert.NoError(v.Err)
			if k == 0 {
				assert.Equal("2", v.Val)
			} else {
				assert.Equal([]any{"3"}, v.Val)
			}
		}
	}
	actualObj, err = client.Do(ctx, "SET", "ddd", &A{A: 1, B: 1.2, C: "hello", D: []any{1, 1.2, "hello", true}})
	if assert.NoError(err) {
		assert.Equal("OK", actualObj)
	}
	actualObj, err = client.Do(ctx, "GET", "ddd")
	if assert.NoError(err) {
		val := &A{}
		err = nconv.ToStructE(actualObj, &val)
		if assert.NoError(err) {
			assert.IsType(&A{}, val)
			assert.Equal(&A{A: 1, B: 1.2, C: "hello", D: []any{float64(1), 1.2, "hello", true}}, val)
			assert.Equal(map[string]any{"A": float64(1), "B": 1.2, "C": "hello", "D": []any{float64(1), 1.2, "hello", true}}, nconv.ToStringMap(actualObj))
		}
	}
	actualObj, err = client.Do(ctx, "SET", "eee", []any{1, 1.2, "hello", true})
	if assert.NoError(err) {
		assert.Equal("OK", actualObj)
	}
	actualObj, err = client.Do(ctx, "GET", "eee")
	if assert.NoError(err) {
		assert.ElementsMatch([]any{float64(1), 1.2, "hello", true}, nconv.ToSlice(actualObj))
	}

	err = client.ScriptLoad(ctx, "lua_script/test1.lua")
	if assert.Error(err) {
		t.Log("lua script load failed: ", err)
	}
	err = client.ScriptLoad(ctx, "lua_script/test.lua")
	if assert.NoError(err) {
		t.Log("lua script load succ")
	}
	actualObj, err = client.EvalSha(ctx, "test", []string{"lua_key"}, 1)
	if assert.NoError(err) {
		assert.Equal(1, nconv.ToInt(actualObj))
	}
}
