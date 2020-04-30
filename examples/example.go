package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/common"
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

	info, err := auth.GetSessionKeyUser(sessionKey, clientCode, httpCli)
	cli, err := api.NewClient(sessionKey, clientCode, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	endpoints, err := cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(endpoints)

	partnerCli, err := api.NewPartnerClient(sessionKey, clientCode, partnerKey, nil)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
