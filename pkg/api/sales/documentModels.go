package sales

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	SaleDocument struct {
		ID           int    `json:"id"`
		CurrencyRate string `json:"currencyRate"`
		WarehouseID  int    `json:"warehouseID"`
		Number       string `json:"number"`
		Date         string `json:"date"`
		Time         string `json:"time"`

		//Payer if invoice_client_is_payer = 1
		ClientID int `json:"clientID"`
		//Recipient if invoice_client_is_payer = 1
		ShipToID int `json:"shipToID"`
		//Recipient if invoice_client_is_payer = 0
		CustomerID int `json:"customerID"`
		//Payer if invoice_client_is_payer = 0
		PayerID int `json:"payerID"`

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

	InvoiceRow struct {
		RowID             string  `json:"rowID"`
		ProductID         string  `json:"productID"`
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
	VatTotalsByTaxRates []VatTotalsByTaxRate
	VatTotalsByTaxRate  struct {
		VatrateID int     `json:"vatrateID"`
		Total     float64 `json:"total"`
	}
	BaseDocument struct {
		ID     int    `json:"id"`
		Number string `json:"number"`
		Type   string `json:"type"`
		Date   string `json:"date"`
	}

	PostSalesDocumentResponse struct {
		Status        common.Status        `json:"status"`
		ImportReports SaleDocImportReports `json:"records"`
	}

	GetSalesDocumentResponse struct {
		Status         common.Status  `json:"status"`
		SalesDocuments []SaleDocument `json:"records"`
	}
	SaleDocImportReports []SaleDocImportReport
	SaleDocImportReport  struct {
		InvoiceID    int     `json:"invoiceID"`
		CustomNumber string  `json:"customNumber"`
		Rounding     float64 `json:"rounding"`
		Total        float64 `json:"total"`
	}
	PurchaseDocImportReports []PurchaseDocImportReport
	PurchaseDocImportReport  struct {
		InvoiceID    int    `json:"invoiceID"`
		CustomNumber string `json:"customNumber"`
		Rounding     int    `json:"rounding"`
		Total        int    `json:"total"`
	}

	SavePurchaseDocumentResponse struct {
		Status        common.Status            `json:"status"`
		ImportReports PurchaseDocImportReports `json:"records"`
	}
)
