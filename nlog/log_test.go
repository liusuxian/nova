/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 22:15:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 16:16:54
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/log_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog_test

import (
	"context"
	"github.com/liusuxian/nova/nlog"
	"testing"
)

func TestLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ctxKey", "ctxValue")
	nlog.Debug(ctx, "I am Debug")
	nlog.Info(ctx, "I am Info")
	nlog.Warn(ctx, "I am Warn")
	nlog.Error(ctx, "I am Error")
	nlog.DPanic(ctx, "I am DPanic")
	nlog.Panic(ctx, "I am Panic")
	nlog.Fatal(ctx, "I am Fatal")
}
