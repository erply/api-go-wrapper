package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"time"
)

func main() {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	suppliers, err := GetSupplierBulk(*username, *password, *clientCode)
	if err != nil {
		panic(err)
	}

	fmt.Println(suppliers)
}

func GetSupplierBulk(username, password, clientCode string) ([]customers.Supplier, error){
	httpCli := common.GetDefaultHTTPClient()
	sessionKey, err := auth.VerifyUser(username, password, clientCode, httpCli)
	if err != nil {
		return []customers.Supplier{}, err
	}

	commonClient := common.NewClient(sessionKey, clientCode, "", nil, nil)
	supplierCli := customers.NewClient(commonClient)

	ctx := context.WithValue(context.Background(), "bulk", []map[string]string{
		{
			"recordsOnPage": "2",
			"pageNo":"1",
		},
		{
			"recordsOnPage": "2",
			"pageNo":"2",
		},
	})
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return supplierCli.GetSuppliers(ctx, map[string]string{})
}
