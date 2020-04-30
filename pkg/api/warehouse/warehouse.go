package warehouse

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

//GetWarehouses ...
func (cli *Client) GetWarehouses(ctx context.Context) (Warehouses, error) {

	resp, err := cli.SendRequest(ctx, "getWarehouses", map[string]string{"warehouseID": "0"})
	if err != nil {
		return nil, err
	}

	res := &GetWarehousesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetWarehousesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Warehouses) == 0 {
		return nil, nil
	}

	return res.Warehouses, nil
}
