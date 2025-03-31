package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Phaneendraaa/loadbalancer-go/serverHandler"
)

func main() {
	fmt.Println("Started load balancer")

	
	var serverIPs = []string{
		"3.84.189.215",
		"34.238.251.205",
		"18.206.250.71",
	}

	
	for i := 0; i < len(serverIPs); i++ {
		// Initializing servers with their urls
		serverHandler.IntializeServers(&url.URL{
			Scheme: "http",
			Host:   serverIPs[i],
		})
	}

	
	http.HandleFunc("/", serverHandler.ProxyHandler)

	
	go serverHandler.HealthCheckLoop()

	fmt.Println("Server running on 8082")
	
	http.ListenAndServe(":8082", nil)


}
