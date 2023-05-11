/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 21:04:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:24:11
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
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

// GetContents 获取文件的内容
func GetContents(path string) (str string) {
	return string(GetBytes(path))
}

// GetBytes 获取文件的内容
func GetBytes(path string) (buf []byte) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// Name 返回路径的最后一个元素，但不包括文件扩展名
func Name(path string) (str string) {
	base := filepath.Base(path)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		return base[:i]
	}
	return base
}
