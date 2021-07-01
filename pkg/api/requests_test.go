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
)

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
