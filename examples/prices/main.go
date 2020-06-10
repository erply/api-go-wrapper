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

	productsInPrice, err := GetProductsInPriceList(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetProductsInPriceList:\n%+v\n", productsInPrice)

	bulkPrices, err := GetProductsInPriceListBulk(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetProductsInPriceListBulk:\n%+v\n", bulkPrices)

	supplierBulkPrices, err := GetSupplierPriceListsBulk(apiClient)
	if err != nil {
		panic(err)
	}

	supplierPrices, err := GetSupplierPriceLists(apiClient, "4644")
	if err != nil {
		panic(err)
	}

	fmt.Printf("BulkPrices:\n%+v\n", supplierBulkPrices)
	fmt.Printf("SupplierPrices:\n%+v\n", supplierPrices)

	bulkProductPrices, err := GetProductsInSupplierPriceListBulk(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("BulkProductPrices:\n%+v\n", bulkProductPrices)

	productPrices, err := GetProductsInSupplierPriceList(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ProductPrices:\n%+v\n", productPrices)

	res, err := AddProductToSupplierPriceList(apiClient, "65661", "1", "100.23")
	if err != nil {
		panic(err)
	}
	fmt.Printf("AddProductToSupplierPriceList:\n%+v\n", res)

	bulkRes, err := ChangeProductToSupplierPriceListBulk(apiClient, []string{"65659", "65660"}, []string{"1", "1"}, []string{"10.22", "111.00"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ChangeProductToSupplierPriceListBulk:\n%+v\n", bulkRes)

	bulkResDel, err := DeleteProductsFromSupplierPriceListBulk(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeleteProductsFromSupplierPriceListBulk:\n%+v\n", bulkResDel)

	saveSupPriceResp, err := SaveSupplierPriceList(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SaveSupplierPriceList:\n%+v\n", saveSupPriceResp)

	saveSupPriceRespBulk, err := SaveSupplierPriceListBulk(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SaveSupplierPriceListBulk:\n%+v\n", saveSupPriceRespBulk)
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

func GetSupplierPriceListsBulk(cl *api.Client) (prics []prices.PriceList, err error) {
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

func GetProductsInSupplierPriceListBulk(cl *api.Client) (prodPrices []prices.ProductsInSupplierPriceList, err error) {
	cli := cl.PricesManager

	bulkFilters := []map[string]interface{}{
		{
			"supplierPriceListID": "1",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := cli.GetProductsInSupplierPriceListBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, pr := range bulkItem.ProductsInSupplierPriceList {
			prodPrices = append(prodPrices, pr)
		}
	}
	return
}

func GetProductsInSupplierPriceList(cl *api.Client) (prodPrices []prices.ProductsInSupplierPriceList, err error) {
	cli := cl.PricesManager

	filters := map[string]string{
		"supplierPriceListID": "1",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	prodPrices, err = cli.GetProductsInSupplierPriceList(ctx, filters)

	return
}

func GetProductsInPriceList(cl *api.Client) (prodPrices []prices.ProductsInPriceList, err error) {
	cli := cl.PricesManager

	filters := map[string]string{
		"priceListID":   "1",
		"pageNo":        "0",
		"recordsOnPage": "10",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	prodPrices, err = cli.GetProductsInPriceList(ctx, filters)

	return
}

func GetProductsInPriceListBulk(cl *api.Client) (prodPricesBulk prices.GetProductsInPriceListResponseBulk, err error) {
	cli := cl.PricesManager

	filters := []map[string]interface{}{
		{
			"priceListID":   "1",
			"pageNo":        "0",
			"recordsOnPage": "10",
		},
		{
			"priceListID":   "1",
			"pageNo":        "0",
			"recordsOnPage": "10",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	prodPricesBulk, err = cli.GetProductsInPriceListBulk(ctx, filters, map[string]string{})

	return
}

func GetSupplierPriceLists(cl *api.Client, supplierID string) (prics []prices.PriceList, err error) {
	cli := cl.PricesManager

	filters := map[string]string{
		"supplierID":    supplierID,
		"recordsOnPage": "3",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return cli.GetSupplierPriceLists(ctx, filters)
}

func DeleteProductsFromSupplierPriceListBulk(cl *api.Client) (prices.DeleteProductsFromSupplierPriceListResponseBulk, error) {
	cli := cl.PricesManager

	bulkFilters := []map[string]interface{}{
		{
			"supplierPriceListID":         "1",
			"supplierPriceListProductIDs": "8",
		},
		{
			"supplierPriceListID":         "3",
			"supplierPriceListProductIDs": "2,3",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return cli.DeleteProductsFromSupplierPriceListBulk(ctx, bulkFilters, map[string]string{})
}

func SaveSupplierPriceList(cl *api.Client) (*prices.SaveSupplierPriceListResult, error) {
	cli := cl.PricesManager

	filter := map[string]string{
		"name":       "some price 1",
		"supplierID": "4866",
		"active":     "0",
		"productID0": "34213",
		"price0":     "100",
		"amount0":    "10",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.SaveSupplierPriceList(ctx, filter)
	return resp, err
}

func SaveSupplierPriceListBulk(cl *api.Client) (prices.SaveSupplierPriceListResponseBulk, error) {
	cli := cl.PricesManager

	bulkItems := []map[string]interface{}{
		{
			"supplierPriceListID": 5711,
			"name":                "Some other name",
			"supplierID":          "4866",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.SaveSupplierPriceListBulk(ctx, bulkItems, map[string]string{})
	return resp, err
}
