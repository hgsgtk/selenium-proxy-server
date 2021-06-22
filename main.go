package main

import (
	"log"
	"net/http"
)

var (
	// revision is assumed to be overwritten by the ldflags option.
	revision = "default"
)

func main() {
	// Fixme temporary server implementation
	log.Fatal(http.ListenAndServe(":8080", nil))
}
