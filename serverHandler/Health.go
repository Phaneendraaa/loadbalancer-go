package serverHandler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	
)

func Healthcheck(u *url.URL, index int) {
	healthURL := fmt.Sprintf("%s/", u.String())
	resp, _ := http.Get(healthURL)
	if resp.StatusCode != http.StatusOK {
		BackendHealth[index] = false
		fmt.Println("Server", BackendServers[index].String(), "is down")
	} else {
		BackendHealth[index] = true
		fmt.Printf("Server %s is healthy\n", u.String())
	}
}

func HealthCheckLoop() {
	for {
		// Check health of each backend every 10 seconds
		for i, backend := range BackendServers {
			Healthcheck(backend, i)
		}

		time.Sleep(20 * time.Second)
	}
}
