/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 15:15:07
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-10 20:07:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_map.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/empty"
	"github.com/liusuxian/nova/internal/json"
	"github.com/liusuxian/nova/internal/utils"
	"reflect"
	"strings"
)

// doMapConvertInput
type doMapConvertInput struct {
	IsRoot    bool     // 如果不是根节点，且无需进行递归转换，则直接返回
	Value     any      // 当前操作的值
	Recursive bool     // 是否对当前操作进行递归转换
	Tags      []string // Map 键映射
}

// ToMap 将 any 转换为 map[string]any 类型
//
//	如果参数`value`不是`map/struct/*struct`类型，则转换失败并返回`nil`
func ToMap(value any, tags ...string) (mVal map[string]any) {
	return mapConvert(value, false, tags...)
}

// ToMapDeep 递归地对 value 进行 ToMap 函数操作
//
//	如果`value`的属性也是一个`struct/*struct`，它会对该属性调用`ToMap`函数，将其转换为`map[string]any`类型的变量
func ToMapDeep(value any, tags ...string) (mVal map[string]any) {
	return mapConvert(value, true, tags...)
}

// mapConvert 实现了 map 类型的转换
//
//	如果`value`是`string/[]byte`类型，它会自动检查并将其转换为`map`
func mapConvert(value any, recursive bool, tags ...string) (mVal map[string]any) {
	if value == nil {
		return nil
	}
	newTags := StructTagPriority
	switch len(tags) {
	case 0:
		// 不需要处理
	case 1:
		newTags = append(strings.Split(tags[0], ","), StructTagPriority...)
	default:
		newTags = append(tags, StructTagPriority...)
	}
	dataMap := make(map[string]any)
	switch r := value.(type) {
	case string:
		// 如果它是 JSON 字符串，自动反序列化它
		if len(r) > 0 && r[0] == '{' && r[len(r)-1] == '}' {
			if err := json.UnmarshalUseNumber([]byte(r), &dataMap); err != nil {
				return nil
			}
		} else {
			return nil
		}
	case []byte:
		// 如果它是 JSON 字符串，自动反序列化它
		if len(r) > 0 && r[0] == '{' && r[len(r)-1] == '}' {
			if err := json.UnmarshalUseNumber(r, &dataMap); err != nil {
				return nil
			}
		} else {
			return nil
		}
	case map[any]any:
		for k, v := range r {
			dataMap[ToString(k)] = doMapConvert(doMapConvertInput{
				IsRoot:    false,
				Value:     v,
				Recursive: recursive,
				Tags:      newTags,
			})
		}
	case map[any]string:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[any]int:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[any]uint:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[any]float32:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[any]float64:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[string]bool:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]int:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]uint:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]float32:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]float64:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]string:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]any:
		if recursive {
			for k, v := range r {
				dataMap[k] = doMapConvert(doMapConvertInput{
					IsRoot:    false,
					Value:     v,
					Recursive: recursive,
					Tags:      newTags,
				})
			}
		} else {
			return r
		}
	case map[int]any:
		for k, v := range r {
			dataMap[ToString(k)] = doMapConvert(doMapConvertInput{
				IsRoot:    false,
				Value:     v,
				Recursive: recursive,
				Tags:      newTags,
			})
		}
	case map[int]string:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	case map[uint]string:
		for k, v := range r {
			dataMap[ToString(k)] = v
		}
	default:
		var reflectValue reflect.Value
		if v, ok := value.(reflect.Value); ok {
			reflectValue = v
		} else {
			reflectValue = reflect.ValueOf(value)
		}
		reflectKind := reflectValue.Kind()
		// 如果它是一个指针，我们需要找到它的实际数据类型
		for reflectKind == reflect.Ptr {
			reflectValue = reflectValue.Elem()
			reflectKind = reflectValue.Kind()
		}
		switch reflectKind {
		// 如果 value 是数组类型，它将把偶数索引的值转换为其键，奇数索引的值作为其对应的值
		case reflect.Slice, reflect.Array:
			length := reflectValue.Len()
			for i := 0; i < length; i += 2 {
				if i+1 < length {
					dataMap[ToString(reflectValue.Index(i).Interface())] = reflectValue.Index(i + 1).Interface()
				} else {
					dataMap[ToString(reflectValue.Index(i).Interface())] = nil
				}
			}
		case reflect.Map, reflect.Struct, reflect.Interface:
			convertedValue := doMapConvert(doMapConvertInput{
				IsRoot:    true,
				Value:     value,
				Recursive: recursive,
				Tags:      newTags,
			},
			)
			if m, ok := convertedValue.(map[string]any); ok {
				return m
			}
			return nil
		default:
			return nil
		}
	}
	return dataMap
}

// ToMapStrStr 将 value 转换为 map[string]string 类型
//
//	注意，在进行此`map`类型转换时可能会发生数据复制
func ToMapStrStr(value any, tags ...string) (mVal map[string]string) {
	if r, ok := value.(map[string]string); ok {
		return r
	}
	m := ToMap(value, tags...)
	if len(m) > 0 {
		vMap := make(map[string]string, len(m))
		for k, v := range m {
			vMap[k] = ToString(v)
		}
		return vMap
	}
	return nil
}

// ToMapStrStrDeep 递归地将 value 转换为 map[string]string 类型
//
//	注意，这种`map`类型的转换可能会进行数据复制
func ToMapStrStrDeep(value any, tags ...string) (mVal map[string]string) {
	if r, ok := value.(map[string]string); ok {
		return r
	}
	m := ToMapDeep(value, tags...)
	if len(m) > 0 {
		vMap := make(map[string]string, len(m))
		for k, v := range m {
			vMap[k] = ToString(v)
		}
		return vMap
	}
	return nil
}

// doMapConvert
func doMapConvert(in doMapConvertInput) (val any) {
	if !in.IsRoot && !in.Recursive {
		return in.Value
	}
	var reflectValue reflect.Value
	if v, ok := in.Value.(reflect.Value); ok {
		reflectValue = v
		in.Value = v.Interface()
	} else {
		reflectValue = reflect.ValueOf(in.Value)
	}
	reflectKind := reflectValue.Kind()
	// 如果它是一个指针，我们需要找到它的实际数据类型
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		mapKeys := reflectValue.MapKeys()
		dataMap := make(map[string]any)
		for _, k := range mapKeys {
			dataMap[ToString(k.Interface())] = doMapConvert(doMapConvertInput{
				IsRoot:    false,
				Value:     reflectValue.MapIndex(k).Interface(),
				Recursive: in.Recursive,
				Tags:      in.Tags,
			})
		}
		return dataMap
	case reflect.Struct:
		dataMap := make(map[string]any)
		// Map 转换接口检查
		if v, ok := in.Value.(iMapStrAny); ok {
			// 值拷贝，以确保并发安全
			for mapK, mapV := range v.MapStrAny() {
				if in.Recursive {
					dataMap[mapK] = doMapConvert(doMapConvertInput{
						IsRoot:    false,
						Value:     mapV,
						Recursive: in.Recursive,
						Tags:      in.Tags,
					})
				} else {
					dataMap[mapK] = mapV
				}
			}
			return dataMap
		}
		// 使用反射进行转换
		var rtField reflect.StructField
		var rvField reflect.Value
		reflectType := reflectValue.Type()
		mapKey := ""
		for i := 0; i < reflectValue.NumField(); i++ {
			rtField = reflectType.Field(i)
			rvField = reflectValue.Field(i)
			// 仅转换公共属性
			fieldName := rtField.Name
			if !utils.IsLetterUpper(fieldName[0]) {
				continue
			}
			mapKey = ""
			fieldTag := rtField.Tag
			for _, tag := range in.Tags {
				if mapKey = fieldTag.Get(tag); mapKey != "" {
					break
				}
			}
			if mapKey == "" {
				mapKey = fieldName
			} else {
				// 支持 json tag 的特性：-, omitempty
				mapKey = strings.TrimSpace(mapKey)
				if mapKey == "-" {
					continue
				}
				array := strings.Split(mapKey, ",")
				if len(array) > 1 {
					switch strings.TrimSpace(array[1]) {
					case "omitempty":
						if empty.IsEmpty(rvField.Interface()) {
							continue
						} else {
							mapKey = strings.TrimSpace(array[0])
						}
					default:
						mapKey = strings.TrimSpace(array[0])
					}
				}
				if mapKey == "" {
					mapKey = fieldName
				}
			}
			if in.Recursive || rtField.Anonymous {
				// 递归执行 map 转换
				rvAttrField := rvField
				rvAttrKind := rvField.Kind()
				if rvAttrKind == reflect.Ptr {
					rvAttrField = rvField.Elem()
					rvAttrKind = rvAttrField.Kind()
				}
				switch rvAttrKind {
				case reflect.Struct:
					// 如果嵌入了一个结构体并且没有字段，则忽略它
					if rvAttrField.Type().NumField() == 0 {
						continue
					}
					hasNoTag := mapKey == fieldName
					// 在这里不要使用 rvAttrField.Interface()，因为它可能从指针变成结构体
					rvInterface := rvField.Interface()
					switch {
					case hasNoTag && rtField.Anonymous:
						// 如果该属性字段没有标签，使用子结构体的属性字段覆盖该属性
						anonymousValue := doMapConvert(doMapConvertInput{
							IsRoot:    false,
							Value:     rvInterface,
							Recursive: in.Recursive,
							Tags:      in.Tags,
						})
						if m, ok := anonymousValue.(map[string]any); ok {
							for k, v := range m {
								dataMap[k] = v
							}
						} else {
							dataMap[mapKey] = rvInterface
						}
					// 这个属性字段包含所需的标签
					case !hasNoTag && rtField.Anonymous:
						dataMap[mapKey] = doMapConvert(doMapConvertInput{
							IsRoot:    false,
							Value:     rvInterface,
							Recursive: in.Recursive,
							Tags:      in.Tags,
						})
					default:
						dataMap[mapKey] = doMapConvert(doMapConvertInput{
							IsRoot:    false,
							Value:     rvInterface,
							Recursive: in.Recursive,
							Tags:      in.Tags,
						})
					}
				// 这个结构体属性是一个切片类型
				case reflect.Array, reflect.Slice:
					length := rvAttrField.Len()
					if length == 0 {
						dataMap[mapKey] = rvAttrField.Interface()
						break
					}
					array := make([]any, length)
					for arrayIndex := 0; arrayIndex < length; arrayIndex++ {
						array[arrayIndex] = doMapConvert(doMapConvertInput{
							IsRoot:    false,
							Value:     rvAttrField.Index(arrayIndex).Interface(),
							Recursive: in.Recursive,
							Tags:      in.Tags,
						})
					}
					dataMap[mapKey] = array
				case reflect.Map:
					mapKeys := rvAttrField.MapKeys()
					nestedMap := make(map[string]any)
					for _, k := range mapKeys {
						nestedMap[ToString(k.Interface())] = doMapConvert(doMapConvertInput{
							IsRoot:    false,
							Value:     rvAttrField.MapIndex(k).Interface(),
							Recursive: in.Recursive,
							Tags:      in.Tags,
						})
					}
					dataMap[mapKey] = nestedMap
				default:
					if rvField.IsValid() {
						dataMap[mapKey] = reflectValue.Field(i).Interface()
					} else {
						dataMap[mapKey] = nil
					}
				}
			} else {
				// 不进行递归的 map 值转换
				if rvField.IsValid() {
					dataMap[mapKey] = reflectValue.Field(i).Interface()
				} else {
					dataMap[mapKey] = nil
				}
			}
		}
		if len(dataMap) == 0 {
			return in.Value
		}
		return dataMap
	// 给定的值是切片类型
	case reflect.Array, reflect.Slice:
		length := reflectValue.Len()
		if length == 0 {
			break
		}
		array := make([]any, reflectValue.Len())
		for i := 0; i < length; i++ {
			array[i] = doMapConvert(doMapConvertInput{
				IsRoot:    false,
				Value:     reflectValue.Index(i).Interface(),
				Recursive: in.Recursive,
				Tags:      in.Tags,
			})
		}
		return array
	}
	return in.Value
}
