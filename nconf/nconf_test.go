/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-21 22:01:24
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 11:01:35
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconf/nconf_test.go
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
	Name             string // 服务器应用名称
	Host             string // 服务器IP
	Port             uint16 // 服务器监听端口（uint16）
	MaxConn          uint32 // 允许的客户端连接最大数量（uint32）
	WorkerPoolSize   uint32 // 工作任务池最大工作Goroutine数量（uint32）
	PackageHeaderLen uint8  // 包头的长度（字节 uint8）
	MaxPacketSize    uint32 // 数据包的最大值（字节 uint32）
	MaxMsgChanLen    uint32 // SendBuffMsg发送消息的缓冲最大长度（字节 uint32）
}

// LogConfig 日志配置
type LogConfig struct {
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

	confSlice := []LogConfig{}
	if err = nconf.StructKey("logger", &confSlice); err != nil {
		t.Log("StructKey confSlice Error:", err)
		return
	}
	t.Logf("confSlice: %+v\n", confSlice)
}
