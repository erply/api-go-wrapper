package customers

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
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
		statusBulk := common.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := getSuppliersResponseBulk{
			Status: common.Status{ResponseStatus:    "ok"},
			BulkItems: []getSuppliersResponseBulkItem{
				{
					Status:    statusBulk,
					Suppliers: []Supplier{
						{
							SupplierId:          123,
							FullName:            "Some Supplier123",
						},
						{
							SupplierId:          124,
							FullName:            "Some Supplier124",
						},
					},
				},
				{
					Status:    statusBulk,
					Suppliers: []Supplier{
						{
							SupplierId:          125,
							FullName:            "Some Supplier125",
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
	ctx := context.WithValue(context.Background(), "bulk", []map[string]string{
		{
			"recordsOnPage": "2",
			"pageNo":"1",
		},
		{
			"recordsOnPage": "2",
			"pageNo":"2",
		},
	})
	suppliers, err := suppliersClient.GetSuppliers(ctx, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []Supplier{
		{
			SupplierId:          123,
			FullName:            "Some Supplier123",
		},
		{
			SupplierId:          124,
			FullName:            "Some Supplier124",
		},
		{
			SupplierId:          125,
			FullName:            "Some Supplier125",
		},
	}, suppliers)
}
