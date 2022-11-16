package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

type Server interface {
	Address() string
	isAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

// SimpleServer STRUCT
type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// SERVER STRUCT INITIALIZER
func newServer(addr string) *SimpleServer {
	serverUrl, err := url2.Parse(addr)
	handleError(err)
	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

// LoadBalancer STRUCT
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (s *SimpleServer) Address() string {
	return s.addr
}

func (s *SimpleServer) isAlive() bool {
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.isAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	log.Printf("Forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	servers := []Server{
		newServer("https://www.facebook.com"),
		newServer("https://www.bing.com"),
		newServer("https://www.duckduckgo.com"),
	}

	lb := newLoadBalancer("8000", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving at localhost : %v\n", lb.port)

	http.ListenAndServe(":"+lb.port, nil)
}
