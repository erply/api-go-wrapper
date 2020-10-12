package prices

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type PriceListRule struct {
	ProductID int     `json:"productID"`
	Price     float32 `json:"price,string"`
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

type ProductsInSupplierPriceList struct {
	SupplierPriceListProductID int     `json:"supplierPriceListProductID"`
	ProductID                  int     `json:"productID"`
	Price                      float32 `json:"price"`
	Amount                     int     `json:"amount"`
	CountryID                  int     `json:"countryID"`
	ProductSupplierCode        string  `json:"supplierCode"`
	ImportCode                 string  `json:"importCode"`
	MasterPackQuantity         int     `json:"masterPackQuantity"`
	MinimumOrderQuantity       int     `json:"minimumOrderQuantity"`
}

type GetPriceListsResponseBulkItem struct {
	Status     sharedCommon.StatusBulk `json:"status"`
	PriceLists []PriceList             `json:"records"`
}

type GetPriceListsResponseBulk struct {
	Status    sharedCommon.Status             `json:"status"`
	BulkItems []GetPriceListsResponseBulkItem `json:"requests"`
}

type ProductsInPriceList struct {
	PriceListProductID int     `json:"priceListProductID"`
	ProductID          int     `json:"productID"`
	Price              float32 `json:"price"`
	Amount             int     `json:"amount"`
	Subsidy            float32 `json:"subsidy"`
	SubsidyTypeID      int     `json:"subsidyTypeID"`
	Page               int     `json:"page"`
	ForecastUnits      int     `json:"forecastUnits"`
}

type GetProductsInPriceListResponseBulkItem struct {
	Status     sharedCommon.StatusBulk `json:"status"`
	PriceLists []ProductsInPriceList   `json:"records"`
}

type GetProductsInPriceListResponseBulk struct {
	Status    sharedCommon.Status                      `json:"status"`
	BulkItems []GetProductsInPriceListResponseBulkItem `json:"requests"`
}

type GetProductsInPriceListResponse struct {
	Status     sharedCommon.Status   `json:"status"`
	PriceLists []ProductsInPriceList `json:"records"`
}

type GetPriceListsResponse struct {
	Status     sharedCommon.Status `json:"status"`
	PriceLists []PriceList         `json:"records"`
}

type ProductsInSupplierPriceListResponseBulkItem struct {
	Status                      sharedCommon.StatusBulk       `json:"status"`
	ProductsInSupplierPriceList []ProductsInSupplierPriceList `json:"records"`
}

type ProductsInSupplierPriceListResponseBulk struct {
	Status    sharedCommon.Status                           `json:"status"`
	BulkItems []ProductsInSupplierPriceListResponseBulkItem `json:"requests"`
}

type ProductsInSupplierPriceListResponse struct {
	Status                      sharedCommon.Status           `json:"status"`
	ProductsInSupplierPriceList []ProductsInSupplierPriceList `json:"records"`
}

type ChangeProductToSupplierPriceListResult struct {
	SupplierPriceListProductID int `json:"supplierPriceListProductID"`
}

type ChangeProductToSupplierPriceListResponse struct {
	Status                                 sharedCommon.Status                      `json:"status"`
	ChangeProductToSupplierPriceListResult []ChangeProductToSupplierPriceListResult `json:"records"`
}

type ChangeProductToSupplierPriceListResultBulkItem struct {
	Status  sharedCommon.StatusBulk                  `json:"status"`
	Records []ChangeProductToSupplierPriceListResult `json:"records"`
}

type ChangeProductToSupplierPriceListResponseBulk struct {
	Status    sharedCommon.Status                              `json:"status"`
	BulkItems []ChangeProductToSupplierPriceListResultBulkItem `json:"requests"`
}

type DeleteProductsFromSupplierPriceListResult struct {
	DeletedIDs     string `json:"deletedIDs"`
	NonExistingIDs string `json:"nonExistingIDs"`
}

type DeleteProductsFromSupplierPriceListResponse struct {
	Status                                    sharedCommon.Status                         `json:"status"`
	DeleteProductsFromSupplierPriceListResult []DeleteProductsFromSupplierPriceListResult `json:"records"`
}

type DeleteProductsFromSupplierPriceListBulkItem struct {
	Status  sharedCommon.StatusBulk                     `json:"status"`
	Records []DeleteProductsFromSupplierPriceListResult `json:"records"`
}

type DeleteProductsFromSupplierPriceListResponseBulk struct {
	Status    sharedCommon.Status                           `json:"status"`
	BulkItems []DeleteProductsFromSupplierPriceListBulkItem `json:"requests"`
}

type SaveSupplierPriceListResult struct {
	SupplierPriceListID int `json:"supplierPriceListID"`
}

type SaveSupplierPriceListResultResponse struct {
	Status                      sharedCommon.Status           `json:"status"`
	SaveSupplierPriceListResult []SaveSupplierPriceListResult `json:"records"`
}

type SaveSupplierPriceListBulkItem struct {
	Status  sharedCommon.StatusBulk       `json:"status"`
	Records []SaveSupplierPriceListResult `json:"records"`
}

type SaveSupplierPriceListResponseBulk struct {
	Status    sharedCommon.Status             `json:"status"`
	BulkItems []SaveSupplierPriceListBulkItem `json:"requests"`
}

type NotAddedPriceItem struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

type SavePriceListResult struct {
	PriceListID              int                 `json:"pricelistID"`
	ItemsNotAddedToPriceList []NotAddedPriceItem `json:"itemsNotAddedToPriceList"`
}

type SavePriceListResultResponse struct {
	Status               sharedCommon.Status   `json:"status"`
	SavePriceListResults []SavePriceListResult `json:"records"`
}

type SavePriceListBulkItem struct {
	Status  sharedCommon.StatusBulk `json:"status"`
	Records []SavePriceListResult   `json:"records"`
}

type SavePriceListResponseBulk struct {
	Status    sharedCommon.Status     `json:"status"`
	BulkItems []SavePriceListBulkItem `json:"requests"`
}

type ChangeProductToPriceListResult struct {
	PriceListProductID int `json:"priceListProductID"`
}

type ChangeProductToPriceListResponse struct {
	Status                          sharedCommon.Status              `json:"status"`
	ChangeProductToPriceListResults []ChangeProductToPriceListResult `json:"records"`
}

type ChangeProductToPriceListResultBulkItem struct {
	Status  sharedCommon.StatusBulk          `json:"status"`
	Records []ChangeProductToPriceListResult `json:"records"`
}

type ChangeProductToPriceListResponseBulk struct {
	Status    sharedCommon.Status                      `json:"status"`
	BulkItems []ChangeProductToPriceListResultBulkItem `json:"requests"`
}

type DeleteProductsFromPriceListResult struct {
	DeletedIDs     string `json:"deletedIDs"`
	NonExistingIDs string `json:"nonExistingIDs"`
}

type DeleteProductsFromPriceListResponse struct {
	Status                             sharedCommon.Status                 `json:"status"`
	DeleteProductsFromPriceListResults []DeleteProductsFromPriceListResult `json:"records"`
}

type DeleteProductsFromPriceListBulkItem struct {
	Status  sharedCommon.StatusBulk             `json:"status"`
	Records []DeleteProductsFromPriceListResult `json:"records"`
}

type DeleteProductsFromPriceListResponseBulk struct {
	Status    sharedCommon.Status                   `json:"status"`
	BulkItems []DeleteProductsFromPriceListBulkItem `json:"requests"`
}
