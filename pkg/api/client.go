package api

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/erply/api-go-wrapper/pkg/api/documents"
	"github.com/erply/api-go-wrapper/pkg/api/prices"

	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api/addresses"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/api/company"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/erply/api-go-wrapper/pkg/api/pos"
	"github.com/erply/api-go-wrapper/pkg/api/products"
	"github.com/erply/api-go-wrapper/pkg/api/sales"
	"github.com/erply/api-go-wrapper/pkg/api/servicediscovery"
	"github.com/erply/api-go-wrapper/pkg/api/warehouse"
)

type Client struct {
	commonClient *common.Client
	//Address requests
	AddressProvider addresses.Manager
	//Token requests
	AuthProvider auth.Provider
	//Company and Conf parameter requests
	CompanyManager company.Manager
	//Customers and suppliers requests
	CustomerManager customers.Manager
	//POS related requests
	PosManager pos.Manager
	//ListingDataProvider related requests
	ProductManager products.Manager
	//SalesDocuments, Payments, Projects, ShoppingCart, VatRates
	SalesManager sales.Manager
	//Warehouse requests
	WarehouseManager warehouse.Manager
	//Prices requests
	PricesManager prices.Manager
	//Documents requests
	DocumentsManager documents.Manager
	//Service Discovery
	ServiceDiscoverer servicediscovery.ServiceDiscoverer
}

//NewUnvalidatedClient returns a new Client without validating any of the incoming parameters giving the
//developer more flexibility
func NewUnvalidatedClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {
	comCli := common.NewClient(sk, cc, partnerKey, httpCli, nil)
	return newErplyClient(comCli)
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode erply customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
//The headersSetToEveryRequest function will be executed on every request and supplied with the request name. There is an example in the /examples of you to use it
func NewClient(sessionKey string, clientCode string, customCli *http.Client) (*Client, error) {
	if sessionKey == "" || clientCode == "" {
		return nil, errors.New("sessionKey and clientCode are required")
	}
	comCli := common.NewClient(sessionKey, clientCode, "", customCli, nil)
	return newErplyClient(comCli), nil
}

//NewClientWithCustomHeaders enables defining the function that will set headers to every request by your own
func NewClientWithCustomHeaders(customHTTPCli *http.Client, headersSetToEveryRequest func(requestName string) url.Values) (*Client, error) {
	if headersSetToEveryRequest == nil {
		return nil, errors.New("the function that will set headers to every request is a required argument")
	}
	return newErplyClient(common.NewClient("", "", "", customHTTPCli, headersSetToEveryRequest)), nil
}

//NewClientWithURL creates a new Client which can have a static URL which is not affected by clientCode
// nor the headersSetToEveryRequest function if set. If the url parameter is set to an empty string, the URL
// is still resolved normally. This allows creating clients which have a static url in your unit tests but function
// normally in the rest of your code
func NewClientWithURL(sk, cc, partnerKey, url string, httpCli *http.Client, headersSetToEveryRequest func(requestName string) url.Values) *Client {
	comCli := common.NewClientWithURL(sk, cc, partnerKey, url, httpCli, headersSetToEveryRequest)
	return newErplyClient(comCli)
}

func newErplyClient(c *common.Client) *Client {
	return &Client{
		commonClient:      c,
		AddressProvider:   addresses.NewClient(c),
		AuthProvider:      auth.NewClient(c),
		CompanyManager:    company.NewClient(c),
		CustomerManager:   customers.NewClient(c),
		PosManager:        pos.NewClient(c),
		ProductManager:    products.NewClient(c),
		SalesManager:      sales.NewClient(c),
		WarehouseManager:  warehouse.NewClient(c),
		ServiceDiscoverer: servicediscovery.NewClient(c),
		PricesManager:     prices.NewClient(c),
		DocumentsManager:  documents.NewClient(c),
	}
}
