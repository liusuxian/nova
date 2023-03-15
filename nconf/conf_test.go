/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 11:04:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-15 16:26:12
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconf/conf_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconf_test

import (
	"github.com/liusuxian/nova/nconf"
	"testing"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Name           string // 服务器应用名称，默认"Nova"
	Network        string // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	Port           uint16 // 服务器监听端口（uint16）
	MaxHeartbeat   uint32 // 最长心跳检测间隔时间（单位: 毫秒 uint32），默认 5000
	MaxConn        uint32 // 允许的客户端连接最大数量，默认 3（uint32）
	WorkerPoolSize uint32 // 工作任务池最大工作 Goroutine 数量，默认 10（uint32）
	MaxPacketSize  uint32 // 数据包的最大值，默认 4096（单位:字节 uint32）
	PacketMethod   uint8  // 封包和拆包方式，默认 1，1: 消息ID(2字节)-消息体长度(4字节)-消息内容（单位:字节 uint8）
	Endian         uint8  // 字节存储次序，默认小端，1: 小端 2: 大端（单位:字节 uint8）
	MaxMsgChanLen  uint32 // SendBuffMsg发送消息的缓冲最大长度，默认 1024（单位:字节 uint32）
}

// LogConfig 日志配置
type LogConfig struct {
	CtxKeys []string          // 自定义 Context 上下文变量名称，自动打印 Context 的变量到日志中。默认为空
	Details []LogDetailConfig // 日志详细配置
}

// LogDetailConfig 日志详细配置
type LogDetailConfig struct {
	Level      string // 日志打印级别 debug、info、warn、error、dpanic、panic、fatal
	Format     string // 输出日志格式 logfmt、json
	Path       string // 输出日志文件路径
	Filename   string // 输出日志文件名称
	MaxSize    int    // 单个日志文件最多存储量，单位(mb)
	MaxBackups int    // 日志备份文件最多数量
	MaxAge     int    // 日志保留时间，单位:天(day)
	Compress   bool   // 是否压缩日志
	Stdout     bool   // 是否输出到控制台
}

func TestConfig(t *testing.T) {
	var err error
	serverConf := ServerConfig{}
	if err = nconf.StructKey("server", &serverConf); err != nil {
		t.Log("StructKey ServerConfig Error:", err)
		return
	}
	t.Logf("serverConf: %+v\n", serverConf)

	logConfig := LogConfig{}
	if err = nconf.StructKey("logger", &logConfig); err != nil {
		t.Log("StructKey confSlice Error:", err)
		return
	}
	t.Logf("logConfig: %+v\n", logConfig)

	var cfg *nconf.Config
	if cfg, err = nconf.NewRemoteConfig("consul", "127.0.0.1:8500", "config/test", "json"); err != nil {
		t.Log("NewRemoteConfig Error:", err)
		return
	}
	t.Logf("a: %+v\n", cfg.GetInt("a"))
	t.Logf("b: %+v\n", cfg.GetInt("b"))
	t.Logf("c: %+v\n", cfg.GetIntSlice("c"))
}
