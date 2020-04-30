package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"net/http"
)

type (
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
	Client struct {
		*common.Client
	}
)

func NewClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {

	cli := &Client{
		common.NewClient(sk, cc, partnerKey, httpCli),
	}
	return cli
}
