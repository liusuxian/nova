/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-18 13:07:45
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-18 13:21:19
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/decoder.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "github.com/mitchellh/mapstructure"

// DecoderConfigOption 解码配置选项
type DecoderConfigOption func(*mapstructure.DecoderConfig)

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
