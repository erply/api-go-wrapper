package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/prices"
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

	bulkPrices, err := GetPricesBulk(apiClient)
	if err != nil {
		panic(err)
	}

	supplierPrices, err := GetSupplierPrices(apiClient, "4644")
	if err != nil {
		panic(err)
	}

	fmt.Printf("BulkPrices:\n%+v\n", bulkPrices)
	fmt.Printf("SupplierPrices:\n%+v\n", supplierPrices)

	bulkProductPrices, err := GetProductPricesBulk(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("BulkProductPrices:\n%+v\n", bulkProductPrices)

	productPrices, err := GetProductPrices(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ProductPrices:\n%+v\n", productPrices)
}

func GetPricesBulk(cl *api.Client) (prics []prices.PriceList, err error) {
	cli := cl.PricesManager

	bulkFilters := []map[string]interface{}{
		{
			"recordsOnPage": "3",
			"pageNo":        "1",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := cli.GetSupplierPriceListsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, pr := range bulkItem.PriceLists {
			prics = append(prics, pr)
		}
	}
	return
}

func GetProductPricesBulk(cl *api.Client) (prodPrices []prices.ProductPriceList, err error) {
	cli := cl.PricesManager

	bulkFilters := []map[string]interface{}{
		{
			"supplierPriceListID": "3",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := cli.GetProductPriceListsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, pr := range bulkItem.ProductPriceList {
			prodPrices = append(prodPrices, pr)
		}
	}
	return
}

func GetProductPrices(cl *api.Client) (prodPrices []prices.ProductPriceList, err error) {
	cli := cl.PricesManager

	filters := map[string]string{
			"supplierPriceListID": "3",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	prodPrices, err = cli.GetProductPriceLists(ctx, filters)

	return
}

func GetSupplierPrices(cl *api.Client, supplierID string) (prics []prices.PriceList, err error) {
	cli := cl.PricesManager

	filters := map[string]string{
		"supplierID":    supplierID,
		"recordsOnPage": "3",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return cli.GetSupplierPriceLists(ctx, filters)
}
