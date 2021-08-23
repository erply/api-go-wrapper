package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

func (cli *Client) SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error) {
	resp, err := cli.SendRequest(ctx, "saveSalesDocument", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("PostSalesDocument request failed", err, 0)
	}
	res := &PostSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling PostSalesDocumentResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

func (cli *Client) SaveSalesDocumentBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SaveSalesDocumentResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveSalesDocument",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return respBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBulk, err
	}

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveSalesDocumentResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&respBulk.Status) {
		return respBulk, sharedCommon.NewErplyError(respBulk.Status.ErrorCode.String(), respBulk.Status.Request+": "+respBulk.Status.ResponseStatus, respBulk.Status.ErrorCode)
	}

	for _, bulkRespItem := range respBulk.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return respBulk, sharedCommon.NewErplyError(bulkRespItem.Status.ErrorCode.String(), bulkRespItem.Status.Request+": "+bulkRespItem.Status.ResponseStatus, respBulk.Status.ErrorCode)
		}
	}

	return respBulk, nil
}

func (cli *Client) GetPurchaseDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error) {
	resp, err := cli.SendRequest(ctx, "getSalesDocuments", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetPurchaseDocuments request failed", err, 0)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetPurchaseDocumentsResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res.SalesDocuments, nil
}

func (cli *Client) SavePurchaseDocument(ctx context.Context, filters map[string]string) (resp PurchaseDocImportReports, err error) {
	res := &SavePurchaseDocumentResponse{}
	err = cli.Scan(ctx, "savePurchaseDocument", filters, res)
	if err != nil {
		return
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

func (cli *Client) SavePurchaseDocumentBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SavePurchaseDocumentResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "savePurchaseDocument",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return respBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBulk, err
	}

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SavePurchaseDocumentResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&respBulk.Status) {
		return respBulk, sharedCommon.NewErplyError(
			respBulk.Status.ErrorCode.String(),
			respBulk.Status.Request+": "+respBulk.Status.ResponseStatus,
			respBulk.Status.ErrorCode,
		)
	}

	for _, bulkRespItem := range respBulk.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return respBulk, sharedCommon.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				bulkRespItem.Status.Request+": "+bulkRespItem.Status.ResponseStatus,
				respBulk.Status.ErrorCode,
			)
		}
	}

	return respBulk, nil
}

func (cli *Client) GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error) {
	resp, err := cli.SendRequest(ctx, "getSalesDocuments", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetSalesDocument request failed", err, 0)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetSalesDocumentResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res.SalesDocuments, nil
}

func (cli *Client) GetSalesDocumentsWithStatus(ctx context.Context, filters map[string]string) (*GetSalesDocumentResponse, error) {
	resp, err := cli.SendRequest(ctx, "getSalesDocuments", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetSalesDocument request failed", err, 0)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetSalesDocumentResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res, nil
}

func (cli *Client) GetSalesDocumentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSaleDocumentResponseBulk, error) {
	var bulkResp GetSaleDocumentResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSalesDocuments",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetSaleDocumentResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, sharedCommon.NewErplyError(
			bulkResp.Status.ErrorCode.String(),
			bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus,
			bulkResp.Status.ErrorCode,
		)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, sharedCommon.NewErplyError(
				bulkItem.Status.ErrorCode.String(),
				bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus,
				bulkItem.Status.ErrorCode,
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) DeleteDocument(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteSalesDocument", filters)
	if err != nil {
		return sharedCommon.NewFromError("DeleteDocumentsByIds request failed", err, 0)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError("unmarshaling DeleteDocumentsByIds failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return nil
}

func (cli *Client) DeleteDocumentsBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk DeleteResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteSalesDocument",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return respBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBulk, err
	}

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal DeleteDocumentsBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&respBulk.Status) {
		return respBulk, sharedCommon.NewErplyError(respBulk.Status.ErrorCode.String(), respBulk.Status.Request+": "+respBulk.Status.ResponseStatus, respBulk.Status.ErrorCode)
	}

	return respBulk, nil
}
