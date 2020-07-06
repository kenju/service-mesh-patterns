.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := backend-service
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.revision=$(REVISION)'

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup: update-submodule
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go

## Update go modules
update:
	go get -u
	go mod tidy

## Update git submodules
update-submodule:
	git submodule update --remote --merge

## Compile protoc
protoc: update-submodule
	protoc \
		-I ../protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		../protobuf-definitions/backend/services/v1/hello.proto
	protoc \
		-I ../grpc-proto \
		--go_out=plugins=grpc:. \
		../grpc-proto/grpc/health/v1/health.proto

## Run server
run-server:
	PORT=127.0.0.1:8080 go run main.go

## Run client
run-client:
	go run bin/client.go
