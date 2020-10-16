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

//works
func TestPaymentManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""

		documentID   = ""
		paymentType  = ""
		currencyCode = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test save payment", func(t *testing.T) {
		params := map[string]string{
			"documentID":   documentID,
			"type":         paymentType,
			"currencyCode": currencyCode,
		}
		paymentID, err := cli.SavePayment(ctx, params)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(paymentID)
	})
	t.Run("test get payments", func(t *testing.T) {
		resp, err := cli.GetPayments(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}

func TestGetPaymentsBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "getPayments",
				"paymentID":   "1",
			},
			{
				"requestName": "getPayments",
				"paymentID":   "2",
			},
			{
				"requestName": "getPayments",
				"paymentID":   "3",
			},
		})

		bulkResp := GetPaymentsResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetPaymentsBulkItem{
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 1,
							Type:       "Some type",
						},
					},
				},
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 2,
							Type:       "Some type",
						},
					},
				},
				{
					Status: statusBulk,
					PaymentInfos: []PaymentInfo{
						{
							DocumentID: 3,
							Type:       "Some type",
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

	bulkResp, err := cl.GetPaymentsBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"paymentID": "1",
			},
			{
				"paymentID": "2",
			},
			{
				"paymentID": "3",
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

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 1,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[0].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 2,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[1].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []PaymentInfo{
		{
			DocumentID: 3,
			Type:       "Some type",
		},
	}, bulkResp.BulkItems[2].PaymentInfos)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}
