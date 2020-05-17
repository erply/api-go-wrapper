package warehouse

import (
	"github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
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
