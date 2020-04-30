package company

import "github.com/erply/api-go-wrapper/pkg/common"

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
