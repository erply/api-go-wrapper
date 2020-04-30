package pos

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

// GetPointsOfSale will list points of sale according to specified filters.
func (cli *Client) GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error) {
	resp, err := cli.SendRequest(ctx, "getPointsOfSale", filters)
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
