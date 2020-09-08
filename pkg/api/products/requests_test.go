package products

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//works
func TestProductManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	//and here
	var (
		ctx = context.Background()
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))

	t.Run("test GetProducts", func(t *testing.T) {
		products, err := cli.GetProducts(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(products)
	})
	t.Run("test get product units", func(t *testing.T) {
		units, err := cli.GetProductUnits(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(units)
	})

	t.Run("test get product categories", func(t *testing.T) {
		cats, err := cli.GetProductCategories(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(cats)
	})
	t.Run("test get product brands", func(t *testing.T) {
		brands, err := cli.GetProductCategories(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(brands)
	})
	t.Run("test get product groups", func(t *testing.T) {
		groups, err := cli.GetProductGroups(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(groups)
	})
}

func TestGetProductsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetProductsResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductsResponseBulkItem{
				{
					Status: statusBulk,
					Products: []Product{
						{
							ProductID: 123,
							Code:      "Some Payload 123",
						},
						{
							ProductID: 124,
							Code:      "Some Payload 124",
						},
					},
				},
				{
					Status: statusBulk,
					Products: []Product{
						{
							ProductID: 125,
							Code:      "Some Payload 125",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(supplierResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	productClient := NewClient(cli)

	productsBulk, err := productClient.GetProductsBulk(
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

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, productsBulk.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, []Product{
		{
			ProductID: 123,
			Code:      "Some Payload 123",
		},
		{
			ProductID: 124,
			Code:      "Some Payload 124",
		},
	}, productsBulk.BulkItems[0].Products)

	assert.Equal(t, expectedStatus, productsBulk.BulkItems[0].Status)

	assert.Equal(t, []Product{
		{
			ProductID: 125,
			Code:      "Some Payload 125",
		},
	}, productsBulk.BulkItems[1].Products)
	assert.Equal(t, expectedStatus, productsBulk.BulkItems[1].Status)
}

func TestGetSuppliersBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	productsClient := NewClient(cli)

	_, err := productsClient.GetProductsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetProductsResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestGetProductsStockFileBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])
		assert.Equal(t, ResponseTypeCSV, parsedRequest["responseType"])

		requests := parsedRequest["requests"].([]map[string]interface{})
		assert.Len(t, requests, 2)
		assert.Equal(t, float64(1), requests[0]["warehouseID"])
		assert.Equal(t, "getProductStock", requests[0]["requestName"])
		assert.Equal(t, float64(2), requests[1]["warehouseID"])
		assert.Equal(t, "getProductStock", requests[1]["requestName"])

		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetProductStockFileResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductStockFileResponseBulkItem{
				{
					Status: statusBulk,
					GetProductStockFiles: []GetProductStockFile{
						{
							ReportLink: "some link 1",
						},
					},
				},
				{
					Status: statusBulk,
					GetProductStockFiles: []GetProductStockFile{
						{
							ReportLink: "some link 2",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(supplierResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	productClient := NewClient(cli)

	productsBulk, err := productClient.GetProductStockFileBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"warehouseID": 1,
			},
			{
				"warehouseID": 2,
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, productsBulk.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, []GetProductStockFile{
		{
			ReportLink: "some link 1",
		},
	}, productsBulk.BulkItems[0].GetProductStockFiles)

	assert.Equal(t, expectedStatus, productsBulk.BulkItems[0].Status)

	assert.Equal(t, []GetProductStockFile{
		{
			ReportLink: "some link 2",
		},
	}, productsBulk.BulkItems[1].GetProductStockFiles)
	assert.Equal(t, expectedStatus, productsBulk.BulkItems[1].Status)
}

func TestTestGetProductsStockFileBulkFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	productsClient := NewClient(cli)

	_, err := productsClient.GetProductStockFileBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"warehouseID": 1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetProductStockFileResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}
