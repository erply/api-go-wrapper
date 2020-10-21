package addresses

import (
	"context"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type Manager interface {
	GetAddresses(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error)
	GetAddressesBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetAddressesResponseBulk, error)
	SaveAddress(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error)
	SaveAddressesBulk(ctx context.Context, addrMap []map[string]interface{}, attrs map[string]string) (SaveAddressesResponseBulk, error)
	DeleteAddress(ctx context.Context, filters map[string]string) error
	DeleteAddressBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (DeleteAddressResponseBulk, error)
}
