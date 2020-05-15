package company

import "context"

type Manager interface {
	GetCompanyInfo(ctx context.Context) (*Info, error)
	GetConfParameters(ctx context.Context) (*ConfParameter, error)
}
