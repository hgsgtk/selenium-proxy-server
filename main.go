package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
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
		fmt.Fprintf(w, "{\"status\": \"0\", \"releaseId\": %s}", revision)
	}

	// Names it hub to avoid confusion with Http proxy,
	// which Selenium provides as an option.
	hubHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		defer r.Body.Close()

		reqBodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"status\": \"27\", \"error\": %s}", err)
			return
		}

		// logging request body
		// formatting JSON to parse it from another program system
		type requestLog struct {
			Method string `json:"method`
			URL    string `json:"url`
			Body   string `json:"body"`
		}
		rql := requestLog{
			Method: r.Method,
			URL:    r.URL.Path,
			Body:   string(reqBodyBytes),
		}
		if err := json.NewEncoder(os.Stdout).Encode(rql); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"status\": \"27\",\"error\": %s}", err)
			return
		}

		// determine url with command
		vars := mux.Vars(r)
		command := vars["command"]
		sessionID := vars["sessionID"]
		subCommand := vars["subCommand"]

		// Todo refactoring the url logic to be testable
		url := seleniumServerURL + "/wd/hub"
		if command != "" {
			url = url + "/" + command
		}
		if sessionID != "" {
			url = url + "/" + sessionID
			if subCommand != "" {
				url = url + "/" + subCommand
			}
		}
		// proxy to selenium server hub endpoint
		req, err := http.NewRequestWithContext(
			r.Context(),
			r.Method,
			url,
			strings.NewReader(string(reqBodyBytes)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"status\": \"27\", \"error\": %s}", err)
			return
		}
		client := &http.Client{
			// to aviod client timeout
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"status\": \"27\", \"error\": %s}", err)
			return
		}
		resBodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"status\": \"27\", \"error\": %s}", err)
			return
		}
		defer resp.Body.Close()

		// logging response body
		// formatting JSON to parse it from another program system
		type responseLog struct {
			Status int    `json:"status`
			Body   string `json:"body`
		}
		rsl := responseLog{
			Status: resp.StatusCode,
			Body:   string(resBodyBytes),
		}
		if err := json.NewEncoder(os.Stdout).Encode(rsl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}

		// response to client
		w.WriteHeader(resp.StatusCode)
		fmt.Fprint(w, string(resBodyBytes))
	}

	notFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		defer r.Body.Close()

		// logging unknown request
		reqBodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": %s}", err)
			return
		}

		fmt.Fprintf(
			os.Stdout,
			"unknown request url path: '%s', method: '%s', body: '%s'\n",
			r.URL.Path,
			r.Method,
			string(reqBodyBytes),
		)

		// Fixme change the response format to one that can be interpreted by selenium
		// ex. KeyError: 'value'
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\": \"unknown url path '%s'\"}", r.URL.Path)
	}

	// use gorilla/mux to parse the URI Path variable to determine the Selenium server command.
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/.healthcheck", hcHandler)
	// Todo understand URI variation
	// ex. /wd/hub/session
	r.HandleFunc("/wd/hub/{command}", hubHandler)
	// ex. /wd/hub/session/71c7df563ed4fa1b57bc8d29c7b30edb/url
	r.HandleFunc("/wd/hub/{command}/{sessionID}", hubHandler)
	r.HandleFunc("/wd/hub/{command}/{sessionID}/{subCommand}", hubHandler)

	// Fixme graceful shutdown implements
	log.Fatal(http.ListenAndServe(":8080", r))
}
