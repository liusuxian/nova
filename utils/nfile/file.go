/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 21:04:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-04-02 01:23:00
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/nfile/file.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nfile

import (
	"os"
	"path/filepath"
	"strings"
)

// PathExists 判断文件或者目录是否存在
func PathExists(path string) (isExist bool) {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// ExtName 获取文件扩展名
func ExtName(path string) (extName string) {
	return strings.TrimLeft(filepath.Ext(path), ".")
}
