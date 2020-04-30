package warehouse

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	GetWarehousesResponse struct {
		Status     common.Status `json:"status"`
		Warehouses Warehouses    `json:"records"`
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
)
