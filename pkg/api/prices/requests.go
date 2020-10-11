package prices

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

func (cli *Client) GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error) {
	resp, err := cli.SendRequest(ctx, "getSupplierPriceLists", filters)
	if err != nil {
		return nil, err
	}
	var res GetPriceListsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []PriceList{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetPriceListsResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	return res.PriceLists, nil
}

func (cli *Client) AddProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error) {
	return cli.persistProductToSupplierPriceList(ctx, "addProductToSupplierPriceList", filters)
}

func (cli *Client) EditProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error) {
	return cli.persistProductToSupplierPriceList(ctx, "editProductInSupplierPriceList", filters)
}

func (cli *Client) persistProductToSupplierPriceList(ctx context.Context, method string, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error) {
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, erro.NewFromError(method+" request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &ChangeProductToSupplierPriceListResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal ChangeProductToSupplierPriceListResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.ChangeProductToSupplierPriceListResult) == 0 {
		return nil, nil
	}

	return &res.ChangeProductToSupplierPriceListResult[0], nil
}

//ChangeProductToSupplierPriceListBulk wraps both additions and edits as addProductToSupplierPriceList or editProductInSupplierPriceList
func (cli *Client) ChangeProductToSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToSupplierPriceListResponseBulk, error) {
	var bulkResp ChangeProductToSupplierPriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot add more than %d products to price list in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, prodPrice := range bulkRequest {
		_, isEditMode := prodPrice["supplierPriceListProductID"]
		methodName := "addProductToSupplierPriceList"
		if isEditMode {
			methodName = "editProductInSupplierPriceList"
		}
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: methodName,
			Filters:    prodPrice,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal ChangeProductToSupplierPriceListResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error) {
	var bulkResp GetPriceListsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getSupplierPriceLists",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetPriceListsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) GetProductsInSupplierPriceList(ctx context.Context, filters map[string]string) ([]ProductsInSupplierPriceList, error) {
	resp, err := cli.SendRequest(ctx, "getProductsInSupplierPriceList", filters)
	if err != nil {
		return nil, err
	}
	var res ProductsInSupplierPriceListResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []ProductsInSupplierPriceList{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal ProductsInSupplierPriceListResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	return res.ProductsInSupplierPriceList, nil
}

func (cli *Client) GetProductsInSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (ProductsInSupplierPriceListResponseBulk, error) {
	var bulkResp ProductsInSupplierPriceListResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductsInSupplierPriceList",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal ProductsInSupplierPriceListResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) GetProductsInPriceList(ctx context.Context, filters map[string]string) ([]ProductsInPriceList, error) {
	resp, err := cli.SendRequest(ctx, "getProductsInPriceList", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductsInPriceListResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []ProductsInPriceList{}, err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal GetProductsInPriceListResponse from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	return res.PriceLists, nil
}

func (cli *Client) GetProductsInPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsInPriceListResponseBulk, error) {
	var bulkResp GetProductsInPriceListResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductsInPriceList",
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal GetProductsInPriceListResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return bulkResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
		}
	}

	return bulkResp, nil
}

func (cli *Client) DeleteProductsFromSupplierPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromSupplierPriceListResult, error) {
	resp, err := cli.SendRequest(ctx, "deleteProductsFromSupplierPriceList", filters)
	if err != nil {
		return nil, erro.NewFromError("deleteProductsFromSupplierPriceList request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &DeleteProductsFromSupplierPriceListResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductsFromSupplierPriceListResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.DeleteProductsFromSupplierPriceListResult) == 0 {
		return nil, nil
	}

	return &res.DeleteProductsFromSupplierPriceListResult[0], nil
}

func (cli *Client) DeleteProductsFromSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (DeleteProductsFromSupplierPriceListResponseBulk, error) {
	var bulkResp DeleteProductsFromSupplierPriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot delete more than %d products from price list in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteProductsFromSupplierPriceList",
			Filters:    bulkInput,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductsFromSupplierPriceListResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) SaveSupplierPriceList(ctx context.Context, filters map[string]string) (*SaveSupplierPriceListResult, error) {
	resp, err := cli.SendRequest(ctx, "saveSupplierPriceList", filters)
	if err != nil {
		return nil, erro.NewFromError("saveSupplierPriceList request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SaveSupplierPriceListResultResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal SaveSupplierPriceListResultResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.SaveSupplierPriceListResult) == 0 {
		return nil, nil
	}

	return &res.SaveSupplierPriceListResult[0], nil
}

func (cli *Client) SaveSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveSupplierPriceListResponseBulk, error) {
	var bulkResp SaveSupplierPriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d price lists in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveSupplierPriceList",
			Filters:    bulkInput,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveSupplierPriceListResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) SavePriceList(ctx context.Context, filters map[string]string) (*SavePriceListResult, error) {
	resp, err := cli.SendRequest(ctx, "savePriceList", filters)
	if err != nil {
		return nil, erro.NewFromError("savePriceList request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &SavePriceListResultResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal SavePriceListResultResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.SavePriceListResults) == 0 {
		return nil, nil
	}

	return &res.SavePriceListResults[0], nil
}

func (cli *Client) SavePriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SavePriceListResponseBulk, error) {
	var bulkResp SavePriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot save more than %d price lists in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "savePriceList",
			Filters:    bulkInput,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal SavePriceListResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) AddProductToPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToPriceListResult, error) {
	return cli.persistProductToPriceList(ctx, "addProductToPriceList", filters)
}

func (cli *Client) EditProductToPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToPriceListResult, error) {
	return cli.persistProductToPriceList(ctx, "editProductInPriceList", filters)
}

func (cli *Client) persistProductToPriceList(ctx context.Context, method string, filters map[string]string) (*ChangeProductToPriceListResult, error) {
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, erro.NewFromError(method+" request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &ChangeProductToPriceListResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal ChangeProductToPriceListResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.ChangeProductToPriceListResults) == 0 {
		return nil, nil
	}

	return &res.ChangeProductToPriceListResults[0], nil
}

func (cli *Client) ChangeProductToPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToPriceListResponseBulk, error) {
	var bulkResp ChangeProductToPriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot add more than %d products to price list in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, prodPrice := range bulkRequest {
		_, isEditMode := prodPrice["priceListProductID"]
		methodName := "addProductToPriceList"
		if isEditMode {
			methodName = "editProductInPriceList"
		}
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: methodName,
			Filters:    prodPrice,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal ChangeProductToPriceListBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}

func (cli *Client) DeleteProductsFromPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromPriceListResult, error) {
	resp, err := cli.SendRequest(ctx, "deleteProductInPriceList", filters)
	if err != nil {
		return nil, erro.NewFromError("deleteProductInPriceList request failed", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &DeleteProductsFromPriceListResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductsFromPriceListResponse from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}

	if len(res.DeleteProductsFromPriceListResults) == 0 {
		return nil, nil
	}

	return &res.DeleteProductsFromPriceListResults[0], nil
}


func (cli *Client) DeleteProductsFromPriceListBulk(
	ctx context.Context,
	bulkRequest []map[string]interface{}, baseFilters map[string]string,
) (DeleteProductsFromPriceListResponseBulk, error) {
	var bulkResp DeleteProductsFromPriceListResponseBulk

	if len(bulkRequest) > common2.MaxBulkRequestsCount {
		return bulkResp, fmt.Errorf("cannot delete more than %d products from price list in one bulk request", common2.MaxBulkRequestsCount)
	}

	bulkInputs := make([]common.BulkInput, 0, len(bulkRequest))
	for _, bulkInput := range bulkRequest {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteProductInPriceList",
			Filters:    bulkInput,
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
		return bulkResp, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductsFromPriceListResponseBulk from '%s': %v", string(body), err)
	}

	if !common.IsJSONResponseOK(&bulkResp.Status) {
		return bulkResp, erro.NewErplyError(bulkResp.Status.ErrorCode.String(), bulkResp.Status.Request+": "+bulkResp.Status.ResponseStatus)
	}

	for _, bulkRespItem := range bulkResp.BulkItems {
		if !common.IsJSONResponseOK(&bulkRespItem.Status.Status) {
			return bulkResp, erro.NewErplyError(
				bulkRespItem.Status.ErrorCode.String(),
				fmt.Sprintf("%+v", bulkRespItem.Status),
			)
		}
	}

	return bulkResp, nil
}
