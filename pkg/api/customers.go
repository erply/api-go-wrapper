package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
	"strings"
)

type Customer struct {
	ID                   int            `json:"id"`
	CustomerID           int            `json:"customerID"`
	TypeID               string         `json:"type_id"`
	FullName             string         `json:"fullName"`
	CompanyName          string         `json:"companyName"`
	FirstName            string         `json:"firstName"`
	LastName             string         `json:"lastName"`
	GroupID              int            `json:"groupID"`
	EDI                  string         `json:"EDI"`
	IsPOSDefaultCustomer int            `json:"isPOSDefaultCustomer"`
	CountryID            string         `json:"countryID"`
	Phone                string         `json:"phone"`
	EInvoiceEmail        string         `json:"eInvoiceEmail"`
	Email                string         `json:"email"`
	Fax                  string         `json:"fax"`
	Code                 string         `json:"code"`
	ReferenceNumber      string         `json:"referenceNumber"`
	VatNumber            string         `json:"vatNumber"`
	BankName             string         `json:"bankName"`
	BankAccountNumber    string         `json:"bankAccountNumber"`
	BankIBAN             string         `json:"bankIBAN"`
	BankSWIFT            string         `json:"bankSWIFT"`
	PaymentDays          int            `json:"paymentDays"`
	Notes                string         `json:"notes"`
	LastModified         int            `json:"lastModified"`
	CustomerType         string         `json:"customerType"`
	Address              string         `json:"address"`
	CustomerAddresses    Addresses      `json:"addresses"`
	Street               string         `json:"street"`
	Address2             string         `json:"address2"`
	City                 string         `json:"city"`
	PostalCode           string         `json:"postalCode"`
	Country              string         `json:"country"`
	State                string         `json:"state"`
	ContactPersons       ContactPersons `json:"contactPersons"`

	// Web-shop related fields
	Username  string `json:"webshopUsername"`
	LastLogin string `json:"webshopLastLogin"`
}
type Customers []Customer

//Attribute field
type Attribute struct {
	Name  string `json:"attributeNam"`
	Type  string `json:"attributeType"`
	Value string `json:"attributeValue"`
}

type CustomerConstructor struct {
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

//GetCustomersResponse ...
type GetCustomersResponse struct {
	Status    Status    `json:"status"`
	Customers Customers `json:"records"`
}

type PostCustomerResponse struct {
	Status                Status                `json:"status"`
	CustomerImportReports CustomerImportReports `json:"records"`
}

/*func (cli *erplyClient) PostCustomer(in *CustomerConstructor) (*CustomerImportReport, error) {
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
}*/
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
