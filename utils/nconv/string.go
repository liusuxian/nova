/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-16 01:11:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-16 03:26:23
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nconv/string.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconv

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
)

// ToStringE 将 any 转换为 string 类型
func ToStringE(i any) (iv string, err error) {
	i = indirectToStringerOrError(i)

	switch val := i.(type) {
	case nil:
		return "", nil
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case bool:
		return strconv.FormatBool(val), nil
	case json.Number:
		return val.String(), nil
	case template.HTML:
		return string(val), nil
	case template.URL:
		return string(val), nil
	case template.JS:
		return string(val), nil
	case template.CSS:
		return string(val), nil
	case template.HTMLAttr:
		return string(val), nil
	case fmt.Stringer:
		return val.String(), nil
	case error:
		return val.Error(), nil
	default:
		// 使用 json.Marshal 函数进行转换
		jsonContent, err := json.Marshal(val)
		if err == nil {
			return string(jsonContent), nil
		}
		return "", convertError(i, "string")
	}
}
