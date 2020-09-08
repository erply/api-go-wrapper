package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/stretchr/testify/assert"
	http2 "net/http"
	"net/http/httptest"
	"testing"
)

//works
func TestSalesDocuments(t *testing.T) {
	const (
		//fill your data here
		sk              = ""
		cc              = ""
		invoiceNoToSave = ""
		supplierID      = ""
		vatrateID       = ""
		amount          = ""
		price           = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test get sales doc", func(t *testing.T) {
		saleDocs, err := cli.GetSalesDocuments(ctx, map[string]string{
			"id": "",
		})
		if err != nil {
			t.Error(err)
			return
		}

		for _, r := range saleDocs[0].InvoiceRows {
			t.Logf("row's code2: %s", r.Code2)
			t.Logf(r.StableRowID)
		}
	})

	t.Run("test save purchase", func(t *testing.T) {
		resp, err := cli.SavePurchaseDocument(ctx, map[string]string{
			"currencyCode": "EUR",
			"no":           invoiceNoToSave,
			"supplierID":   supplierID,
			"vatrateID":    vatrateID,
			"amount":       amount,
			"price":        price,
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	t.Run("test save sales doc", func(t *testing.T) {
		reports, err := cli.SaveSalesDocument(ctx, map[string]string{
			"id":         "57",
			"productID1": "4",
			"amount1":    "2",
			"price1":     "20",
		})
		if err != nil {
			t.Error(err)
			return
		}
		for _, r := range reports {
			t.Log(r.InvoiceID)
		}
	})
}

func TestGetPurchaseDocumentsBulk(t *testing.T) {
	srv := httptest.NewServer(http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		bulkResp := GetSaleDocumentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetSaleDocumentBulkItem{
				{
					Status: statusBulk,
					SaleDocuments: []SaleDocument{
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
					SaleDocuments: []SaleDocument{
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

	bulkResp, err := cl.GetSalesDocumentsBulk(
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

	assert.Equal(t, []SaleDocument{
		{
			ID: 123,
		},
		{
			ID: 124,
		},
	}, bulkResp.BulkItems[0].SaleDocuments)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []SaleDocument{
		{
			ID: 125,
		},
	}, bulkResp.BulkItems[1].SaleDocuments)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
