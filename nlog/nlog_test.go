package nlog_test

import (
	"context"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/utils/ctxglobal"
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
	ctx := ctxglobal.SetCtxGlobalVal(context.Background(), Context{
		User: ContextUser{
			Id:     1,
			Appid:  "111",
			Openid: "222",
		},
		Data: map[string]any{"traceId": "333", "reqId": "444"},
	})
	nlog.Debug(ctx, "I am Debug")
	nlog.Info(ctx, "I am Info")
	nlog.Warn(ctx, "I am Warn")
	nlog.Error(ctx, "I am Error")
	nlog.DPanic(ctx, "I am DPanic")
	nlog.Panic(ctx, "I am Panic")
	nlog.Fatal(ctx, "I am Fatal")
}
