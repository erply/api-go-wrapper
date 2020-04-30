package pos

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

/*
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
		Status       common.Status    `json:"status"`
		PointsOfSale []PointOfSale `json:"records"`
	}

	PointOfSaleManager interface {
		GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error)
	}
)

// GetPointsOfSale will list points of sale according to specified filters.
func (cli *Client) GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error) {
	resp, err := cli.SendRequest(ctx, api.GetPointsOfSaleMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetPointsOfSaleResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetPointsOfSaleResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.PointsOfSale, nil
}
*/
