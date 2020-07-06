package main

import (
	"context"
	"log"
	"time"

	backend_service "github.com/kenju/service-mesh-patterns/containers/go-grpc-api/backend/services/v1"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	serverAddr := "127.0.0.1:8080"

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	backendCheck(conn)
}

func backendCheck(conn *grpc.ClientConn) {
	client := backend_service.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &backend_service.HelloRequest{
		Message: "hello from client",
	}
	message, err := client.Hello(ctx, req)
	if err != nil {
		panic(err)
	}

	log.Printf("backend.Hello() message=%+v\n", message)
}
