/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 21:24:06
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-08 15:18:14
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/npack/defaultpack_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package npack_test

import (
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/npack"
	"io"
	"net"
	"testing"
	"time"
)

// 只测试封包拆包功能
func TestDataPack(t *testing.T) {
	// 创建socket TCP Server
	var listener net.Listener
	var err error
	if listener, err = net.Listen("tcp", "127.0.0.1:8888"); err != nil {
		t.Log("Server Listen Error:", err)
		return
	}
	// 创建服务器gotoutine，负责从客户端goroutine读取粘包的数据，然后进行解析
	go func() {
		for {
			var conn net.Conn
			if conn, err = listener.Accept(); err != nil {
				t.Log("Server Accept Error:", err)
			}
			// 处理客户端请求
			go func(conn net.Conn) {
				// 创建封包拆包对象
				dp := npack.Factory().NewPack()
				for {
					// 先读出流中的head部分
					headData := make([]byte, dp.GetHeadLen())
					// ReadFull会把msg填充满为止
					if _, err = io.ReadFull(conn, headData); err != nil {
						t.Log("Read Head Error")
						return
					}
					// 将headData字节流拆包到msg中
					var msgHead niface.IMessage
					if msgHead, err = dp.Unpack(headData); err != nil {
						t.Log("Server Unpack Error: ", err)
						return
					}
					if msgHead.GetDataLen() > 0 {
						// msg是有data数据的，需要再次读取data数据
						msg := msgHead.(*npack.Message)
						msg.Data = make([]byte, msg.GetDataLen())
						// 根据dataLen从io中读取字节流
						if _, err = io.ReadFull(conn, msg.Data); err != nil {
							t.Log("Server Unpack Data Error: ", err)
							return
						}
						t.Log("==> Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	// 客户端goroutine，负责模拟粘包的数据，然后进行发送
	go func() {
		var conn net.Conn
		if conn, err = net.Dial("tcp", "127.0.0.1:8888"); err != nil {
			t.Log("Client Dial Error: ", err)
			return
		}
		// 创建一个封包对象
		dp := npack.Factory().NewPack()
		// 封装一个msg1包
		msg1 := npack.NewMsgPackage(1, []byte("hello"))
		var sendData1 []byte
		if sendData1, err = dp.Pack(msg1); err != nil {
			t.Log("Client Pack Msg1 Error: ", err)
			return
		}
		// 封装一个msg2包
		msg2 := npack.NewMsgPackage(2, []byte("world!!"))
		var sendData2 []byte
		if sendData2, err = dp.Pack(msg2); err != nil {
			t.Log("Client Pack Msg2 Error: ", err)
			return
		}
		// 将sendData1和sendData2拼接一起，组成粘包
		sendData1 = append(sendData1, sendData2...)
		// 向服务器端写数据
		conn.Write(sendData1)
	}()
	// 客户端阻塞
	select {
	case <-time.After(time.Second):
		return
	}
}
