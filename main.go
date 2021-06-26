package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	// revision is assumed to be overwritten by the ldflags option.
	revision = "default"
)

func main() {
	// Fixme split to another file
	hcHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"releaseId\": %s}", revision)
	}

	proxyHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello proxy server")
	}

	http.HandleFunc("/.healthcheck", hcHandler)
	http.HandleFunc("/proxy", proxyHandler)
	// Fixme graceful shutdown implements
	log.Fatal(http.ListenAndServe(":8080", nil))
}
