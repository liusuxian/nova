/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-06 13:08:08
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:23:23
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nenv

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// All 所有环境变量的字符串副本，形式为`key=value`
func All() (envs []string) {
	return os.Environ()
}

// Map 所有环境变量的字符串副本`Map`
func Map() (m map[string]string) {
	return MapFromEnv(os.Environ())
}

// Get 获取环境变量
func Get(key string, def ...string) (value string) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		return ""
	}
	return v
}

// Set 设置环境变量
func Set(key, value string) (err error) {
	err = os.Setenv(key, value)
	if err != nil {
		err = errors.Wrapf(err, `set environment key-value failed with key "%s", value "%s"`, key, value)
	}
	return
}

// SetMap 使用`Map`设置环境变量
func SetMap(m map[string]string) (err error) {
	for k, v := range m {
		if err = Set(k, v); err != nil {
			return
		}
	}
	return
}

// Contains 检查是否存在名为`key`的环境变量
func Contains(key string) (isExist bool) {
	_, ok := os.LookupEnv(key)
	return ok
}

// Remove 删除一个或多个环境变量
func Remove(key ...string) (err error) {
	for _, v := range key {
		if err = os.Unsetenv(v); err != nil {
			err = errors.Wrapf(err, `delete environment key failed with key "%s"`, v)
			return
		}
	}
	return
}

// MapFromEnv 将环境变量从`Slice`转换为`Map`
func MapFromEnv(envs []string) (m map[string]string) {
	m = make(map[string]string, len(envs))
	i := 0
	for _, env := range envs {
		i = strings.IndexByte(env, '=')
		m[env[0:i]] = env[i+1:]
	}
	return
}

// MapToEnv 将环境变量从`Map`转换为`Slice`
func MapToEnv(m map[string]string) (envs []string) {
	envs = make([]string, 0, len(m))
	for k, v := range m {
		envs = append(envs, fmt.Sprintf(`%s=%s`, k, v))
	}
	return
}

// Filter 筛选给定环境变量中的重复项
func Filter(oldEnvs []string) (newEnvs []string) {
	return MapToEnv(MapFromEnv(oldEnvs))
}
