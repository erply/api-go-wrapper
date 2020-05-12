package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
	erro "github.com/tarmo-randma/api-go-wrapper/internal/errors"
)

func (cli *Client) CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error) {

	resp, err := cli.SendRequest(ctx, "calculateShoppingCart", filters)
	if err != nil {
		return nil, erro.NewFromError("CalculateShoppingCart: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, erro.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  common.Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, erro.NewFromError("CalculateShoppingCart: unmarshaling response failed", err)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(respData.Status.ErrorCode), respData.Status.Request+": "+respData.Status.ResponseStatus)
	}
	if len(respData.Records) < 1 {
		return nil, erro.NewFromError("CalculateShoppingCart: no records in response", nil)
	}

	return respData.Records[0], nil
}
