package products

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
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetProductsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
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

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, productsBulk.Status)

	expectedStatus := sharedCommon.StatusBulk{}
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

		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetProductStockFileResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
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

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, productsBulk.Status)

	expectedStatus := sharedCommon.StatusBulk{}
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

func TestSaveProduct(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveProductResponse{
			Status:             sharedCommon.Status{ResponseStatus: "ok"},
			SaveProductResults: []SaveProductResult{{ProductID: 123}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	inpt := map[string]string{
		"groupID": "4",
		"code":    "code10",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SaveProduct(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 123, resp.ProductID)
}

func TestSaveProductBulk(t *testing.T) {
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
				"requestName": "saveProduct",
				"groupID":     "4",
				"code":        "code1",
			},
			{
				"requestName": "saveProduct",
				"groupID":     "4",
				"code":        "code2",
			},
			{
				"requestName": "saveProduct",
				"groupID":     "4",
				"code":        "code3",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := SaveProductResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveProductResponseBulkItem{
				{
					Status:   statusBulk,
					Products: []SaveProductResult{{ProductID: 123}},
				},
				{
					Status:   statusBulk,
					Products: []SaveProductResult{{ProductID: 124}},
				},
				{
					Status:   statusBulk,
					Products: []SaveProductResult{{ProductID: 125}},
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	inpt := []map[string]interface{}{
		{
			"requestName": "saveProduct",
			"groupID":     "4",
			"code":        "code1",
		},
		{
			"requestName": "saveProduct",
			"groupID":     "4",
			"code":        "code2",
		},
		{
			"requestName": "saveProduct",
			"groupID":     "4",
			"code":        "code3",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveProductBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)
	assert.Equal(t, []SaveProductResult{{ProductID: 123}}, bulkResp.BulkItems[0].Products)
	assert.Equal(t, []SaveProductResult{{ProductID: 124}}, bulkResp.BulkItems[1].Products)
	assert.Equal(t, []SaveProductResult{{ProductID: 125}}, bulkResp.BulkItems[2].Products)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestDeleteProduct(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := DeleteProductResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		clientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", clientCode)

		sessKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessKey)

		requestName := r.FormValue("request")
		assert.Equal(t, "deleteProduct", requestName)

		productToDelete := r.FormValue("productID")
		assert.Equal(t, "100001021", productToDelete)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	inpt := map[string]string{
		"productID": "100001021",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	err := cl.DeleteProduct(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
}

func TestDeleteProductBulk(t *testing.T) {
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
				"requestName": "deleteProduct",
				"productID":   "123",
			},
			{
				"requestName": "deleteProduct",
				"productID":   "124",
			},
			{
				"requestName": "deleteProduct",
				"productID":   "125",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := DeleteProductResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteProductResponseBulkItem{
				{
					Status: statusBulk,
				},
				{
					Status: statusBulk,
				},
				{
					Status: statusBulk,
				},
			},
		}
		jsonRaw, err := json.Marshal(bulkResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	inpt := []map[string]interface{}{
		{
			"productID": "123",
		},
		{
			"productID": "124",
		},
		{
			"productID": "125",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.DeleteProductBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}

func TestSaveAssortment(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveAssortmentResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			SaveAssortmentResults: []SaveAssortmentResult{
				{
					AssortmentID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		clientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", clientCode)

		sessKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessKey)

		requestName := r.FormValue("request")
		assert.Equal(t, "saveAssortment", requestName)

		assortmentID := r.FormValue("assortmentID")
		assert.Equal(t, "100001021", assortmentID)

		name := r.FormValue("name")
		assert.Equal(t, "some name", name)

		code := r.FormValue("code")
		assert.Equal(t, "some code", code)

		assert.Equal(t, "attributeName 1", r.FormValue("attributeName1"))
		assert.Equal(t, "attributeType 1", r.FormValue("attributeType1"))
		assert.Equal(t, "attributeValue 1", r.FormValue("attributeValue1"))

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	inpt := map[string]string{
		"assortmentID":    "100001021",
		"name":            "some name",
		"code":            "some code",
		"attributeName1":  "attributeName 1",
		"attributeType1":  "attributeType 1",
		"attributeValue1": "attributeValue 1",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedAssortmentRes := SaveAssortmentResult{
		AssortmentID: 123,
	}
	actualAssortmentRes, err := cl.SaveAssortment(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedAssortmentRes, actualAssortmentRes)
}

func TestSaveAssortmentBulk(t *testing.T) {
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
				"assortmentID": "123",
				"name":         "name 123",
				"requestName":  "saveAssortment",
			},
			{
				"assortmentID": "",
				"name":         "name 124",
				"requestName":  "saveAssortment",
			},
			{
				"assortmentID": "",
				"name":         "name 125",
				"code":         "code 125",
				"requestName":  "saveAssortment",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := SaveAssortmentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveAssortmentResponseBulkItem{
				{
					Status: statusBulk,
					SaveAssortmentResults: []SaveAssortmentResult{
						{
							AssortmentID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					SaveAssortmentResults: []SaveAssortmentResult{
						{
							AssortmentID: 124,
						},
					},
				},
				{
					Status: statusBulk,
					SaveAssortmentResults: []SaveAssortmentResult{
						{
							AssortmentID: 125,
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

	inpt := []map[string]interface{}{
		{
			"assortmentID": "123",
			"name":         "name 123",
		},
		{
			"assortmentID": "",
			"name":         "name 124",
		},
		{
			"assortmentID": "",
			"name":         "name 125",
			"code":         "code 125",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveAssortmentBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].SaveAssortmentResults, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].SaveAssortmentResults[0].AssortmentID)
	assert.Len(t, bulkResp.BulkItems[1].SaveAssortmentResults, 1)
	assert.Equal(t, 124, bulkResp.BulkItems[1].SaveAssortmentResults[0].AssortmentID)
	assert.Len(t, bulkResp.BulkItems[2].SaveAssortmentResults, 1)
	assert.Equal(t, 125, bulkResp.BulkItems[2].SaveAssortmentResults[0].AssortmentID)
}
