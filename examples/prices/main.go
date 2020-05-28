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

	res, err := AddProductToSupplierPriceList(apiClient, "65538", "1", "100.23")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ChangeProductToSupplierPriceList:\n%+v\n", res)

	bulkRes, err := ChangeProductToSupplierPriceListBulk(apiClient, []string{"65539", "65540"}, []string{"1", "1"}, []string{"10.22", "111.00"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ChangeProductToSupplierPriceListBulk:\n%+v\n", bulkRes)
}

func ChangeProductToSupplierPriceListBulk(cl *api.Client, productIds, priceIds, prices []string) (prices.ChangeProductToSupplierPriceListResponseBulk, error) {
	cli := cl.PricesManager

	bulkItems := []map[string]interface{}{}
	for i, prodID := range productIds {
		bulkItems = append(bulkItems, map[string]interface{}{
			"productID":           prodID,
			"supplierPriceListID": priceIds[i],
			"price":               prices[i],
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.ChangeProductToSupplierPriceListBulk(ctx, bulkItems, map[string]string{})
	return resp, err
}

func AddProductToSupplierPriceList(cl *api.Client, productId, priceId, price string) (*prices.ChangeProductToSupplierPriceListResult, error) {
	cli := cl.PricesManager

	filter := map[string]string{
		"productID":           productId,
		"supplierPriceListID": priceId,
		"price":               price,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.AddProductToSupplierPriceList(ctx, filter)
	return resp, err
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
