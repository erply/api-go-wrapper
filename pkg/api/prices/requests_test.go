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
							ID:   123,
							Name: "Some Price 123",
						},
						{
							ID:   124,
							Name: "Some Price 124",
						},
					},
				},
				{
					Status: statusBulk,
					PriceLists: []PriceList{
						{
							ID:   125,
							Name: "Some Price 125",
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

	defer srv.Close()

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
			ID:   123,
			Name: "Some Price 123",
		},
		{
			ID:   124,
			Name: "Some Price 124",
		},
	}, bulkResp.BulkItems[0].PriceLists)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []PriceList{
		{
			ID:   125,
			Name: "Some Price 125",
		},
	}, bulkResp.BulkItems[1].PriceLists)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestAddProductToSupplierPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ChangeProductToSupplierPriceListResponse{
			Status:                                 sharedCommon.Status{ResponseStatus: "ok"},
			ChangeProductToSupplierPriceListResult: []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 123}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"productID":           "123",
		"supplierPriceListID": "10",
		"price":               "10.00",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.AddProductToSupplierPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 123, resp.SupplierPriceListProductID)
}

func TestEditProductToSupplierPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "editProductInSupplierPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "1234", r.URL.Query().Get("supplierPriceListProductID"))
		assert.Equal(t, "20.23", r.URL.Query().Get("price"))

		resp := ChangeProductToSupplierPriceListResponse{
			Status:                                 sharedCommon.Status{ResponseStatus: "ok"},
			ChangeProductToSupplierPriceListResult: []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 1234}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"supplierPriceListProductID": "1234",
		"price":                      "20.23",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.EditProductToSupplierPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 1234, resp.SupplierPriceListProductID)
}

func TestChangeProductToSupplierPriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		clientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", clientCode)

		sessKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessKey)

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName":         "addProductToSupplierPriceList",
				"productID":           "123",
				"supplierPriceListID": "10",
				"price":               "10.00",
			},
			{
				"requestName":         "addProductToSupplierPriceList",
				"productID":           "124",
				"supplierPriceListID": "10",
				"price":               "20.01",
			},
			{
				"requestName":                "editProductInSupplierPriceList",
				"supplierPriceListProductID": "777",
				"price":                      "22.01",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := ChangeProductToSupplierPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []ChangeProductToSupplierPriceListResultBulkItem{
				{
					Status:  statusBulk,
					Records: []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 123}},
				},
				{
					Status:  statusBulk,
					Records: []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 124}},
				},
				{
					Status:  statusBulk,
					Records: []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 125}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"productID":           "123",
			"supplierPriceListID": "10",
			"price":               "10.00",
		},
		{
			"productID":           "124",
			"supplierPriceListID": "10",
			"price":               "20.01",
		},
		{
			"supplierPriceListProductID": "777",
			"price":                      "22.01",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.ChangeProductToSupplierPriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)
	assert.Equal(t, []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 123}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 124}}, bulkResp.BulkItems[1].Records)
	assert.Equal(t, []ChangeProductToSupplierPriceListResult{{SupplierPriceListProductID: 125}}, bulkResp.BulkItems[2].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestGetSupplierPriceLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := GetPriceListsResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			PriceLists: []PriceList{
				{
					ID:   123,
					Name: "Some Price 123",
				},
				{
					ID:   124,
					Name: "Some Price 124",
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

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
			ID:   123,
			Name: "Some Price 123",
		},
		{
			ID:   124,
			Name: "Some Price 124",
		},
	}, actualPrices)
}

func TestGetSupplierPriceListsBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	defer srv.Close()

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

	defer srv.Close()

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
		bulkResp := ProductsInSupplierPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []ProductsInSupplierPriceListResponseBulkItem{
				{
					Status: statusBulk,
					ProductsInSupplierPriceList: []ProductsInSupplierPriceList{
						{
							SupplierPriceListProductID: 123,
							Price:                      100,
						},
					},
				},
				{
					Status: statusBulk,
					ProductsInSupplierPriceList: []ProductsInSupplierPriceList{
						{
							SupplierPriceListProductID: 124,
							Price:                      200,
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

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetProductsInSupplierPriceListBulk(
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

	assert.Equal(t, []ProductsInSupplierPriceList{
		{
			SupplierPriceListProductID: 123,
			Price:                      100,
		},
	}, bulkResp.BulkItems[0].ProductsInSupplierPriceList)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []ProductsInSupplierPriceList{
		{
			SupplierPriceListProductID: 124,
			Price:                      200,
		},
	}, bulkResp.BulkItems[1].ProductsInSupplierPriceList)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestGetProductPriceListsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetProductsInPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductsInPriceListResponseBulkItem{
				{
					Status: statusBulk,
					PriceLists: []ProductsInPriceList{
						{
							PriceListProductID: 123,
							Price:              100,
						},
					},
				},
				{
					Status: statusBulk,
					PriceLists: []ProductsInPriceList{
						{
							PriceListProductID: 124,
							Price:              200,
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

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetProductsInPriceListBulk(
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

	assert.Equal(t, []ProductsInPriceList{
		{
			PriceListProductID: 123,
			Price:              100,
		},
	}, bulkResp.BulkItems[0].PriceLists)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []ProductsInPriceList{
		{
			PriceListProductID: 124,
			Price:              200,
		},
	}, bulkResp.BulkItems[1].PriceLists)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestGetProductSupplierPriceLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := ProductsInSupplierPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			ProductsInSupplierPriceList: []ProductsInSupplierPriceList{
				{
					SupplierPriceListProductID: 123,
					Price:                      100,
				},
				{
					SupplierPriceListProductID: 124,
					Price:                      200,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	actualProductPriceItems, err := cl.GetProductsInSupplierPriceList(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []ProductsInSupplierPriceList{
		{
			SupplierPriceListProductID: 123,
			Price:                      100,
		},
		{
			SupplierPriceListProductID: 124,
			Price:                      200,
		},
	}, actualProductPriceItems)
}

func TestGetProductPriceLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := GetProductsInPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			PriceLists: []ProductsInPriceList{
				{
					PriceListProductID: 123,
					Price:              100,
				},
				{
					PriceListProductID: 124,
					Price:              200,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	actualProductPriceItems, err := cl.GetProductsInPriceList(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []ProductsInPriceList{
		{
			PriceListProductID: 123,
			Price:              100,
		},
		{
			PriceListProductID: 124,
			Price:              200,
		},
	}, actualProductPriceItems)
}

func TestGetProductPriceListsBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetProductsInSupplierPriceListBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal ProductsInSupplierPriceListResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestGetProductPriceListsResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	pricesCl := NewClient(cli)

	_, err := pricesCl.GetProductsInSupplierPriceList(
		context.Background(),
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal ProductsInSupplierPriceListResponse from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestDeleteProductsFromSupplierPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "deleteProductsFromSupplierPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "2223", r.URL.Query().Get("supplierPriceListID"))
		assert.Equal(t, "3444,3445", r.URL.Query().Get("supplierPriceListProductIDs"))

		resp := DeleteProductsFromSupplierPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			DeleteProductsFromSupplierPriceListResult: []DeleteProductsFromSupplierPriceListResult{{DeletedIDs: "3444", NonExistingIDs: "3445"}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"supplierPriceListID":         "2223",
		"supplierPriceListProductIDs": "3444,3445",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.DeleteProductsFromSupplierPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, "3444", resp.DeletedIDs)
	assert.Equal(t, "3445", resp.NonExistingIDs)
}

func TestDeleteProductsFromSupplierPriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		clientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", clientCode)

		sessKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessKey)

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName":                 "deleteProductsFromSupplierPriceList",
				"supplierPriceListID":         "22",
				"supplierPriceListProductIDs": "3456",
			},
			{
				"requestName":                 "deleteProductsFromSupplierPriceList",
				"supplierPriceListID":         "23",
				"supplierPriceListProductIDs": "3457",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := DeleteProductsFromSupplierPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteProductsFromSupplierPriceListBulkItem{
				{
					Status:  statusBulk,
					Records: []DeleteProductsFromSupplierPriceListResult{{DeletedIDs: "3456", NonExistingIDs: ""}},
				},
				{
					Status:  statusBulk,
					Records: []DeleteProductsFromSupplierPriceListResult{{DeletedIDs: "3457", NonExistingIDs: ""}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"supplierPriceListID":         "22",
			"supplierPriceListProductIDs": "3456",
		},
		{
			"supplierPriceListID":         "23",
			"supplierPriceListProductIDs": "3457",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.DeleteProductsFromSupplierPriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []DeleteProductsFromSupplierPriceListResult{{DeletedIDs: "3456", NonExistingIDs: ""}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []DeleteProductsFromSupplierPriceListResult{{DeletedIDs: "3457", NonExistingIDs: ""}}, bulkResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestSaveSupplierPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "saveSupplierPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "Some Price Name 1", r.URL.Query().Get("name"))
		assert.Equal(t, "34456", r.URL.Query().Get("supplierID"))
		assert.Equal(t, "1", r.URL.Query().Get("productID"))
		assert.Equal(t, "100", r.URL.Query().Get("price"))
		assert.Equal(t, "10", r.URL.Query().Get("amount"))

		resp := SaveSupplierPriceListResultResponse{
			Status:                      sharedCommon.Status{ResponseStatus: "ok"},
			SaveSupplierPriceListResult: []SaveSupplierPriceListResult{{SupplierPriceListID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"name":       "Some Price Name 1",
		"supplierID": "34456",
		"productID":  "1",
		"price":      "100",
		"amount":     "10",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SaveSupplierPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, resp.SupplierPriceListID)
}

func TestSaveSupplierPriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		clientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", clientCode)

		sessKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessKey)

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName": "saveSupplierPriceList",
				"name":        "Some Price 1",
				"supplierID":  "1",
				"productID":   "3456",
				"price":       "10.22",
				"amount":      "23",
			},
			{
				"requestName": "saveSupplierPriceList",
				"name":        "Some Price 2",
				"supplierID":  "1",
				"productID":   "3457",
				"price":       "230.22",
				"amount":      "1",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := SaveSupplierPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveSupplierPriceListBulkItem{
				{
					Status:  statusBulk,
					Records: []SaveSupplierPriceListResult{{SupplierPriceListID: 3456}},
				},
				{
					Status:  statusBulk,
					Records: []SaveSupplierPriceListResult{{SupplierPriceListID: 3457}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"name":       "Some Price 1",
			"supplierID": "1",
			"productID":  "3456",
			"price":      "10.22",
			"amount":     "23",
		},
		{
			"name":       "Some Price 2",
			"supplierID": "1",
			"productID":  "3457",
			"price":      "230.22",
			"amount":     "1",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveSupplierPriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []SaveSupplierPriceListResult{{SupplierPriceListID: 3456}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []SaveSupplierPriceListResult{{SupplierPriceListID: 3457}}, bulkResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestSavePriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "savePriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "Some Price Name 1", r.URL.Query().Get("name"))
		assert.Equal(t, "34456", r.URL.Query().Get("pricelistID"))
		assert.Equal(t, "BASE_PRICE_LIST", r.URL.Query().Get("type"))

		resp := SavePriceListResultResponse{
			Status:               sharedCommon.Status{ResponseStatus: "ok"},
			SavePriceListResults: []SavePriceListResult{{PriceListID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"name":        "Some Price Name 1",
		"pricelistID": "34456",
		"type":        "BASE_PRICE_LIST",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SavePriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, resp.PriceListID)
}

func TestSavePriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "savePriceList",
				"name":        "Some Price Name 1",
				"pricelistID": "34456",
				"type":        "BASE_PRICE_LIST",
			},
			{
				"requestName": "savePriceList",
				"productID":   "124",
				"name":        "Some Price Name 2",
				"type":        "BASE_PRICE_LIST",
			},
			{
				"requestName": "savePriceList",
				"name":        "Some Price Name 3",
				"type":        "STORE_PRICE_LIST",
			},
		})

		bulkResp := SavePriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SavePriceListBulkItem{
				{
					Status:  statusBulk,
					Records: []SavePriceListResult{{PriceListID: 3456}},
				},
				{
					Status:  statusBulk,
					Records: []SavePriceListResult{{PriceListID: 3457}},
				},
				{
					Status: statusBulk,
					Records: []SavePriceListResult{
						{
							PriceListID: 3458,
							ItemsNotAddedToPriceList: []NotAddedPriceItem{
								{
									Type: "Product",
									ID:   123,
								},
							},
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

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"name":        "Some Price Name 1",
			"pricelistID": "34456",
			"type":        "BASE_PRICE_LIST",
		},
		{
			"productID": "124",
			"name":      "Some Price Name 2",
			"type":      "BASE_PRICE_LIST",
		},
		{
			"name": "Some Price Name 3",
			"type": "STORE_PRICE_LIST",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SavePriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)
	assert.Equal(t, []SavePriceListResult{{PriceListID: 3456}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []SavePriceListResult{{PriceListID: 3457}}, bulkResp.BulkItems[1].Records)
	assert.Equal(
		t,
		[]SavePriceListResult{
			{
				PriceListID: 3458,
				ItemsNotAddedToPriceList: []NotAddedPriceItem{
					{
						Type: "Product",
						ID:   123,
					},
				}},
		}, bulkResp.BulkItems[2].Records,
	)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestChangeProductInPriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName": "addProductToPriceList",
				"priceListID": "123",
				"productID":   "993",
				"price":       "10.00",
			},
			{
				"requestName": "addProductToPriceList",
				"priceListID": "123",
				"productID":   "994",
				"price":       "22.00",
			},
			{
				"requestName":        "editProductInPriceList",
				"priceListProductID": "777",
				"price":              "22.01",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := ChangeProductToPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []ChangeProductToPriceListResultBulkItem{
				{
					Status:  statusBulk,
					Records: []ChangeProductToPriceListResult{{PriceListProductID: 123}},
				},
				{
					Status:  statusBulk,
					Records: []ChangeProductToPriceListResult{{PriceListProductID: 124}},
				},
				{
					Status:  statusBulk,
					Records: []ChangeProductToPriceListResult{{PriceListProductID: 125}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"priceListID": "123",
			"productID":   "993",
			"price":       "10.00",
		},
		{
			"priceListID": "123",
			"productID":   "994",
			"price":       "22.00",
		},
		{
			"priceListProductID": "777",
			"price":              "22.01",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.ChangeProductToPriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)
	assert.Equal(t, []ChangeProductToPriceListResult{{PriceListProductID: 123}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []ChangeProductToPriceListResult{{PriceListProductID: 124}}, bulkResp.BulkItems[1].Records)
	assert.Equal(t, []ChangeProductToPriceListResult{{PriceListProductID: 125}}, bulkResp.BulkItems[2].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestAddProductToPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "addProductToPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "3333", r.URL.Query().Get("priceListID"))
		assert.Equal(t, "342314", r.URL.Query().Get("productID"))
		assert.Equal(t, "22.22", r.URL.Query().Get("price"))

		resp := ChangeProductToPriceListResponse{
			Status:                          sharedCommon.Status{ResponseStatus: "ok"},
			ChangeProductToPriceListResults: []ChangeProductToPriceListResult{{PriceListProductID: 123}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"priceListID": "3333",
		"productID":   "342314",
		"price":       "22.22",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.AddProductToPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 123, resp.PriceListProductID)
}

func TestEditProductInPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "editProductInPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "1234", r.URL.Query().Get("priceListProductID"))
		assert.Equal(t, "20.23", r.URL.Query().Get("price"))

		resp := ChangeProductToPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			ChangeProductToPriceListResults: []ChangeProductToPriceListResult{
				{
					PriceListProductID: 1234,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"priceListProductID": "1234",
		"price":              "20.23",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.EditProductToPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 1234, resp.PriceListProductID)
}

func TestDeleteProductsFromPriceList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "deleteProductInPriceList", r.URL.Query().Get("request"))
		assert.Equal(t, "2223", r.URL.Query().Get("priceListID"))
		assert.Equal(t, "3444,3445", r.URL.Query().Get("priceListProductIDs"))

		resp := DeleteProductsFromPriceListResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			DeleteProductsFromPriceListResults: []DeleteProductsFromPriceListResult{
				{
					DeletedIDs:     "3444",
					NonExistingIDs: "3445",
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"priceListID":         "2223",
		"priceListProductIDs": "3444,3445",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.DeleteProductsFromPriceList(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, "3444", resp.DeletedIDs)
	assert.Equal(t, "3445", resp.NonExistingIDs)
}

func TestDeleteProductsFromPriceListBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName":         "deleteProductInPriceList",
				"priceListID":         "22",
				"priceListProductIDs": "3456",
			},
			{
				"requestName":         "deleteProductInPriceList",
				"priceListID":         "23",
				"priceListProductIDs": "3457",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := DeleteProductsFromPriceListResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteProductsFromPriceListBulkItem{
				{
					Status:  statusBulk,
					Records: []DeleteProductsFromPriceListResult{{DeletedIDs: "3456", NonExistingIDs: ""}},
				},
				{
					Status:  statusBulk,
					Records: []DeleteProductsFromPriceListResult{{DeletedIDs: "3457", NonExistingIDs: ""}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"priceListID":         "22",
			"priceListProductIDs": "3456",
		},
		{
			"priceListID":         "23",
			"priceListProductIDs": "3457",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.DeleteProductsFromPriceListBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []DeleteProductsFromPriceListResult{{DeletedIDs: "3456", NonExistingIDs: ""}}, bulkResp.BulkItems[0].Records)
	assert.Equal(t, []DeleteProductsFromPriceListResult{{DeletedIDs: "3457", NonExistingIDs: ""}}, bulkResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
