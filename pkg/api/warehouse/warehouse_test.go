package warehouse

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
func TestErplyClient_GetWarehouses(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	resp, err := cli.GetWarehouses(context.Background(), map[string]string{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestGetWarehousesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetWarehousesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetWarehousesBulkItem{
				{
					Status: statusBulk,
					Warehouses: Warehouses{
						{
							WarehouseID: "123",
						},
						{
							WarehouseID: "124",
						},
					},
				},
				{
					Status: statusBulk,
					Warehouses: Warehouses{
						{
							WarehouseID: "125",
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

	cl := NewClient(cli)

	bulkResp, err := cl.GetWarehousesBulk(
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

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, []string{"123", "124"}, collectWarehouseIDs(bulkResp, 0))

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []string{"125"}, collectWarehouseIDs(bulkResp, 1))
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func collectWarehouseIDs(resp GetWarehousesResponseBulk, index int) []string {
	res := make([]string, 0)
	for _, warehouse := range resp.BulkItems[index].Warehouses {
		res = append(res, warehouse.WarehouseID)
	}

	return res
}

func TestSaveWarehouse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"request":    "saveWarehouse",
			"code":       "12345",
			"name":       "someWarehouse123",
		})

		resp := SaveWarehouseResponse{
			Status:  sharedCommon.Status{ResponseStatus: "ok"},
			Results: []SaveWarehouseResult{{WarehouseID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"name": "someWarehouse123",
		"code": "12345",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SaveWarehouse(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, resp.WarehouseID)
}

func TestSaveWarehouseBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName": "saveWarehouse",
				"name":        "www1",
				"code":        "1",
			},
			{
				"requestName": "saveWarehouse",
				"name":        "www2",
				"code":        "2",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := SaveWarehouseResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveWarehouseBulkItem{
				{
					Status:  statusBulk,
					Results: []SaveWarehouseResult{{WarehouseID: 3456}},
				},
				{
					Status:  statusBulk,
					Results: []SaveWarehouseResult{{WarehouseID: 3457}},
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
			"name":        "www1",
			"code":        "1",
		},
		{
			"name":        "www2",
			"code":        "2",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveWarehouseBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []SaveWarehouseResult{{WarehouseID: 3456}}, bulkResp.BulkItems[0].Results)
	assert.Equal(t, []SaveWarehouseResult{{WarehouseID: 3457}}, bulkResp.BulkItems[1].Results)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
