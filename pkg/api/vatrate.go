package api

import (
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type NetTotalsByRates []NetTotalsByRate
type NetTotalsByRate struct {
	//Num1 float64 `json:"1"`
}
type VatTotalsByRates []VatTotalsByRate
type VatTotalsByRate struct {
	//Num1 float64 `json:"1"`
}
type VatRate struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rate   string `json:"rate"`
	Code   string `json:"code"`
	Active string `json:"active"`
	//Added        string `json:"added"`
	LastModified string `json:"lastModified"`
	//IsReverseVat int    `json:"isReverseVat"`
	//ReverseRate int `json:"reverseRate"`
}

type VatRates []VatRate
type VatTotalsByTaxRates []VatTotalsByTaxRate

type VatTotalsByTaxRate struct {
	VatrateID int     `json:"vatrateID"`
	Total     float64 `json:"total"`
}
type NetTotalsByTaxRates []NetTotalsByTaxRate
type NetTotalsByTaxRate struct {
	VatrateID int     `json:"vatrateID"`
	Total     float64 `json:"total"`
}

//GetVatRatesResponse ...
type GetVatRatesResponse struct {
	Status   Status    `json:"status"`
	VatRates []VatRate `json:"records"`
}

//GetVatRatesByVatRateID ...
func (cli *erplyClient) GetVatRatesByID(vatRateID string) (VatRates, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetVatRates request", err)
	}

	params := getMandatoryParameters(cli, GetVatRatesMethod)
	params.Add("searchAttributeName", "id")
	params.Add("searchAttributeValue", vatRateID)
	params.Add("active", "1")
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetVatRates request failed", err)
	}
	res := &GetVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetVatRatesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.VatRates, nil
}
