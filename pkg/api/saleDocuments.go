package api

//SaleDocument model
type SaleDocument struct {
	ID           int    `json:"id"`
	CurrencyRate string `json:"currencyRate"`
	WarehouseID  int    `json:"warehouseID"`
	Number       string `json:"number"`
	Date         string `json:"date"`
	Time         string `json:"time"`

	/*
		Parties block
	*/
	//Payer if invoice_client_is_payer = 1
	ClientID int `json:"clientID"`
	//Recipient if invoice_client_is_payer = 1
	ShipToID int `json:"shipToID"`
	ShipTo   *Customer
	//Recipient if invoice_client_is_payer = 0
	CustomerID int `json:"customerID"`
	//Payer if invoice_client_is_payer = 0
	PayerID int `json:"payerID"`
	Payer   *Customer
	//Buyer represents IClient if invoice_client_is_payer = 1 OR Customer if invoice_client_is_payer = 0
	Buyer *Customer

	//Location for additional fields from getWarehouses request to be used in seller party address.
	Location *Warehouse

	AddressID               int                 `json:"addressID"`
	PayerAddressID          int                 `json:"payerAddressID"`
	ShipToAddressID         string              `json:"shipToAddressID"`
	ContactID               int                 `json:"contactID"`
	EmployeeID              int                 `json:"employeeID"`
	PaymentDays             string              `json:"paymentDays"`
	Confirmed               string              `json:"confirmed"`
	Notes                   string              `json:"notes"`
	LastModified            int                 `json:"lastModified"`
	PackingUnitsDescription string              `json:"packingUnitsDescription"`
	CurrencyCode            string              `json:"currencyCode"`
	ContactName             string              `json:"contactName"`
	Type                    string              `json:"type"`
	InvoiceState            string              `json:"invoiceState"`
	PaymentType             string              `json:"paymentType"`
	BaseDocuments           []BaseDocument      `json:"baseDocuments"`
	NetTotal                float64             `json:"netTotal"`
	VatTotal                float64             `json:"vatTotal"`
	VatTotalsByTaxRates     VatTotalsByTaxRates `json:"vatTotalsByTaxRate"`
	Rounding                float64             `json:"rounding"`
	Total                   float64             `json:"total"`
	Paid                    string              `json:"paid"`
	PrintDiscounts          int                 `json:"printDiscounts"`
	ReferenceNumber         string              `json:"referenceNumber"`
	CustomReferenceNumber   string              `json:"customReferenceNumber"`
	PaymentStatus           string              `json:"paymentStatus"`
	Penalty                 string              `json:"penalty"`
	InvoiceLink             string              `json:"invoiceLink"`
	InvoiceRows             []InvoiceRow        `json:"rows"`
}

type BaseDocument struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Type   string `json:"type"`
	Date   string `json:"date"`
}

//SaleDocumentConstructor ..
type SaleDocumentConstructor struct {
	DocumentData  *DocumentData
	Attributes    []*Attribute
	PaymentParty  *Customer
	DeliveryParty *Customer
	SellerParty   *Customer
	VatRates      VatRates
}
