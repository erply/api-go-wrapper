package api

//IClient interface for cached and simple client
type IClient interface {
	GetConfParameters() (*ConfParameter, error)
	GetWarehouses() (Warehouses, error)
	GetUserName() (string, error)
	GetSalesDocumentByID(id string) ([]SaleDocument, error)
	GetSalesDocumentsByIDs(id []string) ([]SaleDocument, error)
	GetCustomersByIDs(customerID []string) (Customers, error)
	GetCustomerByRegNumber(regNumber string) (*Customer, error)
	GetCustomerByGLN(gln string) (*Customer, error)
	GetSupplierByName(name string) (*Customer, error)
	GetVatRatesByID(vatRateID string) (VatRates, error)
	GetCompanyInfo() (*CompanyInfo, error)
	GetProductUnits() ([]ProductUnit, error)
	GetProductsByIDs(ids []string) ([]Product, error)
	GetProductsByCode3(code3 string) (*Product, error)
	GetAddresses() (*Address, error)
	PostPurchaseDocument(in *PurchaseDocumentConstructor, provider string) (PurchaseDocImportReports, error)
	PostSalesDocumentFromWoocomm(in *SaleDocumentConstructor, shopOrderID string) (SaleDocImportReports, error)
	PostSalesDocument(in *SaleDocumentConstructor, provider string) (SaleDocImportReports, error)
	PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error)
	PostSupplier(in *CustomerConstructor) (*CustomerImportReport, error)
	DeleteDocumentsByID(id string) error
	GetPointsOfSaleByID(posID string) (*PointOfSale, error)
	VerifyIdentityToken(jwt string) (*SessionInfo, error)
	GetIdentityToken() (*IdentityToken, error)
	Close()
}
