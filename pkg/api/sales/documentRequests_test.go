package sales

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

	defer srv.Close()

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

	assert.Equal(t, 123, bulkResp.BulkItems[0].SaleDocuments[0].ID)
	assert.Equal(t, 124, bulkResp.BulkItems[0].SaleDocuments[1].ID)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, 125, bulkResp.BulkItems[1].SaleDocuments[0].ID)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}

func TestSavePurchaseDocument(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := SavePurchaseDocumentResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			ImportReports: []PurchaseDocImportReport{
				{
					InvoiceID: 123,
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode":  "someclient",
			"sessionKey":  "somesess",
			"request":     "savePurchaseDocument",
			"warehouseID": "1",
			"no":          "123",
		})

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"warehouseID": "1",
		"no":          "123",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	expectedRes := PurchaseDocImportReports{
		{
			InvoiceID: 123,
		},
	}
	actualRes, err := cl.SavePurchaseDocument(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, expectedRes, actualRes)
}

func TestSavePurchaseDocumentBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"warehouseID":  "12",
				"currencyCode": "code 123",
				"no":           "123",
				"requestName":  "savePurchaseDocument",
			},
			{
				"warehouseID":  "12",
				"currencyCode": "code 124",
				"no":           "124",
				"requestName":  "savePurchaseDocument",
			},
		})

		bulkResp := SavePurchaseDocumentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SavePurchaseDocumentBulkItem{
				{
					Status: statusBulk,
					Records: PurchaseDocImportReports{
						{
							InvoiceID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					Records: PurchaseDocImportReports{
						{
							InvoiceID: 124,
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

	inpt := []map[string]interface{}{
		{
			"warehouseID":  "12",
			"currencyCode": "code 123",
			"no":           "123",
		},
		{
			"warehouseID":  "12",
			"currencyCode": "code 124",
			"no":           "124",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SavePurchaseDocumentBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, 123, bulkResp.BulkItems[0].Records[0].InvoiceID)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, 124, bulkResp.BulkItems[1].Records[0].InvoiceID)
}

func TestSaveSalesDocumentBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"warehouseID":  "12",
				"currencyCode": "code 123",
				"type":         SaleDocumentTypeCASHINVOICE,
				"requestName":  "saveSalesDocument",
			},
			{
				"warehouseID":  "13",
				"currencyCode": "code 124",
				"type":         SaleDocumentTypeCASHINVOICE,
				"requestName":  "saveSalesDocument",
			},
		})

		bulkResp := SaveSalesDocumentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveSalesDocumentBulkItem{
				{
					Status: statusBulk,
					Records: SaleDocImportReports{
						{
							InvoiceID: json.Number("123"),
						},
					},
				},
				{
					Status: statusBulk,
					Records: SaleDocImportReports{
						{
							InvoiceID: json.Number("124"),
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

	inpt := []map[string]interface{}{
		{
			"warehouseID":  "12",
			"currencyCode": "code 123",
			"type":         "CASHINVOICE",
		},
		{
			"warehouseID":  "13",
			"currencyCode": "code 124",
			"type":         "CASHINVOICE",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveSalesDocumentBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Len(t, bulkResp.BulkItems[0].Records, 1)
	assert.Equal(t, json.Number("123"), bulkResp.BulkItems[0].Records[0].InvoiceID)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
	assert.Len(t, bulkResp.BulkItems[1].Records, 1)
	assert.Equal(t, json.Number("124"), bulkResp.BulkItems[1].Records[0].InvoiceID)
}
