# https://www.gnu.org/software/make/manual/make.html#index-_002eEXPORT_005fALL_005fVARIABLES
.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := backend-service
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Run built binaries
run: build
	./backend-service

## Run tests
test:
	gotest ./... -v

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup: update-submodule
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

## Format source codes
fmt:
	for d in $$(go list ./...); do \
		goimports -w ${GOPATH}/src/$$d; \
	done

## Run server
run-server:
	PORT=127.0.0.1:8080 go run main.go

## Run client
run-client:
	go run bin/client/client.go

## Run docker container
run-docker:
	docker build .
