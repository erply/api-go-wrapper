package addresses

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
)

type Manager interface {
	GetAddresses(ctx context.Context, filters map[string]string) ([]common.Address, error)
	SaveAddress(ctx context.Context, filters map[string]string) ([]common.Address, error)
}
