package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
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
		Status  common2.Status
		Records []struct {
			PaymentID int64 `json:"paymentID"`
		} `json:"records"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, erro.NewFromError("SavePayment: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, erro.NewFromError(fmt.Sprintf("SavePayment: API error %s", respData.Status.ErrorCode), nil)
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
		Status  common2.Status
		Records []PaymentInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erro.NewFromError("GetPayments: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, erro.NewFromError(fmt.Sprintf("GetPayments: API error %s", respData.Status.ErrorCode), nil)
	}

	return respData.Records, nil
}

func (cli *Client) GetPaymentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPaymentsResponseBulk, error) {
	var bulkResp GetPaymentsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getPayments",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return bulkResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bulkResp, err
	}

	if err := json.Unmarshal(body, &bulkResp); err != nil {
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetPaymentsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}
