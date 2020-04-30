package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+res.Status.ErrorField+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.SalesDocuments) == 0 {
		//intentionally, otherwise when the documents are cached the error will be triggered.
		return nil, nil
	}

	return res.SalesDocuments, nil
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
		return erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return nil
}
