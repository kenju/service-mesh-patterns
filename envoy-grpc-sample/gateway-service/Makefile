# https://www.gnu.org/software/make/manual/make.html#index-_002eEXPORT_005fALL_005fVARIABLES
.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := gateway-service
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Run built binaries
run: build
	./gateway-service

## Run tests
test:
	gotest ./... -v

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup:
	# get development tools
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/rakyll/gotest
	# protobuf
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
	# gRPC gateway
	# @see https://github.com/grpc-ecosystem/grpc-gateway#installation:wq
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

## Compile protoc
protoc:
	# application definition
	protoc \
		-I ../protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		../protobuf-definitions/backend/services/v1/hello.proto
	# gateway definition
	protoc \
		-I ../protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		../protobuf-definitions/backend/services/v1/hello.proto

## Swagger
## @see https://github.com/grpc-ecosystem/grpc-gateway#installation
swagger:
	protoc \
		-I ../protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		../protobuf-definitions/backend/services/v1/hello.proto

## Update go modules
update:
	go get -u
	go mod tidy

## Format source codes
fmt:
	for d in $$(go list ./...); do \
		goimports -w ${GOPATH}/src/$$d; \
	done
