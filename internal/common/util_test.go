package common

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSendRequestBulk(t *testing.T) {
	calledTimes := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledTimes++

		AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
			"someKey":    "someValue",
		})

		AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName":   "getSuppliers",
				"recordsOnPage": "10",
				"pageNo":        "1",
			},
			{
				"requestName":   "getSuppliers",
				"recordsOnPage": "10",
				"pageNo":        "2",
			},
		})
	}))

	defer srv.Close()

	cli := NewClientWithURL(
		"somesess",
		"someclient",
		"",
		srv.URL,
		&http.Client{
			Timeout: 5 * time.Second,
		},
		nil,
	)

	resp, err := cli.SendRequestBulk(
		context.Background(),
		[]BulkInput{
			{
				MethodName: "getSuppliers",
				Filters: map[string]interface{}{
					"recordsOnPage": "10",
					"pageNo":        "1",
				},
			},
			{
				MethodName: "getSuppliers",
				Filters: map[string]interface{}{
					"recordsOnPage": "10",
					"pageNo":        "2",
				},
			},
		},
		map[string]string{"someKey": "someValue"},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, calledTimes)
}
