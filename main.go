package main

import (
	"net/http"
	"time"
)

func startServer() {
	router := CreateRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func main() {
	startServer()
}
