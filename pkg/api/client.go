package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api/addresses"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/company"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/erply/api-go-wrapper/pkg/api/pos"
	"github.com/erply/api-go-wrapper/pkg/api/products"
	"github.com/erply/api-go-wrapper/pkg/api/sales"
	"github.com/erply/api-go-wrapper/pkg/api/servicediscovery"
	"github.com/erply/api-go-wrapper/pkg/api/warehouse"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"net/url"
)

//IClient interface for cached and simple client
type IClient interface {
	/*	servicediscovery.ServiceDiscoverer
		auth.TokenProvider
		configuration.ConfManager
		warehouse.WarehouseManager

		GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error)

		//sales document requests
		sales.SalesDocumentManager

		//customer requests
		customers.CustomerManager

		//supplier requests
		customers.SupplierManager

		sales.VatRateManager
		company.CompanyManager

		products.ProductManager

		GetCountries(ctx context.Context, filters map[string]string) ([]Country, error)
		GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error)
		GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error)

		//project requests
		sales.ProjectManager

		GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error)
		SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error)

		pos.PointOfSaleManager
		sales.PaymentManager*/

	//CalculateShoppingCart(in *DocumentData) (*sales.ShoppingCartTotals, error)
	//
	//Close()
}

type IPartnerClient interface {
	IClient
	auth.PartnerTokenProvider
}

type erplyClient struct {
	AddressProvider addresses.Manager
	//Token requests
	AuthProvider auth.Provider
	//Company and Conf parameter requests
	CompanyManager company.Manager
	//Customers and suppliers requests
	CustomerManager customers.Manager
	PosManager      pos.Manager
	ProductManager  products.Manager

	//SalesDocuments, Payments, Projects, ShoppingCart, VatRates
	SalesManager sales.Manager

	WarehouseManager  warehouse.Manager
	ServiceDiscoverer servicediscovery.ServiceDiscoverer
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
		return "", erro.NewFromError("failed to build HTTP request", err)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", erro.NewFromError("failed to build VerifyUser request", err)
	}

	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erro.NewFromError("failed to decode VerifyUserResponse", err)
	}
	if len(res.Records) < 1 {
		return "", erro.NewFromError("VerifyUser: no records in response", nil)
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

	//declare short getters
	var (
		//sessionKey
		s = sessionKey
		//clientCode
		c = clientCode
		h = customCli
	)
	cli := &erplyClient{
		AddressProvider:   addresses.NewClient(s, c, "", h),
		AuthProvider:      auth.NewClient(s, c, "", h),
		CompanyManager:    company.NewClient(s, c, "", h),
		CustomerManager:   customers.NewClient(s, c, "", h),
		PosManager:        pos.NewClient(s, c, "", h),
		ProductManager:    products.NewClient(s, c, "", h),
		SalesManager:      sales.NewClient(s, c, "", h),
		WarehouseManager:  warehouse.NewClient(s, c, "", h),
		ServiceDiscoverer: servicediscovery.NewClient(s, c, "", h),
	}

	return cli, nil
}
