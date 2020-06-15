package addresses

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

func (cli *Client) GetAddresses(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error) {
	resp, err := cli.SendRequest(ctx, "getAddresses", filters)
	if err != nil {
		return nil, erro.NewFromError("GetAddresses request failed", err)
	}

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetAddressesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Addresses, nil
}

// GetAddressesBulk will list addresses according to specified filters sending a bulk request to fetch more addresses than the default limit
func (cli *Client) GetAddressesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetAddressesResponseBulk, error) {
	var addrResp GetAddressesResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getAddresses",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return addrResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return addrResp, err
	}

	if err := json.Unmarshal(body, &addrResp); err != nil {
		return addrResp, fmt.Errorf("ERPLY API: failed to unmarshal GetAddressesResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&addrResp.Status) {
		return addrResp, erro.NewErplyError(addrResp.Status.ErrorCode.String(), addrResp.Status.Request+": "+addrResp.Status.ResponseStatus)
	}

	for _, addrBulkItem := range addrResp.BulkItems {
		if !common.IsJSONResponseOK(&addrBulkItem.Status.Status) {
			return addrResp, erro.NewErplyError(addrBulkItem.Status.ErrorCode.String(), addrBulkItem.Status.Request+": "+addrBulkItem.Status.ResponseStatus)
		}
	}

	return addrResp, nil
}

func (cli *Client) SaveAddress(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error) {
	method := "saveAddress"
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, erro.NewFromError(method+": request failed", err)
	}
	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError(method+": JSON unmarshal failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Addresses) == 0 {
		return nil, erro.NewFromError(method+": no records in response", nil)
	}

	return res.Addresses, nil
}

func (cli *Client) SaveAddressesBulk(ctx context.Context, addrMap []map[string]interface{}, attrs map[string]string) (SaveAddressesResponseBulk, error) {
	var saveAddressesResponseBulk SaveAddressesResponseBulk

	if len(addrMap) > sharedCommon.MaxBulkRequestsCount {
		return saveAddressesResponseBulk, fmt.Errorf("cannot save more than %d addresses in one request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(addrMap))
	for _, addr := range addrMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveAddress",
			Filters:    addr,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return saveAddressesResponseBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return saveAddressesResponseBulk, err
	}

	if err := json.Unmarshal(body, &saveAddressesResponseBulk); err != nil {
		return saveAddressesResponseBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveAddressesResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&saveAddressesResponseBulk.Status) {
		return saveAddressesResponseBulk, erro.NewErplyError(saveAddressesResponseBulk.Status.ErrorCode.String(), saveAddressesResponseBulk.Status.Request+": "+saveAddressesResponseBulk.Status.ResponseStatus)
	}

	for _, addrBulkItem := range saveAddressesResponseBulk.BulkItems {
		if !common.IsJSONResponseOK(&addrBulkItem.Status.Status) {
			return saveAddressesResponseBulk, erro.NewErplyError(
				addrBulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", addrBulkItem.Status),
			)
		}
	}

	return saveAddressesResponseBulk, nil
}

