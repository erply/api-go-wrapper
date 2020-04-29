package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

//IClient interface for cached and simple client
type IClient interface {
	ServiceDiscoverer
	TokenProvider

	ConfManager
	WarehouseManager
	GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error)

	//sales document requests
	GetSalesDocumentByID(id string) ([]SaleDocument, error)
	GetSalesDocumentsByIDs(id []string) ([]SaleDocument, error)
	PostSalesDocumentFromWoocomm(in *SaleDocumentConstructor, shopOrderID string) (SaleDocImportReports, error)
	PostSalesDocument(in *SaleDocumentConstructor, provider string) (SaleDocImportReports, error)
	DeleteDocumentsByID(id string) error

	//customer requests
	GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
	GetCustomersByIDs(customerID []string) (Customers, error)
	GetCustomerByRegNumber(regNumber string) (*Customer, error)
	GetCustomerByGLN(gln string) (*Customer, error)
	//PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error)

	//supplier requests
	GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
	GetSupplierByName(name string) (*Customer, error)
	PostSupplier(in *CustomerConstructor) (*CustomerImportReport, error)

	VatRateManager
	CompanyManager

	//product requests
	GetProductUnits() ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductsByIDs(ids []string) ([]Product, error)
	GetProductsByCode3(code3 string) (*Product, error)

	GetCountries(ctx context.Context, filters map[string]string) ([]Country, error)
	GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error)
	GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error)

	//project requests
	GetProjects(ctx context.Context, filters map[string]string) ([]Project, error)
	GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error)

	GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error)
	PostPurchaseDocument(in *PurchaseDocumentConstructor, provider string) (PurchaseDocImportReports, error)

	PointOfSaleManager

	//payment requests
	SavePayment(in *PaymentInfo) (int64, error)
	GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)

	AddressManager
	VerifyCustomerUser(username, password string) (*WebshopClient, error)
	CalculateShoppingCart(in *DocumentData) (*ShoppingCartTotals, error)
	IsCustomerUsernameAvailable(username string) (bool, error)
	Close()
}

type IPartnerClient interface {
	IClient
	PartnerTokenProvider
}

type erplyClient struct {
	sessionKey string
	clientCode string
	partnerKey string
	url        string
	httpClient *http.Client
}

//VerifyUser will give you session key
func VerifyUser(username, password, clientCode string, client *http.Client) (string, error) {
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
	if len(res.Records) < 1 {
		return "", erplyerr("VerifyUser: no records in response", nil)
	}
	return res.Records[0].SessionKey, nil
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode erply customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
func NewClient(sessionKey, clientCode string, customCli *http.Client) (IClient, error) {

	if sessionKey == "" || clientCode == "" {
		return nil, errors.New("sessionKey and clientCode are required")
	}

	cli := erplyClient{
		sessionKey: sessionKey,
		clientCode: clientCode,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: getDefaultHTTPClient(),
	}
	if customCli != nil {
		cli.httpClient = customCli
	}
	return &cli, nil
}

func NewPartnerClient(sessionKey, clientCode, partnerKey string, customCli *http.Client) (IPartnerClient, error) {
	if sessionKey == "" || clientCode == "" || partnerKey == "" {
		return nil, errors.New("sessionKey, clientCode and partnerKey are required")
	}

	cli := erplyClient{
		sessionKey: sessionKey,
		clientCode: clientCode,
		partnerKey: partnerKey,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: getDefaultHTTPClient(),
	}
	if customCli != nil {
		cli.httpClient = customCli
	}
	return &cli, nil
}

func (cli *erplyClient) Close() {
	cli.httpClient.CloseIdleConnections()
}

func getDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,

			ExpectContinueTimeout: 4 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,

			MaxIdleConns:    MaxIdleConns,
			MaxConnsPerHost: MaxConnsPerHost,
		},
		Timeout: 5 * time.Second,
	}
}
