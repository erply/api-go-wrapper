package sales

import "context"

type Manager interface {
	ProjectManager
	DocumentManager
	VatRateManager
	//payment requests
	SavePayment(ctx context.Context, filters map[string]string) (int64, error)
	GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)

	//shopping cart
	CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error)
}
