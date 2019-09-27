package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	backend_services_v1 "github.com/kenju/envoy-grpc-sample/backend-service/backend/services/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	defaultAddr = ":8080"
)

func init() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func main() {
	port := getEnv("ADDR", defaultAddr)
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	severOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	server := grpc.NewServer(severOpts...)

	grpc_health_v1.RegisterHealthServer(server, newHealthServer())
	backend_services_v1.RegisterHelloServiceServer(server, newBackendServer())

	server.Serve(listenPort)
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
// health check server
//--------------------------------

type healthServer struct {
}

func newHealthServer() *healthServer {
	return &healthServer{}
}

// Check is an implementation of grpc_health_v1.HealthServer interface
// https://godoc.org/google.golang.org/grpc/health/grpc_health_v1#HealthServer
func (s *healthServer) Check(
	ctx context.Context,
	in *grpc_health_v1.HealthCheckRequest,
) (*grpc_health_v1.HealthCheckResponse, error) {

	log.WithFields(log.Fields{
		"time": time.Now().Format(time.RFC3339),
	}).Debug("Check()")

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch is an implementation of grpc_health_v1.HealthServer interface
// https://godoc.org/google.golang.org/grpc/health/grpc_health_v1#HealthServer
func (s *healthServer) Watch(
	in *grpc_health_v1.HealthCheckRequest,
	w grpc_health_v1.Health_WatchServer,
) error {
	return status.Errorf(codes.Unimplemented, "(c *healthServer) Watch() is not implemented")
}

//--------------------------------
// backend application server
//--------------------------------
type backendServer struct {
}

func newBackendServer() *backendServer {
	return &backendServer{}
}

func (bs *backendServer) Hello(
	ctx context.Context,
	req *backend_services_v1.HelloRequest,
) (*backend_services_v1.HelloResponse, error) {
	log.WithFields(log.Fields{
		"request": req,
	}).Info("Hello()")

	return &backend_services_v1.HelloResponse{
		Result: &backend_services_v1.HelloResponse_Success_{
			Success: &backend_services_v1.HelloResponse_Success{
				StatusCode: fmt.Sprintf("%d", codes.OK),
			},
		},
	}, nil
}
