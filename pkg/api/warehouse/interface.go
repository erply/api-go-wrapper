package warehouse

import "context"

type (
	Manager interface {
		GetWarehouses(ctx context.Context) (Warehouses, error)
	}
)
