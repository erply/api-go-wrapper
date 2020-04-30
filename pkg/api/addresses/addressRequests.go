package addresses

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

func (cli *Client) GetAddresses(ctx context.Context, filters map[string]string) ([]common.Address, error) {
	resp, err := cli.SendRequest(ctx, "getAddresses", filters)
	if err != nil {
		return nil, erro.NewFromError("GetAddresses request failed", err)
	}

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetAddressesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Addresses, nil
}
func (cli *Client) SaveAddress(ctx context.Context, filters map[string]string) ([]common.Address, error) {
	method := "saveAddress"
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, erro.NewFromError(method+": request failed", err)
	}
	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError(method+": JSON unmarshal failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Addresses) == 0 {
		return nil, erro.NewFromError(method+": no records in response", nil)
	}

	return res.Addresses, nil
}
