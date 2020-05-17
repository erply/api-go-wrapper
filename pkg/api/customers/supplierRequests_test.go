package customers

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api/common"
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
		suppliers, status, err := cli.GetSuppliers(ctx, map[string]string{})
		assert.NoError(t, err)
		assert.Equal(t, "ok", status.ResponseStatus)
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
		statusBulk := common.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetSuppliersResponseBulk{
			Status: common.Status{ResponseStatus: "ok"},
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
	bulkFilters := []map[string]string{
		{
			"recordsOnPage": "2",
			"pageNo":        "1",
		},
		{
			"recordsOnPage": "2",
			"pageNo":        "2",
		},
	}
	bulkResp, err := suppliersClient.GetSuppliersBulk(context.Background(), bulkFilters, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	expectedStatusBulk := common.StatusBulk{}
	expectedStatusBulk.ResponseStatus = "ok"

	assert.Equal(t, []Supplier{
		{
			SupplierId: 123,
			FullName:   "Some Supplier123",
		},
		{
			SupplierId: 124,
			FullName:   "Some Supplier124",
		},
	}, bulkResp.BulkItems[0].Suppliers)

	assert.Equal(t, []Supplier{
		{
			SupplierId: 125,
			FullName:   "Some Supplier125",
		},
	}, bulkResp.BulkItems[1].Suppliers)

	assert.Equal(t, expectedStatusBulk, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatusBulk, bulkResp.BulkItems[1].Status)
}
