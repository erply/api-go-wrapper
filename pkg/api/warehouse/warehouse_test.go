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
	resp, err := cli.GetWarehouses(context.Background())
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

	assert.Equal(t, Warehouses{
		{
			WarehouseID: "123",
		},
		{
			WarehouseID: "124",
		},
	}, bulkResp.BulkItems[0].Warehouses)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, Warehouses{
		{
			WarehouseID: "125",
		},
	}, bulkResp.BulkItems[1].Warehouses)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
