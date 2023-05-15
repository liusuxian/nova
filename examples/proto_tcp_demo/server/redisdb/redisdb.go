/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-09 10:57:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:07:57
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package redisdb

import (
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nredis"
)

type Redis struct {
	niface.IRedisClient
}

var redis *Redis

// Start 启动 redis
func Start() {
	client := nredis.NewClient(func(cc *nredis.ClientConfig) {
		cc.Addr = nconf.GetString("redis.addr")
		cc.Password = nconf.GetString("redis.password")
		cc.DB = nconf.GetInt("redis.db")
	})
	redis = &Redis{client}
	nlog.Info("Redis Start")
}

// Instance redis 实例
func Instance() (r *Redis) {
	return redis
}

// Close 关闭 redis
func Close() (err error) {
	nlog.Info("Redis Close")
	return redis.Close()
}
