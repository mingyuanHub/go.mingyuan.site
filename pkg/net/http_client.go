package net

import (
	"net/http"
	"net"
	"time"
)

var (
	HttpClient1000      *http.Client
	HttpClient2000      *http.Client
)



func init() {
	HttpClient1000 = createHTTPClient(1000)
	HttpClient2000 = createHTTPClient(2000)
}

func createHTTPClient(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     5 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	return client
}