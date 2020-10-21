package addresses

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
func TestAddressManager(t *testing.T) {
	const (
		//fill your data here
		sk      = ""
		cc      = ""
		ownerID = ""
	)
	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	resp, err := cli.GetAddresses(ctx, map[string]string{
		"ownerID": ownerID,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
	t.Run("test save address", func(t *testing.T) {
		filters := map[string]string{
			"ownerID": "", //put your value here
			"typeID":  "", //put your value here
		}
		resp, err := cli.SaveAddress(ctx, filters)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}

func TestGetAddressesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetAddressesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetAddressesResponseBulkItem{
				{
					Status: statusBulk,
					Addresses: sharedCommon.Addresses{
						{
							AddressID: 123,
							Address:   "Some Address 123",
						},
						{
							AddressID: 124,
							Address:   "Some Address 124",
						},
					},
				},
				{
					Status: statusBulk,
					Addresses: sharedCommon.Addresses{
						{
							AddressID: 125,
							Address:   "Some Address 125",
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(supplierResp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	addressClient := NewClient(cli)

	suppliersBulk, err := addressClient.GetAddressesBulk(
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

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, suppliersBulk.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, sharedCommon.Addresses{
		{
			AddressID: 123,
			Address:   "Some Address 123",
		},
		{
			AddressID: 124,
			Address:   "Some Address 124",
		},
	}, suppliersBulk.BulkItems[0].Addresses)

	assert.Equal(t, expectedStatus, suppliersBulk.BulkItems[0].Status)

	assert.Equal(t, sharedCommon.Addresses{
		{
			AddressID: 125,
			Address:   "Some Address 125",
		},
	}, suppliersBulk.BulkItems[1].Addresses)
	assert.Equal(t, expectedStatus, suppliersBulk.BulkItems[1].Status)
}

func TestGetAddressesBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk`))
		assert.NoError(t, err)
	}))
	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	addressClient := NewClient(cli)

	_, err := addressClient.GetAddressesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"recordsOnPage": 1,
				"pageNo":        1,
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal GetAddressesResponseBulk from 'some junk': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestSaveAddressesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := SaveAddressesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveAddressesResponseBulkItem{
				{
					Status: statusBulk,
					Records: []SaveAddressResp{
						{
							AddressID: 123,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveAddressResp{
						{
							AddressID: 124,
						},
					},
				},
			},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	addressesClient := NewClient(cli)

	saveResp, err := addressesClient.SaveAddressesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"addressID": 123,
				"street":    "Some street 123",
			},
			{
				"street": "Some street new",
			},
		},
		map[string]string{},
	)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, saveResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, saveResp.BulkItems, 2)

	assert.Equal(t, []SaveAddressResp{
		{
			AddressID: 123,
		},
	}, saveResp.BulkItems[0].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[0].Status)

	assert.Equal(t, []SaveAddressResp{
		{
			AddressID: 124,
		},
	}, saveResp.BulkItems[1].Records)

	assert.Equal(t, expectedStatus, saveResp.BulkItems[1].Status)
}

func TestSaveAddressesBulkResponseFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`some junk`))
		assert.NoError(t, err)
	}))

	defer srv.Close()

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	addressesClient := NewClient(cli)

	_, err := addressesClient.SaveAddressesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"addressID": 123,
				"street":    "Some street 123",
			},
		},
		map[string]string{},
	)
	assert.EqualError(t, err, `ERPLY API: failed to unmarshal SaveAddressesResponseBulk from 'some junk': invalid character 's' looking for beginning of value`)
	if err == nil {
		return
	}
}

func TestDeleteAddresses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "deleteAddress", r.URL.Query().Get("request"))
		assert.Equal(t, "2223", r.URL.Query().Get("addressID"))

		resp := DeleteAddressResponse{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"addressID":         "2223",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	err := cl.DeleteAddress(context.Background(), inpt)
	assert.NoError(t, err)
}

func TestDeleteAddressesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		err := r.ParseForm()
		assert.NoError(t, err)
		if err != nil {
			return
		}

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		bulkRequestsRaw := r.FormValue("requests")

		bulkRequests := []map[string]interface{}{}
		err = json.Unmarshal([]byte(bulkRequestsRaw), &bulkRequests)
		if err != nil {
			return
		}
		expectedBulkRequests := []map[string]interface{}{
			{
				"requestName":         "deleteAddress",
				"addressID": "3456",
			},
			{
				"requestName":         "deleteAddress",
				"addressID": "3457",
			},
		}
		assert.Equal(t, expectedBulkRequests, bulkRequests)

		bulkResp := DeleteAddressResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []DeleteAddressBulkItem{
				{
					Status:  statusBulk,
				},
				{
					Status:  statusBulk,
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
			"addressID": "3456",
		},
		{
			"addressID": "3457",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.DeleteAddressBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)
}
