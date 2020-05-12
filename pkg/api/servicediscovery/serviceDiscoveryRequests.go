package servicediscovery

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/tarmo-randma/api-go-wrapper/internal/common"
	erro "github.com/tarmo-randma/api-go-wrapper/internal/errors"
)

func (cli *Client) GetServiceEndpoints(ctx context.Context) (*ServiceEndpoints, error) {
	const method = "getServiceEndpoints"
	resp, err := cli.SendRequest(ctx, method, map[string]string{})
	if err != nil {
		return nil, err
	}

	res := &getServiceEndpointsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to decode %s response", method))
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Records) < 1 {
		return nil, errors.New("no records in response")
	}
	return &res.Records[0], nil
}
