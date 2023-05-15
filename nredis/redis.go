/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 02:58:43
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 13:35:35
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nredis

import (
	"context"
	"encoding/json"
	"github.com/liusuxian/nova/internal/reflection"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/utils/nfile"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"reflect"
	"strings"
)

// ClientConfig redis 客户端配置
type ClientConfig = redis.Options

// ClientConfigOption redis 客户端配置选项
type ClientConfigOption func(cc *ClientConfig)

// RedisClient redis 客户端结构
type RedisClient struct {
	redis        *redis.Client     // redis 客户端
	luaScriptMap map[string]string // lua 脚本
}

// NewClient 创建 redis 客户端
func NewClient(opts ...ClientConfigOption) (client niface.IRedisClient) {
	ro := &redis.Options{}
	for _, opt := range opts {
		opt(ro)
	}
	client = &RedisClient{
		redis:        redis.NewClient(ro),
		luaScriptMap: make(map[string]string),
	}
	return
}

// Do 执行 redis 命令
func (rc *RedisClient) Do(ctx context.Context, cmd string, args ...any) (value any, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// 处理 redis 命令参数
	for k, v := range args {
		reflectInfo := reflection.OriginTypeAndKind(v)
		switch reflectInfo.OriginKind {
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			// 忽略切片类型为 []byte 的情况
			if _, ok := v.([]byte); !ok {
				if args[k], err = json.Marshal(v); err != nil {
					return
				}
			}
		}
	}
	// 执行 redis 命令
	cmdArgs := make([]any, 0, len(args)+1)
	cmdArgs = append(cmdArgs, cmd)
	cmdArgs = append(cmdArgs, args...)
	value, err = rc.redis.Do(ctx, cmdArgs...).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

// Pipeline 执行 redis 管道命令
func (rc *RedisClient) Pipeline(ctx context.Context, cmdArgsList ...[]any) (results []*niface.PipelineResult, err error) {
	if len(cmdArgsList) == 0 {
		err = errors.New("pipeline cmd args list is empty")
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}
	// 执行 redis 管道命令
	p := rc.redis.Pipeline()
	// 处理redis命令参数
	for _, cmdArgs := range cmdArgsList {
		for k, v := range cmdArgs {
			if k > 0 {
				reflectInfo := reflection.OriginTypeAndKind(v)
				switch reflectInfo.OriginKind {
				case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
					// 忽略切片类型为 []byte 的情况
					if _, ok := v.([]byte); !ok {
						if cmdArgs[k], err = json.Marshal(v); err != nil {
							return
						}
					}
				}
			}
		}
		// 执行 redis 命令
		p.Do(ctx, cmdArgs...)
	}
	var resList []redis.Cmder
	resList, err = p.Exec(ctx)
	if err == redis.Nil {
		err = nil
	}
	if err != nil {
		return
	}
	// 处理返回结果
	results = make([]*niface.PipelineResult, 0, len(resList))
	for _, v := range resList {
		results = append(results, &niface.PipelineResult{
			Val: v.(*redis.Cmd).Val(),
			Err: v.Err(),
		})
	}
	return
}

// ScriptLoad 加载 lua 脚本
func (rc *RedisClient) ScriptLoad(ctx context.Context, scriptFilePath string) (err error) {
	script := nfile.GetContents(scriptFilePath)
	if strings.EqualFold("", script) {
		return errors.Errorf("[%s] script not found", scriptFilePath)
	}
	evalsha, err := rc.redis.ScriptLoad(ctx, script).Result()
	if err != nil {
		return err
	}
	scriptFileName := nfile.Name(scriptFilePath)
	rc.luaScriptMap[scriptFileName] = evalsha
	return nil
}

// EvalSha 执行 lua 脚本
func (rc *RedisClient) EvalSha(ctx context.Context, scriptFileName string, keys []string, args ...any) (value any, err error) {
	evalsha, ok := rc.luaScriptMap[scriptFileName]
	if !ok {
		return nil, errors.Errorf("[%s] Script Not Found", scriptFileName)
	}
	value, err = rc.redis.EvalSha(ctx, evalsha, keys, args...).Result()
	if err == redis.Nil {
		err = nil
	}
	return
}

// Close 关闭 redis
func (rc *RedisClient) Close() (err error) {
	return rc.redis.Close()
}
