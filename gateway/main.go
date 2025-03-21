package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sandi2156/oms-microservices-using-go/common"
)

var (
	addr = common.GetEnv("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()

	handler := NewHandler()
	handler.RegisterRoutes(mux)

	log.Printf("Started the server on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Error starting the server")
	}
}