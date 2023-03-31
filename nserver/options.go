/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-15 13:26:56
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-31 20:35:45
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
func WithPacket(packet niface.IDataPack) (opt Option) {
	return func(s *Server) {
		s.SetPacket(packet)
	}
}

// WithMulticore 表示引擎是否将有效地创建为多核，如果是这样的话，
// 那么您必须注意在所有事件回调之间同步内存，否则，
// 它将使用单线程运行引擎。引擎中的线程数将自动
// 被分配为当前进程可用的逻辑 CPU 数
func WithMulticore(multicore bool) (opt Option) {
	return func(s *Server) {
		s.options.Multicore = multicore
	}
}

// WithNumEventLoop 设置要启动的事件循环goroutine的数量
//
//	注意：设置 NumEventLoop 将覆盖 Multicore
func WithNumEventLoop(numEventLoop int) (opt Option) {
	return func(s *Server) {
		s.options.NumEventLoop = numEventLoop
	}
}

// WithLoadBalancing 表示分配新连接时使用的负载均衡算法
func WithLoadBalancing(lb LoadBalancing) (opt Option) {
	return func(s *Server) {
		s.options.LB = gnet.LoadBalancing(lb)
	}
}

// WithReuseAddr 表示是否设置 SO_REUSEADDR 套接字选项
func WithReuseAddr(reuseAddr bool) (opt Option) {
	return func(s *Server) {
		s.options.ReuseAddr = reuseAddr
	}
}

// WithReusePort 表示是否设置 SO_REUSEPORT 套接字选项
func WithReusePort(reusePort bool) (opt Option) {
	return func(s *Server) {
		s.options.ReusePort = reusePort
	}
}

// WithMulticastInterfaceIndex 是多播 UDP 地址将绑定到的接口名称的索引
func WithMulticastInterfaceIndex(idx int) (opt Option) {
	return func(s *Server) {
		s.options.MulticastInterfaceIndex = idx
	}
}

// WithReadBufferCap 是在可读事件到来时从对端读取的最大字节数。
// 默认值为64KB，可以减小它以避免影响后续连接，也可以增加它以从套接字读取更多数据。
//
// 请注意，ReadBufferCap 将始终转换为大于或等于其实际值的最小的2的幂整数值
func WithReadBufferCap(readBufferCap int) (opt Option) {
	return func(s *Server) {
		s.options.ReadBufferCap = readBufferCap
	}
}

// WithWriteBufferCap 是静态出站缓冲区可以容纳的最大字节数，
// 如果数据超过此值，溢出将存储在弹性链表缓冲区中。默认值为64KB。
//
// 请注意，WriteBufferCap 将始终转换为大于或等于其实际值的最小的2的幂整数值
func WithWriteBufferCap(writeBufferCap int) (opt Option) {
	return func(s *Server) {
		s.options.WriteBufferCap = writeBufferCap
	}
}

// WithLockOSThread 用于确定每个 I/O 事件循环是否与一个 OS 线程关联，它在需要某些机制时非常有用，
// 如线程本地存储或调用某些需要通过 cgo 进行线程级操作的 C 库（如图形库：GLib），
// 或希望所有 I/O 事件循环实际上并行运行以提高性能
func WithLockOSThread(lockOSThread bool) (opt Option) {
	return func(s *Server) {
		s.options.LockOSThread = lockOSThread
	}
}

// WithTicker 表示是否已设置定时器
func WithTicker(ticker bool) (opt Option) {
	return func(s *Server) {
		s.options.Ticker = ticker
	}
}

// WithTCPKeepAlive 设置（SO_KEEPALIVE）套接字选项的持续时间
func WithTCPKeepAlive(tcpKeepAlive time.Duration) (opt Option) {
	return func(s *Server) {
		s.options.TCPKeepAlive = tcpKeepAlive
	}
}

// WithTCPNoDelay 控制操作系统是否应该延迟数据包传输以期望发送较少的数据包（Nagle 算法）。
//
// 默认值为 true（无延迟），意味着数据在写操作后尽快发送
func WithTCPNoDelay(tcpNoDelay TCPSocketOpt) (opt Option) {
	return func(s *Server) {
		s.options.TCPNoDelay = gnet.TCPSocketOpt(tcpNoDelay)
	}
}

// WithSocketRecvBuffer 设置套接字接收缓冲区的最大字节数
func WithSocketRecvBuffer(recvBuf int) (opt Option) {
	return func(s *Server) {
		s.options.SocketRecvBuffer = recvBuf
	}
}

// WithSocketSendBuffer 设置套接字发送缓冲区的最大字节数
func WithSocketSendBuffer(sendBuf int) (opt Option) {
	return func(s *Server) {
		s.options.SocketSendBuffer = sendBuf
	}
}
