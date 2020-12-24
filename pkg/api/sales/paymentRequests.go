package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
	"net/http"
)

func (cli *Client) SavePayment(ctx context.Context, filters map[string]string) (int64, error) {
	resp, err := cli.SendRequest(ctx, "savePayment", filters)
	if err != nil {
		return 0, sharedCommon.NewFromError("SavePayment: error sending POST request", err, 0)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, sharedCommon.NewFromError(fmt.Sprintf("SavePayment: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []struct {
			PaymentID int64 `json:"paymentID"`
		} `json:"records"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, sharedCommon.NewFromError("SavePayment: error decoding JSON response body", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, sharedCommon.NewFromError(fmt.Sprintf("SavePayment: API error %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return 0, sharedCommon.NewFromError("SavePayment: no records in response", nil, respData.Status.ErrorCode)
	}

	return respData.Records[0].PaymentID, nil
}

func (cli *Client) SavePaymentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SavePaymentsResponseBulk, error) {
	var bulkResp SavePaymentsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "savePayment",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SavePaymentsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, bulkResp.Status.ErrorCode)
		}
	}

	return bulkResp, nil
}

func (cli *Client) GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error) {
	resp, err := cli.SendRequest(ctx, "getPayments", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetPayments: error sending request", err, 0)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("GetPayments: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []PaymentInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetPayments: error decoding JSON response body", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("GetPayments: API error %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
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
		return bulkResp, sharedCommon.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus, bulkResp.Status.ErrorCode)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus, bulkResp.Status.ErrorCode)
		}
	}

	return bulkResp, nil
}
