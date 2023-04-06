/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-04 11:38:10
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-06 19:29:08
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/niface/redis.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package niface

import "context"

// IRedisClient redis 客户端接口
type IRedisClient interface {
	Do(ctx context.Context, cmd *RedisCmd) (result *RedisResult)                        // 执行 redis 命令
	Pipeline(ctx context.Context, cmds []*RedisCmd) (results []*RedisResult, err error) // 执行 redis 管道命令
	Close() (err error)                                                                 // 关闭 redis
}

// RedisCmd redis 命令参数
type RedisCmd struct {
	Cmd  string
	Args []any
}

// RedisResult redis 执行结果
type RedisResult struct {
	Val any
	Err error
}
