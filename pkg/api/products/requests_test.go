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
	"time"
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

	defer srv.Close()

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

	defer srv.Close()

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
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":   "someclient",
			"sessionKey":   "somesess",
			"responseType": ResponseTypeCSV,
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"warehouseID": float64(1),
				"requestName": "getProductStock",
			},
			{
				"warehouseID": float64(2),
				"requestName": "getProductStock",
			},
		})

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

	defer srv.Close()

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

	defer srv.Close()

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

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
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
		})

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

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"request":    "deleteProduct",
			"productID":  "100001021",
		})

		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
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
		})

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

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":      "someclient",
			"sessionKey":      "somesess",
			"request":         "saveAssortment",
			"assortmentID":    "100001021",
			"name":            "some name",
			"code":            "some code",
			"attributeName1":  "attributeName 1",
			"attributeType1":  "attributeType 1",
			"attributeValue1": "attributeValue 1",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
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
		})

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

	defer srv.Close()

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

func TestAddAssortmentProducts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":   "someclient",
			"sessionKey":   "somesess",
			"request":      "addAssortmentProducts",
			"productIDs":   "1,2,3",
			"assortmentID": "123",
			"status":       "ACTIVE",
		})

		resp := AddAssortmentProductsResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			AddAssortmentProductsResults: []AddAssortmentProductsResult{
				{
					ProductsAlreadyInAssortment: "123",
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
		"productIDs":   "1,2,3",
		"assortmentID": "123",
		"status":       "ACTIVE",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedAssortmentRes := AddAssortmentProductsResult{
		ProductsAlreadyInAssortment: "123",
	}
	actualAssortmentRes, err := cl.AddAssortmentProducts(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedAssortmentRes, actualAssortmentRes)
}

func TestAddAssortmentProductsBulk(t *testing.T) {
	bulkRequestWasMade := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bulkRequestWasMade = true
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"productIDs":   "1,2,3",
				"assortmentID": "assortment 123",
				"status":       "ACTIVE",
				"requestName":  "addAssortmentProducts",
			},
			{
				"productIDs":   "4",
				"assortmentID": "assortment 123",
				"status":       "NOT_FOR_SALE",
				"requestName":  "addAssortmentProducts",
			},
		})

		bulkResp := AddAssortmentProductsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []AddAssortmentProductsResponseBulkItem{
				{
					Status: statusBulk,
					AddAssortmentProductsResults: []AddAssortmentProductsResult{
						{},
					},
				},
				{
					Status: statusBulk,
					AddAssortmentProductsResults: []AddAssortmentProductsResult{
						{},
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
			"productIDs":   "1,2,3",
			"assortmentID": "assortment 123",
			"status":       "ACTIVE",
		},
		{
			"productIDs":   "4",
			"assortmentID": "assortment 123",
			"status":       "NOT_FOR_SALE",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.AddAssortmentProductsBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.True(t, bulkRequestWasMade)

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].AddAssortmentProductsResults, 1)
	assert.Equal(t, "", bulkResp.BulkItems[0].AddAssortmentProductsResults[0].ProductsAlreadyInAssortment)
	assert.Len(t, bulkResp.BulkItems[0].AddAssortmentProductsResults, 1)
	assert.Equal(t, "", bulkResp.BulkItems[0].AddAssortmentProductsResults[0].NonExistingIDs)
	assert.Len(t, bulkResp.BulkItems[1].AddAssortmentProductsResults, 1)
	assert.Equal(t, "", bulkResp.BulkItems[1].AddAssortmentProductsResults[0].ProductsAlreadyInAssortment)
	assert.Len(t, bulkResp.BulkItems[1].AddAssortmentProductsResults, 1)
	assert.Equal(t, "", bulkResp.BulkItems[1].AddAssortmentProductsResults[0].NonExistingIDs)
}

func TestEditAssortmentProducts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := EditAssortmentProductsResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			EditAssortmentProductsResults: []EditAssortmentProductsResult{
				{
					ProductsNotInAssortment: "123",
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":   "someclient",
			"sessionKey":   "somesess",
			"request":      "editAssortmentProducts",
			"productIDs":   "1,2,3",
			"assortmentID": "123",
			"status":       "ACTIVE",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"productIDs":   "1,2,3",
		"assortmentID": "123",
		"status":       "ACTIVE",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedAssortmentRes := EditAssortmentProductsResult{
		ProductsNotInAssortment: "123",
	}
	actualAssortmentRes, err := cl.EditAssortmentProducts(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedAssortmentRes, actualAssortmentRes)
}

func TestEditAssortmentProductsBulk(t *testing.T) {
	bulkRequestWasMade := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bulkRequestWasMade = true
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"productIDs":   "1,2,3",
				"assortmentID": "assortment 123",
				"status":       "ACTIVE",
				"requestName":  "editAssortmentProducts",
			},
			{
				"productIDs":   "4",
				"assortmentID": "assortment 123",
				"requestName":  "editAssortmentProducts",
			},
		})

		bulkResp := EditAssortmentProductsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []EditAssortmentProductsResponseBulkItem{
				{
					Status: statusBulk,
					EditAssortmentProductsResults: []EditAssortmentProductsResult{
						{
							ProductsNotInAssortment: "1",
						},
					},
				},
				{
					Status: statusBulk,
					EditAssortmentProductsResults: []EditAssortmentProductsResult{
						{},
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
			"productIDs":   "1,2,3",
			"assortmentID": "assortment 123",
			"status":       "ACTIVE",
			"requestName":  "editAssortmentProducts",
		},
		{
			"productIDs":   "4",
			"assortmentID": "assortment 123",
			"requestName":  "addAssortmentProducts",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.EditAssortmentProductsBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.True(t, bulkRequestWasMade)

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].EditAssortmentProductsResults, 1)
	assert.Equal(t, "1", bulkResp.BulkItems[0].EditAssortmentProductsResults[0].ProductsNotInAssortment)
	assert.Len(t, bulkResp.BulkItems[1].EditAssortmentProductsResults, 1)
	assert.Equal(t, "", bulkResp.BulkItems[1].EditAssortmentProductsResults[0].ProductsNotInAssortment)
}

func TestRemoveAssortmentProducts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := RemoveAssortmentProductResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			RemoveAssortmentProductResults: []RemoveAssortmentProductResult{
				{
					ProductsNotInAssortment: "123,124",
					DeletedIDs:              "1,2,3",
				},
			},
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":   "someclient",
			"sessionKey":   "somesess",
			"request":      "removeAssortmentProducts",
			"productIDs":   "1,2,3",
			"assortmentID": "123",
		})

		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"productIDs":   "1,2,3",
		"assortmentID": "123",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedAssortmentRes := RemoveAssortmentProductResult{
		ProductsNotInAssortment: "123,124",
		DeletedIDs:              "1,2,3",
	}

	actualAssortmentRes, err := cl.RemoveAssortmentProducts(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedAssortmentRes, actualAssortmentRes)
}

func TestRemoveAssortmentProductsBulk(t *testing.T) {
	bulkRequestWasMade := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bulkRequestWasMade = true
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"productIDs":   "1,2,3",
				"assortmentID": "assortment 123",
				"requestName":  "removeAssortmentProducts",
			},
			{
				"productIDs":   "4,5",
				"assortmentID": "assortment 123",
				"requestName":  "removeAssortmentProducts",
			},
		})

		bulkResp := RemoveAssortmentProductResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []RemoveAssortmentProductResponseBulkItem{
				{
					Status: statusBulk,
					RemoveAssortmentProductResults: []RemoveAssortmentProductResult{
						{
							DeletedIDs: "1,2,3",
						},
					},
				},
				{
					Status: statusBulk,
					RemoveAssortmentProductResults: []RemoveAssortmentProductResult{
						{
							DeletedIDs:              "4",
							ProductsNotInAssortment: "5",
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
			"productIDs":   "1,2,3",
			"assortmentID": "assortment 123",
		},
		{
			"productIDs":   "4,5",
			"assortmentID": "assortment 123",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.RemoveAssortmentProductsBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.True(t, bulkRequestWasMade)

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].RemoveAssortmentProductResults, 1)
	assert.Equal(t, "1,2,3", bulkResp.BulkItems[0].RemoveAssortmentProductResults[0].DeletedIDs)
	assert.Equal(t, "", bulkResp.BulkItems[0].RemoveAssortmentProductResults[0].ProductsNotInAssortment)

	assert.Len(t, bulkResp.BulkItems[1].RemoveAssortmentProductResults, 1)
	assert.Equal(t, "4", bulkResp.BulkItems[1].RemoveAssortmentProductResults[0].DeletedIDs)
	assert.Equal(t, "5", bulkResp.BulkItems[1].RemoveAssortmentProductResults[0].ProductsNotInAssortment)
}

func TestSaveProductCategory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveProductCategoryResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			SaveProductCategoryResults: []SaveProductCategoryResult{
				{
					ProductCategoryID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":        "someclient",
			"sessionKey":        "somesess",
			"request":           "saveProductCategory",
			"productCategoryID": "100001021",
			"name":              "some name",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"productCategoryID": "100001021",
		"name":              "some name",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedProdCategoryResp := SaveProductCategoryResult{
		ProductCategoryID: 123,
	}
	prodCategoryRes, err := cl.SaveProductCategory(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedProdCategoryResp, prodCategoryRes)
}

func TestSaveProductCategoryBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"productCategoryID": "123",
				"name":              "name 123",
				"requestName":       "saveProductCategory",
			},
			{
				"name":        "name 124",
				"requestName": "saveProductCategory",
			},
			{
				"name":        "name 125",
				"requestName": "saveProductCategory",
			},
		})

		bulkResp := SaveProductCategoryResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveProductCategoryResponseBulkItem{
				{
					Status: statusBulk,
					Records: []SaveProductCategoryResult{
						{
							ProductCategoryID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveProductCategoryResult{
						{
							ProductCategoryID: 124,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveProductCategoryResult{
						{
							ProductCategoryID: 125,
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
			"productCategoryID": "123",
			"name":              "name 123",
		},
		{
			"name": "name 124",
		},
		{
			"name": "name 125",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveProductCategoryBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].Records[0].ProductCategoryID)
	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 124, bulkResp.BulkItems[1].Records[0].ProductCategoryID)
	assert.Len(t, bulkResp.BulkItems[2].Records, 1)
	assert.Equal(t, 125, bulkResp.BulkItems[2].Records[0].ProductCategoryID)
}

func TestSaveBrand(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveBrandResultResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			SaveBrandResults: []SaveBrandResult{
				{
					BrandID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"request":    "saveBrand",
			"brandID":    "100001021",
			"name":       "some name",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"brandID": "100001021",
		"name":    "some name",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedRes := SaveBrandResult{
		BrandID: 123,
	}
	actualRes, err := cl.SaveBrand(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedRes, actualRes)
}

func TestSaveBrandBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"brandID":     "123",
				"name":        "name 123",
				"requestName": "saveBrand",
			},
		})

		bulkResp := SaveBrandResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveBrandResultResponseBulkItem{
				{
					Status: statusBulk,
					Records: []SaveBrandResult{
						{
							BrandID: 123,
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
			"brandID": "123",
			"name":    "name 123",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveBrandBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 1)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].Records[0].BrandID)
}

func TestSaveProductPriorityGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveProductPriorityGroupResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			SaveProductPriorityGroupResults: []SaveProductPriorityGroupResult{
				{
					PriorityGroupID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"request":    "saveProductPriorityGroup",
			"name":       "some name",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"name": "some name",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedRes := SaveProductPriorityGroupResult{
		PriorityGroupID: 123,
	}
	actualRes, err := cl.SaveProductPriorityGroup(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedRes, actualRes)
}

func TestSaveProductPriorityGroupBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"priorityGroupID": "123",
				"name":            "name 123",
				"requestName":     "saveProductPriorityGroup",
			},
			{
				"name":        "name 124",
				"requestName": "saveProductPriorityGroup",
			},
		})

		bulkResp := SaveProductPriorityGroupResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveProductPriorityGroupBulkItem{
				{
					Status: statusBulk,
					Records: []SaveProductPriorityGroupResult{
						{
							PriorityGroupID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveProductPriorityGroupResult{
						{
							PriorityGroupID: 124,
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
			"priorityGroupID": "123",
			"name":            "name 123",
		},
		{
			"name": "name 124",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveProductPriorityGroupBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].Records[0].PriorityGroupID)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 124, bulkResp.BulkItems[1].Records[0].PriorityGroupID)
}

func TestSaveProductGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveProductGroupResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			SaveProductGroupResults: []SaveProductGroupResult{
				{
					ProductGroupID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"request":    "saveProductGroup",
			"name":       "some name",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"name": "some name",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedRes := SaveProductGroupResult{
		ProductGroupID: 123,
	}
	actualRes, err := cl.SaveProductGroup(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedRes, actualRes)
}

func TestSaveProductGroupBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"productGroupID": "123",
				"name":           "name 123",
				"requestName":    "saveProductGroup",
			},
			{
				"name":        "name 124",
				"requestName": "saveProductGroup",
			},
		})

		bulkResp := SaveProductGroupResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveProductGroupBulkItem{
				{
					Status: statusBulk,
					Records: []SaveProductGroupResult{
						{
							ProductGroupID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveProductGroupResult{
						{
							ProductGroupID: 124,
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
			"productGroupID": "123",
			"name":           "name 123",
		},
		{
			"name": "name 124",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveProductGroupBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].Records[0].ProductGroupID)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 124, bulkResp.BulkItems[1].Records[0].ProductGroupID)
}

func TestDeleteProductGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := DeleteProductGroupResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":     "someclient",
			"sessionKey":     "somesess",
			"request":        "deleteProductGroup",
			"productGroupID": "100001021",
		})

		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"productGroupID": "100001021",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	err := cl.DeleteProductGroup(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
}

func TestDeleteProductGroupBulk(t *testing.T) {
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
				"requestName":    "deleteProductGroup",
				"productGroupID": "123",
			},
			{
				"requestName":    "deleteProductGroup",
				"productGroupID": "124",
			},
			{
				"requestName":    "deleteProductGroup",
				"productGroupID": "125",
			},
		})

		bulkResp := DeleteProductGroupResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteProductGroupResponseBulkItem{
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

	defer srv.Close()

	inpt := []map[string]interface{}{
		{
			"productGroupID": "123",
		},
		{
			"productGroupID": "124",
		},
		{
			"productGroupID": "125",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.DeleteProductGroupBulk(context.Background(), inpt, map[string]string{})
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

func TestGetProductPriorityGroupBulk(t *testing.T) {
	nowTimeStamp := time.Now().Unix()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"recordsOnPage": "10",
				"pageNo":        "1",
				"requestName":"getProductPriorityGroups",
			},
			{
				"recordsOnPage": "10",
				"pageNo":        "2",
				"requestName":"getProductPriorityGroups",
			},
		})

		bulkResp := GetProductPriorityGroupResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductPriorityGroupBulkItem{
				{
					Status: statusBulk,
					Records: []ProductPriorityGroup{
						{
							PriorityGroupID:   1,
							PriorityGroupName: "Some group",
							Added:             nowTimeStamp,
							LastModified:      nowTimeStamp,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []ProductPriorityGroup{
						{
							PriorityGroupID:   2,
							PriorityGroupName: "Some group 2",
							Added:             nowTimeStamp,
							LastModified:      nowTimeStamp,
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
			"recordsOnPage": "10",
			"pageNo":        "1",
		},
		{
			"recordsOnPage": "10",
			"pageNo":        "2",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetProductPriorityGroupBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 1, bulkResp.BulkItems[0].Records[0].PriorityGroupID)
	assert.Equal(t, "Some group", bulkResp.BulkItems[0].Records[0].PriorityGroupName)
	assert.Equal(t, nowTimeStamp, bulkResp.BulkItems[0].Records[0].Added)
	assert.Equal(t, nowTimeStamp, bulkResp.BulkItems[0].Records[0].LastModified)

	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 2, bulkResp.BulkItems[1].Records[0].PriorityGroupID)
	assert.Equal(t, "Some group 2", bulkResp.BulkItems[1].Records[0].PriorityGroupName)
	assert.Equal(t, nowTimeStamp, bulkResp.BulkItems[1].Records[0].Added)
	assert.Equal(t, nowTimeStamp, bulkResp.BulkItems[1].Records[0].LastModified)
}

func TestGetProductGroupsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"recordsOnPage": "10",
				"pageNo":        "1",
				"requestName":"getProductGroups",
			},
			{
				"recordsOnPage": "10",
				"pageNo":        "2",
				"requestName":"getProductGroups",
			},
		})

		bulkResp := GetProductGroupResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductGroupBulkItem{
				{
					Status: statusBulk,
					Records: []ProductGroup{
						{
							ID:   1,
							NameLanguages: NameLanguages{
								Name: "Prod Group 1",
							},
						},
					},
				},
				{
					Status: statusBulk,
					Records: []ProductGroup{
						{
							ID:   2,
							NameLanguages: NameLanguages{
								Name: "Prod Group 2",
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
			"recordsOnPage": "10",
			"pageNo":        "1",
		},
		{
			"recordsOnPage": "10",
			"pageNo":        "2",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetProductGroupsBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 1, bulkResp.BulkItems[0].Records[0].ID)
	assert.Equal(t, "Prod Group 1", bulkResp.BulkItems[0].Records[0].Name)

	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 2, bulkResp.BulkItems[1].Records[0].ID)
	assert.Equal(t, "Prod Group 2", bulkResp.BulkItems[1].Records[0].Name)
}

func TestGetProductCategoriesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"recordsOnPage": "10",
				"pageNo":        "1",
				"requestName":"getProductCategories",
			},
			{
				"recordsOnPage": "10",
				"pageNo":        "2",
				"requestName":"getProductCategories",
			},
		})

		bulkResp := GetProductCategoryResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetProductCategoryBulkItem{
				{
					Status: statusBulk,
					Records: []ProductCategory{
						{
							ProductCategoryID:   1,
							ProductCategoryName: "Product category 1",
						},
					},
				},
				{
					Status: statusBulk,
					Records: []ProductCategory{
						{
							ProductCategoryID:   2,
							ProductCategoryName: "Product category 2",
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
			"recordsOnPage": "10",
			"pageNo":        "1",
		},
		{
			"recordsOnPage": "10",
			"pageNo":        "2",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.GetProductCategoriesBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 1, bulkResp.BulkItems[0].Records[0].ProductCategoryID)
	assert.Equal(t, "Product category 1", bulkResp.BulkItems[0].Records[0].ProductCategoryName)

	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 2, bulkResp.BulkItems[1].Records[0].ProductCategoryID)
	assert.Equal(t, "Product category 2", bulkResp.BulkItems[1].Records[0].ProductCategoryName)
}
