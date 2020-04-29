package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
	"strings"
)

//GetProductsResponse ...
type GetProductsResponse struct {
	Status   Status    `json:"status"`
	Products []Product `json:"records"`
}

type getProductCategoriesResponse struct {
	Status            Status            `json:"status"`
	ProductCategories []ProductCategory `json:"records"`
}

type getProductBrandsResponse struct {
	Status        Status         `json:"status"`
	ProductBrands []ProductBrand `json:"records"`
}

type getProductGroupsResponse struct {
	Status        Status         `json:"status"`
	ProductGroups []ProductGroup `json:"records"`
}

type ProductDimensions struct {
	Name             string `json:"name"`
	Value            string `json:"value"`
	Order            int    `json:"order"`
	DimensionID      int    `json:"dimensionID"`
	DimensionValueID int    `json:"dimensionValueID"`
}

type ProductVariaton struct {
	ProductID  string              `json:"productID"`
	Name       string              `json:"name"`
	Code       string              `json:"code"`
	Code2      string              `json:"code2"`
	Dimensions []ProductDimensions `json:"dimensions"`
}

//Product ...
type Product struct {
	ProductID          int                `json:"productID"`
	ParentProductID    int                `json:"parentProductID"`
	Type               string             `json:"type"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	DescriptionLong    string             `json:"longdesc"`
	Status             string             `json:"status"`
	Code               string             `json:"code"`
	Code2              string             `json:"code2"`
	Code3              *string            `json:"code3"`
	Price              float64            `json:"price"`
	PriceWithVat       float32            `json:"priceWithVat"`
	UnitName           *string            `json:"unitName"`
	Images             []ProductImage     `json:"images"`
	DisplayedInWebshop byte               `json:"displayedInWebshop"`
	CategoryId         uint               `json:"categoryID"`
	CategoryName       string             `json:"categoryName"`
	BrandID            uint               `json:"brandID"`
	BrandName          string             `json:"brandName"`
	GroupID            uint               `json:"groupID"`
	GroupName          string             `json:"groupName"`
	Warehouses         map[uint]StockInfo `json:"warehouses"`
	RelatedProducts    []string           `json:"relatedProducts"`
	Vatrate            float64            `json:"vatrate"`
	ProductVariations  []string           `json:"productVariations"` // Variations of matrix product
	VariationList      []ProductVariaton  `json:"variationList"`
}

type StockInfo struct {
	WarehouseID   uint    `json:"warehouseID"`
	Free          int     `json:"free"`
	OrderPending  int     `json:"orderPending"`
	ReorderPoint  int     `json:"reorderPoint"`
	RestockLevel  int     `json:"restockLevel"`
	FifoCost      float32 `json:"FIFOCost"`
	PurchasePrice float32 `json:"purchasePrice"`
}

type ProductImage struct {
	PictureID       string  `json:"pictureID"`
	Name            string  `json:"name"`
	ThumbURL        string  `json:"thumbURL"`
	SmallURL        string  `json:"smallURL"`
	LargeURL        string  `json:"largeURL"`
	FullURL         string  `json:"fullURL"`
	External        byte    `json:"external"`
	HostingProvider string  `json:"hostingProvider"`
	Hash            *string `json:"hash"`
	Tenant          *string `json:"tenant"`
}

type ProductCategory struct {
	ProductCategoryID   uint   `json:"productCategoryID"`
	ParentCategoryID    uint   `json:"parentCategoryID"`
	ProductCategoryName string `json:"productCategoryName"`
	Added               uint64 `json:"added"`
	LastModified        uint64 `json:"lastModified"`
}

type ProductBrand struct {
	ID           uint   `json:"brandID"`
	Name         string `json:"name"`
	Added        uint64 `json:"added"`
	LastModified uint64 `json:"lastModified"`
}

type ProductGroup struct {
	ID              uint           `json:"productGroupID"`
	Name            string         `json:"name"`
	ShowInWebshop   string         `json:"showInWebshop"`
	NonDiscountable byte           `json:"nonDiscountable"`
	PositionNo      int            `json:"positionNo"`
	ParentGroupID   string         `json:"parentGroupID"`
	Added           uint64         `json:"added"`
	LastModified    uint64         `json:"lastModified"`
	SubGroups       []ProductGroup `json:"subGroups"`
}

//GetProductUnitsResponse ...
type GetProductUnitsResponse struct {
	Status       Status        `json:"status"`
	ProductUnits []ProductUnit `json:"records"`
}

//ProductUnit ...
type ProductUnit struct {
	UnitID string `json:"unitID"`
	Name   string `json:"name"`
}

func (cli *erplyClient) GetProductUnits() ([]ProductUnit, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProductUnits request", err)
	}

	params := getMandatoryParameters(cli, GetProductUnitsMethod)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProductUnits request failed", err)
	}

	res := &GetProductUnitsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductUnitsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.ProductUnits, nil
}

func (cli *erplyClient) GetProducts(ctx context.Context, filters map[string]string) ([]Product, error) {
	resp, err := cli.sendRequest(ctx, GetProductsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProductsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Products, nil
}

func (cli *erplyClient) GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error) {
	resp, err := cli.sendRequest(ctx, GetProductCategoriesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductCategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductCategoriesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductCategories, nil
}

func (cli *erplyClient) GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error) {
	resp, err := cli.sendRequest(ctx, GetProductBrandsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductBrandsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductBrandsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductBrands, nil
}

func (cli *erplyClient) GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error) {
	resp, err := cli.sendRequest(ctx, GetProductGroupsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res getProductGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal getProductGroupsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProductGroups, nil
}

//GetProductsByIDs - NOTE: if product's id is 0 - the product is not in the database. It was created during the sales document creation
func (cli *erplyClient) GetProductsByIDs(ids []string) ([]Product, error) {
	if len(ids) == 0 {
		return nil, erplyerr("No ids provided for products request", nil)
	}

	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProducts request", err)
	}

	params := getMandatoryParameters(cli, GetProductsMethod)
	params.Add("productIDs", strings.Join(ids, ","))
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProducts request failed", err)
	}

	res := &GetProductsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Products, nil
}

func (cli *erplyClient) GetProductsByCode3(code3 string) (*Product, error) {
	if code3 == "" {
		return nil, erplyerr("No code3 provided for product request", nil)
	}

	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetProducts request", err)
	}

	params := getMandatoryParameters(cli, GetProductsMethod)
	params.Add("code3", code3)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetProducts request failed", err)
	}

	res := &GetProductsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetProductsResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Products) == 0 {
		return nil, erplyerr("no such product found", err)
	}

	return &res.Products[0], nil
}
