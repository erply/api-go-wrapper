package sales

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

/*
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
		Status   common.Status `json:"status"`
		VatRates []VatRate  `json:"records"`
	}
	VatRateManager interface {
		GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error)
	}
)

//GetVatRatesByVatRateID ...
func (cli *Client) GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error) {

	resp, err := cli.SendRequest(ctx, api.GetVatRatesMethod, filters)
	if err != nil {
		return nil, err
	}
	res := &getVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetVatRatesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if res.VatRates == nil {
		return nil, errors.New("no vat rates in response")
	}
	return res.VatRates, nil
}
*/
