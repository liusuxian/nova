/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-13 11:04:59
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-06-22 14:12:37
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconf_test

import (
	"github.com/liusuxian/nova/nconf"
	"testing"
	"time"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Name                   string        // 服务器应用名称，默认"Nova"
	Network                string        // 服务器网络协议 tcp、tcp4、tcp6、udp、udp4、udp6、unix
	Port                   int           // 服务器监听端口
	HeartBeat              time.Duration // 心跳发送间隔时间（一定要小于 maxHeartBeat 配置），默认 10秒
	MaxHeartBeat           time.Duration // 最长心跳检测间隔时间（一定要大于 heartBeat 配置），默认 15秒
	MaxConn                int           // 允许的客户端连接最大数量，默认 3
	WorkerPoolSize         int           // 工作任务池最大工作 Goroutine 数量，默认 10
	WorkerPoolSizeOverflow int           // 当处理任务超过工作任务池的容量时，增加的 Goroutine 数量，默认 5
	MaxPacketSize          int           // 数据包的最大值（单位:字节），默认 4096
	PacketMethod           int           // 封包和拆包方式，1: 消息ID(2字节)-消息体长度(4字节)-消息内容，默认 1
	Endian                 int           // 字节存储次序，1: 小端 2: 大端，默认 1
	SlowThreshold          time.Duration // 处理请求或执行操作时的慢速阈值
}

// LogConfig 日志配置
type LogConfig struct {
	Path    string            // 输出日志文件路径
	Details []LogDetailConfig // 日志详细配置
}

// LogDetailConfig 日志详细配置
type LogDetailConfig struct {
	Type       int    // 日志类型 0:打印所有级别 1:打印 DEBUG、INFO 级别 2:打印 WARN、ERROR、DPANIC、PANIC、FATAL 级别，默认0
	Level      int    // 日志打印级别 0:DEBUG 1:INFO 2:WARN 3:ERROR 4:DPANIC、5:PANIC、6:FATAL，默认0
	Format     int    // 输出日志格式 0:logfmt 1:json，默认1
	Filename   string // 输出日志文件名称
	MaxSize    int    // 单个日志文件最多存储量（单位:MB）
	MaxBackups int    // 日志备份文件最多数量
	MaxAge     int    // 日志保留时间（单位:天）
	Compress   bool   // 是否压缩日志
	Stdout     bool   // 是否输出到控制台
}

// TestCfg 测试配置
type TestCfg struct {
	ID        int64  `json:"id" dc:"id"`
	MoneyReal int64  `json:"money_real" dc:"money_real"`
	TotalTime uint32 `json:"total_time" dc:"total_time"`
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

	var localCfg *nconf.Config
	if localCfg, err = nconf.NewConfig("config/test.json"); err != nil {
		t.Log("NewConfig Error:", err)
		return
	}
	testCfg1 := TestCfg{}
	if err = localCfg.StructKey("test", &testCfg1); err != nil {
		t.Log("StructKey TestCfg1 Error:", err)
		return
	}
	t.Logf("testCfg1: %+v\n", testCfg1)
	testCfg2 := TestCfg{}
	if err = localCfg.StructKey("test", &testCfg2, func(dc *nconf.DecoderConfig) {
		dc.TagName = "json"
	}); err != nil {
		t.Log("StructKey TestCfg2 Error:", err)
		return
	}
	t.Logf("testCfg2: %+v\n", testCfg2)

	var cfg *nconf.Config
	if cfg, err = nconf.NewRemoteConfig("consul", "127.0.0.1:8500", "config/test", "json"); err != nil {
		t.Log("NewRemoteConfig Error:", err)
		return
	}
	t.Logf("a: %+v\n", cfg.GetInt("a"))
	t.Logf("b: %+v\n", cfg.GetInt("b"))
	t.Logf("c: %+v\n", cfg.GetIntSlice("c"))
}
