package nclient

import (
	"github.com/liusuxian/nova/niface"
	"github.com/panjf2000/gnet/v2"
	"time"
)

// Client 的服务 Option
type Option func(c *Client)

// TCPSocketOpt 是TCP套接字选项的类型。用于设置 TCP 连接的 NoDelay 选项，该选项表示是否禁用 Nagle 算法。启用 NoDelay 选项可以降低延迟，但会增加网络负载
type TCPSocketOpt int

const (
	TCPNoDelay TCPSocketOpt = iota
	TCPDelay
)

// WithPacket 只要实现 Packet，接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(packet niface.IDataPack) Option {
	return func(c *Client) {
		c.SetPacket(packet)
	}
}

// WithHeartbeat 设置心跳发送间隔时间
func WithHeartbeat(heartbeat time.Duration) Option {
	return func(c *Client) {
		c.heartbeat = heartbeat
	}
}

// WithMaxHeartbeat 设置最长心跳检测间隔时间
func WithMaxHeartbeat(maxHeartbeat time.Duration) Option {
	return func(c *Client) {
		c.maxHeartbeat = maxHeartbeat
	}
}

// WithReadBufferCap 是在可读事件到来时从对端读取的最大字节数。
// 默认值为64KB，可以减小它以避免影响后续连接，也可以增加它以从套接字读取更多数据。
//
// 请注意，ReadBufferCap 将始终转换为大于或等于其实际值的最小的2的幂整数值
func WithReadBufferCap(readBufferCap int) Option {
	return func(c *Client) {
		c.options.ReadBufferCap = readBufferCap
	}
}

// WithWriteBufferCap 是静态出站缓冲区可以容纳的最大字节数，
// 如果数据超过此值，溢出将存储在弹性链表缓冲区中。默认值为64KB。
//
// 请注意，WriteBufferCap 将始终转换为大于或等于其实际值的最小的2的幂整数值
func WithWriteBufferCap(writeBufferCap int) Option {
	return func(c *Client) {
		c.options.WriteBufferCap = writeBufferCap
	}
}

// WithLockOSThread 用于确定每个 I/O 事件循环是否与一个 OS 线程关联，它在需要某些机制时非常有用，
// 如线程本地存储或调用某些需要通过 cgo 进行线程级操作的 C 库（如图形库：GLib），
// 或希望所有 I/O 事件循环实际上并行运行以提高性能
func WithLockOSThread(lockOSThread bool) Option {
	return func(c *Client) {
		c.options.LockOSThread = lockOSThread
	}
}

// WithTicker 表示是否已设置定时器
func WithTicker(ticker bool) Option {
	return func(c *Client) {
		c.options.Ticker = ticker
	}
}

// WithTCPKeepAlive 设置（SO_KEEPALIVE）套接字选项的持续时间
func WithTCPKeepAlive(tcpKeepAlive time.Duration) Option {
	return func(c *Client) {
		c.options.TCPKeepAlive = tcpKeepAlive
	}
}

// WithTCPNoDelay 控制操作系统是否应该延迟数据包传输以期望发送较少的数据包（Nagle 算法）。
//
// 默认值为 true（无延迟），意味着数据在写操作后尽快发送
func WithTCPNoDelay(tcpNoDelay TCPSocketOpt) Option {
	return func(c *Client) {
		c.options.TCPNoDelay = gnet.TCPSocketOpt(tcpNoDelay)
	}
}

// WithSocketRecvBuffer 设置套接字接收缓冲区的最大字节数
func WithSocketRecvBuffer(recvBuf int) Option {
	return func(c *Client) {
		c.options.SocketRecvBuffer = recvBuf
	}
}

// WithSocketSendBuffer 设置套接字发送缓冲区的最大字节数
func WithSocketSendBuffer(sendBuf int) Option {
	return func(c *Client) {
		c.options.SocketSendBuffer = sendBuf
	}
}
