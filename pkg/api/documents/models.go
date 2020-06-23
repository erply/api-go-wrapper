package documents

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"time"
)

type PurchaseOrderType string

const (
	PurchaseOrder          PurchaseOrderType = "PRCORDER"
	PurchaseInvoiceWaybill PurchaseOrderType = "PRCINVOICE"
	PurchaseReceipt        PurchaseOrderType = "CASHPRCINVOICE"
	PurchaseReturn         PurchaseOrderType = "PRCRETURN"
	PurchaseWaybill        PurchaseOrderType = "PRCWAYBILL"
	PurchaseInvoice        PurchaseOrderType = "PRCINVOICEONLY"
)

type DocumentStatus string

const (
	Pending           DocumentStatus = "PENDING"
	PartiallyReceived DocumentStatus = "PARTIALLY_RECEIVED"
	Received          DocumentStatus = "RECEIVED"
	Ready             DocumentStatus = "READY"
)

type VatRate struct {
	VatrateID int     `json:"vatrateID"`
	Total     float64 `json:"total"`
}

type ReferencedPurchaseDocument struct {
	ID        int               `json:"id"`
	Number    string            `json:"number"`
	RegNumber string            `json:"regnumber"`
	Type      PurchaseOrderType `json:"type"`
	Date      time.Time         `json:"date"`
}

type PurchaseDocument struct {
	ID                       int                          `json:"id"`
	Type                     PurchaseOrderType            `json:"type"`
	Status                   DocumentStatus               `json:"status"`
	CurrencyCode             string                       `json:"currencyCode"`
	CurrencyRate             float64                      `json:"currencyRate"`
	WarehouseID              int                          `json:"warehouseID"`
	WarehouseName            string                       `json:"warehouseName"`
	Number                   string                       `json:"number"`
	RegNumber                string                       `json:"regnumber"`
	Date                     time.Time                    `json:"date"`
	InventoryTransactionDate time.Time                    `json:"inventoryTransactionDate"`
	Time                     string                       `json:"time"`
	SupplierID               int                          `json:"supplierID"`
	SupplierName             string                       `json:"supplierName"`
	SupplierGroupID          int                          `json:"supplierGroupID"`
	AddressID                int                          `json:"addressID"`
	Address                  string                       `json:"address"`
	ContactID                int                          `json:"contactID"`
	ContactName              string                       `json:"contactName"`
	EmployeeID               int                          `json:"employeeID"`
	EmployeeName             string                       `json:"employeeName"`
	SupplierID2              int                          `json:"supplierID2"`
	SupplierName2            string                       `json:"supplierName2"`
	StateID                  int                          `json:"stateID"`
	PaymentDays              int                          `json:"paymentDays"`
	Paid                     int                          `json:"paid"`
	TransactionTypeID        int                          `json:"transactionTypeID"`
	TransportTypeID          int                          `json:"transportTypeID"`
	DeliveryTermsID          int                          `json:"deliveryTermsID"`
	DeliveryTermsLocation    string                       `json:"deliveryTermsLocation"`
	TriangularTransaction    int                          `json:"triangularTransaction"`
	ProjectID                int                          `json:"projectID"`
	Confirmed                int                          `json:"confirmed"`
	ReferenceNumber          string                       `json:"referenceNumber"`
	Notes                    string                       `json:"notes"`
	Rounding                 float64                      `json:"rounding"`
	NetTotal                 float64                      `json:"netTotal"`
	VatTotal                 float64                      `json:"vatTotal"`
	Total                    float64                      `json:"total"`
	NetTotalsByTaxRate       []VatRate                    `json:"netTotalsByTaxRate"`
	VatTotalsByTaxRate       []VatRate                    `json:"vatTotalsByTaxRate"`
	InvoiceLink              string                       `json:"invoiceLink"`
	ShipDate                 time.Time                    `json:"shipDate"`
	Cost                     float64                      `json:"cost"`
	NetTotalForAccounting    float64                      `json:"netTotalForAccounting"`
	TotalForAccounting       float64                      `json:"totalForAccounting"`
	BaseToDocuments          []ReferencedPurchaseDocument `json:"baseToDocuments"`
	BaseDocuments            []ReferencedPurchaseDocument `json:"baseDocuments"`
	LastModified             int64                        `json:"lastModified"`
	Rows                     []PurchaseDocumentRow        `json:"rows"`
	Attributes               []sharedCommon.ObjAttribute  `json:"attributes"`
}

type PurchaseDocumentRow struct {
	ProductID        int       `json:"productID"`
	ServiceID        int       `json:"serviceID"`
	ItemName         string    `json:"itemName"`
	Code             string    `json:"code"`
	Code2            string    `json:"code2"`
	VatrateID        int       `json:"vatrateID"`
	Amount           float64   `json:"amount"`
	Price            float64   `json:"price"`
	Discount         float64   `json:"discount"`
	DeliveryDate     time.Time `json:"deliveryDate"`
	UnitCost         float64   `json:"unitCost"`
	CostTotal        float64   `json:"costTotal"`
	PackageID        int       `json:"packageID"`
	AmountOfPackages float64   `json:"amountOfPackages"`
	AmountInPackage  float64   `json:"amountInPackage"`
	PackageType      string    `json:"packageType"`
	PackageTypeID    int       `json:"packageTypeID"`
}

type GetPurchaseDocumentBulkItem struct {
	Status            sharedCommon.StatusBulk `json:"status"`
	PurchaseDocuments []PurchaseDocument      `json:"records"`
}

type GetPurchaseDocumentResponseBulk struct {
	Status    sharedCommon.Status `json:"status"`
	BulkItems []GetPurchaseDocumentBulkItem  `json:"requests"`
}

type GetPurchaseDocumentsResponse struct {
	Status            sharedCommon.Status `json:"status"`
	PurchaseDocuments []PurchaseDocument  `json:"records"`
}
