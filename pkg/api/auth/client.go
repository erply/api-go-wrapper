package auth

import "github.com/erply/api-go-wrapper/pkg/api/common"

type (
	Client        struct{ *common.Client }
	PartnerClient struct{ *common.Client }
)

func NewClient(client *common.Client) *Client {

	cli := &Client{
		client,
	}
	return cli
}

func NewPartnerClient(client *common.Client) *PartnerClient {
	return &PartnerClient{client}
}
