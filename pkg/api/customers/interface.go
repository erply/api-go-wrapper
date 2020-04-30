package customers

import "context"

type Manager interface {
	SaveCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
	GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
	VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error)
	ValidateCustomerUsername(ctx context.Context, username string) (bool, error)
	GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
	SaveSupplier(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
}
