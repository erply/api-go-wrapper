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
func TestCustomerManager(t *testing.T) {
	const (
		//fill your data here
		sk                    = ""
		cc                    = ""
		someAvailableUsername = ""
	)
	//and here
	var (
		testingCustomer = &CustomerRequest{
			CompanyName: "",
			Username:    "",
			Password:    "",
		}
		ctx = context.Background()
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test get customers", func(t *testing.T) {
		resp, err := cli.GetCustomers(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test post customer", func(t *testing.T) {

		params := map[string]string{
			"companyName": testingCustomer.CompanyName,
		}
		params["username"] = testingCustomer.Username
		params["password"] = testingCustomer.Password
		report, err := cli.SaveCustomer(ctx, params)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(report)
	})
	t.Run("test verifyCustomerUser", func(t *testing.T) {

		isAvailable, err := cli.VerifyCustomerUser(ctx, testingCustomer.Username, testingCustomer.Password)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(isAvailable)
	})
	t.Run("test validation of the username", func(t *testing.T) {
		isAvailable, err := cli.ValidateCustomerUsername(ctx, someAvailableUsername)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(isAvailable)
	})
}

func TestGetCustomersBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetCustomersResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []GetCustomersResponseBulkItem{
				{
					Status: statusBulk,
					Customers: []Customer{
						{
							ID:          123,
							CompanyName: "Customer 123",
						},
						{
							ID:          124,
							CompanyName: "Customer 124",
						},
					},
				},
				{
					Status: statusBulk,
					Customers: []Customer{
						{
							ID:          125,
							CompanyName: "Customer 125",
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

	customerClient := NewClient(cli)

	customersBulk, err := customerClient.GetCustomersBulk(
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

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, customersBulk.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, Customers{
		{
			ID:          123,
			CompanyName: "Customer 123",
		},
		{
			ID:          124,
			CompanyName: "Customer 124",
		},
	}, customersBulk.BulkItems[0].Customers)

	assert.Equal(t, expectedStatus, customersBulk.BulkItems[0].Status)

	assert.Equal(t, Customers{
		{
			ID:          125,
			CompanyName: "Customer 125",
		},
	}, customersBulk.BulkItems[1].Customers)
	assert.Equal(t, expectedStatus, customersBulk.BulkItems[1].Status)
}

func TestGetCustomersBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk value`))
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	customersClient := NewClient(cli)

	_, err := customersClient.GetCustomersBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetCustomersResponseBulk from 'some junk value': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}
