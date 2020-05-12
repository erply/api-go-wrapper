package addresses

import (
	"context"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
)

type Manager interface {
	GetAddresses(ctx context.Context, filters map[string]string) ([]common.Address, error)
	SaveAddress(ctx context.Context, filters map[string]string) ([]common.Address, error)
}
