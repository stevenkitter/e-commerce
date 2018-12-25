#!/usr/bin/env bash
SERVER_PATH=$PWD/services
protoc -I${SERVER_PATH} ${SERVER_PATH}/*/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src