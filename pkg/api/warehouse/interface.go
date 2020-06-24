package warehouse

import "context"

type (
	Manager interface {
		GetWarehouses(ctx context.Context) (Warehouses, error)
		GetWarehousesBulk(
			ctx context.Context,
			bulkRequest []map[string]interface{},
			baseFilters map[string]string) (
			GetWarehousesResponseBulk,
			error,
		)
	}
)
