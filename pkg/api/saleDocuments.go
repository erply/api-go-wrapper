package api

import (
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
	"strings"
)

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
type PostSalesDocumentResponse struct {
	Status        Status               `json:"status"`
	ImportReports SaleDocImportReports `json:"records"`
}

type GetSalesDocumentResponse struct {
	Status         Status         `json:"status"`
	SalesDocuments []SaleDocument `json:"records"`
}
type SaleDocImportReports []SaleDocImportReport
type SaleDocImportReport struct {
	InvoiceID    int     `json:"invoiceID"`
	CustomNumber string  `json:"customNumber"`
	Rounding     float64 `json:"rounding"`
	Total        float64 `json:"total"`
}

func (cli *erplyClient) PostSalesDocumentFromWoocomm(in *SaleDocumentConstructor, shopOrderID string) (SaleDocImportReports, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build PostSalesDocument request", err)
	}
	params := getMandatoryParameters(cli, saveSalesDocumentMethod)

	params.Add("currencyCode", in.DocumentData.CurrencyCode)
	params.Add("type", in.DocumentData.Type)
	params.Add("date", in.DocumentData.Date)
	params.Add("deliveryDate", in.DocumentData.DeliveryDate)
	// set to POS owner or company info if seller is omitted
	if in.SellerParty != nil {
		params.Add("customerID", strconv.Itoa(in.SellerParty.ID))
		if in.SellerParty.ContactPersons != nil && len(in.SellerParty.ContactPersons) != 0 {
			params.Add("contactID", strconv.Itoa(in.SellerParty.ContactPersons[0].ContactPersonID))
		} else if in.SellerParty.CustomerAddresses != nil && len(in.SellerParty.CustomerAddresses) != 0 {
			params.Add("addressID", strconv.Itoa(in.SellerParty.CustomerAddresses[0].AddressID))
		}
	}
	if in.PaymentParty != nil {
		params.Add("payerID", strconv.Itoa(in.PaymentParty.ID))
		if in.PaymentParty.CustomerAddresses != nil && len(in.PaymentParty.CustomerAddresses) != 0 {
			params.Add("payerAddressID", strconv.Itoa(in.PaymentParty.CustomerAddresses[0].AddressID))
		}
	}
	if in.DeliveryParty != nil {
		params.Add("shipToID", strconv.Itoa(in.DeliveryParty.ID))
	}
	params.Add("confirmInvoice", "1")
	//params.Add("customNumber", fmt.Sprintf("%s-%s-%s", shopOrderID, in.DocumentData.CustomNumber, in.DocumentData.InvoiceNumber))
	params.Add("customReferenceNumber", in.DocumentData.PaymentReferenceNumber)
	params.Add("notes", in.DocumentData.Notes)
	params.Add("internalNotes", shopOrderID)

	switch in.DocumentData.PaymentMethod {
	case BankTransfer, PayPal:
		params.Add("paymentType", Transfer)
		params.Add("paymentStatus", Paid)
	case CheckPayment:
		params.Add("paymentType", Check)
		params.Add("paymentStatus", Unpaid)
	case CashOnDelivery:
		params.Add("paymentType", Cash)
		params.Add("paymentStatus", Unpaid)
	}

	params.Add("paymentDays", in.DocumentData.PaymentDays)
	params.Add("paymentInfo", in.DocumentData.InvoiceContentText)

	for id, item := range in.DocumentData.ProductRows {
		params.Add(fmt.Sprintf("productID%d", id), item.ProductID)
		params.Add(fmt.Sprintf("itemName%d", id), item.ItemName)
		params.Add(fmt.Sprintf("amount%d", id), item.Amount)
		params.Add(fmt.Sprintf("price%d", id), item.Price)
		if len(in.VatRates) != 0 {
			params.Add(fmt.Sprintf("vatrateID%d", id), in.VatRates[0].ID)
		}
	}

	for id, attr := range in.Attributes {
		params.Add(fmt.Sprintf("attributeName%d", id), attr.Name)
		params.Add(fmt.Sprintf("attributeType%d", id), attr.Type)
		params.Add(fmt.Sprintf("attributeValue%d", id), attr.Value)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("PostSalesDocument request failed", err)
	}
	res := &PostSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling PostSalesDocumentResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

func (cli *erplyClient) PostSalesDocument(in *SaleDocumentConstructor, provider string) (SaleDocImportReports, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build PostSalesDocument request", err)
	}
	params := getMandatoryParameters(cli, saveSalesDocumentMethod)

	params.Add("currencyCode", in.DocumentData.CurrencyCode)
	params.Add("type", in.DocumentData.Type)
	params.Add("date", in.DocumentData.Date)
	params.Add("deliveryDate", in.DocumentData.DeliveryDate)
	//params.Add("time", in.InvoiceInformation.)
	// set to POS owner or company info if seller is omitted
	if in.SellerParty != nil {
		params.Add("customerID", strconv.Itoa(in.SellerParty.ID))
		if in.SellerParty.ContactPersons != nil && len(in.SellerParty.ContactPersons) != 0 {
			params.Add("contactID", strconv.Itoa(in.SellerParty.ContactPersons[0].ContactPersonID))
		} else if in.SellerParty.CustomerAddresses != nil && len(in.SellerParty.CustomerAddresses) != 0 {
			params.Add("addressID", strconv.Itoa(in.SellerParty.CustomerAddresses[0].AddressID))
		}
	}
	if in.PaymentParty != nil {
		params.Add("payerID", strconv.Itoa(in.PaymentParty.ID))
		if in.PaymentParty.CustomerAddresses != nil && len(in.PaymentParty.CustomerAddresses) != 0 {
			params.Add("payerAddressID", strconv.Itoa(in.PaymentParty.CustomerAddresses[0].AddressID))
		}
	}
	if in.DeliveryParty != nil {
		params.Add("shipToID", strconv.Itoa(in.DeliveryParty.ID))
	}
	params.Add("confirmInvoice", "1")
	params.Add("customNumber", fmt.Sprintf("%s-%s-%s", provider, in.DocumentData.CustomNumber, in.DocumentData.InvoiceNumber))
	params.Add("customReferenceNumber", in.DocumentData.PaymentReferenceNumber)
	params.Add("notes", in.DocumentData.Notes)
	params.Add("internalNotes", in.DocumentData.Text)
	params.Add("paymentType", in.DocumentData.PaymentMethod)
	params.Add("paymentDays", in.DocumentData.PaymentDays)
	params.Add("paymentInfo", in.DocumentData.InvoiceContentText)
	params.Add("paymentStatus", "UNPAID")
	params.Add("customerID", fmt.Sprint(in.DocumentData.CustomerId))

	fmt.Println("customerId", fmt.Sprint(in.DocumentData.CustomerId))
	for id, item := range in.DocumentData.ProductRows {
		params.Add(fmt.Sprintf("productID%d", id), item.ProductID)
		params.Add(fmt.Sprintf("itemName%d", id), item.ItemName)
		params.Add(fmt.Sprintf("amount%d", id), item.Amount)
		params.Add(fmt.Sprintf("price%d", id), item.Price)
		if len(in.VatRates) != 0 {
			params.Add(fmt.Sprintf("vatrateID%d", id), in.VatRates[0].ID)
		}
	}

	for id, attr := range in.Attributes {
		params.Add(fmt.Sprintf("attributeName%d", id), attr.Name)
		params.Add(fmt.Sprintf("attributeType%d", id), attr.Type)
		params.Add(fmt.Sprintf("attributeValue%d", id), attr.Value)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("PostSalesDocument request failed", err)
	}
	res := &PostSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling PostSalesDocumentResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

//GetSalesDocument erply API request
func (cli *erplyClient) GetSalesDocumentsByIDs(ids []string) ([]SaleDocument, error) {
	if len(ids) == 0 {
		return nil, erplyerr("No ids provided for sales documents request", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("No ids provided for sales documents request", err)
	}
	params := getMandatoryParameters(cli, GetSalesDocumentsMethod)
	params.Add("ids", strings.Join(ids, ","))
	params.Add("getRowsForAllInvoices", "1")
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetSalesDocument request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetSalesDocumentResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res.SalesDocuments, nil
}

//GetSalesDocumentById erply API request
func (cli *erplyClient) GetSalesDocumentByID(id string) ([]SaleDocument, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, err
	}
	params := getMandatoryParameters(cli, GetSalesDocumentsMethod)
	params.Add("id", id)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetSalesDocument request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetSalesDocumentResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.SalesDocuments) == 0 {
		return nil, nil
	}

	return res.SalesDocuments, nil
}

func (cli *erplyClient) DeleteDocumentsByID(id string) error {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return err
	}
	params := getMandatoryParameters(cli, "deleteSalesDocument")
	params.Add("documentID", id)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return erplyerr("DeleteDocumentsByIds request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return erplyerr("unmarshaling DeleteDocumentsByIds failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return nil
}
