package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
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

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Header=%s\n", r.Header)
	log.Printf("Host=%s\n", r.Host)
	fmt.Fprintf(w, "ACK from backend\n")
}
