/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 20:10:04
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 20:55:51
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_maps.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import "github.com/liusuxian/nova/internal/json"

// ToMaps 将 any 转换为 []map[string]any 类型
func ToMaps(value any, tags ...string) (ms []map[string]any) {
	if value == nil {
		return nil
	}
	switch r := value.(type) {
	case string:
		list := make([]map[string]any, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber([]byte(r), &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}
	case []byte:
		list := make([]map[string]any, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber(r, &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}
	case []map[string]any:
		return r
	default:
		array := ToSlice(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]any, len(array))
		for k, v := range array {
			list[k] = ToMap(v, tags...)
		}
		return list
	}
}

// ToMapsDeep 递归地对 value 进行 ToMaps 函数操作
func ToMapsDeep(value any, tags ...string) (ms []map[string]any) {
	if value == nil {
		return nil
	}
	switch r := value.(type) {
	case string:
		list := make([]map[string]any, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber([]byte(r), &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}
	case []byte:
		list := make([]map[string]any, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber(r, &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}
	case []map[string]any:
		list := make([]map[string]any, len(r))
		for k, v := range r {
			list[k] = ToMapDeep(v, tags...)
		}
		return list
	default:
		array := ToSlice(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]any, len(array))
		for k, v := range array {
			list[k] = ToMapDeep(v, tags...)
		}
		return list
	}
}
