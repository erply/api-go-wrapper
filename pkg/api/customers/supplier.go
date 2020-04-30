package customers

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

/*
type (
	Supplier struct {
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
		Code            string             `json:"code"`
		IntegrationCode string             `json:"integrationCode"`
		VatrateID       uint               `json:"vatrateID"`
		CurrencyCode    string             `json:"currencyCode"`
		DeliveryTermsID uint               `json:"deliveryTermsID"`
		CountryId       uint               `json:"countryID"`
		CountryName     string             `json:"countryName"`
		CountryCode     string             `json:"countryCode"`
		Address         string             `json:"address"`
		Gln             string             `json:"GLN"`
		Attributes      []api.ObjAttribute `json:"attributes"`

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
	getSuppliersResponse struct {
		Status    common.Status `json:"status"`
		Suppliers []Supplier `json:"records"`
	}
	SupplierManager interface {
		GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
		PostSupplier(ctx context.Context, filters map[string]string) (*api.CustomerImportReport, error)
	}
)

// GetSuppliers will list suppliers according to specified filters.
func (cli *Client) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error) {
	resp, err := cli.SendRequest(ctx, api.getSuppliersMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getSuppliersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetSuppliersResponse ", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Suppliers, nil
}
func (cli *Client) PostSupplier(ctx context.Context, filters map[string]string) (*api.CustomerImportReport, error) {
	resp, err := cli.SendRequest(ctx, api.saveSupplierMethod, filters)
	if err != nil {
		return nil, erro.NewFromError("PostSupplier request failed", err)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling CustomerImportReport failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, nil
	}

	return &res.CustomerImportReports[0], nil
}
*/
