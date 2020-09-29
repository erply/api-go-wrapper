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
