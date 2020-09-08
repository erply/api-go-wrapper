package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/documents"
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

	purchaseDocuments, err := GetPurchaseDocuments(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetPurchaseDocuments: %+v\n", purchaseDocuments)

	purchaseDocumentsBulk, err := GetPurchaseDocumentsBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetPurchaseDocumentsBulk: %+v\n", purchaseDocumentsBulk)

	purchaseDocumentsParallel, err := GetPurchaseDocumentsInParallel(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetPurchaseDocumentsInParallel: %+v\n", purchaseDocumentsParallel)
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
