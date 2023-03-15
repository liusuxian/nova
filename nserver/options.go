/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-15 13:26:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 14:39:09
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nserver/options.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nserver

import (
	"github.com/liusuxian/nova/niface"
	"github.com/panjf2000/gnet/v2"
	"time"
)

// Server 的服务 Option
type Option func(s *Server)

// LoadBalancing 表示负载均衡算法的类型
type LoadBalancing int

const (
	// RoundRobin 算法通过轮询事件循环列表将下一个被接受的连接分配到事件循环中
	RoundRobin LoadBalancing = iota
	// LeastConnections 算法会将下一个被接受的连接分配到当前活跃连接数最少的事件循环中
	LeastConnections
	// SourceAddrHash 算法通过哈希远程地址将下一个被接受的连接分配给事件循环
	SourceAddrHash
)

// TCPSocketOpt 是TCP套接字选项的类型。用于设置 TCP 连接的 NoDelay 选项，该选项表示是否禁用 Nagle 算法。启用 NoDelay 选项可以降低延迟，但会增加网络负载
type TCPSocketOpt int

const (
	TCPNoDelay TCPSocketOpt = iota
	TCPDelay
)

// WithPacket 只要实现 Packet，接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(pack niface.IDataPack) Option {
	return func(s *Server) {
		s.SetPacket(pack)
	}
}

// 设置是否启用多核支持
func WithMulticore(multicore bool) Option {
	return func(s *Server) {
		s.options.Multicore = multicore
	}
}

// 设置事件循环数量。每个事件循环都是一个 goroutine，默认情况下，会创建与 CPU 核心数相同的事件循环
func WithNumEventLoop(numEventLoop int) Option {
	return func(s *Server) {
		s.options.NumEventLoop = numEventLoop
	}
}

// 设置负载均衡算法
func WithLoadBalancing(lb LoadBalancing) Option {
	return func(s *Server) {
		s.options.LB = gnet.LoadBalancing(lb)
	}
}

// 设置是否启用地址复用
func WithReuseAddr(reuseAddr bool) Option {
	return func(s *Server) {
		s.options.ReuseAddr = reuseAddr
	}
}

// 设置是否启用端口复用
func WithReusePort(reusePort bool) Option {
	return func(s *Server) {
		s.options.ReusePort = reusePort
	}
}

// 设置多播（UDP multicast）套接字绑定的网络接口
func WithMulticastInterfaceIndex(idx int) Option {
	return func(s *Server) {
		s.options.MulticastInterfaceIndex = idx
	}
}

// 设置读取数据时的缓冲区大小
func WithReadBufferCap(readBufferCap int) Option {
	return func(s *Server) {
		s.options.ReadBufferCap = readBufferCap
	}
}

// 设置写入数据时的缓冲区大小
func WithWriteBufferCap(writeBufferCap int) Option {
	return func(s *Server) {
		s.options.WriteBufferCap = writeBufferCap
	}
}

// 设置 I/O 事件循环的 LockOSThread 模式。每个事件循环都是一个 goroutine，LockOSThread 模式可以锁定 goroutine 所在的线程，防止其被调度到其他线程上执行
func WithLockOSThread(lockOSThread bool) Option {
	return func(s *Server) {
		s.options.LockOSThread = lockOSThread
	}
}

// 设置一个定时器，它将以指定的时间间隔触发回调函数。定时器通常用于周期性地执行一些任务，例如定期刷新缓存、定期发送心跳包等
func WithTicker(ticker bool) Option {
	return func(s *Server) {
		s.options.Ticker = ticker
	}
}

// 设置 TCP KeepAlive 选项。用于检测空闲的、无数据传输的 TCP 连接，以便及时释放资源。
func WithTCPKeepAlive(tcpKeepAlive time.Duration) Option {
	return func(s *Server) {
		s.options.TCPKeepAlive = tcpKeepAlive
	}
}

// 设置是否禁用 Nagle 算法
func WithTCPNoDelay(tcpNoDelay TCPSocketOpt) Option {
	return func(s *Server) {
		s.options.TCPNoDelay = gnet.TCPSocketOpt(tcpNoDelay)
	}
}

// 设置 socket 最大接收缓冲区大小（单位为字节）
func WithSocketRecvBuffer(recvBuf int) Option {
	return func(s *Server) {
		s.options.SocketRecvBuffer = recvBuf
	}
}

// 设置 socket 最大发送缓冲区大小（单位为字节）
func WithSocketSendBuffer(sendBuf int) Option {
	return func(s *Server) {
		s.options.SocketSendBuffer = sendBuf
	}
}
