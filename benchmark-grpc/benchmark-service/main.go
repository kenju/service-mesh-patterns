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

// See ghz documentation for each metrics explanation
// @see https://ghz.sh/docs/output
var (
	promGaugeCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_count", prometheusPrefix),
		Help: "The total number of completed requests including successful and failed requests.",
	})
	promGaugeTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_total", prometheusPrefix),
		Help: "The total time spent running the test within ghz from start to finish. This is a single measurement from start of the test run to the completion of the final request of the test run.",
	})
	promGaugeAvg = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_average", prometheusPrefix),
		Help: "The mathematical average computed by taking the sum of the individual response times of all requests and dividing it by the total number of requests.",
	})
	promGaugeFastest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_fastest", prometheusPrefix),
		Help: "The measurement of the fastest request.",
	})
	promGaugeSlowest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_slowest", prometheusPrefix),
		Help: "The measurement of the slowest request.",
	})
	promGaugeRPS = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_rps", prometheusPrefix),
		Help: "Theoretical computed RPS computed by taking the total number of requests (successful and failed) and dividing it by the total duration of the test. That is: count / total.",
	})
	// Use Heatmap panel of Grafana
	// @doc https://grafana.com/docs/features/panels/heatmap/
	promGaugeResponseHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: fmt.Sprintf("%s_response_histogram", prometheusPrefix),
		Help: "The histogram of response time.",
		Buckets: prometheus.LinearBuckets(20, 10, 10),
	})
)

func main() {
	port := getEnv("ADDR", defaultAddr)

	log.Printf("starting benchmarker at %s...\n", port)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/start", startHandler)

	if err := http.ListenAndServe(port, mux); err != nil {
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
func startHandler(w http.ResponseWriter, r *http.Request) {

	serverAddr := getEnv("LOAD_TEST_TARGET_ADDR", defaultLoadTestTargetAddr)

	checkRequest(serverAddr)

	startLoadTest(serverAddr, w)
}

// checkRequest is an assertion to check a connection to the target gRPC application.
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
	// exec load testing
	report, err := runner.Run(
		"backend.services.v1.HelloService.Hello",
		serverAddr,
		runner.WithInsecure(true),
		runner.WithDataFromJSON("{}"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// write report JSON output to buffer
	// TODO: this logic can be written as `prometheus` format in ghz library
	promGaugeCount.Set(float64(report.Count))
	promGaugeTotal.Set(durationToMs(report.Total))
	promGaugeAvg.Set(durationToMs(report.Average))
	promGaugeFastest.Set(durationToMs(report.Fastest))
	promGaugeSlowest.Set(durationToMs(report.Slowest))
	promGaugeRPS.Set(report.Rps)
	// @doc https://prometheus.io/docs/concepts/metric_types/#histogram
	for _, bucket := range report.Histogram {
		promGaugeResponseHistogram.Observe(bucket.Frequency)
	}

	// write response to writer
	printer := printer.ReportPrinter{
		Out: writer,
		Report: report,
	}
	printer.Print("json")
}

func durationToMs(d time.Duration) float64 {
	return float64(d / time.Millisecond)
}