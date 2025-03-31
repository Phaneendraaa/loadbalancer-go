package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var currBackend int
var mu sync.Mutex
var backendServers = []*url.URL{
	{Scheme: "http", Host: "3.84.189.215"},
	{Scheme: "http", Host: "34.238.251.205"},
	{Scheme: "http", Host: "18.206.250.71"},
}

func getNextBackend() *url.URL {
	mu.Lock()
	defer mu.Unlock()
	backend := backendServers[currBackend]
	currBackend = (currBackend + 1) % len(backendServers)
	return backend
}
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	backendserver := getNextBackend()
	fmt.Println(" Sending traffic to  server", backendserver)
	proxy := httputil.NewSingleHostReverseProxy(backendserver)
	proxy.ServeHTTP(w, r)
}
func main() {
	fmt.Println("Started load balancer")
	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(":8080", nil)
}
