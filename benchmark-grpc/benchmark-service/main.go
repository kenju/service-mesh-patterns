package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultAddr = ":8001"
)


func main() {
	port := getEnv("ADDR", defaultAddr)

	fmt.Printf("starting benchmarker at %s...\n", port)
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