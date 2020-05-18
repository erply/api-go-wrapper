package api

import (
	"errors"
	"github.com/breathbath/api-go-wrapper/internal/common"
	"github.com/breathbath/api-go-wrapper/pkg/api/auth"
	"net/http"
	"net/url"
)

type PartnerClient struct {
	Client               *Client
	PartnerTokenProvider auth.PartnerTokenProvider
}

func NewPartnerClient(sessionKey, clientCode, partnerKey string, customCli *http.Client, headersSetToEveryRequest func(requestName string) url.Values) (*PartnerClient, error) {
	if sessionKey == "" || clientCode == "" || partnerKey == "" {
		return nil, errors.New("sessionKey, clientCode and partnerKey are required")
	}
	comCli := common.NewClient(sessionKey, clientCode, partnerKey, customCli, headersSetToEveryRequest)

	return &PartnerClient{
		Client:               newErplyClient(comCli),
		PartnerTokenProvider: auth.NewPartnerClient(comCli),
	}, nil
}
