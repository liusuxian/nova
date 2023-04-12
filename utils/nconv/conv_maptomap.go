/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 21:53:19
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-12 18:30:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_maptomap.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"github.com/liusuxian/nova/internal/json"
	"github.com/pkg/errors"
	"reflect"
)

// MapToMap 使用反射将任何类型的 map 变量 params 转换为另一个类型的 map 变量 pointer
//
//	参数`params`可以是任何类型的`map`
//	参数`pointer`应该是`*map`类型
//	可选参数`mapping`用于将结构体属性映射到`map`键上，仅当原始`map params`中的项是结构体类型时才有意义
func MapToMap(params, pointer any, mapping ...map[string]string) (err error) {
	return doMapToMap(params, pointer, mapping...)
}

// doMapToMap 将任何类型的 map 变量 params 转换为另一个类型的 map 变量 pointer
//
//	参数`params`可以是任何类型的`map`
//	参数`pointer`应该是`*map`类型
//	可选参数`mapping`用于将结构体属性映射到`map`键上，仅当原始`map params`中的项是结构体类型时才有意义
func doMapToMap(params, pointer any, mapping ...map[string]string) (err error) {
	// 如果给定的 params 是 JSON，则使用 json.Unmarshal 进行转换
	switch r := params.(type) {
	case []byte:
		if json.Valid(r) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					return json.UnmarshalUseNumber(r, rv.Interface())
				}
			} else {
				return json.UnmarshalUseNumber(r, pointer)
			}
		}
	case string:
		if paramsBytes := []byte(r); json.Valid(paramsBytes) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					return json.UnmarshalUseNumber(paramsBytes, rv.Interface())
				}
			} else {
				return json.UnmarshalUseNumber(paramsBytes, pointer)
			}
		}
	}
	var paramsRv reflect.Value
	var paramsKind reflect.Kind
	var keyToAttributeNameMapping map[string]string
	if len(mapping) > 0 {
		keyToAttributeNameMapping = mapping[0]
	}
	if v, ok := params.(reflect.Value); ok {
		paramsRv = v
	} else {
		paramsRv = reflect.ValueOf(params)
	}
	paramsKind = paramsRv.Kind()
	if paramsKind == reflect.Ptr {
		paramsRv = paramsRv.Elem()
		paramsKind = paramsRv.Kind()
	}
	if paramsKind != reflect.Map {
		return doMapToMap(ToMap(params), pointer, mapping...)
	}
	if paramsRv.Len() == 0 {
		return nil
	}
	var pointerRv reflect.Value
	if v, ok := pointer.(reflect.Value); ok {
		pointerRv = v
	} else {
		pointerRv = reflect.ValueOf(pointer)
	}
	pointerKind := pointerRv.Kind()
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Map {
		return errors.Errorf("pointer should be type of *map, but got: %s", pointerKind)
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok {
				err = v
			} else {
				err = errors.Errorf("internal error: %+v", exception)
			}
		}
	}()
	paramsKeys := paramsRv.MapKeys()
	pointerKeyType := pointerRv.Type().Key()
	pointerValueType := pointerRv.Type().Elem()
	pointerValueKind := pointerValueType.Kind()
	dataMap := reflect.MakeMapWithSize(pointerRv.Type(), len(paramsKeys))
	// 检索目标映射的真实元素类型
	if pointerValueKind == reflect.Ptr {
		pointerValueKind = pointerValueType.Elem().Kind()
	}
	for _, key := range paramsKeys {
		e := reflect.New(pointerValueType).Elem()
		switch pointerValueKind {
		case reflect.Map, reflect.Struct:
			if err = doStruct(paramsRv.MapIndex(key).Interface(), e, keyToAttributeNameMapping, ""); err != nil {
				return err
			}
		default:
			e.Set(reflect.ValueOf(Convert(paramsRv.MapIndex(key).Interface(), pointerValueType.String())))
		}
		dataMap.SetMapIndex(reflect.ValueOf(Convert(key.Interface(), pointerKeyType.Name())), e)
	}
	pointerRv.Set(dataMap)
	return nil
}
