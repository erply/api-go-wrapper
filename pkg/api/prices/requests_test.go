package prices

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSupplierPriceListsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetPriceListsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetPriceListsResponseBulkItem{
				{
					Status: statusBulk,
					PriceLists: []PriceList{
						{
							ID: 123,
							Name:   "Some Price 123",
						},
						{
							ID: 124,
							Name:   "Some Price 124",
						},
					},
				},
				{
					Status: statusBulk,
					PriceLists: []PriceList{
						{
							ID: 125,
							Name:   "Some Price 125",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetSupplierPriceListsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 2,
				"pageNo":        1,
			},
			{
				"recordsOnPage": 2,
				"pageNo":        2,
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, []PriceList{
		{
			ID: 123,
			Name:   "Some Price 123",
		},
		{
			ID: 124,
			Name:    "Some Price 124",
		},
	}, bulkResp.BulkItems[0].PriceLists)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []PriceList{
		{
			ID: 125,
			Name:    "Some Price 125",
		},
	}, bulkResp.BulkItems[1].PriceLists)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestGetSupplierPriceLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := GetPriceListsResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			PriceLists: []PriceList{
				{
					ID: 123,
					Name:   "Some Price 123",
				},
				{
					ID: 124,
					Name:   "Some Price 124",
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	actualPrices, err := cl.GetSupplierPriceLists(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []PriceList{
		{
			ID: 123,
			Name:   "Some Price 123",
		},
		{
			ID: 124,
			Name:    "Some Price 124",
		},
	}, actualPrices)
}
