/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 22:15:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 14:10:59
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/nlog_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog_test

import (
	"context"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nrequest"
	"testing"
)

// Context 上下文结构
type Context struct {
	User ContextUser    // 上下文用户信息
	Data map[string]any // 自定义KV变量，业务模块根据需要设置，不固定
}

// ContextUser 上下文中的用户信息
type ContextUser struct {
	Id     int64  // 用户ID
	Appid  string // 小程序ID
	Openid string // openid
}

func TestLog(t *testing.T) {
	req := nrequest.NewRequest(nil, nil)
	req.SetCtxVal("aaa", 1)
	req.SetCtxVal("bbb", 2)
	req.SetCtxVal("ccc", 3)
	req.SetCtxVal("ccc", 4)
	req.SetCtxVal("ddd", Context{
		User: ContextUser{
			Id:     1,
			Appid:  "111",
			Openid: "222",
		},
		Data: map[string]any{"reqId": "333"},
	})
	nlog.Debug(req.GetCtx(), "I am Debug")
	nlog.Info(req.GetCtx(), "I am Info")
	nlog.Warn(req.GetCtx(), "I am Warn")
	nlog.Error(req.GetCtx(), "I am Error")
	ctx := context.Background()
	nlog.DPanic(ctx, "I am DPanic")
	nlog.Panic(ctx, "I am Panic")
	nlog.Fatal(ctx, "I am Fatal")
}
