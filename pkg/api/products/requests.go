package products

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Products, nil
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
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
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductGroups, nil
}
