package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	erro "github.com/erply/api-go-wrapper/pkg/errors"
)

//IClient ...
type erplyClient struct {
	sessionKey string
	clientCode string
	partnerKey string
	secret     string
	url        string
	httpClient *http.Client
}

//VerifyUser will give you session key
func VerifyUser(username string, password string, clientCode string) (string, error) {
	client := &http.Client{}

	requestUrl := fmt.Sprintf(baseURL, clientCode)
	params := url.Values{}
	params.Add("username", username)
	params.Add("clientCode", clientCode)
	params.Add("password", password)
	params.Add("request", "verifyUser")

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", erplyerr("failed to build HTTP request", err)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", erplyerr("failed to build VerifyUser request", err)
	}

	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erplyerr("failed to decode VerifyUserResponse", err)
	}

	return res.Records[0].SessionKey, nil
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode erply customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
func NewClient(sessionKey string, clientCode string, customCli *http.Client) IClient {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,

		MaxIdleConns:    MaxIdleConns,
		MaxConnsPerHost: MaxConnsPerHost,
	}

	cli := erplyClient{
		sessionKey: sessionKey,
		clientCode: clientCode,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   5 * time.Second,
		},
	}
	if customCli != nil {
		cli.httpClient = customCli
	}
	return &cli
}

func NewClientV2(partnerKey string, secret string, clientCode string) IClient {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,

		ExpectContinueTimeout: 4 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,

		MaxIdleConns:    MaxIdleConns,
		MaxConnsPerHost: MaxConnsPerHost,
	}

	cli := erplyClient{
		partnerKey: partnerKey,
		secret:     secret,
		clientCode: clientCode,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   5 * time.Second,
		},
	}
	return &cli
}

func (cli *erplyClient) Close() {
	cli.httpClient.CloseIdleConnections()
}

func (cli *erplyClient) GetProductUnits() ([]ProductUnit, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProductUnits request", err)
	}

	params := getMandatoryParameters(cli, GetProductUnitsMethod)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProductUnits request failed", err)
	}

	res := &GetProductUnitsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductUnitsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.ProductUnits, nil
}

//GetProductsByIDs - NOTE: if product's id is 0 - the product is not in the database. It was created during the sales document creation
func (cli *erplyClient) GetProductsByIDs(ids []string) ([]Product, error) {
	if len(ids) == 0 {
		return nil, erplyerr("No ids provided for products request", nil)
	}

	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProducts request", err)
	}

	params := getMandatoryParameters(cli, GetProductsMethod)
	params.Add("productIDs", strings.Join(ids, ","))
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProducts request failed", err)
	}

	res := &GetProductsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Products, nil
}

func (cli *erplyClient) GetProductsByCode3(code3 string) (*Product, error) {
	if code3 == "" {
		return nil, erplyerr("No code3 provided for product request", nil)
	}

	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProducts request", err)
	}

	params := getMandatoryParameters(cli, GetProductsMethod)
	params.Add("code3", code3)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProducts request failed", err)
	}

	res := &GetProductsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Products) == 0 {
		return nil, erplyerr("no such product found", err)
	}

	return &res.Products[0], nil
}

func (cli *erplyClient) GetAddresses() (*Address, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetAddresses request", err)
	}

	params := getMandatoryParameters(cli, GetAddressesMethod)
	params.Add("ownerID", "1")
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetAddresses request failed", err)
	}

	res := &GetAddressesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetAddressesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Addresses) == 0 {
		return nil, erplyerr(fmt.Sprintf("Addresses were not found"), nil)
	}
	return &res.Addresses[0], nil
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

func doRequest(req *http.Request, cli *erplyClient) (*http.Response, error) {
	resp, err := cli.httpClient.Do(req)
	return resp, err
}

// GetCustomers erply API request
// Takes customerID string array, converts to a string separated by comma parameter and adds to params(Example: 1,2,3,4)
func (cli *erplyClient) GetCustomersByIDs(customerID []string) (Customers, error) {
	if len(customerID) == 0 {
		return nil, erplyerr("Customer id list can not be empty", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetCustomers request", err)
	}
	customerIDs := strings.Join(customerID[:], ",")
	params := getMandatoryParameters(cli, GetCustomersMethod)
	params.Add("responseMode", "detail")
	params.Add("getContactPersons", "1")
	params.Add("getAddresses", "1")

	params.Add("customerIDs", customerIDs)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetCustomers request failed", err)
	}
	res := &GetCustomersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetCustomersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Customers) == 0 {
		return nil, nil
	}

	return res.Customers, nil
}

func (cli *erplyClient) GetCustomerByRegNumber(regNumber string) (*Customer, error) {
	if regNumber == "" {
		return nil, erplyerr("Customer reg number can not be empty", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetCustomers request", err)
	}
	//regNumbers := strings.Join(regNumbers[:], ",")
	params := getMandatoryParameters(cli, GetCustomersMethod)
	params.Add("responseMode", "detail")
	params.Add("getContactPersons", "1")
	params.Add("getAddresses", "1")

	params.Add("searchRegistryCode", regNumber)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetCustomers request failed", err)
	}
	res := &GetCustomersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetCustomersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Customers) == 0 {
		return nil, nil
	}
	return &res.Customers[0], nil
}

func (cli *erplyClient) GetCustomerByGLN(gln string) (*Customer, error) {
	if gln == "" {
		return nil, erplyerr("Customer gln can not be empty", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetCustomers request", err)
	}
	//regNumbers := strings.Join(regNumbers[:], ",")
	params := getMandatoryParameters(cli, GetCustomersMethod)
	params.Add("responseMode", "detail")
	params.Add("getContactPersons", "1")
	params.Add("getAddresses", "1")

	params.Add("searchAttributeName", "GLN")
	params.Add("searchAttributeValue", gln)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetCustomers request failed", err)
	}
	res := &GetCustomersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetCustomersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Customers) == 0 {
		return nil, nil
	}
	return &res.Customers[0], nil
}

func (cli *erplyClient) GetSupplierByName(name string) (*Customer, error) {
	if name == "" {
		return nil, erplyerr("Customer reg number can not be empty", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetCustomers request", err)
	}
	//regNumbers := strings.Join(regNumbers[:], ",")
	params := getMandatoryParameters(cli, getSuppliersMethod)

	params.Add("searchName", name)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetSupplierBySupplierName request failed", err)
	}
	res := &GetCustomersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetSuppliersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Customers) == 0 {
		return nil, nil
	}
	return &res.Customers[0], nil
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

//GetVatRatesByVatRateID ...
func (cli *erplyClient) GetVatRatesByID(vatRateID string) (VatRates, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetVatRates request", err)
	}

	params := getMandatoryParameters(cli, GetVatRatesMethod)
	params.Add("searchAttributeName", "id")
	params.Add("searchAttributeValue", vatRateID)
	params.Add("active", "1")
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetVatRates request failed", err)
	}
	res := &GetVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetVatRatesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.VatRates, nil
}

//GetCompanyInfo ...
func (cli *erplyClient) GetCompanyInfo() (*CompanyInfo, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetCompanyInfo request", err)
	}

	params := getMandatoryParameters(cli, GetCompanyInfoMethod)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetCompanyInfo request failed", err)
	}
	res := &GetCompanyInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetCompanyInfoResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CompanyInfos) == 0 {
		return nil, nil
	}

	return &res.CompanyInfos[0], nil
}

//GetPointsOfSale ...
func (cli *erplyClient) GetPointsOfSaleByID(posID string) (*PointOfSale, error) {
	method := GetPointsOfSaleMethod
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}

	params := getMandatoryParameters(cli, method)
	params.Add("pointOfSaleID", posID)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}
	res := &GetPointsOfSaleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.PointsOfSale) == 0 {
		return nil, nil
	}

	return &res.PointsOfSale[0], nil
}

//VerifyIdentityToken ...
func (cli *erplyClient) VerifyIdentityToken(jwt string) (*SessionInfo, error) {
	method := VerifyIdentityTokenMethod
	params := url.Values{}
	params.Add("request", method)
	params.Add("clientCode", cli.clientCode)
	params.Add("setContentType", "1")
	params.Add("jwt", jwt)
	req, err := newPostHTTPRequest(cli, params)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}

	res := &verifyIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

//GetIdentityToken ...
func (cli *erplyClient) GetIdentityToken() (*IdentityToken, error) {
	method := GetIdentityToken

	params := getMandatoryParameters(cli, method)
	queryParams := getMandatoryParameters(cli, method)

	req, err := newPostHTTPRequest(cli, params)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}
	req.URL.RawQuery = queryParams.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}
	res := &getIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

//GetWarehouses ...
func (cli *erplyClient) GetWarehouses() (Warehouses, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetWarehouses request", err)
	}

	params := getMandatoryParameters(cli, GetWarehousesMethod)
	params.Add("warehouseID", "0")
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetWarehouses request failed", err)
	}
	res := &GetWarehousesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetWarehousesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res.Warehouses, nil
}

func (cli *erplyClient) GetConfParameters() (*ConfParameter, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetConfParameters request", err)
	}

	params := url.Values{}
	params.Add("request", GetConfParametersMethod)
	params.Add(sessionKey, cli.sessionKey)
	params.Add(clientCode, cli.clientCode)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetConfParameters request failed", err)
	}
	res := &GetConfParametersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetConfParametersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ConfParameters) == 0 {
		return nil, erplyerr(fmt.Sprint("Conf Parameters were not found", nil), err)
	}

	return &res.ConfParameters[0], nil
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

func (cli *erplyClient) PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error) {
	if in.CompanyName == "" || in.RegistryCode == "" {
		return nil, erplyerr("Can not save customer with empty name or registry number", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build postCustomer request", err)
	}
	params := getMandatoryParameters(cli, saveCustomerMethod)
	params.Add("companyName", in.CompanyName)
	params.Add("fullName", in.FullName)
	params.Add("code", in.RegistryCode)
	params.Add("vatNumber", in.VatNumber)
	params.Add("email", in.Email)
	params.Add("phone", in.Phone)
	params.Add("bankName", in.BankName)
	params.Add("bankAccountNumber", in.BankAccountNumber)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
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
		return nil, nil
	}

	return &res.CustomerImportReports[0], nil
}

func (cli *erplyClient) PostSupplier(in *CustomerConstructor) (*CustomerImportReport, error) {
	if in.CompanyName == "" || in.RegistryCode == "" {
		return nil, erplyerr("Can not save customer with empty name or registry number", nil)
	}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build PostSupplier request", err)
	}
	params := getMandatoryParameters(cli, saveSupplierMethod)
	params.Add("companyName", in.CompanyName)
	params.Add("fullName", in.FullName)
	params.Add("code", in.RegistryCode)
	params.Add("vatNumber", in.VatNumber)
	params.Add("email", in.Email)
	params.Add("phone", in.Phone)
	params.Add("bankName", in.BankName)
	params.Add("bankAccountNumber", in.BankAccountNumber)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("PostSupplier request failed", err)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling CustomerImportReport failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, nil
	}

	return &res.CustomerImportReports[0], nil
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

func isJSONResponseOK(status *Status) bool {
	return strings.EqualFold(status.ResponseStatus, responseStatusOK)
}

func getHTTPRequest(cli *erplyClient) (*http.Request, error) {
	req, err := http.NewRequest("GET", cli.url, nil)
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	return req, err
}

func newPostHTTPRequest(cli *erplyClient, params url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	return req, err
}

func getMandatoryParameters(cli *erplyClient, request string) url.Values {
	params := url.Values{}
	params.Add("request", request)
	params.Add("setContentType", "1")
	if cli.sessionKey != "" && cli.clientCode != "" {
		params.Add(sessionKey, cli.sessionKey)
		params.Add(clientCode, cli.clientCode)
	}
	if cli.partnerKey != "" && cli.secret != "" {
		now := time.Now().Unix()
		params.Add(applicationKey, GenerateToken(cli.partnerKey, now, request, cli.secret))
		params.Add(clientCode, cli.clientCode)
		params.Add("partnerKey", cli.partnerKey)
		params.Add("timestamp", strconv.Itoa(int(now)))
	}
	return params
}

type Status struct {
	Request           string  `json:"request"`
	RequestUnixTime   int     `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	GenerationTime    float64 `json:"generationTime"`
	RecordsTotal      int     `json:"recordsTotal"`
	RecordsInResponse int     `json:"recordsInResponse"`
}

type PostSalesDocumentResponse struct {
	Status        Status               `json:"status"`
	ImportReports SaleDocImportReports `json:"records"`
}

type PostPurchaseDocumentResponse struct {
	Status        Status                   `json:"status"`
	ImportReports PurchaseDocImportReports `json:"records"`
}

//GetVatRatesResponse ...
type GetVatRatesResponse struct {
	Status   Status    `json:"status"`
	VatRates []VatRate `json:"records"`
}

//GetCustomersResponse ...
type GetCustomersResponse struct {
	Status    Status    `json:"status"`
	Customers Customers `json:"records"`
}

type GetSalesDocumentResponse struct {
	Status         Status         `json:"status"`
	SalesDocuments []SaleDocument `json:"records"`
}

//CustomerDataProcessingLog ...
type CustomerDataProcessingLog struct {
	activityType string
	fields       []string
	customerIds  []int
}

//GetCompanyInfoResponse ...
type GetCompanyInfoResponse struct {
	Status       Status       `json:"status"`
	CompanyInfos CompanyInfos `json:"records"`
}

type GetPointsOfSaleResponse struct {
	Status       Status        `json:"status"`
	PointsOfSale []PointOfSale `json:"records"`
}

//GetConfParametersResponse ...
type GetConfParametersResponse struct {
	Status         Status          `json:"status"`
	ConfParameters []ConfParameter `json:"records"`
}

type GetWarehousesResponse struct {
	Status     Status     `json:"status"`
	Warehouses Warehouses `json:"records"`
}

type PostCustomerResponse struct {
	Status                Status                `json:"status"`
	CustomerImportReports CustomerImportReports `json:"records"`
}

func erplyerr(msg string, err error) *erro.ErplyError {
	if err != nil {
		return erro.NewErplyError("Error", errors.Wrap(err, msg).Error())
	}
	return erro.NewErplyError("Error", msg)
}
