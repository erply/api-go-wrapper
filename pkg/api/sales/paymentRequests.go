package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
)

func (cli *Client) SavePayment(ctx context.Context, filters map[string]string) (int64, error) {
	resp, err := cli.SendRequest(ctx, "savePayment", filters)
	if err != nil {
		return 0, erro.NewFromError("SavePayment: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, erro.NewFromError(fmt.Sprintf("SavePayment: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  common.Status
		Records []struct {
			PaymentID int64 `json:"paymentID"`
		} `json:"records"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, erro.NewFromError("SavePayment: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, erro.NewFromError(fmt.Sprintf("SavePayment: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return 0, erro.NewFromError("SavePayment: no records in response", nil)
	}

	return respData.Records[0].PaymentID, nil
}

func (cli *Client) GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error) {
	resp, err := cli.SendRequest(ctx, "getPayments", filters)
	if err != nil {
		return nil, erro.NewFromError("GetPayments: error sending request", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, erro.NewFromError(fmt.Sprintf("GetPayments: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  common.Status
		Records []PaymentInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erro.NewFromError("GetPayments: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, erro.NewFromError(fmt.Sprintf("GetPayments: API error %d", respData.Status.ErrorCode), nil)
	}

	return respData.Records, nil
}
