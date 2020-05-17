package customers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"strconv"
)

// GetSuppliers will list suppliers according to specified filters.
func (cli *Client) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error) {
	bulkInputI := ctx.Value("bulk")
	if bulkInputI != nil {
		bulkFilters, ok := bulkInputI.([]map[string]string)
		if !ok {
			return []Supplier{}, fmt.Errorf("unexpected bulk value provided: %+v", bulkInputI)
		}

		return cli.GetSuppliersBulk(ctx, bulkFilters, filters)
	}

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

// GetSuppliers will list suppliers according to specified filters sending a bulk request to fetch more suppliers than the default limit
func (cli *Client) GetSuppliersBulk(ctx context.Context, bulkFilters []map[string]string, baseFilters map[string]string) ([]Supplier, error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSuppliers",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return nil, err
	}

	var suppliersResp getSuppliersResponseBulk
	if err := json.NewDecoder(resp.Body).Decode(&suppliersResp); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetSuppliersResponseBulk ", err)
	}
	if !common.IsJSONResponseOK(&suppliersResp.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(suppliersResp.Status.ErrorCode), suppliersResp.Status.Request+": "+suppliersResp.Status.ResponseStatus)
	}

	suppliers := make([]Supplier, 0)
	for _, supplierBulkItem := range suppliersResp.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return suppliers, erro.NewErplyError(strconv.Itoa(supplierBulkItem.Status.ErrorCode), supplierBulkItem.Status.Request+": "+supplierBulkItem.Status.ResponseStatus)
		}
		suppliers = append(suppliers, supplierBulkItem.Suppliers...)
	}

	return suppliers, nil
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
