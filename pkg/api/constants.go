package api

const (
	baseURL = "https://%s.erply.com/api/"

	GetSalesDocumentsMethod           = "getSalesDocuments"
	GetCustomersMethod                = "getCustomers"
	getSuppliersMethod                = "getSuppliers"
	GetCountriesMethod                = "getCountries"
	GetEmployeesMethod                = "getEmployees"
	GetBusinessAreasMethod            = "getBusinessAreas"
	GetProjectsMethod                 = "getProjects"
	GetProjectStatusesMethod          = "getProjectStatuses"
	GetCurrenciesMethod               = "getCurrencies"
	GetVatRatesMethod                 = "getVatRates"
	GetPaymentsMethod                 = "getPayments"
	//GetSalesDocumentsMethod ...
	GetSalesDocumentsMethod = "getSalesDocuments"
	GetUserRightsMethod     = "getUserRights"
	//GetCompanyInfoMethod ...
	GetCompanyInfoMethod              = "getCompanyInfo"
	VerifyIdentityTokenMethod         = "verifyIdentityToken"
	GetPointsOfSaleMethod             = "getPointsOfSale"
	GetIdentityToken                  = "getIdentityToken"
	GetJWTTokenMethod                 = "getJwtToken"
	GetConfParametersMethod           = "getConfParameters"
	logProcessingOfCustomerDataMethod = "logProcessingOfCustomerData"
	saveSalesDocumentMethod           = "saveSalesDocument"
	savePurchaseDocumentMethod        = "savePurchaseDocument"
	saveCustomerMethod                = "saveCustomer"
	saveAddressMethod                 = "saveAddress"
	saveSupplierMethod                = "saveSupplier"
	savePaymentMethod                 = "savePayment"
	createInstallationMethod          = "createInstallation"
	GetWarehousesMethod               = "getWarehouses"
	GetAddressesMethod                = "getAddresses"
	GetProductsMethod                 = "getProducts"
	GetProductCategoriesMethod        = "getProductCategories"
	GetProductBrandsMethod            = "getBrands"
	GetProductGroupsMethod            = "getProductGroups"
	GetProductUnitsMethod             = "getProductUnits"
	VerifyCustomerUserMethod          = "verifyCustomerUser"
	calculateShoppingCartMethod       = "calculateShoppingCart"
	validateCustomerUsernameMethod    = "validateCustomerUsername"
	clientCode                        = "clientCode"
	sessionKey                        = "sessionKey"
	applicationKey                    = "applicationKey"
	responseStatusOK                  = "ok"
	Cash                              = "CASH"
	Card                              = "CARD"
	Transfer                          = "TRANSFER"
	Check                             = "CHECK"
	Paid                              = "PAID"
	Unpaid                            = "UNPAID"
	BankTransfer                      = "Direct bank transfer"
	CheckPayment                      = "Check payments"
	PayPal                            = "PayPal"
	CashOnDelivery                    = "Cash on delivery"
	//MaxIdleConns for Erply API
	MaxIdleConns = 25

	//MaxConnsPerHost for Erply API
	MaxConnsPerHost = 25
)
