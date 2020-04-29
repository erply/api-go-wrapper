package api

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type PointOfSale struct {
	PointOfSaleID uint   `json:"pointOfSaleID"`
	Name          string `json:"name"`
	WarehouseID   int    `json:"warehouseID"`
	WarehouseName string `json:"warehouseName"`
	Added         uint64 `json:"added"`
	LastModified  uint64 `json:"lastModified"`
}

type GetPointsOfSaleResponse struct {
	Status       Status        `json:"status"`
	PointsOfSale []PointOfSale `json:"records"`
}

// GetPointsOfSale will list points of sale according to specified filters.
func (cli *erplyClient) GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error) {
	resp, err := cli.sendRequest(ctx, GetPointsOfSaleMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetPointsOfSaleResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetPointsOfSaleResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.PointsOfSale, nil
}

//GetPointsOfSale ...
func (cli *erplyClient) GetPointsOfSaleByID(posID string) (*PointOfSale, error) {
	method := GetPointsOfSaleMethod
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}

	params := getMandatoryParameters(cli, method)
	params.Add("pointOfSaleID", posID)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}
	res := &GetPointsOfSaleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.PointsOfSale) == 0 {
		return nil, nil
	}

	return &res.PointsOfSale[0], nil
}
