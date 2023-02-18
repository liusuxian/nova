/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-19 01:58:09
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/tcp_server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nnet

import (
	"fmt"
	"github.com/liusuxian/nova/niface"
	"net"
)

// TCPServer TCP服务器类
type TCPServer struct {
	Name        string                        // 服务器的名称
	IPVersion   string                        // tcp4 or other
	IP          string                        // 服务绑定的IP地址
	Port        uint16                        // 服务绑定的端口
	msgHandler  niface.IMsgHandle             // 当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	ConnMgr     niface.IConnManager           // 当前Server的连接管理器
	OnConnStart func(conn niface.IConnection) // 当前Server的连接创建时Hook函数
	OnConnStop  func(conn niface.IConnection) // 当前Server的连接断开时的Hook函数
	exitChan    chan struct{}
	packet      niface.IDataPack
}

// Start 启动服务器
func (s *TCPServer) Start() {
	fmt.Printf("TCPServer Listener at IP: %s Port: %d is starting\n", s.IP, s.Port)
	go func() {
		// 获取一个TCP的Addr
		var addr *net.TCPAddr
		var err error
		if addr, err = net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port)); err != nil {
			_ = fmt.Errorf("resolve tcp addr error: %v\n", err)
			return
		}
		// 监听服务器的地址
		var listener *net.TCPListener
		if listener, err = net.ListenTCP(s.IPVersion, addr); err != nil {
			_ = fmt.Errorf("listen %s error: %v\n", s.IPVersion, err)
			return
		}
		fmt.Printf("nova server listening %s succ\n", s.Name)
		// 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			var conn *net.TCPConn
			if conn, err = listener.AcceptTCP(); err != nil {
				_ = fmt.Errorf("accept error: %v\n", err)
				continue
			}
			// 已经与客户端建立连接
			go func() {
				for {
					buf := make([]byte, 512)
					var cnt int
					if cnt, err = conn.Read(buf); err != nil {
						_ = fmt.Errorf("recv buf error: %v\n", err)
						continue
					}
					fmt.Printf("recv client buf: %s cnt: %d\n", buf, cnt)
					// 回复
					if _, err := conn.Write(buf[:cnt]); err != nil {
						_ = fmt.Errorf("write buf error: %v\n", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop 停止服务器
func (s *TCPServer) Stop() {
	// TODO 将一些服务器的资源、状态或者一些已经开辟的连接信息进行停止或回收
}

// Server 开启业务服务
func (s *TCPServer) Server() {
	s.Start()
	// TODO 做一些启动服务之后的额外业务
	// 阻塞
	select {}
}

// NewTCPServer 创建一个TCP服务器句柄
func NewTCPServer(name string, ip string, port uint16) niface.IServer {
	s := &TCPServer{
		Name:      name,
		IPVersion: "tcp4",
		IP:        ip,
		Port:      port,
	}

	return s
}
