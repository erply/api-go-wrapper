package warehouse

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

/*
type (
	GetWarehousesResponse struct {
		Status     common.Status `json:"status"`
		Warehouses Warehouses `json:"records"`
	}

	Warehouse struct {
		WarehouseID            string `json:"warehouseID"`
		PricelistID            string `json:"pricelistID"`
		PricelistID2           string `json:"pricelistID2"`
		PricelistID3           string `json:"pricelistID3"`
		PricelistID4           string `json:"pricelistID4"`
		PricelistID5           string `json:"pricelistID5"`
		Name                   string `json:"name"`
		Code                   string `json:"code"`
		AddressID              int    `json:"addressID"`
		Address                string `json:"address"`
		Street                 string `json:"street"`
		Address2               string `json:"address2"`
		City                   string `json:"city"`
		State                  string `json:"state"`
		Country                string `json:"country"`
		ZIPcode                string `json:"ZIPcode"`
		StoreGroups            string `json:"storeGroups"`
		CompanyName            string `json:"companyName"`
		CompanyCode            string `json:"companyCode"`
		CompanyVatNumber       string `json:"companyVatNumber"`
		Phone                  string `json:"phone"`
		Fax                    string `json:"fax"`
		Email                  string `json:"email"`
		Website                string `json:"website"`
		BankName               string `json:"bankName"`
		BankAccountNumber      string `json:"bankAccountNumber"`
		Iban                   string `json:"iban"`
		Swift                  string `json:"swift"`
		UsesLocalQuickButtons  int    `json:"usesLocalQuickButtons"`
		DefaultCustomerGroupID int    `json:"defaultCustomerGroupID"`
		IsOfflineInventory     int    `json:"isOfflineInventory"`
		TimeZone               string `json:"timeZone"`
		Attributes             []struct {
			AttributeName  string `json:"attributeName"`
			AttributeType  string `json:"attributeType"`
			AttributeValue string `json:"attributeValue"`
		} `json:"attributes"`
	}

	Warehouses []Warehouse

	WarehouseManager interface {
		GetWarehouses(ctx context.Context) (Warehouses, error)
	}
)

//GetWarehouses ...
func (cli *Client) GetWarehouses(ctx context.Context) (Warehouses, error) {

	resp, err := cli.SendRequest(ctx, api.GetWarehousesMethod, map[string]string{"warehouseID": "0"})
	if err != nil {
		return nil, err
	}

	res := &GetWarehousesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetWarehousesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res.Warehouses, nil
}
*/
