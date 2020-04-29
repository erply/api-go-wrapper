package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
)

// GetCountries will list countries according to specified filters.
func (cli *erplyClient) GetCountries(ctx context.Context, filters map[string]string) ([]Country, error) {
	resp, err := cli.sendRequest(ctx, GetCountriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCountriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetCountriesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Countries, nil
}

//GetUserName from GetUserRights erply API request
func (cli *erplyClient) GetUserRights(ctx context.Context, filters map[string]string) ([]UserRights, error) {

	resp, err := cli.sendRequest(ctx, GetUserRightsMethod, filters)
	if err != nil {
		return nil, erplyerr(GetUserRightsMethod+" request failed", err)
	}
	res := &GetUserRightsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetUserRightsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Records) == 0 {
		return nil, errors.New("no records found")
	}

	return res.Records, nil
}

// GetEmployees will list employees according to specified filters.
func (cli *erplyClient) GetEmployees(ctx context.Context, filters map[string]string) ([]Employee, error) {
	resp, err := cli.sendRequest(ctx, GetEmployeesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetEmployeesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetEmployeesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Employees, nil
}

// GetBusinessAreas will list business areas according to specified filters.
func (cli *erplyClient) GetBusinessAreas(ctx context.Context, filters map[string]string) ([]BusinessArea, error) {
	resp, err := cli.sendRequest(ctx, GetBusinessAreasMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetBusinessAreasResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetBusinessAreasResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.BusinessAreas, nil
}

// GetCurrencies will list currencies according to specified filters.
func (cli *erplyClient) GetCurrencies(ctx context.Context, filters map[string]string) ([]Currency, error) {
	resp, err := cli.sendRequest(ctx, GetCurrenciesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetCurrenciesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetCurrenciesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Currencies, nil
}

func (cli *erplyClient) SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error) {
	resp, err := cli.sendRequest(ctx, savePurchaseDocumentMethod, filters)
	if err != nil {
		return nil, erplyerr(savePurchaseDocumentMethod+" request failed", err)
	}
	res := &PostPurchaseDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling savePurchaseDocumentResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

func (cli *erplyClient) logProcessingOfCustomerData(ctx context.Context, filters map[string]string) error {
	resp, err := cli.sendRequest(ctx, logProcessingOfCustomerDataMethod, filters)
	if err != nil {
		return erplyerr("logProcessingOfCustomerData request failed", err)
	}

	if resp.StatusCode != http.StatusOK {
		return erplyerr(fmt.Sprintf("Logging response HTTP status is %d", resp.StatusCode), nil)
	}

	return nil
}
