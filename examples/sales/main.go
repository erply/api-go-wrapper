package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/sales"
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

	saleDocuments, err := GetSalesDocumentsBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetSalesDocumentsBulk: %+v\n", saleDocuments)

	saleDocumentsInParallel, err := GetSalesDocumentsInParallel(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetSalesDocumentsInParallel: %+v\n", saleDocumentsInParallel)
}

func GetSalesDocumentsBulk(cl *api.Client) (docs []sales.SaleDocument, err error) {
	salesCLI := cl.SalesManager

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

	bulkResp, err := salesCLI.GetSalesDocumentsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, doc := range bulkItem.SaleDocuments {
			docs = append(docs, doc)
		}
	}

	return
}

func GetSalesDocumentsInParallel(cl *api.Client) ([]sales.SaleDocument, error) {
	saleDocLister := sales.NewSaleDocumentsListingDataProvider(cl.SalesManager)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        300,
			MaxFetchersCount:          10,
		},
		saleDocLister,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	docsChan := lister.Get(ctx, map[string]interface{}{
		"changedSince": 1578441600,
	})

	salesDocuments := make([]sales.SaleDocument, 0)
	var err error
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range docsChan {
			if prod.Err != nil {
				err = prod.Err
				return
			}
			salesDocuments = append(salesDocuments, prod.Payload.(sales.SaleDocument))
		}
	}()

	<-doneChan
	return salesDocuments, err
}
