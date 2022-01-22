package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

func (cli *Client) GetCoupons(ctx context.Context, filters map[string]string) (*GetCouponsResponse, error) {
	resp, err := cli.SendRequest(ctx, "getCoupons", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("getCoupons request failed", err, 0)
	}
	res := &GetCouponsResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling getCoupons failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res, nil
}
