package common

import (
	"net/http"
	"net/url"
)

type AuthFunc func(string) url.Values

//NewClientWithURL allows creating a new Client with a hardcoded URL. Useful for testing purposes
func NewClientWithURL(sk, cc, partnerKey, url string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
	return newClientFull(sk, cc, partnerKey, url, httpCli, headersForEveryRequestFunc)
}

func NewClient(sk, cc, partnerKey string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
	return newClientFull(sk, cc, partnerKey, "", httpCli, headersForEveryRequestFunc)
}

func newClientFull(sk, cc, partnerKey, url string, httpCli *http.Client, headersForEveryRequestFunc AuthFunc) *Client {
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
	if url == "" {
		if cc != "" {
			cli.Url = GetBaseURL(cc)
		} else {
			cli.Url = GetBaseURLFromAuthFunc(cli.headersFunc)
		}
	} else {
		cli.Url = url
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
