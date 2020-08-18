package company

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
)

//GetCompanyInfo ...
func (cli *Client) GetCompanyInfo(ctx context.Context) (*Info, error) {
	resp, err := cli.SendRequest(ctx, "getCompanyInfo", map[string]string{})
	if err != nil {
		return nil, erro.NewFromError("GetCompanyInfo request failed", err)
	}
	res := &GetCompanyInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetCompanyInfoResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.CompanyInfos) == 0 {
		return nil, nil
	}

	return &res.CompanyInfos[0], nil
}
