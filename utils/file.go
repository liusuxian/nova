/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 21:04:58
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 21:06:26
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/utils/file.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package utils

import "os"

// PathExists 判断文件或者目录是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
