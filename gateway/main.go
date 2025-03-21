package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

func main() {
	mux := http.NewServeMux()

	handler := NewHandler()
	handler.RegisterRoutes(mux)

	fmt.Printf("Started the server")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Error starting the server")
	}
}