package customers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
	"net/http"
)

func (cli *Client) SaveCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error) {
	resp, err := cli.SendRequest(ctx, "saveCustomer", filters)
	if err != nil {
		return nil, erro.NewFromError("PostCustomer request failed", err)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling CustomerImportReport failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.CustomerImportReports) == 0 {
		return nil, errors.New("no report found")
	}

	return &res.CustomerImportReports[0], nil
}

// GetCustomers will list customers according to specified filters.
func (cli *Client) GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error) {
	resp, err := cli.SendRequest(ctx, "getCustomers", filters)
	if err != nil {
		return nil, err
	}
	var res GetCustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetCustomersResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.Customers, nil
}

// GetCustomersBulk will list customers according to specified filters sending a bulk request to fetch more customers than the default limit
func (cli *Client) GetCustomersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetCustomersResponseBulk, error) {
	var customersResponse GetCustomersResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getCustomers",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return customersResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return customersResponse, err
	}

	if err := json.Unmarshal(body, &customersResponse); err != nil {
		return customersResponse, fmt.Errorf("ERPLY API: failed to unmarshal GetCustomersResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&customersResponse.Status) {
		return customersResponse, erro.NewErplyError(customersResponse.Status.ErrorCode.String(), customersResponse.Status.Request+": "+customersResponse.Status.ResponseStatus)
	}

	for _, supplierBulkItem := range customersResponse.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return customersResponse, erro.NewErplyError(supplierBulkItem.Status.ErrorCode.String(), supplierBulkItem.Status.Request+": "+supplierBulkItem.Status.ResponseStatus)
		}
	}

	return customersResponse, nil
}

//username and password are required fields here
func (cli *Client) VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error) {
	filters := map[string]string{
		"username": username,
		"password": password,
	}
	resp, err := cli.SendRequest(ctx, "verifyCustomerUser", filters)
	if err != nil {
		return nil, erro.NewFromError("VerifyCustomerUser: request failed", err)
	}

	var res struct {
		Status  common2.Status
		Records []WebshopClient
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("VerifyCustomerUser: unmarhsalling response failed", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	if len(res.Records) != 1 {
		return nil, erro.NewFromError("VerifyCustomerUser: no records in response", nil)
	}

	return &res.Records[0], nil
}
func (cli *Client) ValidateCustomerUsername(ctx context.Context, username string) (bool, error) {
	method := "validateCustomerUsername"
	params := map[string]string{"username": username}
	resp, err := cli.SendRequest(ctx, method, params)
	if err != nil {
		return false, erro.NewFromError("IsCustomerUsernameAvailable: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return false, erro.NewFromError(fmt.Sprintf(method+": bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status common2.Status
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, erro.NewFromError(method+": unmarshaling response failed", err)
	}
	if respData.Status.ErrorCode != 0 {
		return false, erro.NewFromError(fmt.Sprintf(method+": bad response error code: %s", respData.Status.ErrorCode), nil)
	}
	return true, nil
}


func (cli *Client) AddCustomerRewardPoints(ctx context.Context, filters map[string]string) (AddCustomerRewardPointsResult, error) {
	resp, err := cli.SendRequest(ctx, "addCustomerRewardPoints", filters)
	if err != nil {
		return AddCustomerRewardPointsResult{}, err
	}
	var res AddCustomerRewardPointsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return AddCustomerRewardPointsResult{}, erro.NewFromError("failed to unmarshal AddCustomerRewardPointsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return AddCustomerRewardPointsResult{}, erro.NewFromResponseStatus(&res.Status)
	}
	if len(res.AddCustomerRewardPointsResults) > 0 {
		return res.AddCustomerRewardPointsResults[0], nil
	}

	return AddCustomerRewardPointsResult{}, nil
}

func (cli *Client) AddCustomerRewardPointsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddCustomerRewardPointsResponseBulk, error) {
	var respBulk AddCustomerRewardPointsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "addCustomerRewardPoints",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return respBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBulk, err
	}

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal AddCustomerRewardPointsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&respBulk.Status) {
		return respBulk, erro.NewErplyError(respBulk.Status.ErrorCode.String(), respBulk.Status.Request+": "+respBulk.Status.ResponseStatus)
	}

	for _, bulkItem := range respBulk.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return respBulk, erro.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus)
		}
	}

	return respBulk, nil
}
