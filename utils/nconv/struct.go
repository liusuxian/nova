/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:23:47
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 03:29:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/struct.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "github.com/mitchellh/mapstructure"

// DecoderConfigOption 解码配置选项
type DecoderConfigOption func(*mapstructure.DecoderConfig)

// ToStructE 将 any 转换为 struct 类型
func ToStructE(input, output any, opts ...DecoderConfigOption) (err error) {
	return decode(input, defaultDecoderConfig(output, opts...))
}

// ToStructExactE 将 any 转换为 struct 类型。如果目标结构体中不存在某个字段，则会报错
func ToStructExactE(input, output any, opts ...DecoderConfigOption) (err error) {
	config := defaultDecoderConfig(output, opts...)
	config.ErrorUnused = true
	return decode(input, config)
}

// decode 解码
func decode(input any, config *mapstructure.DecoderConfig) (err error) {
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// defaultDecoderConfig 默认的解码配置
func defaultDecoderConfig(output any, opts ...DecoderConfigOption) (config *mapstructure.DecoderConfig) {
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
