package pos

import "context"

type (
	Manager interface {
		GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error)
		GetClockIns(ctx context.Context, filters map[string]string) ([]Clocking, error)
	}
)
