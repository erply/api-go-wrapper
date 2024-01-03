package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"net/http"
)

func (cli *Client) CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error) {

	resp, err := cli.SendRequest(ctx, "calculateShoppingCart", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, sharedCommon.NewErplyError(respData.Status.ErrorCode.String(), respData.Status.Request+": "+respData.Status.ResponseStatus, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Records[0], nil
}

func (cli *Client) CalculateShoppingCartWithFullRowsResponse(ctx context.Context, filters map[string]string) (*ShoppingCartTotalsWithFullRows, error) {
	resp, err := cli.SendRequest(ctx, "calculateShoppingCart", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []*ShoppingCartTotalsWithFullRows
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, sharedCommon.NewErplyError(respData.Status.ErrorCode.String(), respData.Status.Request+": "+respData.Status.ResponseStatus, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CalculateShoppingCart: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Records[0], nil
}
