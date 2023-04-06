/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-04 12:04:41
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-06 20:03:37
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nredis/redis.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nredis

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/liusuxian/nova/internal/reflection"
	"github.com/liusuxian/nova/niface"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"reflect"
	"time"
)

// RedisClient redis 客户端结构
type RedisClient struct {
	redis        *redis.Client     // redis 客户端
	luaScriptMap map[string]string // lua 脚本
}

// Options redis 客户端选项
type Options struct {
	Addr                  string        // host:port 地址
	Password              string        // 访问授权密码
	DB                    int           // 数据库索引
	DialTimeout           time.Duration // 建立新连接的超时时间，默认值是 5 秒
	ReadTimeout           time.Duration // 读取套接字的超时时间。如果超时，命令将以超时失败而不是阻塞。0:默认超时时间（3秒）1:无超时时间（会一直阻塞）2:禁用SetReadDeadline调用
	WriteTimeout          time.Duration // 写入套接字的超时时间。如果超时，命令将以超时失败而不是阻塞。0:默认超时时间（3秒）1:无超时时间（会一直阻塞）2:禁用SetWriteDeadline调用
	ContextTimeoutEnabled bool          // 控制客户端是否遵守上下文的超时和截止时间。当该选项为 true 时，客户端会尊重上下文的超时和截止时间，否则会忽略上下文超时和截止时间
	PoolFIFO              bool          // 连接池类型。true 表示 FIFO 模式，false 表示 LIFO 模式。FIFO 模式相对于 LIFO 模式会有稍微更高的开销，但它有助于更快地关闭空闲连接，从而减少池的大小
	PoolSize              int           // 最大套接字连接数。根据 runtime.GOMAXPROCS 的结果，默认是每个可用 CPU 有 10 个连接
	PoolTimeout           time.Duration // 如果所有连接都忙碌，在返回错误之前，客户端等待连接的时间。默认值为 ReadTimeout+1 秒
	MinIdleConns          int           // 最小空闲连接数
	MaxIdleConns          int           // 最大空闲连接数
	ConnMaxIdleTime       time.Duration // 连接空闲的最大时间。它应该小于服务器的超时时间。默认为 30 分钟。-1 禁用空闲超时检查
	ConnMaxLifetime       time.Duration // 连接可被重复使用的最大时间。默认不关闭空闲连接
	TLSConfig             *tls.Config   // TLS 配置。当设置时，客户端和服务器将协商使用 TLS 加密协议进行通信
	Limiter               redis.Limiter // Limiter 接口用于实现熔断器或速率限制器
}

// NewClient 创建 redis 客户端
func NewClient(opt *Options) (client niface.IRedisClient) {
	client = &RedisClient{
		redis: redis.NewClient(&redis.Options{
			Addr:                  opt.Addr,
			Password:              opt.Password,
			DB:                    opt.DB,
			DialTimeout:           opt.DialTimeout,
			ReadTimeout:           opt.ReadTimeout,
			WriteTimeout:          opt.WriteTimeout,
			ContextTimeoutEnabled: opt.ContextTimeoutEnabled,
			PoolFIFO:              opt.PoolFIFO,
			PoolSize:              opt.PoolSize,
			PoolTimeout:           opt.PoolTimeout,
			MinIdleConns:          opt.MinIdleConns,
			MaxIdleConns:          opt.MaxIdleConns,
			ConnMaxIdleTime:       opt.ConnMaxIdleTime,
			ConnMaxLifetime:       opt.ConnMaxLifetime,
			TLSConfig:             opt.TLSConfig,
			Limiter:               opt.Limiter,
		}),
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
		err = errors.New("Pipeline CmdArgsList Is Empty")
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

// Close 关闭 redis
func (rc *RedisClient) Close() (err error) {
	return rc.redis.Close()
}
