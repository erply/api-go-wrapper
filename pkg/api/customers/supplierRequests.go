package customers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
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
		return nil, sharedCommon.NewFromError("failed to unmarshal GetSuppliersResponse ", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
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
		return suppliersResp, sharedCommon.NewErplyError(suppliersResp.Status.ErrorCode.String(), suppliersResp.Status.Request+": "+suppliersResp.Status.ResponseStatus, suppliersResp.Status.ErrorCode)
	}

	for _, supplierBulkItem := range suppliersResp.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return suppliersResp, sharedCommon.NewErplyError(supplierBulkItem.Status.ErrorCode.String(), supplierBulkItem.Status.Request+": "+supplierBulkItem.Status.ResponseStatus, suppliersResp.Status.ErrorCode)
		}
	}

	return suppliersResp, nil
}

func (cli *Client) SaveSupplier(ctx context.Context, filters map[string]string) (*CustomerImportReport, error) {
	resp, err := cli.SendRequest(ctx, "saveSupplier", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("PostSupplier request failed", err, 0)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling CustomerImportReport failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, nil
	}

	return &res.CustomerImportReports[0], nil
}

func (cli *Client) SaveSupplierBulk(ctx context.Context, supplierMap []map[string]interface{}, attrs map[string]string) (SaveSuppliersResponseBulk, error) {
	var saveSuppliersResponseBulk SaveSuppliersResponseBulk

	if len(supplierMap) > sharedCommon.MaxBulkRequestsCount {
		return saveSuppliersResponseBulk, fmt.Errorf("cannot save more than %d suppliers in one request", sharedCommon.MaxBulkRequestsCount)
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
		return saveSuppliersResponseBulk, sharedCommon.NewErplyError(
			saveSuppliersResponseBulk.Status.ErrorCode.String(),
			saveSuppliersResponseBulk.Status.Request+": "+saveSuppliersResponseBulk.Status.ResponseStatus,
			saveSuppliersResponseBulk.Status.ErrorCode,
		)
	}

	for _, supplierBulkItem := range saveSuppliersResponseBulk.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return saveSuppliersResponseBulk, sharedCommon.NewErplyError(
				supplierBulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", supplierBulkItem.Status),
				supplierBulkItem.Status.ErrorCode,
			)
		}
	}

	return saveSuppliersResponseBulk, nil
}

// DeleteSupplier https://learn-api.erply.com/requests/deletesupplier/
func (cli *Client) DeleteSupplier(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteSupplier", filters)
	if err != nil {
		return err
	}
	var res DeleteSupplierResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError("failed to unmarshal DeleteSupplierResponse ", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return nil
}

func (cli *Client) DeleteSupplierBulk(ctx context.Context, supplierMap []map[string]interface{}, attrs map[string]string) (DeleteSuppliersResponseBulk, error) {
	var deleteSupplierResponse DeleteSuppliersResponseBulk

	if len(supplierMap) > sharedCommon.MaxBulkRequestsCount {
		return deleteSupplierResponse, fmt.Errorf("cannot delete more than %d suppliers in one request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(supplierMap))
	for _, filter := range supplierMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteSupplier",
			Filters:    filter,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return deleteSupplierResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deleteSupplierResponse, err
	}

	if err := json.Unmarshal(body, &deleteSupplierResponse); err != nil {
		return deleteSupplierResponse, fmt.Errorf("ERPLY API: failed to unmarshal DeleteSuppliersResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&deleteSupplierResponse.Status) {
		return deleteSupplierResponse, sharedCommon.NewErplyError(
			deleteSupplierResponse.Status.ErrorCode.String(),
			deleteSupplierResponse.Status.Request+": "+deleteSupplierResponse.Status.ResponseStatus,
			deleteSupplierResponse.Status.ErrorCode,
		)
	}

	for _, supplierBulkItem := range deleteSupplierResponse.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return deleteSupplierResponse, sharedCommon.NewErplyError(
				supplierBulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", supplierBulkItem.Status),
				deleteSupplierResponse.Status.ErrorCode,
			)
		}
	}

	return deleteSupplierResponse, nil
}

func (cli *Client) GetCompanyTypes(ctx context.Context, filters map[string]string) ([]CompanyType, error) {
	res := &GetCompanyTypesResponse{}

	err := cli.Scan(ctx, "getCompanyTypes", filters, res)
	if err != nil {
		return nil, err
	}

	return res.CompanyTypes, nil
}

func (cli *Client) SaveCompanyType(ctx context.Context, filters map[string]string) (*SaveCompanyTypeResponse, error) {
	resp, err := cli.SendRequest(ctx, "saveCompanyType", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("saveCompanyType request failed", err, 0)
	}
	res := &SaveCompanyTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling SaveCompanyType failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res, nil
}

func (cli *Client) SaveSupplierGroup(ctx context.Context, filters map[string]string) (*SaveSupplierGroupResponse, error) {
	resp, err := cli.SendRequest(ctx, "saveSupplierGroup", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("saveSupplierGroup request failed", err, 0)
	}
	res := &SaveSupplierGroupResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling SaveSupplierGroup failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res, nil
}
