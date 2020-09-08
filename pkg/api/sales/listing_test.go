package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"
)

func sendSaleDocumentsRequest(w http.ResponseWriter, errStatus errors.ApiError, totalCount int, documentsIDBulk [][]int) error {
	bulkResp := GetSaleDocumentResponseBulk{
		Status: sharedCommon.Status{ResponseStatus: "ok"},
	}

	bulkItems := make([]GetSaleDocumentBulkItem, 0, len(documentsIDBulk))
	for _, documentID := range documentsIDBulk {
		documents := make([]SaleDocument, 0, len(documentID))
		for _, id := range documentID {
			documents = append(documents, SaleDocument{
				ID:     id,
				Number: fmt.Sprintf("Doc %d", id),
			})
		}
		statusBulk := sharedCommon.StatusBulk{}
		if errStatus == 0 {
			statusBulk.ResponseStatus = "ok"
		} else {
			statusBulk.ResponseStatus = "not ok"
		}
		statusBulk.RecordsTotal = totalCount
		statusBulk.ErrorCode = errStatus
		statusBulk.RecordsInResponse = len(documentID)

		bulkItems = append(bulkItems, GetSaleDocumentBulkItem{
			Status:        statusBulk,
			SaleDocuments: documents,
		})
	}
	bulkResp.BulkItems = bulkItems

	jsonRaw, err := json.Marshal(bulkResp)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonRaw)
	if err != nil {
		return err
	}
	return nil
}

func TestSaleDocumentsListingCountSuccess(t *testing.T) {
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])
		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Equal(t, float64(1), requests[0]["pageNo"])
		assert.Equal(t, float64(1), requests[0]["recordsOnPage"])
		assert.Equal(t, "getSalesDocuments", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendSaleDocumentsRequest(w, 0, totalCount, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocumentsLister := NewSaleDocumentsListingDataProvider(salesClient)

	actualCount, err := salesDocumentsLister.Count(context.Background(), map[string]interface{}{"somekey": "smeval", "pageNo": 1, "recordsOnPage": 1})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, totalCount, actualCount)
}

func TestSaleDocumentsListingCountError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendSaleDocumentsRequest(w, errors.MalformedRequest, 0, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocLister := NewSaleDocumentsListingDataProvider(salesClient)

	actualCount, err := salesDocLister.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), errors.MalformedRequest.String())
	assert.Equal(t, 0, actualCount)
}

func TestSaleDocumentsListingCountWithNoBulkItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendSaleDocumentsRequest(w, 0, 0, [][]int{})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocLister := NewSaleDocumentsListingDataProvider(salesClient)

	actualCount, err := salesDocLister.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 0, actualCount)
}

func TestSaleDocumentsReadSuccess(t *testing.T) {
	const limit = 2
	const offset = 1
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])

		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Len(t, requests, 1)

		assert.Equal(t, float64(offset), requests[0]["pageNo"])
		assert.Equal(t, float64(limit), requests[0]["recordsOnPage"])

		assert.Equal(t, "getSalesDocuments", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendSaleDocumentsRequest(w, 0, totalCount, [][]int{{1, 2}, {3, 4}, {5}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocLister := NewSaleDocumentsListingDataProvider(salesClient)

	actualSalesIDs := make([]int, 0, 5)
	err := salesDocLister.Read(
		context.Background(),
		[]map[string]interface{}{
			{
				"somekey":       "smeval",
				"pageNo":        1,
				"recordsOnPage": 2,
			},
		},
		func(item interface{}) {
			assert.IsType(t, item, SaleDocument{})
			actualSalesIDs = append(actualSalesIDs, item.(SaleDocument).ID)
		},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, actualSalesIDs)
}

func TestSaleDocumentsReadError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendSaleDocumentsRequest(w, errors.MalformedRequest, 10, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocLister := NewSaleDocumentsListingDataProvider(salesClient)

	err := salesDocLister.Read(
		context.Background(),
		[]map[string]interface{}{{"somekey": "smeval"}},
		func(item interface{}) {},
	)
	assert.Error(t, err)
	if err == nil {
		return
	}

	assert.Contains(t, err.Error(), errors.MalformedRequest.String())
}

func TestSaleDocumentsReadSuccessIntegration(t *testing.T) {
	const totalCount = 11
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Len(t, requests, 1)

		if requests[0]["pageNo"] == float64(1) {
			err = sendSaleDocumentsRequest(w, 0, totalCount, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
		} else {
			err = sendSaleDocumentsRequest(w, 0, totalCount, [][]int{{11}})
		}

		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	salesClient := NewClient(baseClient)
	salesDocLister := NewSaleDocumentsListingDataProvider(salesClient)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        10,
			MaxFetchersCount:          10,
		},
		salesDocLister,
		func(sleepTime time.Duration) {},
	)

	salesDocChan := lister.Get(context.Background(), map[string]interface{}{})

	actualDocIDs := collectSaleDocIDsFromChannel(salesDocChan)
	sort.Ints(actualDocIDs)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, actualDocIDs)
}

func collectSaleDocIDsFromChannel(prodsChan sharedCommon.ItemsStream) []int {
	actualDocIDs := make([]int, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for prod := range prodsChan {
			actualDocIDs = append(actualDocIDs, prod.Payload.(SaleDocument).ID)
		}
	}()

mainLoop:
	for {
		select {
		case <-doneChan:
			break mainLoop
		case <-time.After(time.Second * 5):
			break mainLoop
		}
	}

	return actualDocIDs
}
