/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-20 20:23:24
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/nlog_demo/main.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import "github.com/liusuxian/nova/nlog"

func main() {
	nlog.Debug("Nlog Debug")
	nlog.Debugf("Nlog Debugf: %s", "Debugf")
	nlog.Info("Nlog Info")
	nlog.Infof("Nlog Infof: %s", "Infof")
}
