package prices

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"io/ioutil"
)

func (cli *Client) GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error) {
	resp, err := cli.SendRequest(ctx, "getSupplierPriceLists", filters)
	if err != nil {
		return nil, err
	}
	var res GetPriceListsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []PriceList{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetPriceListsResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.PriceLists, nil
}

func (cli *Client) GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error) {
	var bulkResp GetPriceListsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSupplierPriceLists",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetPriceListsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) GetProductPriceLists(ctx context.Context, filters map[string]string) ([]ProductPriceList, error) {
	resp, err := cli.SendRequest(ctx, "getProductsInSupplierPriceList", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductPriceListResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []ProductPriceList{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetProductPriceListResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.ProductPriceLists, nil
}

func (cli *Client) GetProductPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductPriceListResponseBulk, error) {
	var bulkResp GetProductPriceListResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductsInSupplierPriceList",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetProductPriceListResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}
