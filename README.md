# nova
highly concurrent game server framework(高并发游戏服务器框架)

```
├── README.md
├── examples
│   ├── proto_tcp_demo
│   │   ├── client
│   │   │   ├── client.go
│   │   │   ├── heartbeat
│   │   │   │   └── heartbeat.go
│   │   │   └── serveroverload
│   │   │       └── serveroverload.go
│   │   └── server
│   │       ├── config
│   │       │   └── config.yaml
│   │       ├── heartbeat
│   │       │   └── heartbeat.go
│   │       ├── proto
│   │       │   ├── code.proto
│   │       │   ├── msg.proto
│   │       │   ├── msg_id.proto
│   │       │   └── pb
│   │       │       ├── code.pb.go
│   │       │       ├── msg.pb.go
│   │       │       └── msg_id.pb.go
│   │       ├── proto_build.sh
│   │       ├── server.go
│   │       └── serveroverload
│   │           └── serveroverload.go
│   └── simple_tcp_demo
│       ├── client
│       │   └── client.go
│       └── server
│           ├── config
│           │   └── config.yaml
│           └── server.go
├── go.mod
├── go.sum
├── nclient
│   ├── client.go
│   └── options.go
├── nconf
│   ├── conf.go
│   └── conf_test.go
├── nconn
│   ├── connection.go
│   └── connmanager.go
├── nheartbeat
│   └── heartbeat.go
├── niface
│   ├── client.go
│   ├── connection.go
│   ├── connmanager.go
│   ├── datapack.go
│   ├── heartbeat.go
│   ├── message.go
│   ├── msghandler.go
│   ├── notify.go
│   ├── request.go
│   ├── router.go
│   ├── server.go
│   └── serveroverload.go
├── nlog
│   ├── field.go
│   ├── log.go
│   └── log_test.go
├── nmsghandler
│   └── msghandler.go
├── npack
│   ├── defaultpack.go
│   ├── message.go
│   └── pack.go
├── nrequest
│   └── request.go
├── nrouter
│   └── router.go
├── nserver
│   ├── options.go
│   └── server.go
├── nserveroverload
│   └── serveroverload.go
└── utils
    ├── nfile
    │   ├── file.go
    │   └── file_test.go
    ├── nrandom
    │   ├── random.go
    │   └── random_test.go
    ├── nslice
    │   ├── slice.go
    │   └── slice_test.go
    └── nstr
        ├── str.go
        └── str_test.go

```

## protoc 安装
- 进入 [protobuf release](https://github.com/protocolbuffers/protobuf/releases) 页面，选择适合自己操作系统的压缩包文件
- 解压 protoc-x.x.x-osx-aarch_64.zip 并进入 protoc-x.x.x-osx-aarch_64
  ```
  $ cd protoc-x.x.x-osx-aarch_64/bin
  ```
- 将解压后的 bin 目录中的 protoc 二进制文件复制到 $GOPATH/bin 目录里面，并配置 PATH 环境变量，确保 protoc 可以正常执行
  ```
  $ cp protoc $GOPATH/bin
  ```
  :::tip $GOPATH 为你本机的实际文件夹地址 :::
- 验证安装结果
  ```
  $ protoc --version
  libprotoc x.x.x
  ```
## protoc-gen-go/protoc-gen-go-grpc 安装
- 下载安装 protoc-gen-go
  ```
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```