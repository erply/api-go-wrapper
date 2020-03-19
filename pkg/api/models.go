package api

type PurchaseDocumentConstructors []PurchaseDocumentConstructor

type PurchaseDocumentConstructor struct {
	DocumentData  *DocumentData
	PaymentParty  *Customer
	DeliveryParty *Customer
	SellerParty   *Customer
	VatRates      VatRates
}
type (
	verifyIdentityTokenResponse struct {
		Status Status      `json:"status"`
		Result SessionInfo `json:"records"`
	}

	SessionInfo struct {
		SessionKey string `json:"sessionKey"`
	}

	getIdentityTokenResponse struct {
		Status Status        `json:"status"`
		Result IdentityToken `json:"records"`
	}
	IdentityToken struct {
		Jwt string `json:"identityToken"`
	}
)
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
	PaymentType PaymentType
	// DEB for debit and CRED for credit
	PaymentMethod string
	//By default: system-specific, usually 14.
	//In how many days the invoice is due.
	PaymentDays string
	//Invoice payment status.
	//Possible values: PAID, UNPAID.
	PaymentStatus PaymentStatus
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
	Seller   CustomerConstructor
	Payer    CustomerConstructor
	Buyer    CustomerConstructor
	Delivery CustomerConstructor

	ProductRows ProductRows
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

const (
	PAID    PaymentStatus = "PAID"
	UNPAID  PaymentStatus = "UNPAID"
	PENDING InvoiceState  = "PENDING"
	CARD    PaymentType   = "CARD"
)

type InvoiceState string

type ContactPersons []ContactPerson
type ContactPerson struct {
	ContactPersonID   int    `json:"contactPersonID"`
	FullName          string `json:"fullName"`
	GroupName         string `json:"groupName"`
	CountryID         string `json:"countryID"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Fax               string `json:"fax"`
	Code              string `json:"code"`
	BankName          string `json:"bankName"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankIBAN          string `json:"bankIBAN"`
	BankSWIFT         string `json:"bankSWIFT"`
	Notes             string `json:"notes"`
}

type CustomerImportReports []CustomerImportReport
type CustomerImportReport struct {
	ClientID   int `json:"clientID"`
	CustomerID int `json:"customerID"`
}

type PointOfSale struct {
	WarehouseID int `json:"warehouseID"`
}

type PurchaseDocImportReports []PurchaseDocImportReport
type PurchaseDocImportReport struct {
	InvoiceID    int    `json:"invoiceID"`
	CustomNumber string `json:"customNumber"`
	Rounding     int    `json:"rounding"`
	Total        int    `json:"total"`
}

type SaleDocImportReports []SaleDocImportReport
type SaleDocImportReport struct {
	InvoiceID    int     `json:"invoiceID"`
	CustomNumber string  `json:"customNumber"`
	Rounding     float64 `json:"rounding"`
	Total        float64 `json:"total"`
}

type InvoiceRow struct {
	RowID             string `json:"rowID"`
	ProductID         string `json:"productID"`
	Product           *Product
	ItemName          string  `json:"itemName"`
	Barcode           string  `json:"barcode"`
	VatrateID         string  `json:"vatrateID"`
	Amount            string  `json:"amount"`
	Price             string  `json:"price"`
	Discount          string  `json:"discount"`
	BillingStartDate  string  `json:"billingStartDate"`
	BillingEndDate    string  `json:"billingEndDate"`
	Code              string  `json:"code"`
	FinalNetPrice     float64 `json:"finalNetPrice"`
	FinalPriceWithVAT float64 `json:"finalPriceWithVAT"`
	RowNetTotal       float64 `json:"rowNetTotal"`
	RowVAT            float64 `json:"rowVAT"`
	RowTotal          float64 `json:"rowTotal"`
	CampaignIDs       string  `json:"campaignIDs"`
}

type ConfParameter struct {
	Announcement         string `json:"invoice_announcement_eng"`
	InvoiceClientIsPayer string `json:"invoice_client_is_payer"`
	ReverseVatText       string `json:"reverse_vat_text"`
}

type Warehouse struct {
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

type Warehouses []Warehouse

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
