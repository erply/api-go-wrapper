package warehouse

import "context"

type (
	Manager interface {
		GetWarehouses(ctx context.Context, filters map[string]string) (Warehouses, error)
		GetWarehousesBulk(
			ctx context.Context,
			bulkRequest []map[string]interface{},
			baseFilters map[string]string) (
			GetWarehousesResponseBulk,
			error,
		)
		SaveWarehouse(ctx context.Context, filters map[string]string) (*SaveWarehouseResult, error)
		SaveWarehouseBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveWarehouseResponseBulk, error)
		InventoryManager
	}
)
