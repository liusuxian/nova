/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:28:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 12:33:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nutils/nctx/nctx_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nctx_test

import (
	"context"
	"github.com/liusuxian/nova/nutils/nctx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtx(t *testing.T) {
	assert := assert.New(t)
	ctx := nctx.SetCtxGlobalVal(context.Background(), 1)
	assert.NotNil(ctx)
	assert.Equal(1, nctx.GetCtxGlobalVal(ctx))
	ctx = nctx.SetCtxGlobalVal(ctx, 2)
	assert.Equal(2, nctx.GetCtxGlobalVal(ctx))
}
