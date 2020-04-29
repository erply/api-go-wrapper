package api

import (
	"context"
	"encoding/json"
	"errors"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	VatRate struct {
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

	VatRates            []VatRate
	VatTotalsByTaxRates []VatTotalsByTaxRate

	VatTotalsByTaxRate struct {
		VatrateID int     `json:"vatrateID"`
		Total     float64 `json:"total"`
	}

	NetTotalsByTaxRate struct {
		VatrateID int     `json:"vatrateID"`
		Total     float64 `json:"total"`
	}

	//GetVatRatesResponse ...
	getVatRatesResponse struct {
		Status   Status    `json:"status"`
		VatRates []VatRate `json:"records"`
	}
	VatRateManager interface {
		GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error)
	}
)

//GetVatRatesByVatRateID ...
func (cli *erplyClient) GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error) {

	resp, err := cli.sendRequest(ctx, GetVatRatesMethod, filters)
	if err != nil {
		return nil, err
	}
	res := &getVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetVatRatesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if res.VatRates == nil {
		return nil, errors.New("no vat rates in response")
	}
	return res.VatRates, nil
}
