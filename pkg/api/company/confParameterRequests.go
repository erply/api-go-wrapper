package company

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

func (cli *Client) GetConfParameters(ctx context.Context) (*ConfParameter, error) {

	resp, err := cli.SendRequest(ctx, "getConfParameters", map[string]string{})
	if err != nil {
		return nil, erro.NewFromError("GetConfParameters request failed", err)
	}
	res := &GetConfParametersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetConfParametersResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ConfParameters) == 0 {
		return nil, erro.NewFromError(fmt.Sprint("Conf Parameters were not found", nil), err)
	}

	return &res.ConfParameters[0], nil
}
