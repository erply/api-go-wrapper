package warehouse

import "context"

type InventoryManager interface {
	SaveInventoryRegistration(ctx context.Context, filters map[string]string) (inventoryRegistrationID int, err error)
	SaveInventoryRegistrationBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveInventoryRegistrationResponseBulk, error)
	SaveInventoryWriteOff(ctx context.Context, filters map[string]string) (inventoryWriteOffID int, err error)
	SaveInventoryTransfer(ctx context.Context, filters map[string]string) (inventoryTransferID int, err error)
	GetReasonCodes(ctx context.Context, filters map[string]string)
}
