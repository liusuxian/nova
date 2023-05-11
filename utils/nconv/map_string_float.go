/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-05 15:21:28
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:43:23
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapFloat64E 将 any 转换为 map[string]float64 类型
func ToStringMapFloat64E(i any, opts ...DecoderConfigOption) (iv map[string]float64, err error) {
	if i == nil {
		return map[string]float64{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]float64, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]float64{}, convertError(i, "map[string]float64")
			}
			value, err := ToFloat64E(v)
			if err != nil {
				return map[string]float64{}, convertError(i, "map[string]float64")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]float64, len(val))
		for k, v := range val {
			value, err := ToFloat64E(v)
			if err != nil {
				return map[string]float64{}, convertError(i, "map[string]float64")
			}
			iv[k] = value
		}
		return
	case map[any]float64:
		iv = make(map[string]float64, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]float64{}, convertError(i, "map[string]float64")
			}
			iv[key] = v
		}
		return
	case map[string]float64:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]float64, len(im))
				for k, v := range im {
					value, err := ToFloat64E(v)
					if err != nil {
						return map[string]float64{}, convertError(i, "map[string]float64")
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
				iv = make(map[string]float64, len(im))
				for k, v := range im {
					value, err := ToFloat64E(v)
					if err != nil {
						return map[string]float64{}, convertError(i, "map[string]float64")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]float64{}, convertError(i, "map[string]float64")
	}
	iv = make(map[string]float64, len(im))
	for k, v := range im {
		value, err := ToFloat64E(v)
		if err != nil {
			return map[string]float64{}, convertError(i, "map[string]float64")
		}
		iv[k] = value
	}
	return
}

// ToStringMapFloat32E 将 any 转换为 map[string]float32 类型
func ToStringMapFloat32E(i any, opts ...DecoderConfigOption) (iv map[string]float32, err error) {
	if i == nil {
		return map[string]float32{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]float32, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]float32{}, convertError(i, "map[string]float32")
			}
			value, err := ToFloat32E(v)
			if err != nil {
				return map[string]float32{}, convertError(i, "map[string]float32")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]float32, len(val))
		for k, v := range val {
			value, err := ToFloat32E(v)
			if err != nil {
				return map[string]float32{}, convertError(i, "map[string]float32")
			}
			iv[k] = value
		}
		return
	case map[any]float32:
		iv = make(map[string]float32, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]float32{}, convertError(i, "map[string]float32")
			}
			iv[key] = v
		}
		return
	case map[string]float32:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]float32, len(im))
				for k, v := range im {
					value, err := ToFloat32E(v)
					if err != nil {
						return map[string]float32{}, convertError(i, "map[string]float32")
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
				iv = make(map[string]float32, len(im))
				for k, v := range im {
					value, err := ToFloat32E(v)
					if err != nil {
						return map[string]float32{}, convertError(i, "map[string]float32")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]float32{}, convertError(i, "map[string]float32")
	}
	iv = make(map[string]float32, len(im))
	for k, v := range im {
		value, err := ToFloat32E(v)
		if err != nil {
			return map[string]float32{}, convertError(i, "map[string]float32")
		}
		iv[k] = value
	}
	return
}
