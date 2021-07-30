package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"io/ioutil"
)

func (cli *Client) GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error) {

	resp, err := cli.SendRequest(ctx, "getProductUnits", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("GetProductUnits request failed", err, 0)
	}

	res := &GetProductUnitsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("unmarshaling GetProductUnitsResponse failed", err, 0)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}

	return res.ProductUnits, nil
}

func (cli *Client) GetProducts(ctx context.Context, filters map[string]string) ([]Product, error) {
	resp, err := cli.SendRequest(ctx, "getProducts", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProductsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Products, nil
}

func (cli *Client) GetProductsCount(ctx context.Context, filters map[string]string) (int, error) {
	resp, err := cli.SendRequest(ctx, "getProducts", filters)
	if err != nil {
		return 0, err
	}
	var res GetProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, sharedCommon.NewFromError("failed to unmarshal GetProductsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return 0, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Status.RecordsTotal, nil
}

func (cli *Client) GetProductPriorityGroups(ctx context.Context, filters map[string]string) (GetProductPriorityGroups, error) {
	var res GetProductPriorityGroups
	resp, err := cli.SendRequest(ctx, "getProductPriorityGroups", filters)
	if err != nil {
		return res, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, sharedCommon.NewFromError("failed to unmarshal GetProductPriorityGroups", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return res, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res, nil
}

// GetProductsBulk will list products according to specified filters sending a bulk request to fetch more products than the default limit
func (cli *Client) GetProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsResponseBulk, error) {
	var productsResp GetProductsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProducts",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return productsResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productsResp, err
	}

	if err := json.Unmarshal(body, &productsResp); err != nil {
		return productsResp, fmt.Errorf("ERPLY API: failed to unmarshal GetProductsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&productsResp.Status) {
		return productsResp, sharedCommon.NewErplyError(productsResp.Status.ErrorCode.String(), productsResp.Status.Request+": "+productsResp.Status.ResponseStatus, productsResp.Status.ErrorCode)
	}

	for _, prodBulkItem := range productsResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsResp, sharedCommon.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus, prodBulkItem.Status.ErrorCode)
		}
	}

	return productsResp, nil
}

func (cli *Client) SaveProduct(ctx context.Context, filters map[string]string) (SaveProductResult, error) {
	resp, err := cli.SendRequest(ctx, "saveProduct", filters)
	if err != nil {
		return SaveProductResult{}, err
	}
	var res SaveProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return SaveProductResult{}, sharedCommon.NewFromError("failed to unmarshal SaveProductResult", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return SaveProductResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveProductResults) > 0 {
		return res.SaveProductResults[0], nil
	}

	return SaveProductResult{}, nil
}

func (cli *Client) SaveProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveProductResponseBulk, error) {
	var productsResp SaveProductResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveProduct",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return productsResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productsResp, err
	}

	if err := json.Unmarshal(body, &productsResp); err != nil {
		return productsResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveProductResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&productsResp.Status) {
		return productsResp, sharedCommon.NewErplyError(productsResp.Status.ErrorCode.String(), productsResp.Status.Request+": "+productsResp.Status.ResponseStatus, productsResp.Status.ErrorCode)
	}

	for _, prodBulkItem := range productsResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsResp, sharedCommon.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus, productsResp.Status.ErrorCode)
		}
	}

	return productsResp, nil
}

func (cli *Client) DeleteProduct(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteProduct", filters)
	if err != nil {
		return err
	}
	var res DeleteProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError("failed to unmarshal DeleteProductResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return nil
}

func (cli *Client) DeleteProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductResponseBulk, error) {
	var deleteRespBulk DeleteProductResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteProduct",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return deleteRespBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deleteRespBulk, err
	}

	if err := json.Unmarshal(body, &deleteRespBulk); err != nil {
		return deleteRespBulk, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&deleteRespBulk.Status) {
		return deleteRespBulk, sharedCommon.NewErplyError(deleteRespBulk.Status.ErrorCode.String(), deleteRespBulk.Status.Request+": "+deleteRespBulk.Status.ResponseStatus, deleteRespBulk.Status.ErrorCode)
	}

	for _, prodBulkItem := range deleteRespBulk.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return deleteRespBulk, sharedCommon.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus, deleteRespBulk.Status.ErrorCode)
		}
	}

	return deleteRespBulk, nil
}

func (cli *Client) GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error) {
	resp, err := cli.SendRequest(ctx, "getProductCategories", filters)
	if err != nil {
		return nil, err
	}
	var res getProductCategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal getProductCategoriesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.ProductCategories, nil
}

func (cli *Client) GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error) {
	resp, err := cli.SendRequest(ctx, "getProductBrands", filters)
	if err != nil {
		return nil, err
	}
	var res getProductBrandsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal getProductBrandsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.ProductBrands, nil
}

func (cli *Client) GetBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error) {
	resp, err := cli.SendRequest(ctx, "getBrands", filters)
	if err != nil {
		return nil, err
	}
	var res getProductBrandsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal getBrandsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.ProductBrands, nil
}

func (cli *Client) GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error) {
	resp, err := cli.SendRequest(ctx, "getProductGroups", filters)
	if err != nil {
		return nil, err
	}
	var res getProductGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal getProductGroupsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.ProductGroups, nil
}

func (cli *Client) GetProductStock(ctx context.Context, filters map[string]string) ([]GetProductStock, error) {
	resp, err := cli.SendRequest(ctx, "getProductStock", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductStockResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProductStockResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.GetProductStock, nil
}

func (cli *Client) GetProductStockFile(ctx context.Context, filters map[string]string) ([]GetProductStockFile, error) {
	filters["responseType"] = ResponseTypeCSV
	resp, err := cli.SendRequest(ctx, "getProductStock", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductStockFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProductStockFileResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.GetProductStockFile, nil
}

func (cli *Client) GetProductStockBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductStockResponseBulk, error) {
	var productsStockResp GetProductStockResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductStock",
			Filters:    bulkFilterMap,
		})
	}

	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return productsStockResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productsStockResp, err
	}

	if err := json.Unmarshal(body, &productsStockResp); err != nil {
		return productsStockResp, fmt.Errorf("ERPLY API: failed to unmarshal GetProductStockFileResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&productsStockResp.Status) {
		return productsStockResp, sharedCommon.NewErplyError(productsStockResp.Status.ErrorCode.String(), productsStockResp.Status.Request+": "+productsStockResp.Status.ResponseStatus, productsStockResp.Status.ErrorCode)
	}

	for _, prodBulkItem := range productsStockResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsStockResp, sharedCommon.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus, productsStockResp.Status.ErrorCode)
		}
	}

	return productsStockResp, nil
}

func (cli *Client) GetProductStockFileBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductStockFileResponseBulk, error) {
	var productsStockResp GetProductStockFileResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductStock",
			Filters:    bulkFilterMap,
		})
	}
	baseFilters["responseType"] = ResponseTypeCSV
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return productsStockResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productsStockResp, err
	}

	if err := json.Unmarshal(body, &productsStockResp); err != nil {
		return productsStockResp, fmt.Errorf("ERPLY API: failed to unmarshal GetProductStockFileResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&productsStockResp.Status) {
		return productsStockResp, sharedCommon.NewErplyError(
			productsStockResp.Status.ErrorCode.String(),
			productsStockResp.Status.Request+": "+productsStockResp.Status.ResponseStatus,
			productsStockResp.Status.ErrorCode,
		)
	}

	for _, prodBulkItem := range productsStockResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsStockResp, sharedCommon.NewErplyError(
				prodBulkItem.Status.ErrorCode.String(),
				prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus,
				productsStockResp.Status.ErrorCode,
			)
		}
	}

	return productsStockResp, nil
}

func (cli *Client) SaveAssortment(ctx context.Context, filters map[string]string) (SaveAssortmentResult, error) {
	resp, err := cli.SendRequest(ctx, "saveAssortment", filters)
	if err != nil {
		return SaveAssortmentResult{}, err
	}
	var res SaveAssortmentResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return SaveAssortmentResult{}, sharedCommon.NewFromError("failed to unmarshal SaveAssortmentResult", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return SaveAssortmentResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveAssortmentResults) > 0 {
		return res.SaveAssortmentResults[0], nil
	}

	return SaveAssortmentResult{}, nil
}

func (cli *Client) SaveAssortmentBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveAssortmentResponseBulk, error) {
	var assortmentResp SaveAssortmentResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveAssortment",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return assortmentResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assortmentResp, err
	}

	if err := json.Unmarshal(body, &assortmentResp); err != nil {
		return assortmentResp, fmt.Errorf("ERPLY API: failed to unmarshal SaveAssortmentResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&assortmentResp.Status) {
		return assortmentResp, sharedCommon.NewErplyError(
			assortmentResp.Status.ErrorCode.String(),
			assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus,
			assortmentResp.Status.ErrorCode,
		)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, sharedCommon.NewErplyError(
				assortmentItem.Status.ErrorCode.String(),
				assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus,
				assortmentResp.Status.ErrorCode,
			)
		}
	}

	return assortmentResp, nil
}

func (cli *Client) AddAssortmentProducts(ctx context.Context, filters map[string]string) (AddAssortmentProductsResult, error) {
	resp, err := cli.SendRequest(ctx, "addAssortmentProducts", filters)
	if err != nil {
		return AddAssortmentProductsResult{}, err
	}
	var res AddAssortmentProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return AddAssortmentProductsResult{}, sharedCommon.NewFromError("failed to unmarshal AddAssortmentProductsResult", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return AddAssortmentProductsResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.AddAssortmentProductsResults) > 0 {
		return res.AddAssortmentProductsResults[0], nil
	}

	return AddAssortmentProductsResult{}, nil
}

func (cli *Client) AddAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddAssortmentProductsResponseBulk, error) {
	var assortmentResp AddAssortmentProductsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "addAssortmentProducts",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return assortmentResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assortmentResp, err
	}

	if err := json.Unmarshal(body, &assortmentResp); err != nil {
		return assortmentResp, fmt.Errorf("ERPLY API: failed to unmarshal AddAssortmentProductsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&assortmentResp.Status) {
		return assortmentResp, sharedCommon.NewErplyError(
			assortmentResp.Status.ErrorCode.String(),
			assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus,
			assortmentResp.Status.ErrorCode,
		)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, sharedCommon.NewErplyError(
				assortmentItem.Status.ErrorCode.String(),
				assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus,
				assortmentItem.Status.ErrorCode,
			)
		}
	}

	return assortmentResp, nil
}

func (cli *Client) EditAssortmentProducts(ctx context.Context, filters map[string]string) (EditAssortmentProductsResult, error) {
	resp, err := cli.SendRequest(ctx, "editAssortmentProducts", filters)
	if err != nil {
		return EditAssortmentProductsResult{}, err
	}
	var res EditAssortmentProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return EditAssortmentProductsResult{}, sharedCommon.NewFromError("failed to unmarshal EditAssortmentProductsResult", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return EditAssortmentProductsResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.EditAssortmentProductsResults) > 0 {
		return res.EditAssortmentProductsResults[0], nil
	}

	return EditAssortmentProductsResult{}, nil
}

func (cli *Client) EditAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (EditAssortmentProductsResponseBulk, error) {
	var assortmentResp EditAssortmentProductsResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "editAssortmentProducts",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return assortmentResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assortmentResp, err
	}

	if err := json.Unmarshal(body, &assortmentResp); err != nil {
		return assortmentResp, fmt.Errorf("ERPLY API: failed to unmarshal EditAssortmentProductsResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&assortmentResp.Status) {
		return assortmentResp, sharedCommon.NewErplyError(
			assortmentResp.Status.ErrorCode.String(),
			assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus,
			assortmentResp.Status.ErrorCode,
		)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, sharedCommon.NewErplyError(
				assortmentItem.Status.ErrorCode.String(),
				assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus,
				assortmentItem.Status.ErrorCode,
			)
		}
	}

	return assortmentResp, nil
}

func (cli *Client) RemoveAssortmentProducts(ctx context.Context, filters map[string]string) (RemoveAssortmentProductResult, error) {
	resp, err := cli.SendRequest(ctx, "removeAssortmentProducts", filters)
	if err != nil {
		return RemoveAssortmentProductResult{}, err
	}
	var res RemoveAssortmentProductResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return RemoveAssortmentProductResult{}, sharedCommon.NewFromError("failed to unmarshal RemoveAssortmentProductResults", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return RemoveAssortmentProductResult{}, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.RemoveAssortmentProductResults) > 0 {
		return res.RemoveAssortmentProductResults[0], nil
	}

	return RemoveAssortmentProductResult{}, nil
}

func (cli *Client) RemoveAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (RemoveAssortmentProductResponseBulk, error) {
	var assortmentResp RemoveAssortmentProductResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "removeAssortmentProducts",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return assortmentResp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return assortmentResp, err
	}

	if err := json.Unmarshal(body, &assortmentResp); err != nil {
		return assortmentResp, fmt.Errorf("ERPLY API: failed to unmarshal RemoveAssortmentProductResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&assortmentResp.Status) {
		return assortmentResp, sharedCommon.NewErplyError(
			assortmentResp.Status.ErrorCode.String(),
			assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus,
			assortmentResp.Status.ErrorCode,
		)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, sharedCommon.NewErplyError(
				assortmentItem.Status.ErrorCode.String(),
				assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus,
				assortmentResp.Status.ErrorCode,
			)
		}
	}

	return assortmentResp, nil
}

func (cli *Client) SaveProductCategory(ctx context.Context, filters map[string]string) (result SaveProductCategoryResult, err error) {
	resp, err := cli.SendRequest(ctx, "saveProductCategory", filters)
	if err != nil {
		return result, err
	}
	var res SaveProductCategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return result, sharedCommon.NewFromError("failed to unmarshal SaveAssortmentResult", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return result, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveProductCategoryResults) > 0 {
		return res.SaveProductCategoryResults[0], nil
	}

	return result, nil
}

func (cli *Client) SaveProductCategoryBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SaveProductCategoryResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveProductCategory",
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
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveProductCategoryResponseBulk from '%s': %v", string(body), err)
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

func (cli *Client) SaveBrand(ctx context.Context, filters map[string]string) (result SaveBrandResult, err error) {
	resp, err := cli.SendRequest(ctx, "saveBrand", filters)
	if err != nil {
		return result, err
	}
	var res SaveBrandResultResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return result, sharedCommon.NewFromError("failed to unmarshal SaveBrandResultResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return result, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveBrandResults) > 0 {
		return res.SaveBrandResults[0], nil
	}

	return result, nil
}

func (cli *Client) SaveBrandBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SaveBrandResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveBrand",
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
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveBrandResponseBulk from '%s': %v", string(body), err)
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

func (cli *Client) SaveProductPriorityGroup(ctx context.Context, filters map[string]string) (result SaveProductPriorityGroupResult, err error) {
	resp, err := cli.SendRequest(ctx, "saveProductPriorityGroup", filters)
	if err != nil {
		return result, err
	}
	var res SaveProductPriorityGroupResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return result, sharedCommon.NewFromError("failed to unmarshal SaveProductPriorityGroupResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return result, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveProductPriorityGroupResults) > 0 {
		return res.SaveProductPriorityGroupResults[0], nil
	}

	return result, nil
}

func (cli *Client) SaveProductPriorityGroupBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SaveProductPriorityGroupResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveProductPriorityGroup",
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
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveProductPriorityGroupResponseBulk from '%s': %v", string(body), err)
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

func (cli *Client) GetProductPriorityGroupBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk GetProductPriorityGroupResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductPriorityGroups",
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

	bodyStr := string(body)

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal GetProductPriorityGroupResponseBulk from '%s': %v", bodyStr, err)
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

func (cli *Client) GetProductCategoriesBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk GetProductCategoryResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductCategories",
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

	bodyStr := string(body)

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal GetProductCategoryResponseBulk from '%s': %v", bodyStr, err)
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

func (cli *Client) GetProductGroupsBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk GetProductGroupResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "getProductGroups",
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

	bodyStr := string(body)

	if err := json.Unmarshal(body, &respBulk); err != nil {
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal GetProductGroupResponseBulk from '%s': %v", bodyStr, err)
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

func (cli *Client) SaveProductGroup(ctx context.Context, filters map[string]string) (result SaveProductGroupResult, err error) {
	resp, err := cli.SendRequest(ctx, "saveProductGroup", filters)
	if err != nil {
		return result, err
	}
	var res SaveProductGroupResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return result, sharedCommon.NewFromError("failed to unmarshal SaveProductGroupResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return result, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	if len(res.SaveProductGroupResults) > 0 {
		return res.SaveProductGroupResults[0], nil
	}

	return result, nil
}

func (cli *Client) SaveProductGroupBulk(
	ctx context.Context,
	bulkFilters []map[string]interface{},
	baseFilters map[string]string,
) (respBulk SaveProductGroupResponseBulk, err error) {
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "saveProductGroup",
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
		return respBulk, fmt.Errorf("ERPLY API: failed to unmarshal SaveProductGroupResponseBulk from '%s': %v", string(body), err)
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

func (cli *Client) DeleteProductGroup(ctx context.Context, filters map[string]string) error {
	resp, err := cli.SendRequest(ctx, "deleteProductGroup", filters)
	if err != nil {
		return err
	}
	var res DeleteProductGroupResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return sharedCommon.NewFromError("failed to unmarshal DeleteProductGroupResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return nil
}

func (cli *Client) DeleteProductGroupBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductGroupResponseBulk, error) {
	var deleteRespBulk DeleteProductGroupResponseBulk
	bulkInputs := make([]common.BulkInput, 0, len(bulkFilters))
	for _, bulkFilterMap := range bulkFilters {
		bulkInputs = append(bulkInputs, common.BulkInput{
			MethodName: "deleteProductGroup",
			Filters:    bulkFilterMap,
		})
	}
	resp, err := cli.SendRequestBulk(ctx, bulkInputs, baseFilters)
	if err != nil {
		return deleteRespBulk, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deleteRespBulk, err
	}

	if err := json.Unmarshal(body, &deleteRespBulk); err != nil {
		return deleteRespBulk, fmt.Errorf("ERPLY API: failed to unmarshal DeleteProductGroupResponseBulk from '%s': %v", string(body), err)
	}
	if !common.IsJSONResponseOK(&deleteRespBulk.Status) {
		return deleteRespBulk, sharedCommon.NewErplyError(deleteRespBulk.Status.ErrorCode.String(), deleteRespBulk.Status.Request+": "+deleteRespBulk.Status.ResponseStatus, deleteRespBulk.Status.ErrorCode)
	}

	for _, prodBulkItem := range deleteRespBulk.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return deleteRespBulk, sharedCommon.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus, deleteRespBulk.Status.ErrorCode)
		}
	}

	return deleteRespBulk, nil
}

func (cli *Client) GetProductPictures(ctx context.Context, filters map[string]string) ([]Image, error) {
	resp, err := cli.SendRequest(ctx, "getProductPictures", filters)
	if err != nil {
		return nil, err
	}
	var res GetProductPicturesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProductPicturesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.Records, nil
}
