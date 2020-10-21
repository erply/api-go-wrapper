package warehouse

import (
	"encoding/json"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	GetWarehousesResponse struct {
		Status     sharedCommon.Status `json:"status"`
		Warehouses Warehouses          `json:"records"`
	}

	Warehouse struct {
		WarehouseID            string      `json:"warehouseID"`
		PricelistID            json.Number `json:"pricelistID"`
		PricelistID2           json.Number `json:"pricelistID2"`
		PricelistID3           json.Number `json:"pricelistID3"`
		PricelistID4           json.Number `json:"pricelistID4"`
		PricelistID5           json.Number `json:"pricelistID5"`
		Name                   string      `json:"name"`
		Code                   string      `json:"code"`
		AddressID              int         `json:"addressID"`
		Address                string      `json:"address"`
		Street                 string      `json:"street"`
		Address2               string      `json:"address2"`
		City                   string      `json:"city"`
		State                  string      `json:"state"`
		Country                string      `json:"country"`
		ZIPcode                string      `json:"ZIPcode"`
		StoreGroups            string      `json:"storeGroups"`
		CompanyName            string      `json:"companyName"`
		CompanyCode            string      `json:"companyCode"`
		CompanyVatNumber       string      `json:"companyVatNumber"`
		Phone                  string      `json:"phone"`
		Fax                    string      `json:"fax"`
		Email                  string      `json:"email"`
		Website                string      `json:"website"`
		BankName               string      `json:"bankName"`
		BankAccountNumber      string      `json:"bankAccountNumber"`
		Iban                   string      `json:"iban"`
		Swift                  string      `json:"swift"`
		UsesLocalQuickButtons  int         `json:"usesLocalQuickButtons"`
		DefaultCustomerGroupID int         `json:"defaultCustomerGroupID"`
		IsOfflineInventory     int         `json:"isOfflineInventory"`
		TimeZone               string      `json:"timeZone"`
		Attributes             []struct {
			AttributeName  string `json:"attributeName"`
			AttributeType  string `json:"attributeType"`
			AttributeValue string `json:"attributeValue"`
		} `json:"attributes"`
	}

	Warehouses []Warehouse

	GetWarehousesBulkItem struct {
		Status     sharedCommon.StatusBulk `json:"status"`
		Warehouses Warehouses              `json:"records"`
	}

	GetWarehousesResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []GetWarehousesBulkItem `json:"requests"`
	}

	SaveWarehouseResult struct {
		WarehouseID int `json:"warehouseID"`
	}

	SaveWarehouseResponse struct {
		Status  sharedCommon.Status   `json:"status"`
		Results []SaveWarehouseResult `json:"records"`
	}

	SaveWarehouseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Results []SaveWarehouseResult   `json:"records"`
	}

	SaveWarehouseResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []SaveWarehouseBulkItem `json:"requests"`
	}

	SaveInventoryRegistrationResult struct {
		InventoryRegistrationID int `json:"inventoryRegistrationID"`
	}

	SaveInventoryRegistrationResponse struct {
		Status  sharedCommon.Status               `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryRegistrationBulkItem struct {
		Status  sharedCommon.StatusBulk           `json:"status"`
		Results []SaveInventoryRegistrationResult `json:"records"`
	}

	SaveInventoryRegistrationResponseBulk struct {
		Status    sharedCommon.Status                 `json:"status"`
		BulkItems []SaveInventoryRegistrationBulkItem `json:"requests"`
	}
)
