/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-10 18:51:03
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 12:48:20
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import "github.com/mitchellh/mapstructure"

// DecoderConfig 解码配置
type DecoderConfig = mapstructure.DecoderConfig

// DecoderConfigOption 解码配置选项
type DecoderConfigOption func(dc *DecoderConfig)

// decode 解码
func decode(input any, config *DecoderConfig) (err error) {
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// defaultDecoderConfig 默认的解码配置
func defaultDecoderConfig(output any, opts ...DecoderConfigOption) (config *mapstructure.DecoderConfig) {
	c := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.RecursiveStructToMapHookFunc(),
			mapstructure.StringToIPHookFunc(),
			mapstructure.StringToIPNetHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
			mapstructure.TextUnmarshallerHookFunc(),
		),
		WeaklyTypedInput: true,
		Metadata:         nil,
		Result:           output,
		TagName:          "json",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
