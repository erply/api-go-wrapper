package warehouse

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
)

type (
	Manager interface {
		GetWarehouses(ctx context.Context) (Warehouses, error)
	}
	Client struct {
		*common.Client
	}
)

func NewClient(client *common.Client) *Client {

	cli := &Client{
		client,
	}
	return cli
}
