package servicediscovery

import (
	"context"
	"github.com/breathbath/api-go-wrapper/internal/common"
)

type (
	Manager interface {
		GetServiceEndpoints(ctx context.Context) (*ServiceEndpoints, error)
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
