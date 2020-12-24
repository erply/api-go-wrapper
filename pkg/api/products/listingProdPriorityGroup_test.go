package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"
)

func sendRequestPriorityGroup(w http.ResponseWriter, errStatus sharedCommon.ApiError, totalCount int, productPrioGroupIDs [][]int) error {
	bulkResp := GetProductPriorityGroupResponseBulk{
		Status: sharedCommon.Status{ResponseStatus: "ok"},
	}

	bulkItems := make([]GetProductPriorityGroupBulkItem, 0, len(productPrioGroupIDs))
	for _, groupIDs := range productPrioGroupIDs {
		prodPrioGroups := make([]ProductPriorityGroup, 0, len(groupIDs))
		for _, id := range groupIDs {
			prodPrioGroups = append(prodPrioGroups, ProductPriorityGroup{
				PriorityGroupID:   id,
				PriorityGroupName: fmt.Sprintf("Some Prio Group %d", id),
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
		statusBulk.RecordsInResponse = len(productPrioGroupIDs)

		bulkItems = append(bulkItems, GetProductPriorityGroupBulkItem{
			Status:  statusBulk,
			Records: prodPrioGroups,
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

func TestProdPrioGroupListingCountSuccess(t *testing.T) {
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"recordsOnPage": float64(1),
				"pageNo":        float64(1),
				"requestName":   "getProductPriorityGroups",
			},
		})

		err := sendRequestPriorityGroup(w, 0, totalCount, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	dataProvider := NewPrioGroupListingDataProvider(productsClient)

	actualCount, err := dataProvider.Count(context.Background(), map[string]interface{}{"pageNo": 1, "recordsOnPage": 1})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, totalCount, actualCount)
}

func TestProdPrioGroupListingCountError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequestPriorityGroup(w, sharedCommon.MalformedRequest, 0, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	listingDataProvider := NewPrioGroupListingDataProvider(productsClient)

	actualCount, err := listingDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), sharedCommon.MalformedRequest.String())
	assert.Equal(t, 0, actualCount)
}

func TestProdPrioGroupListingCountWithNoBulkItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequest(w, 0, 0, [][]int{})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	dataProvider := NewPrioGroupListingDataProvider(productsClient)

	actualCount, err := dataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 0, actualCount)
}

func TestProdPrioGroupReadSuccess(t *testing.T) {
	const limit = 2
	const offset = 1
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"recordsOnPage": float64(10),
				"pageNo":        float64(1),
				"requestName":"getProductPriorityGroups",
			},
			{
				"recordsOnPage": float64(10),
				"pageNo":        float64(2),
				"requestName":"getProductPriorityGroups",
			},
		})

		err := sendRequestPriorityGroup(w, 0, totalCount, [][]int{{1, 2}, {3, 4}, {5}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	dataProvider := NewPrioGroupListingDataProvider(productsClient)

	actualProdPrioGroupIDs := make([]int, 0, 5)
	err := dataProvider.Read(
		context.Background(),
		[]map[string]interface{}{
			{
				"pageNo":        1,
				"recordsOnPage": 10,
			},
			{
				"pageNo":        2,
				"recordsOnPage": 10,
			},
		},
		func(item interface{}) {
			assert.IsType(t, item, ProductPriorityGroup{})
			actualProdPrioGroupIDs = append(actualProdPrioGroupIDs, item.(ProductPriorityGroup).PriorityGroupID)
		},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, actualProdPrioGroupIDs)
}

func TestProdPrioGroupReadError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendRequestPriorityGroup(w, sharedCommon.MalformedRequest, 10, [][]int{{1}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	dataProvider := NewPrioGroupListingDataProvider(productsClient)

	err := dataProvider.Read(
		context.Background(),
		[]map[string]interface{}{{"somekey": "smeval"}},
		func(item interface{}) {},
	)
	assert.Error(t, err)
	if err == nil {
		return
	}

	assert.Contains(t, err.Error(), sharedCommon.MalformedRequest.String())
}

func TestProdPrioGroupReadSuccessIntegration(t *testing.T) {
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
			err = sendRequestPriorityGroup(w, 0, totalCount, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
		} else {
			err = sendRequestPriorityGroup(w, 0, totalCount, [][]int{{11}})
		}

		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	dataProvider := NewPrioGroupListingDataProvider(productsClient)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        10,
			MaxFetchersCount:          10,
		},
		dataProvider,
		func(sleepTime time.Duration) {},
	)

	prodsChan := lister.Get(context.Background(), map[string]interface{}{})

	actualProdPrioGroupIDs := collectProdPrioGroupIDsFromChannel(prodsChan)
	sort.Ints(actualProdPrioGroupIDs)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, actualProdPrioGroupIDs)
}

func collectProdPrioGroupIDsFromChannel(itemsChan sharedCommon.ItemsStream) []int {
	actualProdPrioGroupIDs := make([]int, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for item := range itemsChan {
			actualProdPrioGroupIDs = append(actualProdPrioGroupIDs, item.Payload.(ProductPriorityGroup).PriorityGroupID)
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

	return actualProdPrioGroupIDs
}
