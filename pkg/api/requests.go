package api

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
)

func (cli *erplyClient) VerifyCustomerUser(username, password string) (*WebshopClient, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("VerifyCustomerUser: failed to build request", err)
	}

	params := getMandatoryParameters(cli, VerifyCustomerUserMethod)
	params.Set("username", username)
	params.Set("password", password)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
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

// GetCountries will list countries according to specified filters.
func (cli *erplyClient) GetCountries(ctx context.Context, filters map[string]string) ([]Country, error) {
	resp, err := cli.sendRequest(ctx, GetCountriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetCountriesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Countries, nil
}

//GetUserName from GetUserRights erply API request
func (cli *erplyClient) GetUserName() (string, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return "", err
	}
	params := getMandatoryParameters(cli, GetUserRightsMethod)
	params.Add("getRowsForAllInvoices", "1")
	params.Add("getCurrentUser", "1")
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return "", erplyerr("GetUserRights request failed", err)
	}
	res := &GetUserRightsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erplyerr("unmarshaling GetUserRightsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return "", erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Records) == 0 {
		return "", nil
	}

	return res.Records[0].UserName, nil
}

// GetEmployees will list employees according to specified filters.
func (cli *erplyClient) GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error) {
	resp, err := cli.sendRequest(ctx, GetEmployeesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetEmployeesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetEmployeesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Employees, nil
}

// GetBusinessAreas will list business areas according to specified filters.
func (cli *erplyClient) GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error) {
	resp, err := cli.sendRequest(ctx, GetBusinessAreasMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetBusinessAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetBusinessAreasResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.BusinessAreas, nil
}

// GetCurrencies will list currencies according to specified filters.
func (cli *erplyClient) GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error) {
	resp, err := cli.sendRequest(ctx, GetCurrenciesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCurrenciesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetCurrenciesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Currencies, nil
}

func (cli *erplyClient) PostPurchaseDocument(in *PurchaseDocumentConstructor, provider string) (PurchaseDocImportReports, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build PostSalesDocument request", err)
	}
	params := getMandatoryParameters(cli, savePurchaseDocumentMethod)

	params.Add("currencyCode", in.DocumentData.CurrencyCode)
	params.Add("no", in.DocumentData.InvoiceNumber)
	params.Add("type", "PRCINVOICE")
	params.Add("date", in.DocumentData.Date)
	//params.Add("time", in.InvoiceInformation.)
	// set to POS owner or company info if seller is omitted
	if in.SellerParty != nil {
		params.Add("supplierID", strconv.Itoa(in.SellerParty.ID))
		if in.SellerParty.ContactPersons != nil && len(in.SellerParty.ContactPersons) > 0 {
			params.Add("contactID", strconv.Itoa(in.SellerParty.ContactPersons[0].ContactPersonID))
		} else if in.SellerParty.CustomerAddresses != nil && len(in.SellerParty.CustomerAddresses) > 0 {
			params.Add("addressID", strconv.Itoa(in.SellerParty.CustomerAddresses[0].AddressID))
		}
	}
	if in.PaymentParty != nil {
		params.Add("payerID", strconv.Itoa(in.PaymentParty.ID))
		if in.PaymentParty.CustomerAddresses != nil && len(in.PaymentParty.CustomerAddresses) > 0 {
			params.Add("payerAddressID", strconv.Itoa(in.PaymentParty.CustomerAddresses[0].AddressID))
		}
	}

	//params.Add("confirmInvoice", "0")
	params.Add("customNumber", fmt.Sprintf("%s-%s", provider, in.DocumentData.InvoiceNumber))
	params.Add("referenceNumber", in.DocumentData.PaymentReferenceNumber)
	params.Add("notes", in.DocumentData.Notes)
	params.Add("paymentDays", in.DocumentData.PaymentDays)
	params.Add("paid", "0")

	for id, item := range in.DocumentData.ProductRows {
		params.Add(fmt.Sprintf("productID%d", id), item.ProductID)
		params.Add(fmt.Sprintf("itemName%d", id), item.ItemName)
		params.Add(fmt.Sprintf("amount%d", id), item.Amount)
		params.Add(fmt.Sprintf("price%d", id), item.Price)
		params.Add(fmt.Sprintf("TotalSalesTax%d", id), item.VatRate)
	}
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("PostSalesDocument request failed", err)
	}
	res := &PostPurchaseDocumentResponse{}
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

func (cli *erplyClient) IsCustomerUsernameAvailable(username string) (bool, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return false, erplyerr("IsCustomerUsernameAvailable: failed to build request", err)
	}

	params := getMandatoryParameters(cli, validateCustomerUsernameMethod)
	params.Add("username", username)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return false, erplyerr("IsCustomerUsernameAvailable: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return false, erplyerr(fmt.Sprintf("IsCustomerUsernameAvailable: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status Status
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, erplyerr("IsCustomerUsernameAvailable: unmarshaling response failed", err)
	}

	return respData.Status.ErrorCode == 0, nil
}

func (cli *erplyClient) logProcessingOfCustomerData(log *CustomerDataProcessingLog) error {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return erplyerr("failed to build logProcessingOfCustomerData request", err)
	}

	params := url.Values{}
	params.Add("request", logProcessingOfCustomerDataMethod)
	params.Add(sessionKey, cli.sessionKey)
	params.Add(clientCode, cli.clientCode)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return erplyerr("logProcessingOfCustomerData request failed", err)
	}

	if resp.StatusCode != http.StatusOK {
		return erplyerr(fmt.Sprintf("Logging response HTTP status is %d", resp.StatusCode), nil)
	}

	return nil
}
