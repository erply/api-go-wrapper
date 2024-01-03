package api

import (
	"errors"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"net/http"
)

type PartnerClient struct {
	Client               *Client
	PartnerTokenProvider auth.PartnerTokenProvider
}

func NewPartnerClientFromCredentials(username, password, clientCode, partnerKey string, customCli *http.Client) (*PartnerClient, error) {
	if customCli == nil {
		customCli = common.GetDefaultHTTPClient()
	}
	sessionKey, err := auth.VerifyUser(username, password, clientCode, customCli)
	if err != nil {
		return nil, err
	}

	return NewPartnerClient(sessionKey, clientCode, partnerKey, customCli)
}

// Deprecated
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
