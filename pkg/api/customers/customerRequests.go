package customers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Customers, nil
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
		Status  common.Status
		Records []WebshopClient
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("VerifyCustomerUser: unmarhsalling response failed", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		Status common.Status
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, erro.NewFromError(method+": unmarshaling response failed", err)
	}
	if respData.Status.ErrorCode != 0 {
		return false, erro.NewFromError(fmt.Sprintf(method+": bad response error code: %d", respData.Status.ErrorCode), nil)
	}
	return true, nil
}
