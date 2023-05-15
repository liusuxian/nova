/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 22:15:16
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 17:28:56
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nlog_test

import (
	"github.com/liusuxian/nova/nlog"
	"testing"
)

type TestObject struct {
	Key   string
	Value any
}

func (to TestObject) MarshalLogObject(encoder nlog.ObjectEncoder) (err error) {
	encoder.AddString("key", to.Key)
	encoder.AddReflected("value", to.Value)
	return nil
}

type TestArray []any

func (ta TestArray) MarshalLogArray(encoder nlog.ArrayEncoder) (err error) {
	for _, v := range ta {
		err := encoder.AppendReflected(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestLog(t *testing.T) {
	nlog.Debug("I am Debug", nlog.Array("TestArray", TestArray{"apple", 42, struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}}))
	nlog.Debug("I am Debug", nlog.Object("TestObject", TestObject{Key: "hello", Value: true}))
	nlog.Debug("I am Debug", nlog.ObjectValues("TestObject", []TestObject{{Key: "hello", Value: true}, {Key: "world", Value: false}}))
	nlog.Debug("I am Debug", nlog.Objects("TestObject", []TestObject{{Key: "hello", Value: true}, {Key: "world", Value: false}}))
	nlog.Debug("I am Debug", nlog.Int("Int", 1))
	nlog.Info("I am Info", nlog.Any("Array", []int{1, 2, 3}))
	nlog.Warn("I am Warn")
	nlog.Error("I am Error")
	nlog.DPanic("I am DPanic")
	nlog.Panic("I am Panic")
	nlog.Fatal("I am Fatal")
}
