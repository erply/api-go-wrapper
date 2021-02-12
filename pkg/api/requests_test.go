package api

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

//here the "works" indicator will be under each test separately
func TestApiRequests(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("test GetUserRights", func(t *testing.T) {
		resp, err := cli.GetUserRights(context.Background(), map[string]string{
			"getCurrentUser": "1",
		})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetEmployees", func(t *testing.T) {
		resp, err := cli.GetEmployees(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetBusinessAreas", func(t *testing.T) {
		resp, err := cli.GetBusinessAreas(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test GetCurrencies", func(t *testing.T) {
		resp, err := cli.GetCurrencies(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test LogProcessingOfCustomerData", func(t *testing.T) {
		assert.NoError(t, cli.LogProcessingOfCustomerData(context.Background(), map[string]string{}))
	})

	t.Run("test GetCountries", func(t *testing.T) {
		resp, err := cli.GetCountries(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(resp)
	})

	t.Run("test SaveEvent", func(t *testing.T) {
		layout := "2006-01-02 15:04:05"
		d := time.Now
		assert.NoError(t, err)
		eventID, err := cli.SaveEvent(context.Background(), map[string]string{
			"startTime": d().Format(layout),
			"endTime":   d().Format(layout),
			"typeID":    "APPOINTMENT",
		})
		assert.NoError(t, err)
		assert.NotEqual(t, 0, eventID)
	})

	t.Run("test GetEvents", func(t *testing.T) {
		events, err := cli.GetEvents(context.Background(), map[string]string{})
		assert.NoError(t, err)
		t.Log(events)
	})
}

func TestGetEmployeesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "getEmployees",
				"employeeID":  "1",
			},
			{
				"requestName": "getEmployees",
				"employeeID":  "2",
			},
			{
				"requestName": "getEmployees",
				"employeeID":  "3",
			},
		})

		bulkResp := GetEmployeesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetEmployeesResponseBulkItem{
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "1",
							FullName:   "Name 1",
						},
					},
				},
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "2",
							FullName:   "Name 2",
						},
					},
				},
				{
					Status: statusBulk,
					Employees: []Employee{
						{
							EmployeeID: "3",
							FullName:   "Name 3",
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

	c := newErplyClient(cli)

	bulkResp, err := c.GetEmployeesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"employeeID": "1",
			},
			{
				"employeeID": "2",
			},
			{
				"employeeID": "3",
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

	assert.Len(t, bulkResp.BulkItems, 3)

	assert.Equal(t, []Employee{
		{
			EmployeeID: "1",
			FullName:   "Name 1",
		},
	}, bulkResp.BulkItems[0].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []Employee{
		{
			EmployeeID: "2",
			FullName:   "Name 2",
		},
	}, bulkResp.BulkItems[1].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []Employee{
		{
			EmployeeID: "3",
			FullName:   "Name 3",
		},
	}, bulkResp.BulkItems[2].Employees)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}
