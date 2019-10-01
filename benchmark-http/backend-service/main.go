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

	log.Printf("backend-service listenining at %s...\n", addr)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(addr, nil); err != nil {
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
// application logic
//--------------------------------
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from backend-service\n")
}