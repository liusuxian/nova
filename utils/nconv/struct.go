/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-15 13:23:47
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:08:19
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/struct.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// ToStructE 将 any 转换为 struct/[]struct 类型
func ToStructE(input, output any, opts ...DecoderConfigOption) (err error) {
	if input == nil {
		return nil
	}
	if output == nil {
		return errors.New("object output cannot be nil")
	}

	switch val := input.(type) {
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if e := json.Unmarshal(val, output); e == nil {
				return
			}
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if e := json.Unmarshal(anyBytes, output); e == nil {
				return
			}
		}
	}

	return decode(input, defaultDecoderConfig(output, opts...))
}
