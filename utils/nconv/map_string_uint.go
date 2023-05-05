/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-04 14:02:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-05 17:31:56
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/map_string_uint.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "encoding/json"

// ToStringMapUint64E 将 any 转换为 map[string]uint64 类型
func ToStringMapUint64E(i any, opts ...DecoderConfigOption) (iv map[string]uint64, err error) {
	if i == nil {
		return map[string]uint64{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]uint64, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint64{}, convertError(i, "map[string]uint64")
			}
			value, err := ToUint64E(v)
			if err != nil {
				return map[string]uint64{}, convertError(i, "map[string]uint64")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]uint64, len(val))
		for k, v := range val {
			value, err := ToUint64E(v)
			if err != nil {
				return map[string]uint64{}, convertError(i, "map[string]uint64")
			}
			iv[k] = value
		}
		return
	case map[any]uint64:
		iv = make(map[string]uint64, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint64{}, convertError(i, "map[string]uint64")
			}
			iv[key] = v
		}
		return
	case map[string]uint64:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]uint64, len(im))
				for k, v := range im {
					value, err := ToUint64E(v)
					if err != nil {
						return map[string]uint64{}, convertError(i, "map[string]uint64")
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
				iv = make(map[string]uint64, len(im))
				for k, v := range im {
					value, err := ToUint64E(v)
					if err != nil {
						return map[string]uint64{}, convertError(i, "map[string]uint64")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]uint64{}, convertError(i, "map[string]uint64")
	}
	iv = make(map[string]uint64, len(im))
	for k, v := range im {
		value, err := ToUint64E(v)
		if err != nil {
			return map[string]uint64{}, convertError(i, "map[string]uint64")
		}
		iv[k] = value
	}
	return
}

// ToStringMapUint32E 将 any 转换为 map[string]uint32 类型
func ToStringMapUint32E(i any, opts ...DecoderConfigOption) (iv map[string]uint32, err error) {
	if i == nil {
		return map[string]uint32{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]uint32, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint32{}, convertError(i, "map[string]uint32")
			}
			value, err := ToUint32E(v)
			if err != nil {
				return map[string]uint32{}, convertError(i, "map[string]uint32")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]uint32, len(val))
		for k, v := range val {
			value, err := ToUint32E(v)
			if err != nil {
				return map[string]uint32{}, convertError(i, "map[string]uint32")
			}
			iv[k] = value
		}
		return
	case map[any]uint32:
		iv = make(map[string]uint32, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint32{}, convertError(i, "map[string]uint32")
			}
			iv[key] = v
		}
		return
	case map[string]uint32:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]uint32, len(im))
				for k, v := range im {
					value, err := ToUint32E(v)
					if err != nil {
						return map[string]uint32{}, convertError(i, "map[string]uint32")
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
				iv = make(map[string]uint32, len(im))
				for k, v := range im {
					value, err := ToUint32E(v)
					if err != nil {
						return map[string]uint32{}, convertError(i, "map[string]uint32")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]uint32{}, convertError(i, "map[string]uint32")
	}
	iv = make(map[string]uint32, len(im))
	for k, v := range im {
		value, err := ToUint32E(v)
		if err != nil {
			return map[string]uint32{}, convertError(i, "map[string]uint32")
		}
		iv[k] = value
	}
	return
}

// ToStringMapUint16E 将 any 转换为 map[string]uint16 类型
func ToStringMapUint16E(i any, opts ...DecoderConfigOption) (iv map[string]uint16, err error) {
	if i == nil {
		return map[string]uint16{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]uint16, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint16{}, convertError(i, "map[string]uint16")
			}
			value, err := ToUint16E(v)
			if err != nil {
				return map[string]uint16{}, convertError(i, "map[string]uint16")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]uint16, len(val))
		for k, v := range val {
			value, err := ToUint16E(v)
			if err != nil {
				return map[string]uint16{}, convertError(i, "map[string]uint16")
			}
			iv[k] = value
		}
		return
	case map[any]uint16:
		iv = make(map[string]uint16, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint16{}, convertError(i, "map[string]uint16")
			}
			iv[key] = v
		}
		return
	case map[string]uint16:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]uint16, len(im))
				for k, v := range im {
					value, err := ToUint16E(v)
					if err != nil {
						return map[string]uint16{}, convertError(i, "map[string]uint16")
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
				iv = make(map[string]uint16, len(im))
				for k, v := range im {
					value, err := ToUint16E(v)
					if err != nil {
						return map[string]uint16{}, convertError(i, "map[string]uint16")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]uint16{}, convertError(i, "map[string]uint16")
	}
	iv = make(map[string]uint16, len(im))
	for k, v := range im {
		value, err := ToUint16E(v)
		if err != nil {
			return map[string]uint16{}, convertError(i, "map[string]uint16")
		}
		iv[k] = value
	}
	return
}

// ToStringMapUint8E 将 any 转换为 map[string]uint8 类型
func ToStringMapUint8E(i any, opts ...DecoderConfigOption) (iv map[string]uint8, err error) {
	if i == nil {
		return map[string]uint8{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]uint8, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint8{}, convertError(i, "map[string]uint8")
			}
			value, err := ToUint8E(v)
			if err != nil {
				return map[string]uint8{}, convertError(i, "map[string]uint8")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]uint8, len(val))
		for k, v := range val {
			value, err := ToUint8E(v)
			if err != nil {
				return map[string]uint8{}, convertError(i, "map[string]uint8")
			}
			iv[k] = value
		}
		return
	case map[any]uint8:
		iv = make(map[string]uint8, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint8{}, convertError(i, "map[string]uint8")
			}
			iv[key] = v
		}
		return
	case map[string]uint8:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]uint8, len(im))
				for k, v := range im {
					value, err := ToUint8E(v)
					if err != nil {
						return map[string]uint8{}, convertError(i, "map[string]uint8")
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
				iv = make(map[string]uint8, len(im))
				for k, v := range im {
					value, err := ToUint8E(v)
					if err != nil {
						return map[string]uint8{}, convertError(i, "map[string]uint8")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]uint8{}, convertError(i, "map[string]uint8")
	}
	iv = make(map[string]uint8, len(im))
	for k, v := range im {
		value, err := ToUint8E(v)
		if err != nil {
			return map[string]uint8{}, convertError(i, "map[string]uint8")
		}
		iv[k] = value
	}
	return
}

// ToStringMapUintE 将 any 转换为 map[string]uint 类型
func ToStringMapUintE(i any, opts ...DecoderConfigOption) (iv map[string]uint, err error) {
	if i == nil {
		return map[string]uint{}, nil
	}

	switch val := i.(type) {
	case map[any]any:
		iv = make(map[string]uint, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint{}, convertError(i, "map[string]uint")
			}
			value, err := ToUintE(v)
			if err != nil {
				return map[string]uint{}, convertError(i, "map[string]uint")
			}
			iv[key] = value
		}
		return
	case map[string]any:
		iv = make(map[string]uint, len(val))
		for k, v := range val {
			value, err := ToUintE(v)
			if err != nil {
				return map[string]uint{}, convertError(i, "map[string]uint")
			}
			iv[k] = value
		}
		return
	case map[any]uint:
		iv = make(map[string]uint, len(val))
		for k, v := range val {
			key, err := ToStringE(k)
			if err != nil {
				return map[string]uint{}, convertError(i, "map[string]uint")
			}
			iv[key] = v
		}
		return
	case map[string]uint:
		return val, nil
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if json.Valid(val) {
			im := map[string]any{}
			if e := json.Unmarshal(val, &im); e == nil {
				iv = make(map[string]uint, len(im))
				for k, v := range im {
					value, err := ToUintE(v)
					if err != nil {
						return map[string]uint{}, convertError(i, "map[string]uint")
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
				iv = make(map[string]uint, len(im))
				for k, v := range im {
					value, err := ToUintE(v)
					if err != nil {
						return map[string]uint{}, convertError(i, "map[string]uint")
					}
					iv[k] = value
				}
				return
			}
		}
	}

	im := map[string]any{}
	if err := decode(i, defaultDecoderConfig(&im, opts...)); err != nil {
		return map[string]uint{}, convertError(i, "map[string]uint")
	}
	iv = make(map[string]uint, len(im))
	for k, v := range im {
		value, err := ToUintE(v)
		if err != nil {
			return map[string]uint{}, convertError(i, "map[string]uint")
		}
		iv[k] = value
	}
	return
}
