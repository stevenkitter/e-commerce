#!/usr/bin/env bash
COMMON_PATH=$PWD/common/proto
protoc -I${COMMON_PATH} ${COMMON_PATH}/*.proto --go_out=plugins=grpc:${GOPATH}/src
SERVER_PATH=$PWD/services
protoc -I${PWD} ${SERVER_PATH}/account/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src
protoc -I${PWD} ${SERVER_PATH}/message/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src