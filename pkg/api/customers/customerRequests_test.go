package customers

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
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetCustomersResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
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

	defer srv.Close()

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

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, customersBulk.Status)

	expectedStatus := sharedCommon.StatusBulk{}
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

	defer srv.Close()

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

func TestAddCustomerRewardPoints(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "addCustomerRewardPoints", r.URL.Query().Get("request"))
		assert.Equal(t, "1232131", r.URL.Query().Get("customerID"))
		assert.Equal(t, "34456", r.URL.Query().Get("invoiceID"))
		assert.Equal(t, "11", r.URL.Query().Get("points"))

		resp := AddCustomerRewardPointsResponse{
			Status:                         sharedCommon.Status{ResponseStatus: "ok"},
			AddCustomerRewardPointsResults: []AddCustomerRewardPointsResult{{TransactionID: 999, CustomerID: 22}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"customerID": "1232131",
		"invoiceID":  "34456",
		"points":     "11",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.AddCustomerRewardPoints(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, int64(999), resp.TransactionID)
	assert.Equal(t, int64(22), resp.CustomerID)
}

func TestAddCustomerRewardPointsBulk(t *testing.T) {
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
				"requestName": "addCustomerRewardPoints",
				"customerID":  "123",
				"invoiceID":   "34456",
				"points":      "22",
			},
			{
				"requestName": "addCustomerRewardPoints",
				"customerID":  "124",
				"invoiceID":   "34457",
				"points":      "12",
			},
		})

		bulkResp := AddCustomerRewardPointsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []AddCustomerRewardPointsResponseBulkItem{
				{
					Status:  statusBulk,
					AddCustomerRewardPointsResults: []AddCustomerRewardPointsResult{{TransactionID: 3456}},
				},
				{
					Status:  statusBulk,
					AddCustomerRewardPointsResults: []AddCustomerRewardPointsResult{{TransactionID: 3457}},
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
			"customerID":  "123",
			"invoiceID":   "34456",
			"points":      "22",
		},
		{
			"customerID":  "124",
			"invoiceID":   "34457",
			"points":      "12",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.AddCustomerRewardPointsBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []AddCustomerRewardPointsResult{{TransactionID: 3456}}, bulkResp.BulkItems[0].AddCustomerRewardPointsResults)
	assert.Equal(t, []AddCustomerRewardPointsResult{{TransactionID: 3457}}, bulkResp.BulkItems[1].AddCustomerRewardPointsResults)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
