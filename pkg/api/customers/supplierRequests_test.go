package customers

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
func TestSupplierManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	//and here
	var (
		testingCustomer = &CustomerRequest{
			RegistryCode: "",
			CompanyName:  "",
			Username:     "",
			Password:     "",
		}
		ctx = context.Background()
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test get suppliers", func(t *testing.T) {
		suppliers, err := cli.GetSuppliers(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(suppliers)
	})
	t.Run("test post supplier", func(t *testing.T) {
		params := map[string]string{
			"companyName": testingCustomer.CompanyName,
			"code":        testingCustomer.RegistryCode,
		}
		resp, err := cli.SaveSupplier(ctx, params)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}

func TestGetSuppliersBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetSuppliersResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []GetSuppliersResponseBulkItem{
				{
					Status: statusBulk,
					Suppliers: []Supplier{
						{
							SupplierId: 123,
							FullName:   "Some Supplier123",
						},
						{
							SupplierId: 124,
							FullName:   "Some Supplier124",
						},
					},
				},
				{
					Status: statusBulk,
					Suppliers: []Supplier{
						{
							SupplierId: 125,
							FullName:   "Some Supplier125",
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

	suppliersClient := NewClient(cli)

	suppliersBulk, err := suppliersClient.GetSuppliersBulk(
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

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, suppliersBulk.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, []Supplier{
		{
			SupplierId: 123,
			FullName:   "Some Supplier123",
		},
		{
			SupplierId: 124,
			FullName:   "Some Supplier124",
		},
	}, suppliersBulk.BulkItems[0].Suppliers)

	assert.Equal(t, expectedStatus, suppliersBulk.BulkItems[0].Status)

	assert.Equal(t, []Supplier{
		{
			SupplierId: 125,
			FullName:   "Some Supplier125",
		},
	}, suppliersBulk.BulkItems[1].Suppliers)
	assert.Equal(t, expectedStatus, suppliersBulk.BulkItems[1].Status)
}

func TestGetSuppliersBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	_, err := suppliersClient.GetSuppliersBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetSuppliersResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestSaveSupplierBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := SaveSuppliersResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []SaveSuppliersResponseBulkItem{
				{
					Status: statusBulk,
					Records: []SaveSupplierResp{
						{
							SupplierID:    123,
							AlreadyExists: false,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveSupplierResp{
						{
							SupplierID:    124,
							AlreadyExists: true,
						},
					},
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

	suppliersClient := NewClient(cli)

	saveResp, err := suppliersClient.SaveSupplierBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"supplierID": 123,
				"fullName":   "Some name",
			},
			{
				"fullName": "Some other name",
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, saveResp.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, saveResp.BulkItems, 2)

	assert.Equal(t, []SaveSupplierResp{
		{
			SupplierID:    123,
			AlreadyExists: false,
		},
	}, saveResp.BulkItems[0].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[0].Status)

	assert.Equal(t, []SaveSupplierResp{
		{
			SupplierID:    124,
			AlreadyExists: true,
		},
	}, saveResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[1].Status)
}

func TestSaveSupplierBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("some junk value"))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	_, err := suppliersClient.SaveSupplierBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"supplierID": 123,
				"fullName":   "Some name",
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal SaveSuppliersResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestDeleteSupplierBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := DeleteSuppliersResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteSuppliersResponseBulkItem{
				{
					Status: common2.StatusBulk{
						RequestID: "123",
						Status: common2.Status{
							ResponseStatus: "ok",
						},
					},
				},
				{
					Status: common2.StatusBulk{
						RequestID: "124",
						Status: common2.Status{
							ResponseStatus: "ok",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)

		parsedRequest, err := common.ExtractBulkFiltersFromRequest(r)
		assert.NoError(t, err)
		if err != nil {
			return
		}

		assert.Equal(t, "someclient", parsedRequest["clientCode"])
		assert.Equal(t, "somesess", parsedRequest["sessionKey"])

		requests := parsedRequest["requests"].([]map[string]interface{})

		assert.Len(t, requests, 2)
		assert.Equal(t, "deleteSupplier", requests[0]["requestName"])
		assert.Equal(t, "100000046", requests[0]["supplierID"])

		assert.Equal(t, "deleteSupplier", requests[1]["requestName"])
		assert.Equal(t, "100000047", requests[1]["supplierID"])
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	saveResp, err := suppliersClient.DeleteSupplierBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"supplierID": "100000046",
			},
			{
				"supplierID": "100000047",
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, saveResp.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, saveResp.BulkItems, 2)

	requestIDsFromBulkResp := []string{}
	for _, bulkItem := range saveResp.BulkItems {
		requestIDsFromBulkResp = append(requestIDsFromBulkResp, bulkItem.Status.RequestID)
	}

	assert.EqualValues(t, []string{"123", "124"}, requestIDsFromBulkResp)
}

func TestDeleteSupplier(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := DeleteSupplierResponse{
			Status: common2.Status{ResponseStatus: "ok"},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)

		reqItems := make(map[string]interface{})
		for key, vals := range r.URL.Query() {
			reqItems[key] = vals[0]
		}

		assert.Equal(t, map[string]interface{}{
			"setContentType": "1",
			"request":        "deleteSupplier",
			"sessionKey":     "somesess",
			"supplierID":     "100000046",
			"clientCode":     "someclient",
		}, reqItems)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	err := suppliersClient.DeleteSupplier(
		context.Background(),
		map[string]string{
			"supplierID": "100000046",
		},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}
}

func TestDeleteSupplierBulkFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("some junk value"))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	_, err := suppliersClient.DeleteSupplierBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"supplierID": 123,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal DeleteSuppliersResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}
