/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-06 19:05:04
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:16:29
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package niface

import "context"

// IRedisClient redis 客户端接口
type IRedisClient interface {
	Do(ctx context.Context, cmd string, args ...any) (value any, err error)                                // 执行 redis 命令
	Pipeline(ctx context.Context, cmdArgsList ...[]any) (results []*PipelineResult, err error)             // 执行 redis 管道命令
	ScriptLoad(ctx context.Context, scriptFilePath string) (err error)                                     // 加载 lua 脚本
	EvalSha(ctx context.Context, scriptFileName string, keys []string, args ...any) (value any, err error) // 执行 lua 脚本
	Close() (err error)                                                                                    // 关闭 redis
}

// PipelineResult 管道返回值
type PipelineResult struct {
	Val any
	Err error
}
