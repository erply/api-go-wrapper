package common

import (
	"net/http"
	"net/url"
)

type AuthFunc func(string) url.Values

func NewClient(sk, cc, partnerKey string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
	if httpCli == nil {
		httpCli = GetDefaultHTTPClient()
	}
	cli := &Client{
		httpClient:  httpCli,
		sessionKey:  sk,
		clientCode:  cc,
		partnerKey:  partnerKey,
		headersFunc: headersForEveryRequestFunc,
	}

	if cli.headersFunc == nil {
		cli.headersFunc = cli.getDefaultMandatoryHeaders
	}
	if cc != "" {
		cli.Url = GetBaseURL(cc)
	} else {
		cli.Url = GetBaseURLFromAuthFunc(cli.headersFunc)
	}
	return cli
}

type Client struct {
	Url         string
	httpClient  *http.Client
	sessionKey  string
	clientCode  string
	partnerKey  string
	headersFunc func(string) url.Values
}

func (cli *Client) Close() {
	cli.httpClient.CloseIdleConnections()
}
