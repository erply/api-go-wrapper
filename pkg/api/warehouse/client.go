package warehouse

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"net/http"
)

type (
	Manager interface {
		GetWarehouses(ctx context.Context) (Warehouses, error)
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
