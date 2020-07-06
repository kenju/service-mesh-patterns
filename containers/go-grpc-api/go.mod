module github.com/kenju/service-mesh-patterns/containers/go-grpc-api

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/kenju/service-mesh-patterns/benchmark-grpc/backend-service v0.0.0-20200705095907-33a80c441e77 // indirect
	github.com/kenju/service-mesh-patterns/envoy-grpc-gateway/backend-service v0.0.0-20200705095907-33a80c441e77
	github.com/sirupsen/logrus v1.6.0
	google.golang.org/grpc v1.30.0
)
