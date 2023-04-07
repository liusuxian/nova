/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-07 12:25:23
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-07 14:06:25
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/internal/json/json.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package json

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

// RawMessage 是一个原始的 JSON 编码值。
// 它实现了 Marshaler 和 Unmarshaler，可用于延迟 JSON 解码或预计算 JSON 编码
type RawMessage = json.RawMessage

// Marshal 将 Go 对象编码为 JSON 字节数组
func Marshal(v any) (marshaledBytes []byte, err error) {
	if marshaledBytes, err = json.Marshal(v); err != nil {
		err = errors.Wrap(err, `json.Marshal failed`)
	}
	return
}

// MarshalIndent 将 Go 对象编码为带有缩进的 JSON 字节数组
func MarshalIndent(v any, prefix, indent string) (marshaledBytes []byte, err error) {
	if marshaledBytes, err = json.MarshalIndent(v, prefix, indent); err != nil {
		err = errors.Wrap(err, `json.MarshalIndent failed`)
	}
	return
}

// Unmarshal 将 JSON 字节数组解码为 Go 对象
func Unmarshal(data []byte, v any) (err error) {
	if err = json.Unmarshal(data, v); err != nil {
		err = errors.Wrap(err, `json.Unmarshal failed`)
	}
	return
}

// UnmarshalUseNumber 使用 number 选项将 JSON 数据字节解码为目标接口
func UnmarshalUseNumber(data []byte, v any) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err = decoder.Decode(v); err != nil {
		err = errors.Wrap(err, `json.UnmarshalUseNumber failed`)
	}
	return
}

// NewEncoder 创建一个将 JSON 编码写入指定 io.Writer 的 Encoder 对象
func NewEncoder(writer io.Writer) (encoder *json.Encoder) {
	return json.NewEncoder(writer)
}

// NewDecoder 从输入流中读取 JSON 数据并解析成 Go 对象
func NewDecoder(reader io.Reader) (decoder *json.Decoder) {
	return json.NewDecoder(reader)
}

// Valid 判断一个 JSON 字符串是否合法
func Valid(data []byte) bool {
	return json.Valid(data)
}
