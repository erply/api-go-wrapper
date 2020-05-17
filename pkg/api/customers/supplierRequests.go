package customers

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"strconv"
)

// GetSuppliers will list suppliers according to specified filters.
func (cli *Client) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, common.Status, error) {
	resp, err := cli.SendRequest(ctx, "getSuppliers", filters)
	if err != nil {
		return nil, common.Status{}, err
	}
	var res GetSuppliersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, common.Status{}, erro.NewFromError("failed to unmarshal GetSuppliersResponse ", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, common.Status{}, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Suppliers, res.Status, nil
}

// GetSuppliersBulk will list suppliers according to specified filters sending a bulk request to fetch more suppliers than the default limit
func (cli *Client) GetSuppliersBulk(ctx context.Context, bulkFilters []map[string]string, baseFilters map[string]string) (GetSuppliersResponseBulk, error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSuppliers",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return GetSuppliersResponseBulk{}, err
	}

	var suppliersResp GetSuppliersResponseBulk
	if err := json.NewDecoder(resp.Body).Decode(&suppliersResp); err != nil {
		return GetSuppliersResponseBulk{}, erro.NewFromError("failed to unmarshal GetSuppliersResponseBulk ", err)
	}
	if !common.IsJSONResponseOK(&suppliersResp.Status) {
		return GetSuppliersResponseBulk{}, erro.NewErplyError(strconv.Itoa(suppliersResp.Status.ErrorCode), suppliersResp.Status.Request+": "+suppliersResp.Status.ResponseStatus)
	}

	return suppliersResp, nil
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
