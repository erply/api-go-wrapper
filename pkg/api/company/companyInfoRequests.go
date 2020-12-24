package company

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

//GetCompanyInfo ...
func (cli *Client) GetCompanyInfo(ctx context.Context) (*Info, error) {
	resp, err := cli.SendRequest(ctx, "getCompanyInfo", map[string]string{})
	if err != nil {
		return nil, sharedCommon.NewFromError("GetCompanyInfo request failed", err, 0)
	}
	res := &GetCompanyInfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetCompanyInfoResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.CompanyInfos) == 0 {
		return nil, nil
	}

	return &res.CompanyInfos[0], nil
}
