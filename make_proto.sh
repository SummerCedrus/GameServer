#!/bin/bash
protoc protocol/protocol.proto --go_out=./
protoc protocol/rpc.proto --go_out=plugins=grpc:./
go run gen_proto.go
go fmt protocol/cmd.go
go fmt protocol/const.go
go fmt protocol/error.go
go fmt protocol/decode.go
