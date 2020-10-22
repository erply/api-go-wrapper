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

func TestSaveInventoryRegistration(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":              "someclient",
			"sessionKey":              "somesess",
			"request":                 "saveInventoryRegistration",
			"inventoryRegistrationID": "12345",
			"creatorID":               "2234",
		})

		resp := SaveInventoryRegistrationResponse{
			Status:  sharedCommon.Status{ResponseStatus: "ok"},
			Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"inventoryRegistrationID": "12345",
		"creatorID":               "2234",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	regisrationID, err := cl.SaveInventoryRegistration(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, regisrationID)
}

func TestSaveInventoryRegistrationBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"inventoryRegistrationID": "123",
				"creatorID":               "2",
				"requestName":             "saveInventoryRegistration",
			},
			{
				"warehouseID":   "334",
				"stocktakingID": "233",
				"supplierID":    "455",
				"requestName":   "saveInventoryRegistration",
			},
		})

		bulkResp := SaveInventoryRegistrationResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveInventoryRegistrationBulkItem{
				{
					Status:  statusBulk,
					Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3456}},
				},
				{
					Status:  statusBulk,
					Results: []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3457}},
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
			"inventoryRegistrationID": "123",
			"creatorID":               "2",
		},
		{
			"warehouseID":   "334",
			"stocktakingID": "233",
			"supplierID":    "455",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveInventoryRegistrationBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)
	assert.Equal(t, []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3456}}, bulkResp.BulkItems[0].Results)
	assert.Equal(t, []SaveInventoryRegistrationResult{{InventoryRegistrationID: 3457}}, bulkResp.BulkItems[1].Results)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
