package serverHandler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var currBackend int
var mu sync.Mutex
var backendServers = []*url.URL{}
var backendHealth = []bool{}
var backendHealthMu sync.Mutex 

func IntializeServers(Server *url.URL) {
	backendServers = append(backendServers, Server)
	backendHealthMu.Lock()
	backendHealth = append(backendHealth, true)
	backendHealthMu.Unlock()
}

func GetNextBackend() *url.URL {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < len(backendServers); i++ {
		backend := backendServers[currBackend]
		if backendHealth[currBackend] {

			currBackend = (currBackend + 1) % len(backendServers)
			return backend
		}

		currBackend = (currBackend + 1) % len(backendServers)
	}

	return nil
}
func ProxyHandler(w http.ResponseWriter, r *http.Request) {

	backendserver := GetNextBackend()
	if backendserver == nil {
		http.Error(w, "No healthy backend servers available", http.StatusServiceUnavailable)
		return
	}
	fmt.Println("Sending traffic to server", backendserver)
	proxy := httputil.NewSingleHostReverseProxy(backendserver)
	proxy.ServeHTTP(w, r)
}
func Healthcheck(u *url.URL, index int) {
	healthURL := fmt.Sprintf("%s/", u.String())
	resp, _ := http.Get(healthURL)
	if resp.StatusCode != http.StatusOK {
		backendHealth[index] = false
		fmt.Println("Server", backendServers[index].String(), "is down")
	} else {
		backendHealth[index] = true
		fmt.Printf("Server %s is healthy\n", u.String())
	}
}

func HealthCheckLoop() {
	for {
		// Check health of each backend every 10 seconds
		for i, backend := range backendServers {
			Healthcheck(backend, i)
		}

		time.Sleep(20 * time.Second)
	}
}
