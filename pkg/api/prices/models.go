package prices

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type PriceListRule struct {
	ProductID int     `json:"productID"`
	Price     float32 `json:"price"`
	Amount    int     `json:"amount"`
}

type PriceList struct {
	ID                     int                         `json:"supplierPriceListID"`
	SupplierID             int                         `json:"supplierID"`
	SupplierName           string                      `json:"supplierName"`
	Name                   string                      `json:"name"`
	ValidFrom              string                      `json:"startDate"`
	ValidTo                string                      `json:"endDate"`
	Active                 string                      `json:"active"`
	AddedTimestamp         int                         `json:"added"`
	LastModifiedTimestamp  int                         `json:"lastModified"`
	AddedByUserName        string                      `json:"addedByUserName"`
	LastModifiedByUserName string                      `json:"lastModifiedByUserName"`
	Rules                  []PriceListRule             `json:"pricelistRules"`
	Attributes             []sharedCommon.ObjAttribute `json:"attributes"`
}

type GetPriceListsResponseBulkItem struct {
	Status     sharedCommon.StatusBulk `json:"status"`
	PriceLists []PriceList             `json:"records"`
}

type GetPriceListsResponseBulk struct {
	Status    sharedCommon.Status             `json:"status"`
	BulkItems []GetPriceListsResponseBulkItem `json:"requests"`
}

type GetPriceListsResponse struct {
	Status     sharedCommon.Status `json:"status"`
	PriceLists []PriceList         `json:"records"`
}
