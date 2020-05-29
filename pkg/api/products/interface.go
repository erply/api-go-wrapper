package products

import "context"

type Manager interface {
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsResponseBulk, error)
	GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
}
