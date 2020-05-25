package customers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"io/ioutil"
)

// GetSuppliers will list suppliers according to specified filters.
func (cli *Client) GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error) {
	resp, err := cli.SendRequest(ctx, "getSuppliers", filters)
	if err != nil {
		return nil, err
	}
	var res GetSuppliersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetSuppliersResponse ", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Suppliers, nil
}

// GetSuppliersBulk will list suppliers according to specified filters sending a bulk request to fetch more suppliers than the default limit
func (cli *Client) GetSuppliersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSuppliersResponseBulk, error) {
	var suppliersResp GetSuppliersResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSuppliers",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return suppliersResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return suppliersResp, err
	}

	if err := json.Unmarshal(body, &suppliersResp); err != nil {
		return suppliersResp, fmt.Errorf("ERPLY API: failed to unmarshal GetSuppliersResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&suppliersResp.Status) {
		return suppliersResp, erro.NewErplyError(suppliersResp.Status.ErrorCode.String(), suppliersResp.Status.Request+": "+suppliersResp.Status.ResponseStatus)
	}

	for _, supplierBulkItem := range suppliersResp.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return suppliersResp, erro.NewErplyError(supplierBulkItem.Status.ErrorCode.String(), supplierBulkItem.Status.Request+": "+supplierBulkItem.Status.ResponseStatus)
		}
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, nil
	}

	return &res.CustomerImportReports[0], nil
}

func (cli *Client) SaveSupplierBulk(ctx context.Context, supplierMap []map[string]interface{}, attrs map[string]string) (SaveSuppliersResponseBulk, error) {
	var saveSuppliersResponseBulk SaveSuppliersResponseBulk

	if len(supplierMap) > common.MaxBulkRequestsCount {
		return saveSuppliersResponseBulk, fmt.Errorf("cannot save more than %d suppliers in one request", common.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(supplierMap))
	for _, supplier := range supplierMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveSupplier",
			Filters:    supplier,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return saveSuppliersResponseBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return saveSuppliersResponseBulk, err
	}

	if err := json.Unmarshal(body, &saveSuppliersResponseBulk); err != nil {
		return saveSuppliersResponseBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveSuppliersResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&saveSuppliersResponseBulk.Status) {
		return saveSuppliersResponseBulk, erro.NewErplyError(saveSuppliersResponseBulk.Status.ErrorCode.String(), saveSuppliersResponseBulk.Status.Request+": "+saveSuppliersResponseBulk.Status.ResponseStatus)
	}

	for _, supplierBulkItem := range saveSuppliersResponseBulk.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return saveSuppliersResponseBulk, erro.NewErplyError(
				supplierBulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", supplierBulkItem.Status),
			)
		}
	}

	return saveSuppliersResponseBulk, nil
}
