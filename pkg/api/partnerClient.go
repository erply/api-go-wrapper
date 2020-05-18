package api

import (
	"errors"
	"github.com/breathbath/api-go-wrapper/internal/common"
	"github.com/breathbath/api-go-wrapper/pkg/api/auth"
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
	comCli := common.NewClient(sessionKey, clientCode, partnerKey, customCli, nil)

	return &PartnerClient{
		Client:               newErplyClient(comCli),
		PartnerTokenProvider: auth.NewPartnerClient(comCli),
	}, nil
}
