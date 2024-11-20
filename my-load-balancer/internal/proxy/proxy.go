package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// SimpleServer represents a single backend server.
type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// NewSimpleServer initializes a new proxy server for the given address.
func NewSimpleServer(addr string) *SimpleServer {
	serveUrl, err := url.Parse(addr)
	handleErr(err)

	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serveUrl),
	}
}

// HandleErr logs the error and exits the program if the error is not nil.
func handleErr(err error) {
	if err != nil {
		log.Fatalf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}

// Address returns the address of the server.
func (s *SimpleServer) Address() string {
	return s.addr
}

// IsAlive checks if the server is alive (dummy check in this example).
func (s *SimpleServer) IsAlive() bool {
	return true
}

// Serve forwards the HTTP request to the actual server using reverse proxy.
func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}
