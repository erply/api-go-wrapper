package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecurringBilling(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			common.AssertFormValues(t, r, map[string]interface{}{
				"clientCode": "someclient",
				"sessionKey": "somesess",
			})
			_, err := w.Write([]byte(`{
			"status": {
				"request": "processRecurringBilling",
				"requestUnixTime": 1728482531,
				"responseStatus": "ok",
				"errorCode": 0,
				"generationTime": 0.3931558132171630859375,
				"recordsTotal": 0,
				"recordsInResponse": 0
			},
			"records": [
				{
					"processedInvoices": [
						{
							"id": 89,
							"created": true,
							"updated": true
						}
					]
				}
			]
		}`))
			assert.NoError(t, err)
		}))
		defer srv.Close()

		cli := common.NewClient("somesess", "someclient", "", nil, nil)
		cli.Url = srv.URL
		cl := NewClient(cli)
		bulkResp, err := cl.ProcessRecurringBilling(
			context.Background(),
			map[string]string{
				"billingStatementIDs": "1,2",
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(bulkResp))
		assert.Equal(t, 89, bulkResp[0].ID)
		assert.Equal(t, true, bulkResp[0].Created)
		assert.Equal(t, true, bulkResp[0].Updated)
	})
	t.Run("empty", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			common.AssertFormValues(t, r, map[string]interface{}{
				"clientCode": "someclient",
				"sessionKey": "somesess",
			})
			_, err := w.Write([]byte(`{
				"status": {
					"request": "processRecurringBilling",
					"requestUnixTime": 1728485864,
					"responseStatus": "ok",
					"errorCode": 0,
					"generationTime": 0.1155679225921630859375,
					"recordsTotal": 0,
					"recordsInResponse": 0
				},
				"records": [
					{
						"processedInvoices": []
					}
				]
			}`))
			assert.NoError(t, err)
		}))
		defer srv.Close()

		cli := common.NewClient("somesess", "someclient", "", nil, nil)
		cli.Url = srv.URL
		cl := NewClient(cli)
		bulkResp, err := cl.ProcessRecurringBilling(
			context.Background(),
			map[string]string{
				"billingStatementIDs": "1,2",
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(bulkResp))
	})
	t.Run("error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			common.AssertFormValues(t, r, map[string]interface{}{
				"clientCode": "someclient",
				"sessionKey": "somesess",
			})
			_, err := w.Write([]byte(`{
				"status": {
					"request": "processRecurringBilling",
					"requestUnixTime": 1728489814,
					"responseStatus": "error",
					"errorCode": 1016,
					"errorField": "billingStatementIDs",
					"generationTime": 0.0900490283966064453125,
					"recordsTotal": 0,
					"recordsInResponse": 0
				},
				"records": []
			}`))
			assert.NoError(t, err)
		}))
		defer srv.Close()

		cli := common.NewClient("somesess", "someclient", "", nil, nil)
		cli.Url = srv.URL
		cl := NewClient(cli)
		bulkResp, err := cl.ProcessRecurringBilling(
			context.Background(),
			map[string]string{
				"billingStatementIDs": "a",
			},
		)
		assert.Error(t, err)
		assert.Equal(t, "ERPLY API: processRecurringBilling: error, status: [1016] Invalid value.(Attribute \"errorField\" indicates the field that contains an invalid value.), error field: billingStatementIDs, code: 1016", err.Error())
		assert.Equal(t, 0, len(bulkResp))
	})
}
