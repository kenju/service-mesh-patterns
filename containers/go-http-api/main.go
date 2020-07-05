package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultAddr = ":8080"
)

func main() {
	addr := getEnv("ADDR", defaultAddr)

	log.Printf("go-http-api listenining at %s...\n", addr)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("header: %+v\n", r.Header)
	res := fmt.Sprintf("ACK from backend(Host=%s)\n", r.Host)
	fmt.Fprintf(w, res)
}
