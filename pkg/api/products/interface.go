package products

import "context"

type Manager interface {
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
}
