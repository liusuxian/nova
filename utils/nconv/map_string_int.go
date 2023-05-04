/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-03 16:47:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-04 11:27:48
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_int.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapInt64E 将 any 转换为 map[string]int64 类型
func ToStringMapInt64E(i any, opts ...DecoderConfigOption) (iv map[string]int64, err error) {
	if i == nil {
		return map[string]int64{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]int64, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]int64{}, convertError(i, "map[string]int64")
			}
			value, err := ToInt64E(v)
			if err != nil {
				return map[string]int64{}, convertError(i, "map[string]int64")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]int64, len(val))
		for k, v := range val {
			value, err := ToInt64E(v)
			if err != nil {
				return map[string]int64{}, convertError(i, "map[string]int64")
			}
			iv[k] = value
		}
		return
	case map[string]int64:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if err := json.Unmarshal(val, &iv); err != nil {
				return map[string]int64{}, convertError(i, "map[string]int64")
			}
			return
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := json.Unmarshal(anyBytes, &iv); err != nil {
				return map[string]int64{}, convertError(i, "map[string]int64")
			}
			return
		}
	}

	iv = map[string]int64{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]int64{}, convertError(i, "map[string]int64")
	}
	return
}

// ToStringMapInt32E 将 any 转换为 map[string]int32 类型
func ToStringMapInt32E(i any, opts ...DecoderConfigOption) (iv map[string]int32, err error) {
	if i == nil {
		return map[string]int32{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]int32, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]int32{}, convertError(i, "map[string]int32")
			}
			value, err := ToInt32E(v)
			if err != nil {
				return map[string]int32{}, convertError(i, "map[string]int32")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]int32, len(val))
		for k, v := range val {
			value, err := ToInt32E(v)
			if err != nil {
				return map[string]int32{}, convertError(i, "map[string]int32")
			}
			iv[k] = value
		}
		return
	case map[string]int32:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if err := json.Unmarshal(val, &iv); err != nil {
				return map[string]int32{}, convertError(i, "map[string]int32")
			}
			return
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := json.Unmarshal(anyBytes, &iv); err != nil {
				return map[string]int32{}, convertError(i, "map[string]int32")
			}
			return
		}
	}

	iv = map[string]int32{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]int32{}, convertError(i, "map[string]int32")
	}
	return
}

// ToStringMapInt16E 将 any 转换为 map[string]int16 类型
func ToStringMapInt16E(i any, opts ...DecoderConfigOption) (iv map[string]int16, err error) {
	if i == nil {
		return map[string]int16{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]int16, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]int16{}, convertError(i, "map[string]int16")
			}
			value, err := ToInt16E(v)
			if err != nil {
				return map[string]int16{}, convertError(i, "map[string]int16")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]int16, len(val))
		for k, v := range val {
			value, err := ToInt16E(v)
			if err != nil {
				return map[string]int16{}, convertError(i, "map[string]int16")
			}
			iv[k] = value
		}
		return
	case map[string]int16:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if err := json.Unmarshal(val, &iv); err != nil {
				return map[string]int16{}, convertError(i, "map[string]int16")
			}
			return
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := json.Unmarshal(anyBytes, &iv); err != nil {
				return map[string]int16{}, convertError(i, "map[string]int16")
			}
			return
		}
	}

	iv = map[string]int16{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]int16{}, convertError(i, "map[string]int16")
	}
	return
}

// ToStringMapInt8E 将 any 转换为 map[string]int8 类型
func ToStringMapInt8E(i any, opts ...DecoderConfigOption) (iv map[string]int8, err error) {
	if i == nil {
		return map[string]int8{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]int8, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]int8{}, convertError(i, "map[string]int8")
			}
			value, err := ToInt8E(v)
			if err != nil {
				return map[string]int8{}, convertError(i, "map[string]int8")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]int8, len(val))
		for k, v := range val {
			value, err := ToInt8E(v)
			if err != nil {
				return map[string]int8{}, convertError(i, "map[string]int8")
			}
			iv[k] = value
		}
		return
	case map[string]int8:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if err := json.Unmarshal(val, &iv); err != nil {
				return map[string]int8{}, convertError(i, "map[string]int8")
			}
			return
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := json.Unmarshal(anyBytes, &iv); err != nil {
				return map[string]int8{}, convertError(i, "map[string]int8")
			}
			return
		}
	}

	iv = map[string]int8{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]int8{}, convertError(i, "map[string]int8")
	}
	return
}

// ToStringMapIntE 将 any 转换为 map[string]int 类型
func ToStringMapIntE(i any, opts ...DecoderConfigOption) (iv map[string]int, err error) {
	if i == nil {
		return map[string]int{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]int, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]int{}, convertError(i, "map[string]int")
			}
			value, err := ToIntE(v)
			if err != nil {
				return map[string]int{}, convertError(i, "map[string]int")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]int, len(val))
		for k, v := range val {
			value, err := ToIntE(v)
			if err != nil {
				return map[string]int{}, convertError(i, "map[string]int")
			}
			iv[k] = value
		}
		return
	case map[string]int:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			if err := json.Unmarshal(val, &iv); err != nil {
				return map[string]int{}, convertError(i, "map[string]int")
			}
			return
		}
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		anyBytes := []byte(val)
		if json.Valid(anyBytes) {
			if err := json.Unmarshal(anyBytes, &iv); err != nil {
				return map[string]int{}, convertError(i, "map[string]int")
			}
			return
		}
	}

	iv = map[string]int{}
	if err := decode(i, defaultDecoderConfig(&iv, opts...)); err != nil {
		return map[string]int{}, convertError(i, "map[string]int")
	}
	return
}
