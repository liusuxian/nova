/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 12:05:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 01:54:54
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/examples/tcp_demo/server/server.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/nlog"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Name           string // 服务器应用名称
	Host           string // 服务器IP
	Port           uint16 // 服务器监听端口
	MaxConn        int    // 允许的客户端连接最大数量
	WorkerPoolSize int    // 工作任务池最大工作Goroutine数量
}

// TestConfig 测试配置
type TestConfig struct {
	A int
	B string
	C []int
}

func main() {
	var err error
	serverConf := ServerConfig{}
	err = nconf.Sub("server").GetStruct("base", &serverConf)
	nlog.Debugf("serverConf: %v %+v\n", err, serverConf)
	// 监视配置文件的变化
	nconf.WatchConfig()
	// 设置当配置文件更改时调用的事件处理程序
	nconf.OnConfigChange(func(in fsnotify.Event) {
		// 配置发生变化了，执行响应的操作
		nlog.Debug("Default Config File Changed: ", in.Name)
		err = nconf.Sub("server").GetStruct("base", &serverConf)
		nlog.Debugf("serverConf: %v %+v\n", err, serverConf)
	})

	var cfg *nconf.Config
	if cfg, err = nconf.New("config/test.yaml"); err != nil {
		nlog.Fatal("New Config Error: ", err)
	}
	testConf := []TestConfig{}
	err = cfg.GetStructs("test", &testConf)
	nlog.Debugf("testConf: %v %+v\n", err, testConf)
	// 监视配置文件的变化
	cfg.WatchConfig()
	// 设置当配置文件更改时调用的事件处理程序
	cfg.OnConfigChange(func(in fsnotify.Event) {
		// 配置发生变化了，执行响应的操作
		nlog.Debug("Config File Changed: ", in.Name)
		err = cfg.GetStruct("test", &testConf)
		nlog.Debugf("testConf: %v %+v\n", err, testConf)
	})
	nlog.Error("错误")
	select {}
}
