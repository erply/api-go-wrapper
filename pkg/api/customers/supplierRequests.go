package customers

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
	erro "github.com/tarmo-randma/api-go-wrapper/internal/errors"
)

// GetSuppliers will list suppliers according to specified filters.
func (cli *Client) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error) {
	resp, err := cli.SendRequest(ctx, "getSuppliers", filters)
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
func (cli *Client) SaveSupplier(ctx context.Context, filters map[string]string) (*CustomerImportReport, error) {
	resp, err := cli.SendRequest(ctx, "saveSupplier", filters)
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
