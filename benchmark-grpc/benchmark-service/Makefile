# https://www.gnu.org/software/make/manual/make.html#index-_002eEXPORT_005fALL_005fVARIABLES
.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := benchmark-service
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Run built binaries
run: build
	./benchmark-service

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
	# setup protoc
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go

## Update go modules
update:
	go get -u
	go mod tidy

## Compile protoc
protoc:
	protoc \
		-I ../protobuf-definitions \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		../protobuf-definitions/backend/services/v1/hello.proto

## Format source codes
fmt:
	for d in $$(go list ./...); do \
		goimports -w ${GOPATH}/src/$$d; \
	done

## Run server
run-server:
	ADDR=127.0.0.1:8001 go run main.go

## Run client
run-client:
	go run bin/client/client.go

## Run docker container
run-docker:
	docker build .
