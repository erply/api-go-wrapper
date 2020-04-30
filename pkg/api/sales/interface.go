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
		DeleteDocument(ctx context.Context, filters map[string]string) error
		SavePurchaseDocument(ctx context.Context, filters map[string]string) (PurchaseDocImportReports, error)
	}

	VatRateManager interface {
		GetVatRates(ctx context.Context, filters map[string]string) (VatRates, error)
	}

	Manager interface {
		ProjectManager
		DocumentManager
		VatRateManager
		//payment requests
		SavePayment(ctx context.Context, filters map[string]string) (int64, error)
		GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)

		//shopping cart
		CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error)
	}
)
