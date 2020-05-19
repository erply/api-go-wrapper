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
		[]map[string]string{
			{
				"recordsOnPage": "2",
				"pageNo":        "1",
			},
			{
				"recordsOnPage": "2",
				"pageNo":        "2",
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
							SupplierID: 123,
							AlreadyExists: 0,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveSupplierResp{
						{
							SupplierID: 124,
							AlreadyExists: 1,
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
		[]Supplier{
			{
				SupplierId: 123,
				FullName:  "Some name",
			},
			{
				FullName:  "Some other name",
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
			SupplierID: 123,
			AlreadyExists: 0,
		},
	}, saveResp.BulkItems[0].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[0].Status)

	assert.Equal(t, []SaveSupplierResp{
		{
			SupplierID: 124,
			AlreadyExists: 1,
		},
	}, saveResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[1].Status)
}
