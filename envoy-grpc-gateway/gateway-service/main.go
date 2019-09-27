package main

import (
	"context"
	backend_services_v1 "github.com/kenju/service-mesh-patterns/envoy-grpc-gateway/gateway-service/backend/services/v1"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"strconv"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	defaultAddr = ":3000"
	defaultProxyAddr = "front-proxy:8000"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	proxyAddr := getEnv("PROXY_ADDR", defaultProxyAddr)

	gateway, err := newGateway(proxyAddr, ctx)
	if err != nil {
		panic(err)
	}

	endpoint := getEnv("ADDR", defaultAddr)
	if err := http.ListenAndServe(endpoint, gateway); err != nil {
		panic(err)
	}
}

//--------------------------------
// utility
//--------------------------------

func getEnv(key, defaultVal string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultVal
	}
	return v
}

func getEnvAsInt(key string, defaultVal int) (int, error) {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultVal, nil
	}

	i, err := strconv.Atoi(v)
	return i, err
}

//--------------------------------
// grpc-gateway
//--------------------------------
func newGateway(endpoint string, ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}

	// register handler
	err = backend_services_v1.RegisterHelloServiceHandler(ctx, mux, conn)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
