/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 18:14:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 19:08:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/ctxglobal/ctxglobal.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package ctxglobal

import "context"

const ContextKey = "CtxGlobalKey" // 上下文变量存储键名

// GetCtxGlobalVal 获取Context中的全局Value
func GetCtxGlobalVal(ctx context.Context) any {
	return ctx.Value(ContextKey)
}

// SetCtxGlobalVal 设置Context中的全局Value
func SetCtxGlobalVal(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, ContextKey, value)
}
