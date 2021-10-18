package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

func (cli *Client) GetPurchaseDocuments(ctx context.Context, filters map[string]string) ([]PurchaseDocument, error) {
	resp, err := cli.SendRequest(ctx, "getPurchaseDocuments", filters)
	if err != nil {
		return nil, err
	}
	var res GetPurchaseDocumentsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []PurchaseDocument{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetPurchaseDocumentsResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res.PurchaseDocuments, nil
}

func (cli *Client) GetPurchaseDocumentsWithStatus(ctx context.Context, filters map[string]string) (GetPurchaseDocumentsResponse, error) {
	var res GetPurchaseDocumentsResponse

	resp, err := cli.SendRequest(ctx, "getPurchaseDocuments", filters)
	if err != nil {
		return res, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return res, fmt.Errorf("ERPLY API: failed to unmarshal GetPurchaseDocumentsResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return res, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res, nil
}

func (cli *Client) GetPurchaseDocumentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPurchaseDocumentResponseBulk, error) {
	var bulkResp GetPurchaseDocumentResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getPurchaseDocuments",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetPurchaseDocumentResponseBulk from '%s': %v", string(body), err)
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
