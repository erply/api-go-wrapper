package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type Supplier struct {
	SupplierId      uint           `json:"supplierID"`
	SupplierType    string         `json:"supplierType"`
	FullName        string         `json:"fullName"`
	CompanyName     string         `json:"companyName"`
	FirstName       string         `json:"firstName"`
	LstName         string         `json:"lastName"`
	GroupId         uint           `json:"groupID"`
	GroupName       string         `json:"groupName"`
	Phone           string         `json:"phone"`
	Mobile          string         `json:"mobile"`
	Email           string         `json:"email"`
	Fax             string         `json:"fax"`
	Code            string         `json:"code"`
	IntegrationCode string         `json:"integrationCode"`
	VatrateID       uint           `json:"vatrateID"`
	CurrencyCode    string         `json:"currencyCode"`
	DeliveryTermsID uint           `json:"deliveryTermsID"`
	CountryId       uint           `json:"countryID"`
	CountryName     string         `json:"countryName"`
	CountryCode     string         `json:"countryCode"`
	Address         string         `json:"address"`
	Gln             string         `json:"GLN"`
	Attributes      []ObjAttribute `json:"attributes"`

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

//GetSuppliersResponse
type GetSuppliersResponse struct {
	Status    Status     `json:"status"`
	Suppliers []Supplier `json:"records"`
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
