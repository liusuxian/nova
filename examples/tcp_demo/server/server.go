/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 18:18:18
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/nlog"
	"github.com/liusuxian/nova/nutils/nctx"
	"go.uber.org/zap"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Name                 string // 服务器应用名称，默认"Nova"
	Host                 string // 服务器IP
	Port                 uint16 // 服务器监听端口（uint16）
	MaxConn              uint32 // 允许的客户端连接最大数量，默认3（uint32）
	WorkerPoolSize       uint32 // 工作任务池最大工作Goroutine数量，默认10（uint32）
	PackageHeadIDLen     uint8  // 包头中消息ID长度，默认4（单位:字节 uint8）
	PackageHeadDataLen   uint8  // 包头中消息体长度，默认4（单位:字节 uint8）
	MaxPacketSize        uint32 // 数据包的最大值，默认4096（单位:字节 uint32）
	PacketMethod         uint8  // 封包和拆包方式，默认1，1:消息ID-消息体长度-消息内容（单位:字节 uint8）
	PacketProtocolFormat uint8  // 报文协议格式，默认1，1:默认（单位:字节 uint8）
	Endian               uint8  // 字节存储次序，默认小端，1:小端 2:大端（单位:字节 uint8）
	MaxMsgChanLen        uint32 // SendBuffMsg发送消息的缓冲最大长度，默认1024（单位:字节 uint32）
}

// TestConfig 测试配置
type TestConfig struct {
	A int
	B string
	C []int
}

// Context 上下文结构
type Context struct {
	User ContextUser    // 上下文用户信息
	Data map[string]any // 自定义KV变量，业务模块根据需要设置，不固定
}

// ContextUser 上下文中的用户信息
type ContextUser struct {
	Id     int64  // 用户ID
	Appid  string // 小程序ID
	Openid string // openid
}

func main() {
	ctx := nctx.SetCtxGlobalVal(context.Background(), Context{
		User: ContextUser{
			Id:     1,
			Appid:  "111",
			Openid: "222",
		},
		Data: map[string]any{"traceId": "333", "reqId": "444"},
	})
	nlog.Debug(ctx, "Log Level", zap.String("level", nlog.Level().String()))
	var err error
	serverConf := ServerConfig{}
	err = nconf.StructKey("server", &serverConf)
	nlog.Debug(ctx, "serverConf value: ", zap.Reflect("serverConf", serverConf))
	// 监视配置文件的变化
	nconf.WatchConfig()
	// 设置当配置文件更改时调用的事件处理程序
	nconf.OnConfigChange(func(in fsnotify.Event) {
		// 配置发生变化了，执行响应的操作
		nlog.Debug(ctx, "Default Config File Changed", zap.String("name", in.Name))
		err = nconf.StructKey("server.base", &serverConf)
		nlog.Debug(ctx, "serverConf value: ", zap.Reflect("serverConf", serverConf))
	})

	var cfg *nconf.Config
	if cfg, err = nconf.NewConfig("config/test.yaml"); err != nil {
		nlog.Fatal(ctx, "New Config Error: ", zap.Error(err))
	}
	testConf := []TestConfig{}
	err = cfg.StructKey("test", &testConf)
	nlog.Debug(ctx, "testConf value: ", zap.Reflect("testConf", testConf))
	// 监视配置文件的变化
	cfg.WatchConfig()
	// 设置当配置文件更改时调用的事件处理程序
	cfg.OnConfigChange(func(in fsnotify.Event) {
		// 配置发生变化了，执行响应的操作
		nlog.Debug(ctx, "Config File Changed", zap.String("name", in.Name))
		err = cfg.StructKey("test", &testConf)
		nlog.Debug(ctx, "testConf value: ", zap.Reflect("testConf", testConf))
	})
	nlog.Error(ctx, "错误")
	select {}
}
