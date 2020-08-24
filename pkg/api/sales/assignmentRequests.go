package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
)

//GetVatRatesByVatRateID ...
func (cli *Client) SaveAssignment(ctx context.Context, filters map[string]string) (int64, error) {
	resp, err := cli.SendRequest(ctx, "saveAssignment", filters)
	if err != nil {
		return 0, err
	}
	res := &saveAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, erro.NewFromError("unmarshaling saveAssignment failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return 0, erro.NewFromResponseStatus(&res.Status)
	}

	return res.Records[0].AssignmentID, nil
}
