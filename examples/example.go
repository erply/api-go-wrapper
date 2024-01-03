package main

import (
	"context"
	"fmt"
	"time"

	"github.com/erply/api-go-wrapper/pkg/api"
)

func main() {

	//put your credentials here
	const (
		username   = ""
		password   = ""
		clientCode = ""
	)

	cli, err := api.NewClientFromCredentials(username, password, clientCode, nil)
	if err != nil {
		panic(err)
	}

	//indicate to the client that the request should add the data payload in the
	//request body instead of using the query parameters. Using the request body eliminates the query size
	//limitations imposed by the maximum URL length
	cli.SendParametersInRequestBody()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	salesDocs, err := cli.SalesManager.GetSalesDocuments(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(salesDocs)

	const (
		partnerKey = ""
	)

	//partner client example
	partnerCli, err := api.NewPartnerClientFromCredentials(username, password, clientCode, partnerKey, nil)
	if err != nil {
		panic(err)
	}
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
