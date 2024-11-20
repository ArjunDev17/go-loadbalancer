package lb

import (
	"fmt"
	"net/http"
)

// LoadBalancer struct that holds all servers and load balancing logic.
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

// Server interface represents the methods that a server should implement.
type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

// NewLoadBalancer initializes a new load balancer.
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

// GetNextAvailableServer returns the next available server using round-robin method.
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// ServeProxy forwards the request to the next available server.
func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to address: %s\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

// Port returns the port the LoadBalancer is running on.
func (lb *LoadBalancer) Port() string {
	return lb.port
}
