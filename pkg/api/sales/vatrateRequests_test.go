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
func TestVatRateManager(t *testing.T) {
	const (
		//fill your data here
		sk        = ""
		cc        = ""
		vatRateID = ""
	)
	var (
		ctx = context.Background()
	)
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))

	resp, err := cli.GetVatRates(ctx, map[string]string{
		"searchAttributeName":  "id",
		"searchAttributeValue": vatRateID,
		"active":               "1",
	})

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestSaveVatRate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "saveVatRate", r.URL.Query().Get("request"))
		assert.Equal(t, "ID123", r.URL.Query().Get("vatRateID"))
		assert.Equal(t, "VatName", r.URL.Query().Get("name"))
		assert.Equal(t, "0.19", r.URL.Query().Get("rate"))
		assert.Equal(t, "vatCode123", r.URL.Query().Get("code"))

		resp := SaveVatRateResultResponse{
			Status:            sharedCommon.Status{ResponseStatus: "ok"},
			SaveVatRateResult: []SaveVatRateResult{{VatRateID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"vatRateID": "ID123",
		"name":      "VatName",
		"rate":      "0.19",
		"code":      "vatCode123",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SaveVatRate(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, resp.VatRateID)
}

func TestSaveVatRateBulk(t *testing.T) {
	bulkRequestWasMade := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bulkRequestWasMade = true
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"name":        "VatName1",
				"rate":        "0.19",
				"code":        "vatCode123",
				"requestName": "saveVatRate",
			},
			{
				"name":        "VatName2",
				"rate":        "0.18",
				"code":        "vatCode124",
				"requestName": "saveVatRate",
			},
		})

		bulkResp := SaveVatRateResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveVatRateBulkItem{
				{
					Status: statusBulk,
					Records: []SaveVatRateResult{
						{
							VatRateID: 999,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveVatRateResult{
						{
							VatRateID: 998,
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
			"name": "VatName1",
			"rate": "0.19",
			"code": "vatCode123",
		},
		{
			"name": "VatName2",
			"rate": "0.18",
			"code": "vatCode124",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveVatRateBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.True(t, bulkRequestWasMade)

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, 999, bulkResp.BulkItems[0].Records[0].VatRateID)
	assert.Equal(t, 998, bulkResp.BulkItems[1].Records[0].VatRateID)
}

func TestSaveVatRateComponent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "someclient", r.URL.Query().Get("clientCode"))
		assert.Equal(t, "somesess", r.URL.Query().Get("sessionKey"))
		assert.Equal(t, "saveVatRateComponent", r.URL.Query().Get("request"))
		assert.Equal(t, "ID123", r.URL.Query().Get("vatRateComponentID"))
		assert.Equal(t, "#2333", r.URL.Query().Get("vatRateID"))
		assert.Equal(t, "Some name", r.URL.Query().Get("name"))
		assert.Equal(t, "8.76", r.URL.Query().Get("rate"))

		resp := SaveVatRateComponentResultResponse{
			Status:                     sharedCommon.Status{ResponseStatus: "ok"},
			SaveVatRateComponentResult: []SaveVatRateComponentResult{{VatRateComponentID: 999}},
		}
		jsonRaw, err := json.Marshal(resp)
		assert.NoError(t, err)

		_, err = w.Write(jsonRaw)
		assert.NoError(t, err)
	}))

	defer srv.Close()

	inpt := map[string]string{
		"vatRateComponentID": "ID123",
		"vatRateID":          "#2333",
		"name":               "Some name",
		"rate":               "8.76",
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	resp, err := cl.SaveVatRateComponent(context.Background(), inpt)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, 999, resp.VatRateComponentID)
}

func TestSaveVatRateComponentBulk(t *testing.T) {
	bulkRequestWasMade := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bulkRequestWasMade = true
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"vatRateComponentID": "ID123",
				"vatRateID":          "#2333",
				"name":               "Some name",
				"rate":               "8.76",
				"requestName":        "saveVatRateComponent",
			},
			{
				"vatRateComponentID": "ID124",
				"vatRateID":          "#2334",
				"name":               "Some name 2",
				"rate":               "8.77",
				"requestName":        "saveVatRateComponent",
			},
		})

		bulkResp := SaveVatRateComponentResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []SaveVatRateComponentBulkItem{
				{
					Status: statusBulk,
					Records: []SaveVatRateComponentResult{
						{
							VatRateComponentID: 999,
						},
					},
				},
				{
					Status: statusBulk,
					Records: []SaveVatRateComponentResult{
						{
							VatRateComponentID: 998,
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
			"vatRateComponentID": "ID123",
			"vatRateID":          "#2333",
			"name":               "Some name",
			"rate":               "8.76",
		},
		{
			"vatRateComponentID": "ID124",
			"vatRateID":          "#2334",
			"name":               "Some name 2",
			"rate":               "8.77",
		},
	}

	cli := common.NewClient("somesess", "someclient", "", nil, nil)
	cli.Url = srv.URL

	cl := NewClient(cli)

	bulkResp, err := cl.SaveVatRateComponentBulk(context.Background(), inpt, map[string]string{})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.True(t, bulkRequestWasMade)

	assert.Equal(t, sharedCommon.Status{ResponseStatus: "ok"}, bulkResp.Status)

	expectedStatus := sharedCommon.StatusBulk{}
	expectedStatus.ResponseStatus = "ok"

	assert.Len(t, bulkResp.BulkItems, 2)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)
	assert.Equal(t, 999, bulkResp.BulkItems[0].Records[0].VatRateComponentID)
	assert.Equal(t, 998, bulkResp.BulkItems[1].Records[0].VatRateComponentID)
}

func TestGetVatRatesBulk(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusBulk := sharedCommon.StatusBulk{}
		statusBulk.ResponseStatus = "ok"

		common.AssertFormValues(t, r, map[string]interface{}{
			"clientCode": "someclient",
			"sessionKey": "somesess",
		})

		common.AssertRequestBulk(t, r, []map[string]interface{}{
			{
				"requestName": "getVatRates",
				"id":          "1",
			},
			{
				"requestName": "getVatRates",
				"id":          "2",
			},
			{
				"requestName": "getVatRates",
				"id":          "3",
			},
		})

		bulkResp := GetVatRatesResponseBulk{
			Status: sharedCommon.Status{ResponseStatus: "ok"},
			BulkItems: []GetVatRatesBulkItem{
				{
					Status: statusBulk,
					VatRates: []VatRate{
						{
							ID:   "1",
							Name: "Name 1",
						},
					},
				},
				{
					Status: statusBulk,
					VatRates: []VatRate{
						{
							ID:   "2",
							Name: "Name 2",
						},
					},
				},
				{
					Status: statusBulk,
					VatRates: []VatRate{
						{
							ID:   "3",
							Name: "Name 3",
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

	bulkResp, err := cl.GetVatRatesBulk(
		context.Background(),
		[]map[string]interface{}{
			{
				"id": "1",
			},
			{
				"id": "2",
			},
			{
				"id": "3",
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

	assert.Equal(t, []VatRate{
		{
			ID: "1",
			Name: "Name 1",
		},
	}, bulkResp.BulkItems[0].VatRates)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[0].Status)

	assert.Equal(t, []VatRate{
		{
			ID: "2",
			Name: "Name 2",
		},
	}, bulkResp.BulkItems[1].VatRates)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[1].Status)

	assert.Equal(t, []VatRate{
		{
			ID: "3",
			Name: "Name 3",
		},
	}, bulkResp.BulkItems[2].VatRates)

	assert.Equal(t, expectedStatus, bulkResp.BulkItems[2].Status)
}
