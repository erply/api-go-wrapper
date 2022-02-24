package customers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
	"net/http"
)

func (cli *Client) SaveCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error) {
	resp, err := cli.SendRequest(ctx, "saveCustomer", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("PostCustomer request failed", err, 0)
	}
	res := &PostCustomerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling CustomerImportReport failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
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
		return nil, sharedCommon.NewFromError("failed to unmarshal GetCustomersResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Customers, nil
}

// GetCustomersWithStatus will list customers according to specified filters.
func (cli *Client) GetCustomersWithStatus(ctx context.Context, filters map[string]string) (*GetCustomersResponse, error) {
	resp, err := cli.SendRequest(ctx, "getCustomers", filters)
	if err != nil {
		return nil, err
	}
	var res GetCustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetCustomersResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return &res, nil
}

// GetCustomerGroups will list customers groups according to specified filters.
func (cli *Client) GetCustomerGroups(ctx context.Context, filters map[string]string) ([]CustomerGroup, error) {
	resp, err := cli.SendRequest(ctx, "getCustomerGroups", filters)
	if err != nil {
		return nil, err
	}
	var res GetCustomerGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetCustomerGroupsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
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
		return customersResponse, sharedCommon.NewErplyError(customersResponse.Status.ErrorCode.String(), customersResponse.Status.Request+": "+customersResponse.Status.ResponseStatus, customersResponse.Status.ErrorCode)
	}

	for _, supplierBulkItem := range customersResponse.BulkItems {
		if !common.IsJSONResponseOK(&supplierBulkItem.Status.Status) {
			return customersResponse, sharedCommon.NewErplyError(supplierBulkItem.Status.ErrorCode.String(), supplierBulkItem.Status.Request+": "+supplierBulkItem.Status.ResponseStatus, customersResponse.Status.ErrorCode)
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
		return nil, sharedCommon.NewFromError("VerifyCustomerUser: request failed", err, 0)
	}

	var res struct {
		Status  sharedCommon.Status
		Records []WebshopClient
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("VerifyCustomerUser: unmarhsalling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.Records) != 1 {
		return nil, sharedCommon.NewFromError("VerifyCustomerUser: no records in response", nil, res.Status.ErrorCode)
	}

	return &res.Records[0], nil
}
func (cli *Client) ValidateCustomerUsername(ctx context.Context, username string) (bool, error) {
	method := "validateCustomerUsername"
	params := map[string]string{"username": username}
	resp, err := cli.SendRequest(ctx, method, params)
	if err != nil {
		return false, sharedCommon.NewFromError("IsCustomerUsernameAvailable: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return false, sharedCommon.NewFromError(fmt.Sprintf(method+": bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status sharedCommon.Status
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, sharedCommon.NewFromError(method+": unmarshaling response failed", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return false, sharedCommon.NewFromError(fmt.Sprintf(method+": bad response error code: %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
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
		return AddCustomerRewardPointsResult{}, sharedCommon.NewFromError("failed to unmarshal AddCustomerRewardPointsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return AddCustomerRewardPointsResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
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
		return respBulk, sharedCommon.NewErplyError(respBulk.Status.ErrorCode.String(), respBulk.Status.Request+": "+respBulk.Status.ResponseStatus, respBulk.Status.ErrorCode)
	}

	for _, bulkItem := range respBulk.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return respBulk, sharedCommon.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, respBulk.Status.ErrorCode)
		}
	}

	return respBulk, nil
}

func (cli *Client) SaveCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (SaveCustomerResponseBulk, error) {
	var saveCustomerResponseBulk SaveCustomerResponseBulk

	if len(customerMap) > sharedCommon.MaxBulkRequestsCount {
		return saveCustomerResponseBulk, fmt.Errorf("cannot save more than %d customers in one request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(customerMap))
	for _, customer := range customerMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveCustomer",
			Filters:    customer,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return saveCustomerResponseBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return saveCustomerResponseBulk, err
	}

	if err := json.Unmarshal(body, &saveCustomerResponseBulk); err != nil {
		return saveCustomerResponseBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveCustomerResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&saveCustomerResponseBulk.Status) {
		return saveCustomerResponseBulk, sharedCommon.NewErplyError(saveCustomerResponseBulk.Status.ErrorCode.String(), saveCustomerResponseBulk.Status.Request+": "+saveCustomerResponseBulk.Status.ResponseStatus, saveCustomerResponseBulk.Status.ErrorCode)
	}

	for _, bulkItem := range saveCustomerResponseBulk.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return saveCustomerResponseBulk, sharedCommon.NewErplyError(
				bulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkItem.Status),
				saveCustomerResponseBulk.Status.ErrorCode,
			)
		}
	}

	return saveCustomerResponseBulk, nil
}

func (cli *Client) DeleteCustomer(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteCustomer", filters)
	if err != nil {
		return err
	}

	var res DeleteCustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError("failed to unmarshal DeleteCustomerResponse ", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return nil
}

func (cli *Client) DeleteCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (DeleteCustomersResponseBulk, error) {
	var deleteCustomersResponse DeleteCustomersResponseBulk

	if len(customerMap) > sharedCommon.MaxBulkRequestsCount {
		return deleteCustomersResponse, fmt.Errorf("cannot delete more than %d customers in one request", sharedCommon.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(customerMap))
	for _, filter := range customerMap {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteCustomer",
			Filters:    filter,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, attrs)
	if err != nil {
		return deleteCustomersResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deleteCustomersResponse, err
	}

	if err := json.Unmarshal(body, &deleteCustomersResponse); err != nil {
		return deleteCustomersResponse, fmt.Errorf("ERPLY API: failed to unmarshal DeleteCustomersResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&deleteCustomersResponse.Status) {
		return deleteCustomersResponse, sharedCommon.NewErplyError(
			deleteCustomersResponse.Status.ErrorCode.String(),
			deleteCustomersResponse.Status.Request+": "+deleteCustomersResponse.Status.ResponseStatus,
			deleteCustomersResponse.Status.ErrorCode,
		)
	}

	for _, bulkItem := range deleteCustomersResponse.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return deleteCustomersResponse, sharedCommon.NewErplyError(
				bulkItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkItem.Status),
				deleteCustomersResponse.Status.ErrorCode,
			)
		}
	}

	return deleteCustomersResponse, nil
}
