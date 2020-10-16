package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/sales"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	common.Die(err)

	saleDocuments, err := GetSalesDocumentsBulk(apiClient)
	common.Die(err)
	fmt.Printf("GetSalesDocumentsBulk: %+v\n", saleDocuments)

	saleDocumentsInParallel, err := GetSalesDocumentsInParallel(apiClient)
	common.Die(err)

	fmt.Printf("GetSalesDocumentsInParallel: %+v\n", saleDocumentsInParallel)

	GetPaymentBulk(apiClient)

	GetVatRatesBulk(apiClient)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
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

func GetPaymentBulk(cl *api.Client) {
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

	bulkResp, err := salesCLI.GetPaymentsBulk(ctx, bulkFilters, map[string]string{})
	common.Die(err)

	fmt.Println("GetPaymentsBulk:")
	fmt.Println(common.ConvertSourceToJsonStrIfPossible(bulkResp))
}

func GetVatRatesBulk(cl *api.Client) {
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

	bulkResp, err := salesCLI.GetVatRatesBulk(ctx, bulkFilters, map[string]string{})
	common.Die(err)

	fmt.Println("GetVatRatesBulk:")
	fmt.Println(common.ConvertSourceToJsonStrIfPossible(bulkResp))
}
