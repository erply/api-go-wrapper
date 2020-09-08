package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"net/http"
	"time"
)

func main() {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	sessionKey, err := auth.VerifyUser(*username, *password, *clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := api.NewClient(sessionKey, *clientCode, nil)
	if err != nil {
		panic(err)
	}

	addresses, err := GetAddressesBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", addresses)

	err = SaveAddressesBulk(apiClient)
	if err != nil {
		panic(err)
	}
}

func GetAddressesBulk(cl *api.Client) (addresses []common.Address, err error) {
	addressCli := cl.AddressProvider

	bulkFilters := []map[string]interface{}{
		{
			"recordsOnPage": 2,
			"pageNo":        1,
		},
		{
			"recordsOnPage": 2,
			"pageNo":        2,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := addressCli.GetAddressesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, supplier := range bulkItem.Addresses {
			addresses = append(addresses, supplier)
		}
	}

	return
}

func SaveAddressesBulk(cl *api.Client) (err error) {
	addressProvider := cl.AddressProvider

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	dataToUpdate := []map[string]interface{}{
		{
			"street":     "Elmstr",
			"city":       "Rome",
			"postalCode": "123456",
			"country":    "Italy",
			"ownerID":    12355,
			"typeID":     1,
		},
	}
	bulkResponse, err := addressProvider.SaveAddressesBulk(ctx, dataToUpdate, map[string]string{})
	if err != nil {
		return
	}

	fmt.Printf("%+v", bulkResponse)

	return
}
