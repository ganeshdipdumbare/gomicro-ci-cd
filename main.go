package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/hello", helloWorld)

	server := &http.Server{Addr: ":8080", Handler: nil}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("error occurred while starting the server: ", err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("error occurred while shutting down the server: ", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from: %v", r.RemoteAddr)
}
