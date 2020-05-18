package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/breathbath/api-go-wrapper/internal/common"
	"github.com/breathbath/api-go-wrapper/pkg/api"
	"github.com/breathbath/api-go-wrapper/pkg/api/auth"
	"net/url"
)

func main() {
	var (
		username   = flag.String("u", "", "username")
		password   = flag.String("p", "", "password")
		clientCode = flag.String("cc", "", "client code")
		partnerKey = flag.String("pk", "", "partner key")
	)

	flag.Parse()
	httpCli := common.GetDefaultHTTPClient()
	sessionKey, err := auth.VerifyUser(*username, *password, *clientCode, httpCli)
	if err != nil {
		panic(err)
	}

	_, err = auth.GetSessionKeyUser(sessionKey, *clientCode, httpCli)

	//this function will receive the request name at every request execution inside the erply api client
	customHeadersSetter := func(requestName string) url.Values {
		params := url.Values{}
		params.Add("setContentType", "1")
		params.Add("request", requestName)
		params.Add("sessionKey", sessionKey)
		params.Add("clientCode", *clientCode)

		fmt.Println(len(params))
		return params
	}

	cli, err := api.NewClientWithCustomHeaders(nil, customHeadersSetter)
	if err != nil {
		panic(err)
	}
	//fmt.Println(info)

	_, err = cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	//fmt.Println(endpoints)

	customHeadersSetterForPartnerClient := func(requestName string) url.Values {
		params := url.Values{}
		params.Add("setContentType", "1")
		params.Add("request", requestName)
		params.Add("sessionKey", sessionKey)
		params.Add("clientCode", *clientCode)
		params.Add("partnerKey", *partnerKey)
		fmt.Println(len(params))
		return params
	}
	partnerCli, err := api.NewClientWithCustomHeaders(nil, customHeadersSetterForPartnerClient)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.AuthProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
