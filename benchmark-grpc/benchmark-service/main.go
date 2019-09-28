package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	backend_service "github.com/kenju/service-mesh-patterns/benchmark-grpc/benchmark-service/backend/services/v1"
)

const (
	defaultAddr = ":8001"
	defaultLoadTestTargetAddr = "127.0.0.1:8080"
)


func main() {
	port := getEnv("ADDR", defaultAddr)

	log.Printf("starting benchmarker at %s...\n", port)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(port, nil); err != nil {
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
// business logic
//--------------------------------
func handler(w http.ResponseWriter, r *http.Request) {

	serverAddr := getEnv("LOAD_TEST_TARGET_ADDR", defaultLoadTestTargetAddr)

	loadTest(serverAddr)

	fmt.Fprintf(w, fmt.Sprintf("sent gRPC load testing request to %s\n", serverAddr))
}

func loadTest(serverAddr string) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := backend_service.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &backend_service.HelloRequest{}
	message, err := client.Hello(ctx, req)
	if err != nil {
		panic(err)
	}

	log.Printf("backend.Hello() message=%+v\n", message)
}
