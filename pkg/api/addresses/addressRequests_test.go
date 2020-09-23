package addresses

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
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
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		supplierResp := GetAddressesResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
			BulkItems: []GetAddressesResponseBulkItem{
				{
					Status: statusBulk,
					Addresses: common2.Addresses{
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
					Addresses: common2.Addresses{
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

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, suppliersBulk.Status)

	expectedStatus := common2.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Equal(t, common2.Addresses{
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

	assert.Equal(t, common2.Addresses{
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
		statusBulk := common2.StatusBulk{}
		statusBulk.ResponseStatus = "ok"
		resp := SaveAddressesResponseBulk{
			Status: common2.Status{ResponseStatus: "ok"},
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

	assert.Equal(t, common2.Status{ResponseStatus: "ok"}, saveResp.Status)

	expectedStatus := common2.StatusBulk{}
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
