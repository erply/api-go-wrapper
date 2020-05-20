package addresses

import (
	"context"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type Manager interface {
	GetAddresses(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error)
	SaveAddress(ctx context.Context, filters map[string]string) ([]sharedCommon.Address, error)
}
