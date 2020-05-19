package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
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

	suppliers, err := GetSupplierBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Println(suppliers)

	err = SaveSupplierBulk(apiClient)
	if err != nil {
		panic(err)
	}
}

func GetSupplierBulk(cl *api.Client) (suppliers []customers.Supplier, err error){
	supplierCli := cl.CustomerManager

	bulkFilters := []map[string]string{
		{
			"recordsOnPage": "2",
			"pageNo":"1",
		},
		{
			"recordsOnPage": "2",
			"pageNo":"2",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bulkResp, err := supplierCli.GetSuppliersBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, supplier := range bulkItem.Suppliers {
			suppliers = append(suppliers, supplier)
		}
	}

	return
}

func SaveSupplierBulk(cl *api.Client) (err error){
	supplierCli := cl.CustomerManager

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sup := customers.Supplier{
		CompanyName: "Some Test Company",
		VatNumber:  "Some VAT",
		Code:  "34132434134",
	}
	bulkResponse, err := supplierCli.SaveSupplierBulk(ctx, []customers.Supplier{sup}, map[string]string{})
	if err != nil {
		return
	}

	fmt.Printf("%+v", bulkResponse)

	return
}
