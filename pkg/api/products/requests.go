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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Products, nil
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(res.Status.ErrorCode.String(), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductGroups, nil
}
