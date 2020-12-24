package warehouse

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

func TestWarehouseListingCountSuccess(t *testing.T) {
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
		assert.Equal(t, "getWarehouses", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendWarehouseRequest(w, 0, totalCount, [][]string{{"1"}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	actualCount, err := warehouseDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval", "pageNo": 1, "recordsOnPage": 1})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, totalCount, actualCount)
}

func TestWarehouseListingCountError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendWarehouseRequest(w, sharedCommon.MalformedRequest, 0, [][]string{{"1"}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	actualCount, err := warehouseDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), sharedCommon.MalformedRequest.String())
	assert.Equal(t, 0, actualCount)
}

func TestWarehouseListingCountWithNoBulkItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendWarehouseRequest(w, 0, 0, [][]string{})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	actualCount, err := warehouseDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 0, actualCount)
}

func TestWarehouseListingReadSuccess(t *testing.T) {
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

		assert.Equal(t, "getWarehouses", requests[0]["requestName"])
		assert.Equal(t, "smeval", requests[0]["somekey"])

		err = sendWarehouseRequest(w, 0, totalCount, [][]string{{"1", "2"}, {"3", "4"}, {"5"}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	actualWarehouseCodes := make([]string, 0, 5)
	err := warehouseDataProvider.Read(
		context.Background(),
		[]map[string]interface{}{
			{
				"somekey":       "smeval",
				"pageNo":        1,
				"recordsOnPage": 2,
			},
		},
		func(item interface{}) {
			assert.IsType(t, item, Warehouse{})
			actualWarehouseCodes = append(actualWarehouseCodes, item.(Warehouse).Code)
		},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, actualWarehouseCodes)
}

func TestWarehouseListingReadError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := sendWarehouseRequest(w, sharedCommon.MalformedRequest, 10, [][]string{{"1"}})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	err := warehouseDataProvider.Read(
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

func TestWarehouseListingReadSuccessIntegration(t *testing.T) {
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
			err = sendWarehouseRequest(w, 0, totalCount, [][]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}})
		} else {
			err = sendWarehouseRequest(w, 0, totalCount, [][]string{{"11"}})
		}

		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	defer srv.Close()

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	warehouseClient := NewClient(baseClient)
	warehouseDataProvider := NewListingDataProvider(warehouseClient)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 0,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        10,
			MaxFetchersCount:          10,
		},
		warehouseDataProvider,
		func(sleepTime time.Duration) {},
	)

	warehouseChan := lister.Get(context.Background(), map[string]interface{}{})

	actualWarehouseCodes := collectWarehouseCodesFromChannel(warehouseChan)
	sort.Strings(actualWarehouseCodes)

	assert.Equal(t, []string{"1", "10", "11", "2", "3", "4", "5", "6", "7", "8", "9"}, actualWarehouseCodes)
}

func collectWarehouseCodesFromChannel(warehouseChan sharedCommon.ItemsStream) []string {
	actualWarehouseCodes := make([]string, 0)
	doneChan := make(chan struct{}, 1)
	go func() {
		defer close(doneChan)
		for warehouse := range warehouseChan {
			actualWarehouseCodes = append(actualWarehouseCodes, warehouse.Payload.(Warehouse).Code)
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

	return actualWarehouseCodes
}

func sendWarehouseRequest(w http.ResponseWriter, errStatus sharedCommon.ApiError, totalCount int, warehouseCodeBulk [][]string) error {
	bulkResp := GetWarehousesResponseBulk{
		Status: sharedCommon.Status{ResponseStatus: "ok"},
	}

	bulkItems := make([]GetWarehousesBulkItem, 0, len(warehouseCodeBulk))
	for _, warehouseCodes := range warehouseCodeBulk {
		warehouses := make(Warehouses, 0, len(warehouseCodes))
		for _, id := range warehouseCodes {
			warehouses = append(warehouses, Warehouse{
				Code: id,
				Name: fmt.Sprintf("Some Warehouse %s", id),
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
		statusBulk.RecordsInResponse = len(warehouseCodes)

		bulkItems = append(bulkItems, GetWarehousesBulkItem{
			Status:     statusBulk,
			Warehouses: warehouses,
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
