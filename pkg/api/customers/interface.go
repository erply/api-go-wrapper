package customers

import "context"

type Manager interface {
	SaveCustomer(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
	SaveCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (SaveCustomerResponseBulk, error)
	GetCustomers(ctx context.Context, filters map[string]string) ([]Customer, error)
	GetCustomersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetCustomersResponseBulk, error)
	DeleteCustomer(ctx context.Context, filters map[string]string) error
	DeleteCustomerBulk(ctx context.Context, customerMap []map[string]interface{}, attrs map[string]string) (DeleteCustomersResponseBulk, error)
	VerifyCustomerUser(ctx context.Context, username, password string) (*WebshopClient, error)
	ValidateCustomerUsername(ctx context.Context, username string) (bool, error)
	GetCustomerGroups(ctx context.Context, filters map[string]string) ([]CustomerGroup, error)
	GetSuppliers(ctx context.Context, filters map[string]string) ([]Supplier, error)
	GetSuppliersBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSuppliersResponseBulk, error)
	SaveSupplier(ctx context.Context, filters map[string]string) (*CustomerImportReport, error)
	SaveSupplierBulk(ctx context.Context, suppliers []map[string]interface{}, attrs map[string]string) (SaveSuppliersResponseBulk, error)
	DeleteSupplier(ctx context.Context, filters map[string]string) error
	DeleteSupplierBulk(ctx context.Context, supplierMap []map[string]interface{}, attrs map[string]string) (DeleteSuppliersResponseBulk, error)
	AddCustomerRewardPoints(ctx context.Context, filters map[string]string) (AddCustomerRewardPointsResult, error)
	AddCustomerRewardPointsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (AddCustomerRewardPointsResponseBulk, error)
	GetCompanyTypes(ctx context.Context, filters map[string]string) ([]CompanyType, error)
	SaveCompanyType(ctx context.Context, filters map[string]string) (*SaveCompanyTypeResponse, error)
	SaveSupplierGroup(ctx context.Context, filters map[string]string) (*SaveSupplierGroupResponse, error)
}
