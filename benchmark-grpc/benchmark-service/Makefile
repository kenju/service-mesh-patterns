.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := benchmark-service
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.revision=$(REVISION)'

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go

## Update go modules
update:
	go get -u
	go mod tidy

## Compile protoc
protoc:
	protoc \
		-I ../../containers/protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		../../containers/protobuf-definitions/backend/services/v1/hello.proto

## Run server
run-server:
	ADDR=127.0.0.1:8001 go run main.go

## Run client
run-client:
	go run bin/client/client.go
