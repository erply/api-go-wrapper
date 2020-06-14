package products

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
	"testing"
)

func extractBulkFiltersFromRequest(r *http.Request) (res map[string]interface{}, err error) {
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return
	}

	var requests []map[string]interface{}
	requestsRaw, ok := res["requests"]
	if ok {
		err = json.Unmarshal(requestsRaw.([]byte), &requests)
		if err != nil {
			return
		}
	}
	res["requests"] = requests
	return
}

func sendRequest(w http.ResponseWriter, errStatus errors.ApiError, totalCount int, productIDsBulk [][]int) error {
	bulkResp := GetProductsResponseBulk{
		Status: sharedCommon.Status{ResponseStatus: "ok"},
	}

	bulkItems := make([]GetProductsResponseBulkItem, 0, len(productIDsBulk))
	for _, productIDs := range productIDsBulk {
		products := make([]Product, 0, len(productIDs))
		for _, id := range productIDs {
			products = append(products, Product{
				ProductID: id,
				Name:      fmt.Sprintf("Some Product %d", id),
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
		statusBulk.RecordsInResponse = len(productIDs)

		bulkItems = append(bulkItems, GetProductsResponseBulkItem{
			Status:   statusBulk,
			Products: products,
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

func TestListingCount(t *testing.T) {
	const totalCount = 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := extractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}
		assert.Equal(t, "", parsedRequest)

		err = sendRequest(w, 0, totalCount, [][]int{})
		assert.NoError(t, err)
		if err != nil {
			return
		}
	}))

	baseClient := common.NewClient("somesess", "someclient", "", nil, nil)
	baseClient.Url = srv.URL
	productsClient := NewClient(baseClient)
	productsDataProvider := NewListingDataProvider(productsClient)

	actualCount, err := productsDataProvider.Count(context.Background(), map[string]interface{}{"somekey": "smeval"})
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, totalCount, actualCount)
}
