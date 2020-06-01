package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/products"
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

	prods, err := GetProductsBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", prods)
}

func GetProductsBulk(cl *api.Client) (prods []products.Product, err error) {
	prodCli := cl.ProductManager

	bulkFilters := []map[string]interface{}{
		{
			"code": "266844",
		},
		{
			"code": "437423",
		},
		{
			"code": "87001",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	bulkResp, err := prodCli.GetProductsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, prod := range bulkItem.Products {
			prods = append(prods, prod)
		}
	}

	return
}
