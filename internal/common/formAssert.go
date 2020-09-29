package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func AssertFormValues(t *testing.T, r *http.Request, expectedValues map[string]interface{}) {
	for expectedKey, expectedValue := range expectedValues{
		actualValue := r.FormValue(expectedKey)
		assert.Equal(t, expectedValue, actualValue)
	}
}

func AssertRequestBulk(t *testing.T, r *http.Request, expectedBulkRequest []map[string]interface{}) {
	requestsRaw := r.FormValue("requests")
	decodedValue, err := url.QueryUnescape(requestsRaw)
	assert.NoError(t, err)

	var actualBulkRequest []map[string]interface{}
	err = json.Unmarshal([]byte(decodedValue), &actualBulkRequest)
	assert.NoError(t, err)

	assert.Equal(t, expectedBulkRequest, actualBulkRequest)
}
