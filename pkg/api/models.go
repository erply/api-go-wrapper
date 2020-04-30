package api

import (
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/erply/api-go-wrapper/pkg/api/sales"
	"github.com/erply/api-go-wrapper/pkg/common"
)

type PostPurchaseDocumentResponse struct {
	Status        common.Status            `json:"status"`
	ImportReports PurchaseDocImportReports `json:"records"`
}

type GetCountriesResponse struct {
	Status    common.Status `json:"status"`
	Countries []Country     `json:"records"`
}

type GetEmployeesResponse struct {
	Status    common.Status `json:"status"`
	Employees []Employee    `json:"records"`
}

type GetBusinessAreasResponse struct {
	Status        common.Status  `json:"status"`
	BusinessAreas []BusinessArea `json:"records"`
}
type GetCurrenciesResponse struct {
	Status     common.Status `json:"status"`
	Currencies []Currency    `json:"records"`
}

//CustomerDataProcessingLog ...
type CustomerDataProcessingLog struct {
	activityType string
	fields       []string
	customerIds  []int
}
type DocumentDatas []DocumentData

type DocumentData struct {
	//Document type
	Type string
	//Currency code: "EUR", "USD" etc. Currency must be defined in your Erply account.
	CurrencyCode string
	//eg. 2010-01-29
	//Each sales document must have a date. If omitted, API applies current date.
	Date string
	//eg. 14:59:00
	//If omitted, API applies current time.
	Time string
	//Assign a custom number to this sales document. As opposed to invoiceNo, this field may contain letters, spacing and punctuation.
	CustomNumber string
	// number of invoice document in provider system
	InvoiceNumber string
	// Invoice content text
	InvoiceContentText string
	///Sales document's custom reference number. This field must be used only if you want to override default reference numbers.
	CustomReferenceNumber string
	//Notes printed on the invoice
	Notes string
	//Additional text
	Text string
	//Status of the document itself.
	//For invoices, possible values: PENDING, READY, MAILED, PRINTED. For orders, possible values are: PENDING, READY, SHIPPED, FULFILLED, CANCELLED
	InvoiceState InvoiceState
	//Expected invoice payment method: eg. CASH, CARD, TRANSFER, CHECK, GIFTCARD.
	PaymentType sales.PaymentType
	// DEB for debit and CRED for credit
	PaymentMethod string
	//By default: system-specific, usually 14.
	//In how many days the invoice is due.
	PaymentDays string
	//Invoice payment status.
	//Possible values: PAID, UNPAID.
	PaymentStatus sales.PaymentStatus
	//Invoice payment information, who paid, when, how.
	//Max 255 characters
	PaymentInfo string
	//Payment reference number
	PaymentReferenceNumber string
	//ISO date (yyyy-mm-dd)
	// Customer requested delivery date (for the whole document). You may also set requested delivery dates for each line individually, see deliveryDate#
	DeliveryDate string
	//ISO date (yyyy-mm-dd)
	ShippingDate string
	//Search by exact warehouse code.
	WarehouseCode string
	//seller company registry code
	Seller   customers.CustomerRequest
	Payer    customers.CustomerRequest
	Buyer    customers.CustomerRequest
	Delivery customers.CustomerRequest

	ProductRows ProductRows
	CustomerId  uint
}

type ProductRows []ProductRow

type ProductRow struct {
	//ID of the product (SKU) sold. Either productID or serviceID can be set, but not both at the same time. Both can be omitted, however - in that case a free-text invoice row will be created.
	ProductID string
	ItemName  string
	//Sold quantity must be a decimal, and can not be zero.
	Amount string
	///Net sales price per item, pre-discount.
	Price string
	//Discount % that WILL BE SUBTRACTED from the price specified in previous parameter.
	Discount string
	//Customer requested delivery date for this specific item. You can also set a requested delivery date for the whole document, see deliveryDate above.
	DeliveryDate string
	//Billing start date. See previous field.
	BillingStartDate string
	//Billing end date. See previous field.
	BillingEndDate string
	// item vat rate
	VatRate string
}

type InvoiceState string

type PurchaseDocImportReports []PurchaseDocImportReport
type PurchaseDocImportReport struct {
	InvoiceID    int    `json:"invoiceID"`
	CustomNumber string `json:"customNumber"`
	Rounding     int    `json:"rounding"`
	Total        int    `json:"total"`
}

type UserCredentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type VerifyUserResponse struct {
	Records []Records `json:"records"`
}

type Records struct {
	SessionKey string `json:"sessionKey"`
}

type GetUserRightsResponse struct {
	Status  Status       `json:"status"`
	Records []UserRights `json:"records"`
}

type UserRights struct {
	UserName string `json:"userName"`
}
type Country struct {
	CountryId             uint   `json:"countryID"`
	CountryName           string `json:"countryName"`
	CountryCode           string `json:"countryCode"`
	MemberOfEuropeanUnion byte   `json:"memberOfEuropeanUnion"`
	LastModified          string `json:"lastModified"`
	Added                 uint64 `json:"added"`
}

type Employee struct {
	EmployeeID             string                `json:"employeeID"`
	FullName               string                `json:"fullName"`
	EmployeeName           string                `json:"employeeName"`
	FirstName              string                `json:"firstName"`
	LastName               string                `json:"lastName"`
	Phone                  string                `json:"phone"`
	Mobile                 string                `json:"mobile"`
	Email                  string                `json:"email"`
	Fax                    string                `json:"fax"`
	Code                   string                `json:"code"`
	Gender                 string                `json:"gender"`
	UserID                 string                `json:"userID"`
	Username               string                `json:"username"`
	UserGroupID            string                `json:"userGroupID"`
	Warehouses             []EmployeeWarehouse   `json:"warehouses"`
	PointsOfSale           string                `json:"pointsOfSale"`
	ProductIDs             []EmployeeProduct     `json:"productIDs"`
	Attributes             []common.ObjAttribute `json:"attributes"`
	LastModified           uint64                `json:"lastModified"`
	LastModifiedByUserName string                `json:"lastModifiedByUserName"`

	// detail fileds
	Skype        string `json:"skype"`
	Birthday     string `json:"birthday"`
	JobTitleID   uint   `json:"jobTitleID"`
	JobTitleName string `json:"jobTitleName"`
	Notes        string `json:"notes"`
	Added        uint64 `json:"added"`
}

type EmployeeWarehouse struct {
	Id uint `json:"id"`
}

type EmployeeProduct struct {
	ProductID    uint   `json:"productID"`
	ProductCode  string `json:"productCode"`
	ProductName  string `json:"productName"`
	ProductGroup uint   `json:"productGroup"`
}

type BusinessArea struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Added        uint64 `json:"added"`
	LastModified uint64 `json:"lastModified"`
}
type Currency struct {
	CurrencyID   string `json:"currencyID"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Default      string `json:"default"`
	NameShort    string `json:"nameShort"`
	NameFraction string `json:"nameFraction"`
	Added        string `json:"added"`
	LastModified string `json:"lastModified"`
}
