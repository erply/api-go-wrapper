package prices

import (
	"github.com/erply/api-go-wrapper/internal/common"
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
