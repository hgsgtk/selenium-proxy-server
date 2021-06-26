package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	// revision is assumed to be overwritten by the ldflags option.
	revision = "default"

	// fixme: set from out of container
	seleniumServerURL = "http://selenium-server:4444"
)

func main() {
	// Fixme split to another file
	hcHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"releaseId\": %s}", revision)
	}

	proxyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		defer r.Body.Close()

		reqBodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}

		// logging request body
		fmt.Fprintf(os.Stdout, "request method: '%s', body: '%s'\n", r.Method, reqBodyBytes)

		// proxy to selenium server hub endpoint
		req, err := http.NewRequestWithContext(
			r.Context(),
			r.Method,
			seleniumServerURL+"/wd/hub",
			strings.NewReader(string(reqBodyBytes)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}
		client := &http.Client{
			Timeout: 3 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}
		resBodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}
		defer resp.Body.Close()

		// logging response body
		fmt.Fprintf(os.Stdout, "response status: '%d', body: '%s'\n", resp.StatusCode, string(resBodyBytes))

		// response to client
		w.WriteHeader(resp.StatusCode)
		fmt.Fprint(w, string(resBodyBytes))
	}

	http.HandleFunc("/.healthcheck", hcHandler)
	http.HandleFunc("/proxy", proxyHandler)
	// Fixme graceful shutdown implements
	log.Fatal(http.ListenAndServe(":8080", nil))
}
