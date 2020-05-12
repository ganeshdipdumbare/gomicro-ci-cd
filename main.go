package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("started server")
	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world testing CI/CD: %v", r.RemoteAddr)
}
