package sales

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
)

//GetVatRatesByVatRateID ...
func (cli *Client) GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error) {
	resp, err := cli.SendRequest(ctx, "getVatRates", filters)
	if err != nil {
		return nil, err
	}
	res := &getVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetVatRatesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if res.VatRates == nil {
		return nil, errors.New("no vat rates in response")
	}
	return res.VatRates, nil
}
