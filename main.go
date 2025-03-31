package main

import (
	"fmt"
	"net/url"
)

var backendServers = []*url.URL{
	{Scheme: "http", Host: "3.84.189.215"},
	{Scheme: "http", Host: "34.238.251.205"},
	{Scheme: "http", Host: "18.206.250.71"},
}

func main() {
	fmt.Println(backendServers)
}
