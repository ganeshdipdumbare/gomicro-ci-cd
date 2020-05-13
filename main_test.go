package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(helloWorld))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("expected: %v, got: %v for error", nil, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected: %v, got: %v for error", http.StatusOK, resp.StatusCode)
	}
}

func TestStartServer(t *testing.T) {
	go startServer()
	time.Sleep(time.Second * 2)

	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatalf("expected: %v, got: %v for error", nil, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected: %v, got: %v for error", http.StatusOK, resp.StatusCode)
	}

}
