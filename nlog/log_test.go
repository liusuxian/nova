/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 22:15:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-11 14:18:01
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nlog_test

import (
	"github.com/liusuxian/nova/nlog"
	"testing"
)

func TestLog(t *testing.T) {
	nlog.Debug("I am Debug", nlog.Int("Int", 1))
	nlog.Info("I am Info", nlog.Any("Array", []int{1, 2, 3}))
	nlog.Warn("I am Warn")
	nlog.Error("I am Error")
	nlog.DPanic("I am DPanic")
	nlog.Panic("I am Panic")
	nlog.Fatal("I am Fatal")
}
