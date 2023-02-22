/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-18 23:25:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 21:01:34
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nnet/tcp_server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nnet

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/niface"
	"github.com/liusuxian/nova/nlog"
	"go.uber.org/zap"
	"net"
)

// TCPServer TCP服务器结构
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

var ctx = context.Background()

// Start 启动服务器
func (s *TCPServer) Start() {
	nlog.Info(ctx, "TCPServer Listener Is Starting", zap.String("Ip", s.IP), zap.Uint16("Port", s.Port))
	go func() {
		// 获取一个TCP的Addr
		var addr *net.TCPAddr
		var err error
		if addr, err = net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port)); err != nil {
			nlog.Error(ctx, "TCPServer Resolve Tcp Addr Error", zap.Error(err))
			return
		}
		// 监听服务器的地址
		var listener *net.TCPListener
		if listener, err = net.ListenTCP(s.IPVersion, addr); err != nil {
			nlog.Error(ctx, "TCPServer Listen Error", zap.String("IPVersion", s.IPVersion), zap.Error(err))
			return
		}
		nlog.Info(ctx, "TCPServer Listen Succeed", zap.String("Name", s.Name))
		// 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			var conn *net.TCPConn
			if conn, err = listener.AcceptTCP(); err != nil {
				nlog.Error(ctx, "TCPServer Accept Error", zap.Error(err))
				continue
			}
			// 已经与客户端建立连接
			go func() {
				for {
					buf := make([]byte, 512)
					var cnt int
					if cnt, err = conn.Read(buf); err != nil {
						nlog.Error(ctx, "TCPServer Recv Buf Error", zap.Error(err))
						continue
					}
					nlog.Debug(ctx, "TCPServer Recv Client", zap.ByteString("Buf", buf), zap.Int("cnt", cnt))
					// 回复
					if _, err := conn.Write(buf[:cnt]); err != nil {
						nlog.Error(ctx, "TCPServer Write Buf Error", zap.Error(err))
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

// NewTCPServer 创建一个TCP服务器
func NewTCPServer(name string, ip string, port uint16) niface.IServer {
	s := &TCPServer{
		Name:      name,
		IPVersion: "tcp4",
		IP:        ip,
		Port:      port,
	}

	return s
}
