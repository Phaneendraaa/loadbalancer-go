package serverHandler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	
)

var currBackend int
var mu sync.Mutex
var BackendServers = []*url.URL{}
var BackendHealth = []bool{}
var backendHealthMu sync.Mutex 

func IntializeServers(Server *url.URL) {
	BackendServers = append(BackendServers, Server)
	backendHealthMu.Lock()
	BackendHealth = append(BackendHealth, true)
	backendHealthMu.Unlock()
}

func GetNextBackend() *url.URL {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < len(BackendServers); i++ {
		backend := BackendServers[currBackend]
		if BackendHealth[currBackend] {
			currBackend = (currBackend + 1) % len(BackendServers)
			return backend
		}

		currBackend = (currBackend + 1) % len(BackendServers)
	}

	return nil
}
func ProxyHandler(w http.ResponseWriter, r *http.Request) {

	backendserver := GetNextBackend()
	if backendserver == nil {
		http.Error(w, "No healthy backend servers available", http.StatusServiceUnavailable)
		return
	}
	fmt.Println(" ---------  Sending traffic to server --------->", backendserver)
	proxy := httputil.NewSingleHostReverseProxy(backendserver)
	proxy.ServeHTTP(w, r)
}

