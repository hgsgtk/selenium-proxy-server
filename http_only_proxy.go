// https://gist.github.com/yowu/f7dc34bd4736a65ff28d
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type proxy struct{}

// how to use it! curl -x http://127.0.0.1:8080 http://example.com/
// if https site, got the following error.
// curl: (56) Received HTTP code 400 from proxy after CONNECT
// This server is not supported https protocol.
func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, " ", r.Method, " ", r.URL)

	if r.URL.Scheme != "http" && r.URL.Scheme != "https" {
		msg := "unsupported protocol scheme " + r.URL.Scheme
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(msg)
		return
	}

	client := &http.Client{}

	// http: Request.RequestURI can't be set in client requests.
	// http://golang.org/src/pkg/net/http/client.go
	r.RequestURI = ""

	delHopHeaders(r.Header)

	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		appendHostToXForwardHeader(r.Header, clientIP)
	}

	resp, err := client.Do(r)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
	}
	defer resp.Body.Close()

	log.Println(r.RemoteAddr, " ", resp.Status)

	delHopHeaders(resp.Header)

	copyHeader(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// Hob-by-hop headers.
// https://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ",") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

func main() {
	var addr = flag.String("addr", "127.0.0.1:8080", "The addr of the application.")
	flag.Parse()

	handler := &proxy{}

	log.Println("Starting proxy server on", *addr)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
