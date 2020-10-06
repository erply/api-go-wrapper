package products

import "context"

type Manager interface {
	GetProducts(ctx context.Context, filters map[string]string) ([]Product, error)
	GetProductsCount(ctx context.Context, filters map[string]string) (int, error)
	GetProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsResponseBulk, error)
	GetProductUnits(ctx context.Context, filters map[string]string) ([]ProductUnit, error)
	GetProductCategories(ctx context.Context, filters map[string]string) ([]ProductCategory, error)
	GetProductBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetBrands(ctx context.Context, filters map[string]string) ([]ProductBrand, error)
	GetProductGroups(ctx context.Context, filters map[string]string) ([]ProductGroup, error)
	GetProductStock(ctx context.Context, filters map[string]string) ([]GetProductStock, error)
	GetProductStockFile(ctx context.Context, filters map[string]string) ([]GetProductStockFile, error)
	GetProductStockFileBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductStockFileResponseBulk, error)
	SaveProduct(ctx context.Context, filters map[string]string) (SaveProductResult, error)
	SaveProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveProductResponseBulk, error)
	DeleteProduct(ctx context.Context, filters map[string]string) error
	DeleteProductBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductResponseBulk, error)
	SaveAssortment(ctx context.Context, filters map[string]string) (SaveAssortmentResult, error)
	SaveAssortmentBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (SaveAssortmentResponseBulk, error)
	AddAssortmentProducts(ctx context.Context, filters map[string]string) (AddAssortmentProductsResult, error)
	AddAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddAssortmentProductsResponseBulk, error)
	EditAssortmentProducts(ctx context.Context, filters map[string]string) (EditAssortmentProductsResult, error)
	EditAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (EditAssortmentProductsResponseBulk, error)
	RemoveAssortmentProducts(ctx context.Context, filters map[string]string) (RemoveAssortmentProductResult, error)
	RemoveAssortmentProductsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (RemoveAssortmentProductResponseBulk, error)
}
