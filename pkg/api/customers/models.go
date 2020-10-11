package customers

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	Customer struct {
		ID                      int                         `json:"id"`
		PayerID                 int                         `json:"payerID,omitempty"`
		CustomerID              int                         `json:"customerID"`
		TypeID                  string                      `json:"type_id"`
		FullName                string                      `json:"fullName"`
		CompanyName             string                      `json:"companyName"`
		FirstName               string                      `json:"firstName"`
		LastName                string                      `json:"lastName"`
		GroupID                 int                         `json:"groupID"`
		EDI                     string                      `json:"EDI"`
		GLN                     string                      `json:"GLN"`
		IsPOSDefaultCustomer    int                         `json:"isPOSDefaultCustomer"`
		CountryID               string                      `json:"countryID"`
		Phone                   string                      `json:"phone"`
		EInvoiceEmail           string                      `json:"eInvoiceEmail"`
		Email                   string                      `json:"email"`
		Fax                     string                      `json:"fax"`
		Code                    string                      `json:"code"`
		ReferenceNumber         string                      `json:"referenceNumber"`
		VatNumber               string                      `json:"vatNumber"`
		BankName                string                      `json:"bankName"`
		BankAccountNumber       string                      `json:"bankAccountNumber"`
		BankIBAN                string                      `json:"bankIBAN"`
		BankSWIFT               string                      `json:"bankSWIFT"`
		PaymentDays             int                         `json:"paymentDays"`
		Notes                   string                      `json:"notes"`
		LastModified            int                         `json:"lastModified"`
		CustomerType            string                      `json:"customerType"`
		Address                 string                      `json:"address"`
		CustomerAddresses       sharedCommon.Addresses      `json:"addresses"`
		Street                  string                      `json:"street"`
		Address2                string                      `json:"address2"`
		City                    string                      `json:"city"`
		PostalCode              string                      `json:"postalCode"`
		Country                 string                      `json:"country"`
		State                   string                      `json:"state"`
		ContactPersons          ContactPersons              `json:"contactPersons"`
		Attributes              []sharedCommon.ObjAttribute `json:"attributes"`
		Credit                  int                         `json:"credit"`
		CompanyTypeID           int                         `json:"companyTypeID"`
		PersonTitleID           int                         `json:"personTitleID"`
		EmailEnabled            int                         `json:"emailEnabled"`
		MailEnabled             int                         `json:"mailEnabled"`
		EInvoiceEnabled         int                         `json:"eInvoiceEnabled"`
		FlagStatus              int                         `json:"flagStatus"`
		OperatorIdentifier      string                      `json:"operatorIdentifier"`
		Gender                  string                      `json:"gender"`
		GroupName               string                      `json:"groupName"`
		Mobile                  string                      `json:"mobile"`
		Birthday                string                      `json:"birthday"`
		IntegrationCode         string                      `json:"integrationCode"`
		ColorStatus             string                      `json:"colorStatus"`
		FactoringContractNumber string                      `json:"factoringContractNumber"`
		Image                   string                      `json:"image"`
		TwitterID               string                      `json:"twitterID"`
		FacebookName            string                      `json:"facebookName"`
		CreditCardLastNumbers   string                      `json:"creditCardLastNumbers"`
		EuCustomerType          string                      `json:"euCustomerType"`
		CustomerCardNumber      string                      `json:"customerCardNumber"`
		LastModifierUsername    string                      `json:"lastModifierUsername"`
		DefaultAssociationName  string                      `json:"defaultAssociationName"`
		DefaultProfessionalName string                      `json:"defaultProfessionalName"`
		TaxExempt               int                         `json:"taxExempt"`
		PaysViaFactoring        int                         `json:"paysViaFactoring"`
		SalesBlocked            int                         `json:"salesBlocked"`
		RewardPointsDisabled    int                         `json:"rewardPointsDisabled"`
		CustomerBalanceDisabled int                         `json:"customerBalanceDisabled"`
		PosCouponsDisabled      int                         `json:"posCouponsDisabled"`
		EmailOptOut             int                         `json:"emailOptOut"`
		ShipGoodsWithWaybills   int                         `json:"shipGoodsWithWaybills"`
		DefaultAssociationID    int                         `json:"defaultAssociationID"`
		DefaultProfessionalID   int                         `json:"defaultProfessionalID"`

		// Web-shop related fields
		Username  string `json:"webshopUsername"`
		LastLogin string `json:"webshopLastLogin"`
	}
	ContactPersons []ContactPerson
	ContactPerson  struct {
		ContactPersonID   int    `json:"contactPersonID"`
		FullName          string `json:"fullName"`
		GroupName         string `json:"groupName"`
		CountryID         string `json:"countryID"`
		Phone             string `json:"phone"`
		Email             string `json:"email"`
		Fax               string `json:"fax"`
		Code              string `json:"code"`
		BankName          string `json:"bankName"`
		BankAccountNumber string `json:"bankAccountNumber"`
		BankIBAN          string `json:"bankIBAN"`
		BankSWIFT         string `json:"bankSWIFT"`
		Notes             string `json:"notes"`
	}
	Customers []Customer

	//Attribute field
	Attribute struct {
		Name  string `json:"attributeNam"`
		Type  string `json:"attributeType"`
		Value string `json:"attributeValue"`
	}

	CustomerRequest struct {
		CustomerID        int
		CompanyName       string
		Address           string
		PostalCode        string
		AddressTypeID     int
		City              string
		State             string
		Country           string
		FirstName         string
		LastName          string
		FullName          string
		RegistryCode      string
		VatNumber         string
		Email             string
		Phone             string
		BankName          string
		BankAccountNumber string

		// Web-shop related fields
		Username string
		Password string
	}

	WebshopClient struct {
		ClientID        string `json:"clientID"`
		ClientUsername  string `json:"clientUsername"`
		ClientName      string `json:"clientName"`
		ClientFirstName string `json:"clientFirstName"`
		ClientLastName  string `json:"clientLastName"`
		ClientGroupID   string `json:"clientGroupID"`
		ClientGroupName string `json:"clientGroupName"`
		CompanyID       string `json:"companyID"`
		CompanyName     string `json:"companyName"`
	}
	GetCustomersResponse struct {
		Status    sharedCommon.Status `json:"status"`
		Customers Customers           `json:"records"`
	}

	PostCustomerResponse struct {
		Status                sharedCommon.Status   `json:"status"`
		CustomerImportReports CustomerImportReports `json:"records"`
	}
	CustomerImportReports []CustomerImportReport
	CustomerImportReport  struct {
		ClientID   int `json:"clientID"`
		CustomerID int `json:"customerID"`
	}

	GetCustomersResponseBulkItem struct {
		Status    sharedCommon.StatusBulk `json:"status"`
		Customers Customers               `json:"records"`
	}

	GetCustomersResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []GetCustomersResponseBulkItem `json:"requests"`
	}
)

type (
	Supplier struct {
		SupplierId      uint                        `json:"supplierID"`
		SupplierType    string                      `json:"supplierType"`
		FullName        string                      `json:"fullName"`
		CompanyName     string                      `json:"companyName"`
		FirstName       string                      `json:"firstName"`
		LstName         string                      `json:"lastName"`
		GroupId         uint                        `json:"groupID"`
		GroupName       string                      `json:"groupName"`
		Phone           string                      `json:"phone"`
		Mobile          string                      `json:"mobile"`
		Email           string                      `json:"email"`
		Fax             string                      `json:"fax"`
		Code            string                      `json:"code"`
		IntegrationCode string                      `json:"integrationCode"`
		VatrateID       uint                        `json:"vatrateID"`
		CurrencyCode    string                      `json:"currencyCode"`
		DeliveryTermsID uint                        `json:"deliveryTermsID"`
		CountryId       uint                        `json:"countryID"`
		CountryName     string                      `json:"countryName"`
		CountryCode     string                      `json:"countryCode"`
		Address         string                      `json:"address"`
		Gln             string                      `json:"GLN"`
		Attributes      []sharedCommon.ObjAttribute `json:"attributes"`

		// Detail fields
		VatNumber           string `json:"vatNumber"`
		Skype               string `json:"skype"`
		Website             string `json:"website"`
		BankName            string `json:"bankName"`
		BankAccountNumber   string `json:"bankAccountNumber"`
		BankIBAN            string `json:"bankIBAN"`
		BankSWIFT           string `json:"bankSWIFT"`
		Birthday            string `json:"birthday"`
		CompanyID           uint   `json:"companyID"`
		ParentCompanyName   string `json:"parentCompanyName"`
		SupplierManagerID   uint   `json:"supplierManagerID"`
		SupplierManagerName string `json:"supplierManagerName"`
		PaymentDays         uint   `json:"paymentDays"`
		Notes               string `json:"notes"`
		LastModified        string `json:"lastModified"`
		Added               uint64 `json:"added"`
	}

	//SaveSupplierResp
	SaveSupplierResp struct {
		SupplierID    int  `json:"supplierID"`
		AlreadyExists bool `json:"alreadyExists"`
	}

	//GetSuppliersResponse
	GetSuppliersResponse struct {
		Status    sharedCommon.Status `json:"status"`
		Suppliers []Supplier          `json:"records"`
	}

	GetSuppliersResponseBulkItem struct {
		Status    sharedCommon.StatusBulk `json:"status"`
		Suppliers []Supplier              `json:"records"`
	}

	GetSuppliersResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []GetSuppliersResponseBulkItem `json:"requests"`
	}

	SaveSuppliersResponseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveSupplierResp      `json:"records"`
	}

	SaveSuppliersResponseBulk struct {
		Status    sharedCommon.Status             `json:"status"`
		BulkItems []SaveSuppliersResponseBulkItem `json:"requests"`
	}

	DeleteSupplierResponse struct {
		Status sharedCommon.Status `json:"status"`
	}

	DeleteSuppliersResponseBulkItem struct {
		Status sharedCommon.StatusBulk `json:"status"`
	}

	DeleteSuppliersResponseBulk struct {
		Status    sharedCommon.Status               `json:"status"`
		BulkItems []DeleteSuppliersResponseBulkItem `json:"requests"`
	}

	AddCustomerRewardPointsResult struct {
		TransactionID   int64 `json:"transactionID"`
		CustomerID      int64 `json:"customerID"`
		Points          int64 `json:"points"`
		CreatedUnixTime int64 `json:"createdUnixTime"`
		ExpiryUnixTime  int64 `json:"expiryUnixTime"`
	}

	AddCustomerRewardPointsResponse struct {
		Status                         sharedCommon.Status             `json:"status"`
		AddCustomerRewardPointsResults []AddCustomerRewardPointsResult `json:"records"`
	}

	AddCustomerRewardPointsResponseBulkItem struct {
		Status                         sharedCommon.StatusBulk         `json:"status"`
		AddCustomerRewardPointsResults []AddCustomerRewardPointsResult `json:"records"`
	}

	AddCustomerRewardPointsResponseBulk struct {
		Status    sharedCommon.Status                       `json:"status"`
		BulkItems []AddCustomerRewardPointsResponseBulkItem `json:"requests"`
	}
)
