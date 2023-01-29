package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func getPort() (int, error) {
	DEFAULT_PORT := 80
	portString, exists := os.LookupEnv("PORT")
	if !exists {
		return DEFAULT_PORT, nil
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		return 0, errors.New("Environment variable 'PORT' is not an integer.")
	}
	return port, nil
}

func startServer() int {
	port, err := getPort()
	if err != nil {
		return 1
	}
	server := http.NewServeMux()
	server.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {})
	server.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {})
	fmt.Printf("Starting server at port %d\n", port)

	c := make(chan error)
	go func() {
		c <- http.ListenAndServe(fmt.Sprintf(":%d", port), server)
	}()
	err = <- c
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(startServer())
}
