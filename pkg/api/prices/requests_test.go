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

func TestGetSupplierPriceListsBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetSupplierPriceListsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetPriceListsResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestGetSupplierPriceListsResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetSupplierPriceLists(
		context.Background(),
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetPriceListsResponse from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestGetProductSupplierPriceListsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetProductPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductPriceListResponseBulkItem{
				{
					Status: statusBulk,
					ProductPriceList: []ProductPriceList{
						{
							PriceID: 123,
							Price:   100,
						},
					},
				},
				{
					Status: statusBulk,
					ProductPriceList: []ProductPriceList{
						{
							PriceID: 124,
							Price:   200,
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

	bulkResp, err := cl.GetProductPriceListsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 2,
				"pageNo":        1,
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

	assert.Equal(t, []ProductPriceList{
		{
			PriceID: 123,
			Price:   100,
		},
	}, bulkResp.BulkItems[0].ProductPriceList)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []ProductPriceList{
		{
			PriceID: 124,
			Price:   200,
		},
	}, bulkResp.BulkItems[1].ProductPriceList)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestGetProductSupplierPriceLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := GetProductPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			ProductPriceLists: []ProductPriceList{
				{
					PriceID: 123,
					Price:   100,
				},
				{
					PriceID: 124,
					Price:   200,
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

	actualProductPriceItems, err := cl.GetProductPriceLists(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []ProductPriceList{
		{
			PriceID: 123,
			Price:   100,
		},
		{
			PriceID: 124,
			Price:   200,
		},
	}, actualProductPriceItems)
}

func TestGetProductPriceListsBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetProductPriceListsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetProductPriceListResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestGetProductPriceListsResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetProductPriceLists(
		context.Background(),
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetProductPriceListResponse from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}
