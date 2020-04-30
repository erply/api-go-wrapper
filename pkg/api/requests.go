package api

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

//this interface sums up the general requests here
type Manager interface {
	GetCountries(ctx context.Context, filters map[string]string) ([]Country, error)
	GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error)
	GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error)
	GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error)
	GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error)
	LogProcessingOfCustomerData(ctx context.Context, filters map[string]string) error
}

// GetCountries will list countries according to specified filters.
func (cli *Client) GetCountries(ctx context.Context, filters map[string]string) ([]Country, error) {
	resp, err := cli.commonClient.SendRequest(ctx, GetCountriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetCountriesResponse", err)
	}
	if !common.IsJSONResponseOK((*common.Status)(&res.Status)) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Countries, nil
}

//GetUserName from GetUserRights erply API request
func (cli *Client) GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error) {

	resp, err := cli.commonClient.SendRequest(ctx, GetUserRightsMethod, filters)
	if err != nil {
		return nil, erro.NewFromError(GetUserRightsMethod+" request failed", err)
	}
	res := &GetUserRightsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetUserRightsResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Records) == 0 {
		return nil, errors.New("no records found")
	}

	return res.Records, nil
}

// GetEmployees will list employees according to specified filters.
func (cli *Client) GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error) {
	resp, err := cli.commonClient.SendRequest(ctx, GetEmployeesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetEmployeesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetEmployeesResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Employees, nil
}

// GetBusinessAreas will list business areas according to specified filters.
func (cli *Client) GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error) {
	resp, err := cli.commonClient.SendRequest(ctx, GetBusinessAreasMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetBusinessAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetBusinessAreasResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.BusinessAreas, nil
}

// GetCurrencies will list currencies according to specified filters.
func (cli *Client) GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error) {
	resp, err := cli.commonClient.SendRequest(ctx, GetCurrenciesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCurrenciesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetCurrenciesResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Currencies, nil
}

func (cli *Client) LogProcessingOfCustomerData(ctx context.Context, filters map[string]string) error {
	resp, err := cli.commonClient.SendRequest(ctx, logProcessingOfCustomerDataMethod, filters)
	if err != nil {
		return erro.NewFromError("logProcessingOfCustomerData request failed", err)
	}

	if resp.StatusCode != http.StatusOK {
		return erro.NewFromError(fmt.Sprintf("Logging response HTTP status is %d", resp.StatusCode), nil)
	}

	return nil
}
