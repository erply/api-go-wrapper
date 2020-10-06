package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"io/ioutil"
)

func (cli *Client) GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error) {

	resp, err := cli.SendRequest(ctx, "getProductUnits", filters)
	if err != nil {
		return nil, erro.NewFromError("GetProductUnits request failed", err)
	}

	res := &GetProductUnitsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetProductUnitsResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal GetProductsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return 0, erro.NewFromError("failed to unmarshal GetProductsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return 0, erro.NewFromResponseStatus(&res.Status)
	}
	return res.Status.RecordsTotal, nil
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
		return productsResp, erro.NewErplyError(productsResp.Status.ErrorCode.String(), productsResp.Status.Request+": "+productsResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range productsResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
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
		return SaveProductResult{}, erro.NewFromError("failed to unmarshal SaveProductResult", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return SaveProductResult{}, erro.NewFromResponseStatus(&res.Status)
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
		return productsResp, erro.NewErplyError(productsResp.Status.ErrorCode.String(), productsResp.Status.Request+": "+productsResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range productsResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
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
		return erro.NewFromError("failed to unmarshal DeleteProductResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return erro.NewFromResponseStatus(&res.Status)
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
		return deleteRespBulk, erro.NewErplyError(deleteRespBulk.Status.ErrorCode.String(), deleteRespBulk.Status.Request+": "+deleteRespBulk.Status.ResponseStatus)
	}

	for _, prodBulkItem := range deleteRespBulk.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return deleteRespBulk, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
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
		return nil, erro.NewFromError("failed to unmarshal getProductCategoriesResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal getProductBrandsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal getBrandsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal getProductGroupsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal GetProductStockResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
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
		return nil, erro.NewFromError("failed to unmarshal GetProductStockFileResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewFromResponseStatus(&res.Status)
	}
	return res.GetProductStockFile, nil
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
		return productsStockResp, erro.NewErplyError(productsStockResp.Status.ErrorCode.String(), productsStockResp.Status.Request+": "+productsStockResp.Status.ResponseStatus)
	}

	for _, prodBulkItem := range productsStockResp.BulkItems {
		if !common.IsJSONResponseOK(&prodBulkItem.Status.Status) {
			return productsStockResp, erro.NewErplyError(prodBulkItem.Status.ErrorCode.String(), prodBulkItem.Status.Request+": "+prodBulkItem.Status.ResponseStatus)
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
		return SaveAssortmentResult{}, erro.NewFromError("failed to unmarshal SaveAssortmentResult", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return SaveAssortmentResult{}, erro.NewFromResponseStatus(&res.Status)
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
		return assortmentResp, erro.NewErplyError(assortmentResp.Status.ErrorCode.String(), assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, erro.NewErplyError(assortmentItem.Status.ErrorCode.String(), assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus)
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
		return AddAssortmentProductsResult{}, erro.NewFromError("failed to unmarshal AddAssortmentProductsResult", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return AddAssortmentProductsResult{}, erro.NewFromResponseStatus(&res.Status)
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
		return assortmentResp, erro.NewErplyError(assortmentResp.Status.ErrorCode.String(), assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, erro.NewErplyError(assortmentItem.Status.ErrorCode.String(), assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus)
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
		return EditAssortmentProductsResult{}, erro.NewFromError("failed to unmarshal EditAssortmentProductsResult", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return EditAssortmentProductsResult{}, erro.NewFromResponseStatus(&res.Status)
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
		return assortmentResp, erro.NewErplyError(assortmentResp.Status.ErrorCode.String(), assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, erro.NewErplyError(assortmentItem.Status.ErrorCode.String(), assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus)
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
		return RemoveAssortmentProductResult{}, erro.NewFromError("failed to unmarshal RemoveAssortmentProductResults", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return RemoveAssortmentProductResult{}, erro.NewFromResponseStatus(&res.Status)
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
		return assortmentResp, erro.NewErplyError(assortmentResp.Status.ErrorCode.String(), assortmentResp.Status.Request+": "+assortmentResp.Status.ResponseStatus)
	}

	for _, assortmentItem := range assortmentResp.BulkItems {
		if !common.IsJSONResponseOK(&assortmentItem.Status.Status) {
			return assortmentResp, erro.NewErplyError(assortmentItem.Status.ErrorCode.String(), assortmentItem.Status.Request+": "+assortmentItem.Status.ResponseStatus)
		}
	}

	return assortmentResp, nil
}
