package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
)

type (
	Customer struct {
		ID                   int            `json:"id"`
		CustomerID           int            `json:"customerID"`
		TypeID               string         `json:"type_id"`
		FullName             string         `json:"fullName"`
		CompanyName          string         `json:"companyName"`
		FirstName            string         `json:"firstName"`
		LastName             string         `json:"lastName"`
		GroupID              int            `json:"groupID"`
		EDI                  string         `json:"EDI"`
		IsPOSDefaultCustomer int            `json:"isPOSDefaultCustomer"`
		CountryID            string         `json:"countryID"`
		Phone                string         `json:"phone"`
		EInvoiceEmail        string         `json:"eInvoiceEmail"`
		Email                string         `json:"email"`
		Fax                  string         `json:"fax"`
		Code                 string         `json:"code"`
		ReferenceNumber      string         `json:"referenceNumber"`
		VatNumber            string         `json:"vatNumber"`
		BankName             string         `json:"bankName"`
		BankAccountNumber    string         `json:"bankAccountNumber"`
		BankIBAN             string         `json:"bankIBAN"`
		BankSWIFT            string         `json:"bankSWIFT"`
		PaymentDays          int            `json:"paymentDays"`
		Notes                string         `json:"notes"`
		LastModified         int            `json:"lastModified"`
		CustomerType         string         `json:"customerType"`
		Address              string         `json:"address"`
		CustomerAddresses    Addresses      `json:"addresses"`
		Street               string         `json:"street"`
		Address2             string         `json:"address2"`
		City                 string         `json:"city"`
		PostalCode           string         `json:"postalCode"`
		Country              string         `json:"country"`
		State                string         `json:"state"`
		ContactPersons       ContactPersons `json:"contactPersons"`

		// Web-shop related fields
		Username  string `json:"webshopUsername"`
		LastLogin string `json:"webshopLastLogin"`
	}
	Customers []Customer

	//Attribute field
	Attribute struct {
		Name  string `json:"attributeNam"`
		Type  string `json:"attributeType"`
		Value string `json:"attributeValue"`
	}

	CustomerRequest struct {
		CustomerID        int
		CompanyName       string
		Address           string
		PostalCode        string
		AddressTypeID     int
		City              string
		State             string
		Country           string
		FirstName         string
		LastName          string
		FullName          string
		RegistryCode      string
		VatNumber         string
		Email             string
		Phone             string
		BankName          string
		BankAccountNumber string

		// Web-shop related fields
		Username string
		Password string
	}

	WebshopClient struct {
		ClientID        string `json:"clientID"`
		ClientUsername  string `json:"clientUsername"`
		ClientName      string `json:"clientName"`
		ClientFirstName string `json:"clientFirstName"`
		ClientLastName  string `json:"clientLastName"`
		ClientGroupID   string `json:"clientGroupID"`
		ClientGroupName string `json:"clientGroupName"`
		CompanyID       string `json:"companyID"`
		CompanyName     string `json:"companyName"`
	}
	GetCustomersResponse struct {
		Status    Status    `json:"status"`
		Customers Customers `json:"records"`
	}

	PostCustomerResponse struct {
		Status                Status                `json:"status"`
		CustomerImportReports CustomerImportReports `json:"records"`
	}

	CustomerManager interface {
		PostCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
		GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
		VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error)
		ValidateCustomerUsername(ctx context.Context, username string) (bool, error)
	}
)

func (cli *erplyClient) PostCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error) {
	resp, err := cli.sendRequest(ctx, saveCustomerMethod, filters)
	if err != nil {
		return nil, erplyerr("PostCustomer request failed", err)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling CustomerImportReport failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, errors.New("no report found")
	}

	return &res.CustomerImportReports[0], nil
}

// GetCustomers will list customers according to specified filters.
func (cli *erplyClient) GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error) {
	resp, err := cli.sendRequest(ctx, GetCustomersMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetCustomersResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Customers, nil
}

//username and password are required fields here
func (cli *erplyClient) VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error) {
	filters := map[string]string{
		"username": username,
		"password": password,
	}
	resp, err := cli.sendRequest(ctx, VerifyCustomerUserMethod, filters)
	if err != nil {
		return nil, erplyerr("VerifyCustomerUser: request failed", err)
	}

	var res struct {
		Status  Status
		Records []WebshopClient
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("VerifyCustomerUser: unmarhsalling response failed", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Records) != 1 {
		return nil, erplyerr("VerifyCustomerUser: no records in response", nil)
	}

	return &res.Records[0], nil
}
func (cli *erplyClient) ValidateCustomerUsername(ctx context.Context, username string) (bool, error) {

	params := map[string]string{"username": username}
	resp, err := cli.sendRequest(ctx, validateCustomerUsernameMethod, params)
	if err != nil {
		return false, erplyerr("IsCustomerUsernameAvailable: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return false, erplyerr(fmt.Sprintf(validateCustomerUsernameMethod+": bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status Status
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, erplyerr(validateCustomerUsernameMethod+": unmarshaling response failed", err)
	}
	if respData.Status.ErrorCode != 0 {
		return false, erplyerr(fmt.Sprintf(validateCustomerUsernameMethod+": bad response error code: %d", respData.Status.ErrorCode), nil)
	}
	return true, nil
}
