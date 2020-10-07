package api

import (
	"crypto/tls"
	"flag"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"net/http"
	"time"
)

func BuildClient() (*Client, error) {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	connectionTimeout := 60 * time.Second
	transport := &http.Transport{
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: connectionTimeout,
	}
	httpCl := &http.Client{Transport: transport}

	sessionKey, err := auth.VerifyUser(*username, *password, *clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := NewClient(sessionKey, *clientCode, httpCl)
	if err != nil {
		panic(err)
	}

	return apiClient, nil
}
