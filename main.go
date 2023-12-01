package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/", hello)

	// Liveness Probe
	go func() {
		http.HandleFunc("/alive", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "alive\n")
		})
		http.ListenAndServe(":8091", nil)
	}()

	// Readiness Probe
	go func() {
		time.Sleep(10 * time.Second) // Delay the readiness probe for 10 seconds
		http.HandleFunc("/ready", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "ready\n")
		})
		http.ListenAndServe(":8092", nil)
	}()
	fmt.Println("Server started")
	http.ListenAndServe(":8090", nil)
}
