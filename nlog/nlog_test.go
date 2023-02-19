/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 22:25:41
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-20 01:21:53
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/nlog_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog_test

import (
	"github.com/liusuxian/nova/nlog"
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	nlog.Debug("Nlog Debug")
	nlog.Debugf("Nlog Debugf: %s", "Debugf")
	nlog.Info("Nlog Info")
	nlog.Infof("Nlog Infof: %s", "Infof")
	nlog.DebugFields("Nlog DebugFields", zap.String("key", "value"))
	nlog.Error("Nlog Error")
	nlog.ErrorFields("我是一个错误", zap.String("url", "www.baidu.com"))
	nlog.Fatal("Nlog Fatal")
}
