package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	//CompanyInfos ..
	CompanyInfos []CompanyInfo
	//CompanyInfo ..
	CompanyInfo struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		Code               string `json:"code"`
		VAT                string `json:"VAT"`
		Phone              string `json:"phone"`
		Mobile             string `json:"mobile"`
		Fax                string `json:"fax"`
		Email              string `json:"email"`
		Web                string `json:"web"`
		BankAccountNumber  string `json:"bankAccountNumber"`
		BankName           string `json:"bankName"`
		BankSWIFT          string `json:"bankSWIFT"`
		BankIBAN           string `json:"bankIBAN"`
		BankAccountNumber2 string `json:"bankAccountNumber2"`
		BankName2          string `json:"bankName2"`
		BankSWIFT2         string `json:"bankSWIFT2"`
		BankIBAN2          string `json:"bankIBAN2"`
		Address            string `json:"address"`
		Country            string `json:"country"`

		//field for ConfParameters
		ConfParameters ConfParameter
	} //GetCompanyInfoResponse ...
	GetCompanyInfoResponse struct {
		Status       Status       `json:"status"`
		CompanyInfos CompanyInfos `json:"records"`
	}

	CompanyManager interface {
		GetCompanyInfo(ctx context.Context) (*CompanyInfo, error)
	}
)

//GetCompanyInfo ...
func (cli *erplyClient) GetCompanyInfo(ctx context.Context) (*CompanyInfo, error) {

	resp, err := cli.sendRequest(ctx, GetCompanyInfoMethod, map[string]string{})
	if err != nil {
		return nil, erplyerr("GetCompanyInfo request failed", err)
	}
	res := &GetCompanyInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetCompanyInfoResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CompanyInfos) == 0 {
		return nil, nil
	}

	return &res.CompanyInfos[0], nil
}
