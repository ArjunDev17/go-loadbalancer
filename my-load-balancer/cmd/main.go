package main

import (
	"fmt"
	"log"
	"my-load-balancer/internal/lb"
	"my-load-balancer/internal/proxy"
	"net/http"
)

func main() {
	// Initialize the proxy servers
	servers := []lb.Server{
		proxy.NewSimpleServer("https://www.facebook.com/"),
		proxy.NewSimpleServer("https://www.razorpay.com"),
		proxy.NewSimpleServer("http://www.codeneeti.com"),
	}

	// Initialize the load balancer
	lb := lb.NewLoadBalancer("8000", servers)

	// Define the handler for redirecting requests
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// Setup the HTTP server and listen for requests
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Serving request at 'localhost:%s'\n", lb.Port())

	log.Fatal(http.ListenAndServe(":"+lb.Port(), nil))
}
