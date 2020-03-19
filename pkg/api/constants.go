package api

const (
	baseURL = "https://%s.erply.com/api/"
	//GetSalesDocumentsMethod ...
	GetSalesDocumentsMethod = "getSalesDocuments"
	GetUserRightsMethod     = "getUserRights"
	//GetCustomersMethod ...
	GetCustomersMethod = "getCustomers"
	getSuppliersMethod = "getSuppliers"
	//GetVatRatesMethod ...
	GetVatRatesMethod = "getVatRates"
	//GetPaymentsMethod ...
	GetPaymentsMethod = "getPayments"
	//GetCompanyInfoMethod ...
	GetCompanyInfoMethod              = "getCompanyInfo"
	VerifyIdentityTokenMethod         = "verifyIdentityToken"
	GetPointsOfSaleMethod             = "getPointsOfSale"
	GetIdentityToken                  = "getIdentityToken"
	GetConfParametersMethod           = "getConfParameters"
	logProcessingOfCustomerDataMethod = "logProcessingOfCustomerData"
	saveSalesDocumentMethod           = "saveSalesDocument"
	savePurchaseDocumentMethod        = "savePurchaseDocument"
	saveCustomerMethod                = "saveCustomer"
	saveSupplierMethod                = "saveSupplier"
	GetWarehousesMethod               = "getWarehouses"
	GetAddressesMethod                = "getAddresses"
	GetProductsMethod                 = "getProducts"
	GetProductUnitsMethod             = "getProductUnits"
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
	MaxIdleConns = 1

	//MaxConnsPerHost for Erply API
	MaxConnsPerHost = 1
)
