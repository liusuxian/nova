/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-18 00:54:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:41:47
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_any.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapE 将 any 转换为 map[string]any 类型
func ToStringMapE(i any, opts ...DecoderConfigOption) (iv map[string]any, err error) {
	if i == nil {
		return map[string]any{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]any, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]any{}, convertError(i, "map[string]any")
			}
			iv[key] = v
		}
		return
	case map[string]any:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if e := json.Unmarshal(val, &iv); e == nil {
				return
			}
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if e := json.Unmarshal(anyBytes, &iv); e == nil {
				return
			}
		}
	}

	iv = map[string]any{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]any{}, convertError(i, "map[string]any")
	}
	return
}
