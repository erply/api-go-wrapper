package sales

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

//GetVatRatesByVatRateID ...
func (cli *Client) GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error) {
	resp, err := cli.SendRequest(ctx, "getVatRates", filters)
	if err != nil {
		return nil, err
	}
	res := &GetVatRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetVatRatesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	if res.VatRates == nil {
		return nil, errors.New("no vat rates in response")
	}
	return res.VatRates, nil
}

func (cli *Client) GetVatRatesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetVatRatesResponseBulk, error) {
	var bulkResp GetVatRatesResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getVatRates",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetVatRatesResponseBulk from '%s': %v", string(body), err)
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


func (cli *Client) SaveVatRate(ctx context.Context, filters map[string]string) (*SaveVatRateResult, error) {
	resp, err := cli.SendRequest(ctx, "saveVatRate", filters)
	if err != nil {
		return nil, erro.NewFromError("saveVatRate request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SaveVatRateResultResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal SaveVatRateResultResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.SaveVatRateResult) == 0 {
		return nil, nil
	}

	return &res.SaveVatRateResult[0], nil
}

func (cli *Client) SaveVatRateBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateResponseBulk, error) {
	var bulkResp SaveVatRateResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d price lists in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveVatRate",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveVatRateResponseBulk from '%s': %v", string(body), err)
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

func (cli *Client) SaveVatRateComponent(ctx context.Context, filters map[string]string) (*SaveVatRateComponentResult, error) {
	resp, err := cli.SendRequest(ctx, "saveVatRateComponent", filters)
	if err != nil {
		return nil, erro.NewFromError("saveVatRateComponent request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SaveVatRateComponentResultResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal SaveVatRateComponentResultResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.SaveVatRateComponentResult) == 0 {
		return nil, nil
	}

	return &res.SaveVatRateComponentResult[0], nil
}

func (cli *Client) SaveVatRateComponentBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateComponentResponseBulk, error) {
	var bulkResp SaveVatRateComponentResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d price lists in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveVatRateComponent",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveVatRateComponentResponseBulk from '%s': %v", string(body), err)
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
