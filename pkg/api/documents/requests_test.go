package documents

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

func TestGetPurchaseDocumentsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetPurchaseDocumentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetPurchaseDocumentBulkItem{
				{
					Status: statusBulk,
					PurchaseDocuments: []PurchaseDocument{
						{
							ID: 123,
						},
						{
							ID: 124,
						},
					},
				},
				{
					Status: statusBulk,
					PurchaseDocuments: []PurchaseDocument{
						{
							ID: 125,
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

	bulkResp, err := cl.GetPurchaseDocumentsBulk(
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

	assert.Equal(t, []PurchaseDocument{
		{
			ID: 123,
		},
		{
			ID: 124,
		},
	}, bulkResp.BulkItems[0].PurchaseDocuments)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []PurchaseDocument{
		{
			ID: 125,
		},
	}, bulkResp.BulkItems[1].PurchaseDocuments)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestGetPurchaseDocuments(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := GetPurchaseDocumentsResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			PurchaseDocuments: []PurchaseDocument{
				{
					ID: 123,
				},
				{
					ID: 124,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	actualDocuments, err := cl.GetPurchaseDocuments(
		context.Background(),
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, []PurchaseDocument{
		{
			ID: 123,
		},
		{
			ID: 124,
		},
	}, actualDocuments)
}
