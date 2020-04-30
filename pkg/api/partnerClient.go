package api

import (
	"errors"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/erply/api-go-wrapper/pkg/common"
	"net/http"
)

type PartnerClient struct {
	Client               *Client
	PartnerTokenProvider auth.PartnerTokenProvider
}

func NewPartnerClient(sessionKey, clientCode, partnerKey string, customCli *http.Client) (*PartnerClient, error) {
	if sessionKey == "" || clientCode == "" || partnerKey == "" {
		return nil, errors.New("sessionKey, clientCode and partnerKey are required")
	}
	comCli := common.NewClient(sessionKey, clientCode, partnerKey, customCli)

	return &PartnerClient{
		Client:               newErplyClient(comCli),
		PartnerTokenProvider: auth.NewPartnerClient(comCli),
	}, nil
}
