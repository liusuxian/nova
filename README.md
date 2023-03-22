# nova
highly concurrent game server framework(高并发游戏服务器框架)

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