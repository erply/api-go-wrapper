package sales

import (
	"encoding/json"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

const (
	SaleDocumentTypeInvWayBill    = "INVWAYBILL"
	SaleDocumentTypeCASHINVOICE   = "CASHINVOICE"
	SaleDocumentTypeWayBill       = "WAYBILL"
	SaleDocumentTypePrepayment    = "PREPAYMENT"
	SaleDocumentTypeOffer         = "OFFER"
	SaleDocumentTypeExportInvoice = "EXPORTINVOICE"
	SaleDocumentTypeReservation   = "RESERVATION"
	SaleDocumentTypeCreditInvoice = "CREDITINVOICE"
	SaleDocumentTypeOrder         = "ORDER"
	SaleDocumentTypeInvoice       = "INVOICE"
)

type (
	SaleDocument struct {
		ID            int    `json:"id"`
		CurrencyRate  string `json:"currencyRate"`
		WarehouseID   int    `json:"warehouseID"`
		WarehouseName string `json:"warehouseName"`
		Number        string `json:"number"`
		Date          string `json:"date"`
		DeliveryDate  string `json:"deliveryDate"`
		Time          string `json:"time"`

		//Payer if invoice_client_is_payer = 1
		ClientID    int    `json:"clientID"`
		ClientEmail string `json:"clientEmail"`
		//Recipient if invoice_client_is_payer = 1
		ShipToID int `json:"shipToID"`
		//Recipient if invoice_client_is_payer = 0
		CustomerID int `json:"customerID"`
		//Payer if invoice_client_is_payer = 0
		PayerID int `json:"payerID"`

		AddressID                       int                 `json:"addressID"`
		Address                         string              `json:"address"`
		PayerAddressID                  int                 `json:"payerAddressID"`
		ShipToAddressID                 string              `json:"shipToAddressID"`
		ContactID                       int                 `json:"contactID"`
		EmployeeID                      int                 `json:"employeeID"`
		PaymentDays                     string              `json:"paymentDays"`
		Confirmed                       string              `json:"confirmed"`
		Notes                           string              `json:"notes"`
		InternalNotes                   string              `json:"internalNotes"`
		LastModified                    int                 `json:"lastModified"`
		PackingUnitsDescription         string              `json:"packingUnitsDescription"`
		InventoryTransactionDate        string              `json:"inventoryTransactionDate"`
		CurrencyCode                    string              `json:"currencyCode"`
		ContactName                     string              `json:"contactName"`
		ClientName                      string              `json:"clientName"`
		ClientCardNumber                string              `json:"clientCardNumber"`
		Type                            string              `json:"type"`
		InvoiceState                    string              `json:"invoiceState"`
		PaymentType                     string              `json:"paymentType"`
		BaseDocuments                   []BaseDocument      `json:"baseDocuments"`
		FollowUpDocuments               []BaseDocument      `json:"followUpDocuments"`
		NetTotal                        float64             `json:"netTotal"`
		VatTotal                        float64             `json:"vatTotal"`
		VatTotalsByTaxRates             VatTotalsByTaxRates `json:"vatTotalsByTaxRate"`
		Rounding                        float64             `json:"rounding"`
		Total                           float64             `json:"total"`
		Paid                            string              `json:"paid"`
		PrintDiscounts                  int                 `json:"printDiscounts"`
		ReferenceNumber                 string              `json:"referenceNumber"`
		CustomReferenceNumber           string              `json:"customReferenceNumber"`
		PaymentStatus                   string              `json:"paymentStatus"`
		Penalty                         string              `json:"penalty"`
		InvoiceLink                     string              `json:"invoiceLink"`
		EmployeeName                    string              `json:"employeeName"`
		TransportTypeName               string              `json:"transportTypeName"`
		ShipToName                      string              `json:"shipToName"`
		ShippingDate                    string              `json:"shippingDate"`
		InvoiceRows                     []InvoiceRow        `json:"rows"`
		Attributes                      []PaymentAttribute  `json:"attributes"`
		ExportInvoiceType               string              `json:"exportInvoiceType"`
		PointOfSaleID                   int                 `json:"pointOfSaleID"`
		PricelistID                     string              `json:"pricelistID"`
		PointOfSaleName                 string              `json:"pointOfSaleName"`
		ClientFactoringContractNumber   string              `json:"clientFactoringContractNumber"`
		ClientPaysViaFactoring          int                 `json:"clientPaysViaFactoring"`
		PayerName                       string              `json:"payerName"`
		PayerAddress                    string              `json:"payerAddress"`
		PayerFactoringContractNumber    string              `json:"payerFactoringContractNumber"`
		PayerPaysViaFactoring           string              `json:"payerPaysViaFactoring"`
		ShipToAddress                   string              `json:"shipToAddress"`
		ShipToContactID                 int                 `json:"shipToContactID"`
		ShipToContactName               string              `json:"shipToContactName"`
		ProjectID                       int                 `json:"projectID"`
		PreviousReturnsExist            int                 `json:"previousReturnsExist"`
		NetTotalsByTaxRate              VatTotalsByTaxRates `json:"netTotalsByTaxRate"`
		ExternalNetTotal                float64             `json:"externalNetTotal"`
		ExternalVatTotal                float64             `json:"externalVatTotal"`
		ExternalRounding                float64             `json:"externalRounding"`
		ExternalTotal                   float64             `json:"externalTotal"`
		PaymentTypeID                   int                 `json:"paymentTypeID"`
		TaxExemptCertificateNumber      string              `json:"taxExemptCertificateNumber"`
		PackerID                        int                 `json:"packerID"`
		TrackingNumber                  string              `json:"trackingNumber"`
		FulfillmentStatus               string              `json:"fulfillmentStatus"`
		Cost                            float64             `json:"cost"`
		ReserveGoods                    int                 `json:"reserveGoods"`
		ReserveGoodsUntilDate           string              `json:"reserveGoodsUntilDate"`
		DeliveryTypeID                  int                 `json:"deliveryTypeID"`
		DeliveryTypeName                string              `json:"deliveryTypeName"`
		TriangularTransaction           string              `json:"triangularTransaction"`
		PurchaseOrderDone               string              `json:"purchaseOrderDone"`
		TransactionTypeID               int                 `json:"transactionTypeID"`
		TransactionTypeName             string              `json:"transactionTypeName"`
		TransportTypeID                 int                 `json:"transportTypeID"`
		DeliveryTerms                   string              `json:"deliveryTerms"`
		EuInvoiceType                   string              `json:"euInvoiceType"`
		DeliveryTermsLocation           string              `json:"deliveryTermsLocation"`
		DeliveryOnlyWhenAllItemsInStock int                 `json:"deliveryOnlyWhenAllItemsInStock"`
		LastModifierUsername            string              `json:"lastModifierUsername"`
		Added                           int                 `json:"added"`
		ReceiptLink                     string              `json:"receiptLink"`
		AmountAddedToStoreCredit        json.Number         `json:"amountAddedToStoreCredit"`
		AmountPaidWithStoreCredit       json.Number         `json:"amountPaidWithStoreCredit"`
		ApplianceID                     int                 `json:"applianceID"`
		ApplianceReference              string              `json:"applianceReference"`
		AssignmentID                    int                 `json:"assignmentID"`
		VehicleMileage                  int                 `json:"vehicleMileage"`
	}

	InvoiceRow struct {
		RowID             string      `json:"rowID"`
		StableRowID       string      `json:"stableRowID"`
		ProductID         string      `json:"productID"`
		ItemName          string      `json:"itemName"`
		Barcode           string      `json:"barcode"`
		VatrateID         string      `json:"vatrateID"`
		Amount            string      `json:"amount"`
		Price             string      `json:"price"`
		Discount          string      `json:"discount"`
		BillingStartDate  string      `json:"billingStartDate"`
		BillingEndDate    string      `json:"billingEndDate"`
		Code              string      `json:"code"`
		Code2             string      `json:"code2"`
		FinalNetPrice     float64     `json:"finalNetPrice"`
		FinalPriceWithVAT float64     `json:"finalPriceWithVAT"`
		RowNetTotal       float64     `json:"rowNetTotal"`
		RowVAT            float64     `json:"rowVAT"`
		RowTotal          float64     `json:"rowTotal"`
		CampaignIDs       string      `json:"campaignIDs"`
		Jdoc              interface{} `json:"jdoc"`
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
		Status        sharedCommon.Status  `json:"status"`
		ImportReports SaleDocImportReports `json:"records"`
	}

	GetSalesDocumentResponse struct {
		Status         sharedCommon.Status `json:"status"`
		SalesDocuments []SaleDocument      `json:"records"`
	}
	SaleDocImportReports []SaleDocImportReport
	SaleDocImportReport  struct {
		InvoiceID    int     `json:"invoiceID"`
		InvoiceNo    string  `json:"invoiceNo"`
		CustomNumber string  `json:"customNumber"`
		InvoiceLink  string  `json:"invoiceLink"`
		ReceiptLink  string  `json:"receiptLink"`
		Net          float64 `json:"net"`
		Vat          float64 `json:"vat"`
		Rounding     float64 `json:"rounding"`
		Total        float64 `json:"total"`
	}
	PurchaseDocImportReports []PurchaseDocImportReport
	PurchaseDocImportReport  struct {
		InvoiceID    int     `json:"invoiceID"`
		CustomNumber string  `json:"customNumber"`
		Rounding     int     `json:"rounding"`
		Total        float64 `json:"total"`
	}

	SavePurchaseDocumentResponse struct {
		Status        sharedCommon.Status      `json:"status"`
		ImportReports PurchaseDocImportReports `json:"records"`
	}

	GetSaleDocumentBulkItem struct {
		Status        sharedCommon.StatusBulk `json:"status"`
		SaleDocuments []SaleDocument          `json:"records"`
	}

	GetSaleDocumentResponseBulk struct {
		Status    sharedCommon.Status       `json:"status"`
		BulkItems []GetSaleDocumentBulkItem `json:"requests"`
	}
)
