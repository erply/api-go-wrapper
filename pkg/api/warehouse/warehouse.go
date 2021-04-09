package warehouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

//GetWarehouses ...
func (cli *Client) GetWarehouses(ctx context.Context, filters map[string]string) (Warehouses, error) {
	resp, err := cli.SendRequest(ctx, "getWarehouses", filters)
	if err != nil {
		return nil, err
	}

	res := &GetWarehousesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling GetWarehousesResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res.Warehouses, nil
}

//GetWarehousesWithStatus ...
func (cli *Client) GetWarehousesWithStatus(ctx context.Context, filters map[string]string) (*GetWarehousesResponse, error) {
	resp, err := cli.SendRequest(ctx, "getWarehouses", filters)
	if err != nil {
		return nil, err
	}

	res := &GetWarehousesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling GetWarehousesResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res, nil
}

func (cli *Client) GetWarehousesBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (
	GetWarehousesResponseBulk,
	error,
) {
	var bulkResp GetWarehousesResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getWarehouses",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetWarehousesResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, bulkResp.Status.ErrorCode)
		}
	}

	return bulkResp, nil
}

func (cli *Client) SaveWarehouse(ctx context.Context, filters map[string]string) (*SaveWarehouseResult, error) {
	resp, err := cli.SendRequest(ctx, "saveWarehouse", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("saveWarehouse request failed", err, 0)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SaveWarehouseResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal SaveWarehouseResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Results) == 0 {
		return nil, nil
	}

	return &res.Results[0], nil
}

func (cli *Client) SaveWarehouseBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveWarehouseResponseBulk, error) {
	var bulkResp SaveWarehouseResponseBulk

	if len(bulkRequest) > sharedCommon.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d warehouses in one bulk request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveWarehouse",
			Filters:    bulkInput,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveWarehouseResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
				bulkResp.Status.ErrorCode,
			)
		}
	}

	return bulkResp, nil
}
