/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 15:31:00
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:50:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_string.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapStringE 将 any 转换为 map[string]string 类型
func ToStringMapStringE(i any, opts ...DecoderConfigOption) (iv map[string]string, err error) {
	if i == nil {
		return map[string]string{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]string{}, convertError(i, "map[string]string")
			}
			value, err := ToStringE(v)
			if err != nil {
				return map[string]string{}, convertError(i, "map[string]string")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]string, len(val))
		for k, v := range val {
			value, err := ToStringE(v)
			if err != nil {
				return map[string]string{}, convertError(i, "map[string]string")
			}
			iv[k] = value
		}
		return
	case map[any]string:
		iv = make(map[string]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]string{}, convertError(i, "map[string]string")
			}
			iv[key] = v
		}
		return
	case map[string]string:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]string, len(im))
				for k, v := range im {
					value, err := ToStringE(v)
					if err != nil {
						return map[string]string{}, convertError(i, "map[string]string")
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
				iv = make(map[string]string, len(im))
				for k, v := range im {
					value, err := ToStringE(v)
					if err != nil {
						return map[string]string{}, convertError(i, "map[string]string")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]string{}, convertError(i, "map[string]string")
	}
	iv = make(map[string]string, len(im))
	for k, v := range im {
		value, err := ToStringE(v)
		if err != nil {
			return map[string]string{}, convertError(i, "map[string]string")
		}
		iv[k] = value
	}
	return
}

// ToStringMapStringSliceE 将 any 转换为 map[string][]string 类型
func ToStringMapStringSliceE(i any, opts ...DecoderConfigOption) (iv map[string][]string, err error) {
	if i == nil {
		return map[string][]string{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[key] = value
		}
		return
	case map[any][]any:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[key] = value
		}
		return
	case map[any]string:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[key] = value
		}
		return
	case map[any][]string:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[key] = v
		}
		return
	case map[string]any:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[k] = value
		}
		return
	case map[string]string:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[k] = value
		}
		return
	case map[string][]any:
		iv = make(map[string][]string, len(val))
		for k, v := range val {
			value, err := ToStringSliceE(v)
			if err != nil {
				return map[string][]string{}, convertError(i, "map[string][]string")
			}
			iv[k] = value
		}
		return
	case map[string][]string:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string][]string, len(im))
				for k, v := range im {
					value, err := ToStringSliceE(v)
					if err != nil {
						return map[string][]string{}, convertError(i, "map[string][]string")
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
				iv = make(map[string][]string, len(im))
				for k, v := range im {
					value, err := ToStringSliceE(v)
					if err != nil {
						return map[string][]string{}, convertError(i, "map[string][]string")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string][]string{}, convertError(i, "map[string][]string")
	}
	iv = make(map[string][]string, len(im))
	for k, v := range im {
		value, err := ToStringSliceE(v)
		if err != nil {
			return map[string][]string{}, convertError(i, "map[string][]string")
		}
		iv[k] = value
	}
	return
}
