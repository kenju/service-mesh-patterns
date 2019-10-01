package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"github.com/rakyll/hey/requester"
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
		Help: "sample count",
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

	checkConnection(serverAddr)

	startLoadTest(serverAddr, w)
}

func checkConnection(serverAddr string) {
	fmt.Printf("checking connection to %s\n", serverAddr)

	resp, err := http.Get(serverAddr)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func startLoadTest(serverAddr string, writer io.Writer) {
	fmt.Printf("starting load testing\n")

	method := "GET"
	url := serverAddr
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	worker := &requester.Work{
		Request:            req,
		RequestBody:        nil,
		N:                  200,
		C:                  50,
		QPS:                0,
		Timeout:            20,
		DisableCompression: false,
		DisableKeepAlives:  false,
		DisableRedirects:   false,
		H2:                 false,
	}
	worker.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		worker.Stop()
	}()
	worker.Run()

	fmt.Fprintf(writer, "finished load testing\n")
}
