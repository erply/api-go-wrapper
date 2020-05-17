package common

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestSendRequestBulk(t *testing.T) {
	calledTimes := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledTimes++
		actualClientCode := r.FormValue("clientCode")
		assert.Equal(t, "someclient", actualClientCode)

		sessionKey := r.FormValue("sessionKey")
		assert.Equal(t, "somesess", sessionKey)

		requestsRaw := r.FormValue("requests")
		decodedValue, err := url.QueryUnescape(requestsRaw)
		assert.NoError(t, err)

		customKey := r.FormValue("someKey")
		assert.Equal(t, "someValue", customKey)

		var actualMapRequest []map[string]string
		err = json.Unmarshal([]byte(decodedValue), &actualMapRequest)
		assert.NoError(t, err)

		assert.Equal(
			t,
			[]map[string]string{
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
			},
			actualMapRequest,
		)
	}))

	defer srv.Close()

	cli := &Client{
		Url: srv.URL,
		httpClient:  &http.Client{
			Timeout: 5 * time.Second,
		},
		sessionKey:  "somesess",
		clientCode:   "someclient",
		partnerKey:  "",
	}

	resp, err := cli.SendRequestBulk(
		context.Background(),
		[]BulkInput{
			{
				MethodName: "getSuppliers",
				Filters: map[string]string{
					"recordsOnPage": "10",
					"pageNo":        "1",
				},
			},
			{
				MethodName: "getSuppliers",
				Filters: map[string]string{
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

