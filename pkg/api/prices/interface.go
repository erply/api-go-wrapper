package prices

import "context"

type Manager interface {
	GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error)
	AddProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	EditProductToSupplierPriceList(ctx context.Context, filters map[string]string) (*ChangeProductToSupplierPriceListResult, error)
	ChangeProductToSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (ChangeProductToSupplierPriceListResponseBulk, error)
	GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error)
	GetProductPriceLists(ctx context.Context, filters map[string]string) ([]ProductPriceList, error)
	GetProductPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductPriceListResponseBulk, error)
	DeleteProductsFromSupplierPriceList(ctx context.Context, filters map[string]string) (*DeleteProductsFromSupplierPriceListResult, error)
	DeleteProductsFromSupplierPriceListBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (DeleteProductsFromSupplierPriceListResponseBulk, error)
	SaveSupplierPriceList(ctx context.Context, filters map[string]string) (*SaveSupplierPriceListResult, error)
	SaveSupplierPriceListBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveSupplierPriceListResponseBulk, error)
}
