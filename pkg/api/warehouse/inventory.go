package warehouse

import "context"

type InventoryManager interface {
	SaveInventoryRegistration(ctx context.Context, filters map[string]string) (inventoryRegistrationID int, err error)
	SaveInventoryRegistrationBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveInventoryRegistrationResponseBulk, error)
}
