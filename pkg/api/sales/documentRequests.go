package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"io/ioutil"
)

func (cli *Client) SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error) {
	resp, err := cli.SendRequest(ctx, "saveSalesDocument", filters)
	if err != nil {
		return nil, erro.NewFromError("PostSalesDocument request failed", err)
	}
	res := &PostSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling PostSalesDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}

func (cli *Client) SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error) {
	resp, err := cli.SendRequest(ctx, "savePurchaseDocument", filters)
	if err != nil {
		return nil, erro.NewFromError("savePurchaseDocument"+" request failed", err)
	}
	res := &SavePurchaseDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling savePurchaseDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+res.Status.ErrorField+": "+res.Status.ResponseStatus)
	}

	if len(res.ImportReports) == 0 {
		return nil, nil
	}

	return res.ImportReports, nil
}
func (cli *Client) GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error) {
	resp, err := cli.SendRequest(ctx, "getSalesDocuments", filters)
	if err != nil {
		return nil, erro.NewFromError("GetSalesDocument request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetSalesDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("GetSalesDocument request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetSalesDocumentResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(bulkItem.Status.ErrorCode.String(), bulkItem.Status.Request+": "+bulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) DeleteDocument(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteSalesDocument", filters)
	if err != nil {
		return erro.NewFromError("DeleteDocumentsByIds request failed", err)
	}
	res := &GetSalesDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return erro.NewFromError("unmarshaling DeleteDocumentsByIds failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return erro.NewFromResponseStatus(&res.Status)
	}

	return nil
}
