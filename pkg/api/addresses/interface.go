package addresses

import (
	"context"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type Manager interface {
	GetAddresses(ctx context.Context, filters map[string]string) ([]common2.Address, error)
	SaveAddress(ctx context.Context, filters map[string]string) ([]common2.Address, error)
}
