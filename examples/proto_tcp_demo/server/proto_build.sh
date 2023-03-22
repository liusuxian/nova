#!/bin/bash
protoc --go_out=./proto/pb --go-grpc_out=./proto/pb --proto_path=./proto ./proto/*.proto