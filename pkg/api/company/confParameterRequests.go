package company

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

func (cli *Client) GetConfParameters(ctx context.Context) (*ConfParameter, error) {

	resp, err := cli.SendRequest(ctx, "getConfParameters", map[string]string{})
	if err != nil {
		return nil, sharedCommon.NewFromError("GetConfParameters request failed", err, 0)
	}
	res := &GetConfParametersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetConfParametersResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.ConfParameters) == 0 {
		return nil, sharedCommon.NewFromError(fmt.Sprint("Conf Parameters were not found", nil), err, res.Status.ErrorCode)
	}

	return &res.ConfParameters[0], nil
}
