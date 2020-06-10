package prices

import "context"

type Manager interface {
	GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error)
	AddProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	EditProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	ChangeProductToSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToSupplierPriceListResponseBulk, error)
	GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error)
	GetProductsInPriceList(ctx context.Context, filters map[string]string) ([]ProductsInPriceList, error)
	GetProductsInPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductsInPriceListResponseBulk, error)
	GetProductsInSupplierPriceList(ctx context.Context, filters map[string]string) ([]ProductsInSupplierPriceList, error)
	GetProductsInSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (ProductsInSupplierPriceListResponseBulk, error)
	DeleteProductsFromSupplierPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromSupplierPriceListResult, error)
	DeleteProductsFromSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductsFromSupplierPriceListResponseBulk, error)
	SaveSupplierPriceList(ctx context.Context, filters map[string]string) (*SaveSupplierPriceListResult, error)
	SaveSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveSupplierPriceListResponseBulk, error)
}
