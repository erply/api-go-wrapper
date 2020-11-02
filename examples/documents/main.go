package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/documents"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	common.Die(err)

	purchaseDocuments, err := GetPurchaseDocuments(apiClient)
	common.Die(err)

	fmt.Printf("GetPurchaseDocuments: %+v\n", purchaseDocuments)

	purchaseDocumentsBulk, err := GetPurchaseDocumentsBulk(apiClient)
	common.Die(err)

	fmt.Printf("GetPurchaseDocumentsBulk: %+v\n", purchaseDocumentsBulk)

	purchaseDocumentsParallel, err := GetPurchaseDocumentsInParallel(apiClient)
	common.Die(err)

	fmt.Printf("GetPurchaseDocumentsInParallel: %+v\n", purchaseDocumentsParallel)

	SavePurchaseDocument(apiClient)
	SavePurchaseDocumentBulk(apiClient)
}

func GetPurchaseDocumentsBulk(cl *api.Client) (docs []documents.PurchaseDocument, err error) {
	documentsCli := cl.DocumentsManager

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

	bulkResp, err := documentsCli.GetPurchaseDocumentsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, doc := range bulkItem.PurchaseDocuments {
			docs = append(docs, doc)
		}
	}

	return
}

func GetPurchaseDocuments(cl *api.Client) (purchaseDocs []documents.PurchaseDocument, err error) {
	docsCli := cl.DocumentsManager

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filters := map[string]string{
		"recordsOnPage": "2",
		"pageNo":        "1",
	}
	docs, err := docsCli.GetPurchaseDocuments(ctx, filters)
	if err != nil {
		return
	}

	return docs, nil
}

func GetPurchaseDocumentsInParallel(cl *api.Client) ([]documents.PurchaseDocument, error) {
	documentsDataProvider := documents.NewListingDataProvider(cl.DocumentsManager)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        300,
			MaxFetchersCount:          10,
		},
		documentsDataProvider,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	docsChan := lister.Get(ctx, map[string]interface{}{
		"dateFrom": "2020-05-01",
	})

	purchaseDocuments := make([]documents.PurchaseDocument, 0)
	var err error
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range docsChan {
			if prod.Err != nil {
				err = prod.Err
				return
			}
			purchaseDocuments = append(purchaseDocuments, prod.Payload.(documents.PurchaseDocument))
		}
	}()

	<-doneChan
	return purchaseDocuments, err
}

func SaveBrandBulk(cl *api.Client) {
	prodCli := cl.ProductManager

	filter := []map[string]interface{}{
		{
			"name": "onetwothree",
		},
		{
			"name": "onefour",
		},
		{
			"name": "twor",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := prodCli.SaveBrandBulk(ctx, filter, map[string]string{})
	common.Die(err)

	fmt.Println(common.ConvertSourceToJsonStrIfPossible(res))
}

func SavePurchaseDocument(cl *api.Client) {
	docCli := cl.SalesManager

	filter := map[string]string{
		"warehouseID":  "1",
		"currencyCode": "EUR",
		"no":           "123",
		"supplierID":   "13383",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := docCli.SavePurchaseDocument(ctx, filter)
	common.Die(err)

	fmt.Println(common.ConvertSourceToJsonStrIfPossible(res))
}

func SavePurchaseDocumentBulk(cl *api.Client) {
	prodCli := cl.SalesManager

	filter := []map[string]interface{}{
		{
			"warehouseID":  "1",
			"currencyCode": "EUR",
			"no":           "124",
			"supplierID":   "13383",
		},
		{
			"warehouseID":  "2",
			"currencyCode": "EUR",
			"no":           "125",
			"supplierID":   "13383",
		},
		{
			"warehouseID":  "3",
			"currencyCode": "EUR",
			"no":           "126",
			"supplierID":   "13383",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := prodCli.SavePurchaseDocumentBulk(ctx, filter, map[string]string{})
	common.Die(err)

	fmt.Println(common.ConvertSourceToJsonStrIfPossible(res))
}
