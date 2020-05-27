package prices

import "context"

type Manager interface {
	GetSupplierPriceLists(ctx context.Context, filters map[string]string) ([]PriceList, error)
	GetSupplierPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetPriceListsResponseBulk, error)
	GetProductPriceLists(ctx context.Context, filters map[string]string) ([]ProductPriceList, error)
	GetProductPriceListsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetProductPriceListResponseBulk, error)
}
