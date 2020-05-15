package common

import (
	"net/http"
	"net/url"
)

func NewClient(sk, cc, partnerKey string, httpCli *http.Client, headersForEveryRequestFunc func(string) url.Values) *Client {
	if httpCli == nil {
		httpCli = GetDefaultHTTPClient()
	}
	cli := &Client{
		url:         GetBaseURL(cc),
		httpClient:  httpCli,
		sessionKey:  sk,
		clientCode:  cc,
		partnerKey:  partnerKey,
		headersFunc: headersForEveryRequestFunc,
	}

	if cli.headersFunc == nil {
		cli.headersFunc = cli.getDefaultMandatoryHeaders
	}
	return cli
}

type Client struct {
	url         string
	httpClient  *http.Client
	sessionKey  string
	clientCode  string
	partnerKey  string
	headersFunc func(string) url.Values
}

func (cli *Client) Close() {
	cli.httpClient.CloseIdleConnections()
}
