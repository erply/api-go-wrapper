package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	erro "github.com/erply/api-go-wrapper/pkg/errors"
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

func (cli *erplyClient) GetProducts(ctx context.Context, filters map[string]string) ([]Product, error) {
	resp, err := cli.sendRequest(ctx, GetProductsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProductsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Products, nil
}

func (cli *erplyClient) GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error) {
	resp, err := cli.sendRequest(ctx, GetProductCategoriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductCategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductCategoriesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductCategories, nil
}

func (cli *erplyClient) GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error) {
	resp, err := cli.sendRequest(ctx, GetProductBrandsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductBrandsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductBrandsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductBrands, nil
}

func (cli *erplyClient) GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error) {
	resp, err := cli.sendRequest(ctx, GetProductGroupsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductGroupsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductGroups, nil
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

func (cli *erplyClient) GetAddresses(filters map[string]string) ([]Address, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetAddresses request", err)
	}

	params := getMandatoryParameters(cli, GetAddressesMethod)
	for fk, fv := range filters {
		params.Add(fk, fv)
	}

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

	return res.Addresses, nil
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

// GetProjects will list projects according to specified filters.
func (cli *erplyClient) GetProjects(ctx context.Context, filters map[string]string) ([]Project, error) {
	resp, err := cli.sendRequest(ctx, GetProjectsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProjectsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Projects, nil
}

// GetProjectStatus will list projects statuses according to specified filters.
func (cli *erplyClient) GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error) {
	resp, err := cli.sendRequest(ctx, GetProjectStatusesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectStatusesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProjectStatusesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProjectStatuses, nil
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

// GetSuppliers will list suppliers according to specified filters.
func (cli *erplyClient) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error) {
	resp, err := cli.sendRequest(ctx, getSuppliersMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetSuppliersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetSuppliersResponse ", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Suppliers, nil
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

// GetPointsOfSale will list points of sale according to specified filters.
func (cli *erplyClient) GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error) {
	resp, err := cli.sendRequest(ctx, GetPointsOfSaleMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetPointsOfSaleResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetPointsOfSaleResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.PointsOfSale, nil
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

func (cli *erplyClient) GetJWTToken(partnerKey string) (*JwtToken, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, fmt.Errorf("error building GetJWTToken request: %v", err)
	}

	params := getMandatoryParameters(cli, GetJWTTokenMethod)
	params.Set("partnerKey", partnerKey)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("error making request for GetJWTToken", err)
	}

	var res JwtTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, erplyerr("error decoding GetJWTToken response", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Records, nil
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

func (cli *erplyClient) SaveAddress(in *AddressRequest) (int, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return 0, erplyerr("SaveAddress: failed to build request", err)
	}
	params := getMandatoryParameters(cli, saveAddressMethod)
	params.Add("addressID", strconv.Itoa(in.AddressID))
	params.Add("typeID", strconv.Itoa(in.TypeID))
	params.Add("ownerID", strconv.Itoa(in.OwnerID))
	params.Add("street", in.Street)
	params.Add("postalCode", in.PostalCode)
	params.Add("city", in.City)
	params.Add("state", in.State)
	params.Add("country", in.Country)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return 0, erplyerr("SaveAddress: request failed", err)
	}

	var res struct {
		Status  Status
		Records []Address
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, erplyerr("SaveAddress: JSON unmarshal failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return 0, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Records) == 0 {
		return 0, erplyerr("SaveAddress: no records in response", nil)
	}

	return res.Records[0].AddressID, nil
}

func (cli *erplyClient) PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error) {
	//if in.CompanyName == "" || in.RegistryCode == "" {
	//	return nil, erplyerr("Can not save customer with empty name or registry number", nil)
	//}
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build postCustomer request", err)
	}
	params := getMandatoryParameters(cli, saveCustomerMethod)
	params.Add("customerID", strconv.Itoa(in.CustomerID)) // For updating the existing customer
	params.Add("companyName", in.CompanyName)
	params.Add("firstName", in.FirstName)
	params.Add("lastName", in.LastName)
	params.Add("fullName", in.FullName)
	params.Add("code", in.RegistryCode)
	params.Add("vatNumber", in.VatNumber)
	params.Add("email", in.Email)
	params.Add("phone", in.Phone)
	params.Add("bankName", in.BankName)
	params.Add("bankAccountNumber", in.BankAccountNumber)

	if in.Username != "" {
		if in.Password == "" {
			return nil, erplyerr("password for user can not be empty", nil)
		}

		params.Add("username", in.Username)
		params.Add("password", in.Password)
	}

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

	if in.Address != "" {
		var addressID int

		// existing customer is being updated -
		// overwrite first address record of that customer
		// or simply create a new address record if there are no existing addresses
		if in.CustomerID > 0 {
			addrs, err := cli.GetAddresses(map[string]string{
				"ownerID":    strconv.Itoa(in.CustomerID),
				"orderBy":    "addressID",
				"orderByDir": "ASC",
			})

			if err != nil {
				return nil, erplyerr("error getting existing customer addresses: %v", err)
			}
			if len(addrs) > 0 {
				addressID = addrs[0].AddressID
			}
		}

		addr := AddressRequest{
			AddressID:  addressID,
			OwnerID:    res.CustomerImportReports[0].CustomerID,
			TypeID:     1,
			Street:     in.Address,
			PostalCode: in.PostalCode,
			Country:    in.Country,
			City:       in.City,
			State:      in.State,
		}

		_, err := cli.SaveAddress(&addr)
		if err != nil {
			return nil, erplyerr("error adding address to customer", err)
		}
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

func CreateInstallation(baseUrl, partnerKey string, in *InstallationRequest, cli *http.Client) (*InstallationResponse, error) {

	params := url.Values{}
	params.Add("request", createInstallationMethod)
	params.Add("partnerKey", partnerKey)
	params.Add("companyName", in.CompanyName)
	params.Add("firstName", in.FirstName)
	params.Add("lastName", in.LastName)
	params.Add("phone", in.Phone)
	params.Add("email", in.Email)
	params.Add("sendEmail", strconv.Itoa(in.SendEmail))

	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, &erplyClient{httpClient: cli})
	if err != nil {
		return nil, erplyerr("CreateInstallation: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, erplyerr(fmt.Sprintf("CreateInstallation: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erplyerr("CreateInstallation: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		fmt.Println(respData.Status.ErrorField)
		return nil, erplyerr(fmt.Sprintf("CreateInstallation: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return nil, erplyerr("CreateInstallation: no records in response", nil)
	}

	return &respData.Records[0], nil
}

func (cli *erplyClient) SavePayment(in *PaymentInfo) (int64, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return 0, erplyerr("SavePayment: failed to build request", err)
	}

	params := getMandatoryParameters(cli, savePaymentMethod)
	params.Add("documentID", strconv.Itoa(in.DocumentID))
	params.Add("type", in.Type)
	params.Add("currencyCode", in.CurrencyCode)
	params.Add("date", in.Date)
	params.Add("sum", in.Sum)
	params.Add("info", in.Info)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return 0, erplyerr("SavePayment: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, erplyerr(fmt.Sprintf("SavePayment: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []struct {
			PaymentID int64 `json:"paymentID"`
		} `json:"records"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, erplyerr("SavePayment: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, erplyerr(fmt.Sprintf("SavePayment: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return 0, erplyerr("SavePayment: no records in response", nil)
	}

	return respData.Records[0].PaymentID, nil
}

func (cli *erplyClient) GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("GetPayments: failed to build request", err)
	}

	params := getMandatoryParameters(cli, GetPaymentsMethod)
	for k, v := range filters {
		params.Add(k, v)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetPayments: error sending request", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, erplyerr(fmt.Sprintf("GetPayments: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []PaymentInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erplyerr("GetPayments: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, erplyerr(fmt.Sprintf("GetPayments: API error %d", respData.Status.ErrorCode), nil)
	}

	return respData.Records, nil
}

func (cli *erplyClient) CalculateShoppingCart(in *DocumentData) (*ShoppingCartTotals, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("CalculateShoppingCart: failed to build request", err)
	}

	params := getMandatoryParameters(cli, calculateShoppingCartMethod)
	params.Add("customerID", strconv.FormatUint(uint64(in.CustomerId), 10))

	for i, prod := range in.ProductRows {
		params.Add(fmt.Sprintf("productID%d", i), prod.ProductID)
		params.Add(fmt.Sprintf("amount%d", i), prod.Amount)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("CalculateShoppingCart: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, erplyerr(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, erplyerr("CalculateShoppingCart: unmarshaling response failed", err)
	}
	if !isJSONResponseOK(&respData.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(respData.Status.ErrorCode), respData.Status.Request+": "+respData.Status.ResponseStatus)
	}
	if len(respData.Records) < 1 {
		return nil, erplyerr("CalculateShoppingCart: no records in response", nil)
	}

	return respData.Records[0], nil
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

type Status struct {
	Request           string  `json:"request"`
	RequestUnixTime   int     `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	ErrorField        string  `json:"errorField"`
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

//GetSuppliersResponse
type GetSuppliersResponse struct {
	Status    Status     `json:"status"`
	Suppliers []Supplier `json:"records"`
}

type GetCountriesResponse struct {
	Status    Status    `json:"status"`
	Countries []Country `json:"records"`
}

type GetEmployeesResponse struct {
	Status    Status     `json:"status"`
	Employees []Employee `json:"records"`
}

type GetBusinessAreasResponse struct {
	Status        Status         `json:"status"`
	BusinessAreas []BusinessArea `json:"records"`
}

type GetProjectsResponse struct {
	Status   Status    `json:"status"`
	Projects []Project `json:"records"`
}

type GetProjectStatusesResponse struct {
	Status          Status          `json:"status"`
	ProjectStatuses []ProjectStatus `json:"records"`
}

type GetCurrenciesResponse struct {
	Status     Status     `json:"status"`
	Currencies []Currency `json:"records"`
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
