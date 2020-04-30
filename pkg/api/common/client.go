package common

import "net/http"

func NewClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {
	if httpCli == nil {
		httpCli = GetDefaultHTTPClient()
	}
	return &Client{
		url:        GetBaseURL(cc),
		httpClient: httpCli,
		sessionKey: sk,
		clientCode: cc,
		partnerKey: partnerKey,
	}
}

type Client struct {
	url        string
	httpClient *http.Client
	sessionKey string
	clientCode string
	partnerKey string
}

func (cli *Client) Close() {
	cli.httpClient.CloseIdleConnections()
}
