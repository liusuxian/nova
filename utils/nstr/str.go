/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-22 12:17:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-13 21:38:17
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nstr/str.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nstr

import "strings"

// Split 用sep将str拆分为[]string
func Split(str, sep string) (list []string) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}
	list = strings.Split(str, sep)
	return
}
