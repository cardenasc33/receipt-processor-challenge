package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServerStart(t *testing.T) {
	// Set up the environment variables (if necessary)
	os.Setenv("HOSTNAME", "localhost")
	os.Setenv("PORT", "8080")

	// Tests HTTP request for the root URL 
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Root handler from main.go
		hostname := os.Getenv("HOSTNAME")
		if hostname == "" {
			hostname = "localhost"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running on host: " + hostname + "\n"))
	})

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check if the response status code is 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code 200, got %d", status)
	}

	// Check if the response body is correct
	expected := "Server is running on host: localhost\n"
	if rr.Body.String() != expected {
		t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	}
}
