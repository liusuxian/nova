/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 14:27:18
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:42:19
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapBoolE 将 any 转换为 map[string]bool 类型
func ToStringMapBoolE(i any, opts ...DecoderConfigOption) (iv map[string]bool, err error) {
	if i == nil {
		return map[string]bool{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]bool, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]bool{}, convertError(i, "map[string]bool")
			}
			value, err := ToBoolE(v)
			if err != nil {
				return map[string]bool{}, convertError(i, "map[string]bool")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]bool, len(val))
		for k, v := range val {
			value, err := ToBoolE(v)
			if err != nil {
				return map[string]bool{}, convertError(i, "map[string]bool")
			}
			iv[k] = value
		}
		return
	case map[any]bool:
		iv = make(map[string]bool, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]bool{}, convertError(i, "map[string]bool")
			}
			iv[key] = v
		}
		return
	case map[string]bool:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]bool, len(im))
				for k, v := range im {
					value, err := ToBoolE(v)
					if err != nil {
						return map[string]bool{}, convertError(i, "map[string]bool")
					}
					iv[k] = value
				}
				return
			}
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			im := map[string]any{}
			if e := json.Unmarshal(anyBytes, &im); e == nil {
				iv = make(map[string]bool, len(im))
				for k, v := range im {
					value, err := ToBoolE(v)
					if err != nil {
						return map[string]bool{}, convertError(i, "map[string]bool")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]bool{}, convertError(i, "map[string]bool")
	}
	iv = make(map[string]bool, len(im))
	for k, v := range im {
		value, err := ToBoolE(v)
		if err != nil {
			return map[string]bool{}, convertError(i, "map[string]bool")
		}
		iv[k] = value
	}
	return
}
