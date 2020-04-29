package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	PointOfSale struct {
		PointOfSaleID uint   `json:"pointOfSaleID"`
		Name          string `json:"name"`
		WarehouseID   int    `json:"warehouseID"`
		WarehouseName string `json:"warehouseName"`
		Added         uint64 `json:"added"`
		LastModified  uint64 `json:"lastModified"`
	}

	GetPointsOfSaleResponse struct {
		Status       Status        `json:"status"`
		PointsOfSale []PointOfSale `json:"records"`
	}

	PointOfSaleManager interface {
		GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error)
	}
)

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
