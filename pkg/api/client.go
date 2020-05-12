package api

import (
	"errors"
	"net/http"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/addresses"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/auth"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/company"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/customers"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/pos"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/products"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/sales"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/servicediscovery"
	"github.com/tarmo-randma/api-go-wrapper/pkg/api/warehouse"
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
	//Products related requests
	ProductManager products.Manager
	//SalesDocuments, Payments, Projects, ShoppingCart, VatRates
	SalesManager sales.Manager
	//Warehouse requests
	WarehouseManager warehouse.Manager

	//Service Discovery
	ServiceDiscoverer servicediscovery.ServiceDiscoverer
}

//NewUnvalidatedClient returns a new Client without validating any of the incoming parameters giving the
//developer more flexibility
func NewUnvalidatedClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {
	comCli := common.NewClient(sk, cc, partnerKey, httpCli)
	return newErplyClient(comCli)
}

// NewClient Takes three params:
// sessionKey string obtained from credentials or jwt
// clientCode erply customer identification number
// and a custom http Client if needs to be overwritten. if nil will use default http client provided by the SDK
func NewClient(sessionKey string, clientCode string, customCli *http.Client) (*Client, error) {

	if sessionKey == "" || clientCode == "" {
		return nil, errors.New("sessionKey and clientCode are required")
	}
	comCli := common.NewClient(sessionKey, clientCode, "", customCli)
	return newErplyClient(comCli), nil
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
	}
}
