package main

import (
	"context"
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	backend_service "github.com/kenju/service-mesh-patterns/benchmark-grpc/benchmark-service/backend/services/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	defaultAddr = ":8001"
	defaultLoadTestTargetAddr = "127.0.0.1:8080"
	prometheusPrefix = "benchmark_server"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_counter", prometheusPrefix),
		Help: "sample metrics for counter",
	})
)

func main() {
	port := getEnv("ADDR", defaultAddr)

	log.Printf("starting benchmarker at %s...\n", port)

	recordMetrics()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/start", handler)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}

}

func recordMetrics() {
	go func() {
		counter.Inc()
		time.Sleep(5 * time.Second)
	}()
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

	// assertion to check connection to the target gRPC application
	checkRequest(serverAddr)

	startLoadTest(serverAddr, w)
}

func checkRequest(serverAddr string) {
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

func startLoadTest(serverAddr string, writer io.Writer) {
	report, err := runner.Run(
		"backend.services.v1.HelloService.Hello",
		serverAddr,
		runner.WithInsecure(true),
		runner.WithDataFromJSON("{}"),
	)

	if err != nil {
		log.Fatal(err)
	}

	printer := printer.ReportPrinter{
		Out: writer,
		Report: report,
	}

	printer.Print("json")
}