package main

import (
	"fmt"
	"net/http"

	"github.com/Phaneendraaa/loadbalancer-go/serverHandler"
)
func main() {
	fmt.Println("Started load balancer")
	http.HandleFunc("/", serverHandler.ProxyHandler)
	
	http.ListenAndServe(":8080", nil)
}
