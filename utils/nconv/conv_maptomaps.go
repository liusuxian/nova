/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-10 22:20:40
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-12 18:37:09
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/conv_maptomaps.go
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

// MapToMaps 将任何切片类型变量 params 转换为另一个 map 切片类型变量 pointer
//
//	参数`params`可以是`[]map、[]*map、[]struct、[]*struct`类型之一
//	参数`pointer`应该是`[]map、[]*map`类型之一
//	可选参数`mapping`用于将结构体属性映射到`map`键上，仅在`params`的元素是结构体类型时有意义
func MapToMaps(params, pointer any, mapping ...map[string]string) (err error) {
	return doMapToMaps(params, pointer, mapping...)
}

// doMapToMaps 将任何切片类型变量 params 转换为另一个 map 切片类型变量 pointer
//
//	参数`params`可以是`[]map、[]*map、[]struct、[]*struct`类型之一
//	参数`pointer`应该是`[]map、[]*map`类型之一
//	可选参数`mapping`用于将结构体属性映射到`map`键上，仅在`params`的元素是结构体类型时有意义
func doMapToMaps(params, pointer any, mapping ...map[string]string) (err error) {
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
	if paramsKind != reflect.Array && paramsKind != reflect.Slice {
		return errors.New("params should be type of slice, eg: []map/[]*map/[]struct/[]*struct")
	}
	paramsElem := paramsRv.Type().Elem()
	paramsElemKind := paramsElem.Kind()
	if paramsElemKind == reflect.Ptr {
		paramsElem = paramsElem.Elem()
		paramsElemKind = paramsElem.Kind()
	}
	if paramsElemKind != reflect.Map && paramsElemKind != reflect.Struct && paramsElemKind != reflect.Interface {
		return errors.Errorf("params element should be type of map/*map/struct/*struct, but got: %s", paramsElemKind)
	}
	if paramsRv.Len() == 0 {
		return nil
	}
	pointerRv := reflect.ValueOf(pointer)
	pointerKind := pointerRv.Kind()
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Array && pointerKind != reflect.Slice {
		return errors.New("pointer should be type of *[]map/*[]*map")
	}
	pointerElemType := pointerRv.Type().Elem()
	pointerElemKind := pointerElemType.Kind()
	if pointerElemKind == reflect.Ptr {
		pointerElemKind = pointerElemType.Elem().Kind()
	}
	if pointerElemKind != reflect.Map {
		return errors.New("pointer element should be type of map/*map")
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
	pointerSlice := reflect.MakeSlice(pointerRv.Type(), paramsRv.Len(), paramsRv.Len())
	for i := 0; i < paramsRv.Len(); i++ {
		var item reflect.Value
		if pointerElemType.Kind() == reflect.Ptr {
			item = reflect.New(pointerElemType.Elem())
			if err = MapToMap(paramsRv.Index(i).Interface(), item, mapping...); err != nil {
				return err
			}
			pointerSlice.Index(i).Set(item)
		} else {
			item = reflect.New(pointerElemType)
			if err = MapToMap(paramsRv.Index(i).Interface(), item, mapping...); err != nil {
				return err
			}
			pointerSlice.Index(i).Set(item.Elem())
		}
	}
	pointerRv.Set(pointerSlice)
	return
}
