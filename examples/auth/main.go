package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/log"
	"net/http"
	"time"
)

func main() {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	log.Log = log.StdLogger{}

	connectionTimeout := 60 * time.Second
	transport := &http.Transport{
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{},
		ResponseHeaderTimeout: connectionTimeout,
	}
	httpCl := &http.Client{Transport: transport}

	clBldr := api.ClientBuilder{
		UserName:                 *username,
		Password:                 *password,
		ClientCode:               *clientCode,
		DefaultSessionLenSeconds: 10, //default length is smaller than the wait time so the session will expire before the next attempt which should repeat the auth
		HttpCli:                  httpCl,
	}

	erplyClient := clBldr.Build()

	mngr := erplyClient.SalesManager

	for i := 0; i < 10; i++ {
		filter := map[string]string{
			"active": "1",
		}

		res, err := mngr.GetVatRates(context.Background(), filter)
		common.Die(err)

		fmt.Printf("Attempt %d: GetVatRate: %s\n", i+1, common.ConvertSourceToJsonStrIfPossible(res))
		time.Sleep(time.Second * 2)
	}
}
