package sales

import "context"

type (
	ProjectManager interface {
		GetProjects(ctx context.Context, filters map[string]string) ([]Project, error)
		GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error)
	}
	DocumentManager interface {
		SaveSalesDocument(ctx context.Context, filters map[string]string) (SaleDocImportReports, error)
		GetSalesDocuments(ctx context.Context, filters map[string]string) ([]SaleDocument, error)
		GetSalesDocumentsBulk(ctx context.Context, bulkFilters []map[string]interface{}, baseFilters map[string]string) (GetSaleDocumentResponseBulk, error)
		DeleteDocument(ctx context.Context, filters map[string]string) error
		SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error)
	}

	VatRateManager interface {
		GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error)
	}

	AssignmentsManger interface {
		SaveAssignment(ctx context.Context, filters map[string]string) (int64, error)
		SaveVatRate(ctx context.Context, filters map[string]string) (*SaveVatRateResult, error)
		SaveVatRateBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateResponseBulk, error)
		SaveVatRateComponent(ctx context.Context, filters map[string]string) (*SaveVatRateComponentResult, error)
		SaveVatRateComponentBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (SaveVatRateComponentResponseBulk, error)
	}

	Manager interface {
		ProjectManager
		DocumentManager
		VatRateManager
		AssignmentsManger
		//payment requests
		SavePayment(ctx context.Context, filters map[string]string) (int64, error)
		GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)

		//shopping cart
		CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error)
	}
)
