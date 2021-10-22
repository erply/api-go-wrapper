package customers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erply/api-go-wrapper/internal/common"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/stretchr/testify/assert"
)

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

	defer srv.Close()

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

	defer srv.Close()

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

	defer srv.Close()

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

	defer srv.Close()

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

	defer srv.Close()

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

		common.AssertFormValues(t, r, map[string]interface{}{
			"setContentType": "1",
			"request":        "deleteSupplier",
			"sessionKey":     "somesess",
			"supplierID":     "100000046",
			"clientCode":     "someclient",
		})
	}))

	defer srv.Close()

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

	defer srv.Close()

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

func TestGetCompanyTypes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := GetCompanyTypesResponse{
			Status: common2.Status{
				ResponseStatus: "ok",
			},
			CompanyTypes: []CompanyType{
				{
					ID:    1,
					Name:  "name",
					Order: 1,
				},
				{
					ID:    2,
					Name:  "name2",
					Order: 2,
				},
			}}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	companyTypes, err := cl.GetCompanyTypes(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []CompanyType{
		{
			ID:    1,
			Name:  "name",
			Order: 1,
		},
		{
			ID:    2,
			Name:  "name2",
			Order: 2,
		}}, companyTypes)
}
func TestSaveSupplierGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveSupplierGroupResponse{
			Status: common2.Status{ResponseStatus: "ok"},
			Records: []SaveSupplierGroupRecord{
				{
					SupplierGroupID: 1,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"setContentType": "1",
			"request":        "saveSupplierGroup",
			"sessionKey":     "somesess",
			"name":           "100000046",
			"clientCode":     "someclient",
		})
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	res, err := suppliersClient.SaveSupplierGroup(
		context.Background(),
		map[string]string{
			"name": "100000046",
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 1, res.Records[0].SupplierGroupID)
}

func TestSaveCompanyType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SaveCompanyTypeResponse{
			Status: common2.Status{ResponseStatus: "ok"},
			Records: []SaveCompanyTypeRecord{
				{
					CompanyTypeID: 1,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"setContentType": "1",
			"request":        "saveCompanyType",
			"sessionKey":     "somesess",
			"name":           "100000046",
			"clientCode":     "someclient",
		})
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	suppliersClient := NewClient(cli)

	res, err := suppliersClient.SaveCompanyType(
		context.Background(),
		map[string]string{
			"name": "100000046",
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 1, res.Records[0].CompanyTypeID)
}
