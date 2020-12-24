package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

//GetVatRatesByVatRateID ...
func (cli *Client) SaveAssignment(ctx context.Context, filters map[string]string) (int64, error) {
	resp, err := cli.SendRequest(ctx, "saveAssignment", filters)
	if err != nil {
		return 0, err
	}
	res := &saveAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, sharedCommon.NewFromError("unmarshaling saveAssignment failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return 0, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res.Records[0].AssignmentID, nil
}
