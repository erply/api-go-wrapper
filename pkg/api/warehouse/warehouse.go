package warehouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
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
		return nil, erro.NewFromError("unmarshaling GetWarehousesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res.Warehouses, nil
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
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) SaveWarehouse(ctx context.Context, filters map[string]string) (*SaveWarehouseResult, error) {
	resp, err := cli.SendRequest(ctx, "saveWarehouse", filters)
	if err != nil {
		return nil, erro.NewFromError("saveWarehouse request failed", err)
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
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.Results) == 0 {
		return nil, nil
	}

	return &res.Results[0], nil
}

func (cli *Client) SaveWarehouseBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveWarehouseResponseBulk, error) {
	var bulkResp SaveWarehouseResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d warehouses in one bulk request", common2.MaxBulkRequestsCount)
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
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}
