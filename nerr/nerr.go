/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-05-22 20:40:14
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-28 19:29:20
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nerr

import "github.com/pkg/errors"

// ErrIncompletePacket 不完整的包
var ErrIncompletePacket = errors.New("incomplete packet")

// ErrLargeDataReceived 接收到的消息数据太大
var ErrTooLargeMsgReceived = errors.New("too large msg data received")

// ErrConnectionClosed 发送数据时连接关闭
var ErrConnectionClosed = errors.New("connection closed when send data")
