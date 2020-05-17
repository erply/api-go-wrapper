package common

import (
	"net"
	"net/http"
	"time"
)

const (
	//MaxIdleConns for Erply API
	MaxIdleConns = 25

	//MaxConnsPerHost for Erply API
	MaxConnsPerHost = 25
)

func GetDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,

			ExpectContinueTimeout: 4 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,

			MaxIdleConns:    MaxIdleConns,
			MaxConnsPerHost: MaxConnsPerHost,
		},
		Timeout: 5 * time.Second,
	}
}
