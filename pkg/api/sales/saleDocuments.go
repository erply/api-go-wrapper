package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/erply/api-go-wrapper/pkg/api/warehouse"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

/*
//SaleDocument model
type SaleDocument struct {
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
	ShipTo   *customers.Customer
	//Recipient if invoice_client_is_payer = 0
	CustomerID int `json:"customerID"`
	//Payer if invoice_client_is_payer = 0
	PayerID int `json:"payerID"`
	Payer   *customers.Customer
	//Buyer represents IClient if invoice_client_is_payer = 1 OR Customer if invoice_client_is_payer = 0
	Buyer *customers.Customer

	//Location for additional fields from getWarehouses request to be used in seller party address.
	Location *warehouse.Warehouse

	AddressID               int                 `json:"addressID"`
	PayerAddressID          int                 `json:"payerAddressID"`
	ShipToAddressID         string              `json:"shipToAddressID"`
	ContactID               int                 `json:"contactID"`
	EmployeeID              int                 `json:"employeeID"`
	PaymentDays             string                  `json:"paymentDays"`
	Confirmed               string                  `json:"confirmed"`
	Notes                   string                  `json:"notes"`
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
	InvoiceRows             []api.InvoiceRow    `json:"rows"`
}

type BaseDocument struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Type   string `json:"type"`
	Date   string `json:"date"`
}

type PostSalesDocumentResponse struct {
	Status        common.Status           `json:"status"`
	ImportReports SaleDocImportReports `json:"records"`
}

type GetSalesDocumentResponse struct {
	Status         common.Status     `json:"status"`
	SalesDocuments []SaleDocument `json:"records"`
}
type SaleDocImportReports []SaleDocImportReport
type SaleDocImportReport struct {
	InvoiceID    int     `json:"invoiceID"`
	CustomNumber string  `json:"customNumber"`
	Rounding     float64 `json:"rounding"`
	Total        float64 `json:"total"`
}

type SalesDocumentManager interface {
	SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error)
	GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error)
	DeleteDocument(ctx context.Context, filters map[string]string) error
}

func (cli *Client) SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error) {
	resp, err := cli.SendRequest(ctx, api.saveSalesDocumentMethod, filters)
	if err != nil {
		return nil, erro.NewFromError("PostSalesDocument request failed", err)
	}
	res := &PostSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling PostSalesDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

//GetSalesDocument erply API request
func (cli *Client) GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error) {
	resp, err := cli.SendRequest(ctx, api.GetSalesDocumentsMethod, filters)
	if err != nil {
		return nil, erro.NewFromError("GetSalesDocument request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetSalesDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res.SalesDocuments, nil
}

func (cli *Client) DeleteDocument(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteSalesDocument", filters)
	if err != nil {
		return erro.NewFromError("DeleteDocumentsByIds request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return erro.NewFromError("unmarshaling DeleteDocumentsByIds failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return nil
}
*/
