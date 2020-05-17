package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"net/url"
)

func main() {
	const (
		username   = ""
		password   = ""
		clientCode = ""
		partnerKey = ""
	)
	httpCli := common.GetDefaultHTTPClient()
	sessionKey, err := auth.VerifyUser(username, password, clientCode, httpCli)
	if err != nil {
		panic(err)
	}

	_, err = auth.GetSessionKeyUser(sessionKey, clientCode, httpCli)

	//this function will receive the request name at every request execution inside the erply api client
	customHeadersSetter := func(requestName string) url.Values {
		params := url.Values{}
		params.Add("setContentType", "1")
		params.Add("request", requestName)
		params.Add("sessionKey", sessionKey)
		params.Add("clientCode", clientCode)

		fmt.Println(len(params))
		return params
	}

	cli, err := api.NewClient(sessionKey, clientCode, nil, customHeadersSetter)
	if err != nil {
		panic(err)
	}
	//fmt.Println(info)

	_, err = cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	//fmt.Println(endpoints)

	partnerCli, err := api.NewPartnerClient(sessionKey, clientCode, partnerKey, nil, customHeadersSetter)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
