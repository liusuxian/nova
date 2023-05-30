/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-27 01:28:36
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-31 01:37:57
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nnotify_test

import (
	"fmt"
	"github.com/liusuxian/nova/nnotify"
	"testing"
)

type TestOfflineMsg struct {
}

func (tom *TestOfflineMsg) Save() (err error) {
	fmt.Println("TestOfflineMsg Save Func !!!")
	return nil
}

func TestNotify_Notify(t *testing.T) {
	n := nnotify.NewNotify()
	n.Notify(0, 0, func() (buf []byte, err error) { return nil, nil }, nil)
	n.Notify(1, 1, func() (buf []byte, err error) { return nil, nil }, &TestOfflineMsg{})
}
