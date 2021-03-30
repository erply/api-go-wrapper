package company

import (
	"context"
	"encoding/json"
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
		return nil, sharedCommon.NewFromError("Conf Parameters were not found", err, res.Status.ErrorCode)
	}

	return &res.ConfParameters[0], nil
}

func (cli *Client) GetDefaultLanguage(ctx context.Context) (*Language, error) {

	const requestName = "getDefaultLanguage"
	resp, err := cli.SendRequest(ctx, requestName, map[string]string{})
	if err != nil {
		return nil, sharedCommon.NewFromError(requestName+" request failed", err, 0)
	}
	res := &GetDefaultLanguageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshalling "+requestName+" response failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.Languages) == 0 {
		return nil, sharedCommon.NewFromError("languages were not found", err, res.Status.ErrorCode)
	}

	return &res.Languages[0], nil
}
