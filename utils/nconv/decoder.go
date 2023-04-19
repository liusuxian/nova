/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-18 13:07:45
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-19 11:50:27
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/decoder.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "github.com/mitchellh/mapstructure"

// DecoderConfigOption 解码配置选项
type DecoderConfigOption func(*mapstructure.DecoderConfig)

// WithDecoderConfigOption
func WithDecoderConfigOption(config *mapstructure.DecoderConfig) DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		if config.DecodeHook != nil {
			c.DecodeHook = config.DecodeHook
		}
		if c.ErrorUnused != config.ErrorUnused {
			c.ErrorUnused = config.ErrorUnused
		}
		if c.ErrorUnset != config.ErrorUnset {
			c.ErrorUnset = config.ErrorUnset
		}
		if c.ZeroFields != config.ZeroFields {
			c.ZeroFields = config.ZeroFields
		}
		if c.WeaklyTypedInput != config.WeaklyTypedInput {
			c.WeaklyTypedInput = config.WeaklyTypedInput
		}
		if c.Squash != config.Squash {
			c.Squash = config.Squash
		}
		if config.Metadata != nil {
			c.Metadata = config.Metadata
		}
		if config.TagName != "" && c.TagName != config.TagName {
			c.TagName = config.TagName
		}
		if c.IgnoreUntaggedFields != config.IgnoreUntaggedFields {
			c.IgnoreUntaggedFields = config.IgnoreUntaggedFields
		}
		if config.MatchName != nil {
			c.MatchName = config.MatchName
		}
	}
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
