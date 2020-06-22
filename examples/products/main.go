package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/products"
	"net/http"
	"time"
)

func main() {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	connectionTimeout := 60 * time.Second
	transport := &http.Transport{
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: connectionTimeout,
	}
	httpCl := &http.Client{Transport: transport}

	sessionKey, err := auth.VerifyUser(*username, *password, *clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := api.NewClient(sessionKey, *clientCode, httpCl)
	if err != nil {
		panic(err)
	}

	prodGroups, err := GetProductGroups(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetProductGroups:\n%+v\n", prodGroups)

	prods, err := GetProductsBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetProductsBulk:\n%+v\n", prods)

	prods, err = GetProductsInParallel(apiClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetProductsInParallel:\n%+v\n", prods)
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

func GetProductsInParallel(cl *api.Client) ([]products.Product, error) {
	productsDataProvider := products.NewListingDataProvider(cl.ProductManager)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        300,
			MaxFetchersCount:          10,
		},
		productsDataProvider,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	prodsChan := lister.Get(ctx, map[string]interface{}{
		"changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(),
	})

	prods := make([]products.Product, 0)
	var err error
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range prodsChan {
			if prod.Err != nil {
				err = prod.Err
				return
			}
			prods = append(prods, prod.Payload.(products.Product))
		}
	}()

	<-doneChan
	return prods, err
}

func GetProductGroups(cl *api.Client) ([]products.ProductGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	groups, err := cl.ProductManager.GetProductGroups(ctx, map[string]string{
		"productGroupID": "2",
	})

	return groups, err
}
